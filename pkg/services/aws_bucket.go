package services

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"mime"
	"path/filepath"
	"strings"
	"time"

	"Restro/pkg/framework"
	"Restro/pkg/utils"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/bluele/gcache"
)

type FileData interface {
	Key() string
	FileName() string
}

type S3BucketService struct {
	*s3.Client
	uploader *manager.Uploader
	preSign  *s3.PresignClient
	logger   framework.Logger
	env      *framework.Env
	cache    gcache.Cache
}

func NewS3BucketService(
	logger framework.Logger,
	client *s3.Client,
	env *framework.Env,
	preSign *s3.PresignClient,
	uploader *manager.Uploader,
) *S3BucketService {
	cache := gcache.New(2000).ARC().Build()
	return &S3BucketService{
		logger:   logger,
		Client:   client,
		env:      env,
		preSign:  preSign,
		uploader: uploader,
		cache:    cache,
	}
}

// UploadFile uploads the fileupload
func (s *S3BucketService) UploadFile(
	ctx context.Context,
	file io.Reader,
	key string,
) (*manager.UploadOutput, error) {
	input := &s3.PutObjectInput{
		Bucket: aws.String(s.env.S3BucketName),
		Key:    aws.String(key),
		Body:   file,
	}
	s.logger.Info("KEY::", key, s.env.S3BucketName)
	return s.uploader.Upload(ctx, input)
}

func (s *S3BucketService) UploadFilePublicly(
	ctx context.Context,
	file io.Reader,
	key string,
) (*manager.UploadOutput, error) {
	input := &s3.PutObjectInput{
		Bucket: aws.String(s.env.S3BucketName),
		Key:    &key,
		Body:   file,
		ACL:    types.ObjectCannedACLPublicRead,
	}

	return s.uploader.Upload(ctx, input)
}

// GetSignedURL get the signed url for fileupload
func (s *S3BucketService) GetSignedURL(
	ctx context.Context,
	fileData FileData,
	folder string,
	expires *time.Time,
) string {
	key := fmt.Sprintf("%v::%v", folder, fileData.Key())

	if url, err := s.cache.Get(key); err == nil {
		return url.(string)
	}

	if expires == nil {
		e := time.Now().Add(24 * time.Hour)
		expires = &e
	}

	dst := fmt.Sprintf("%v/%v", folder, fileData.Key())

	fileName := ""
	if fileData.FileName() == "" {
		fileName = fileData.Key()
	} else {
		fileName = fileData.FileName()
	}

	extension := strings.ToLower(filepath.Ext(fileData.Key()))
	mimeType := mime.TypeByExtension(extension)

	input := &s3.GetObjectInput{
		Bucket:                     aws.String(s.env.S3BucketName),
		Key:                        aws.String(dst),
		ResponseContentDisposition: aws.String(fmt.Sprintf("inline; filename=\"%v\"", fileName)),
		ResponseContentType:        aws.String(mimeType),
	}
	duration := time.Until(*expires)

	resp, err := s.preSign.PresignGetObject(ctx, input, s3.WithPresignExpires(duration))
	if err != nil {
		s.logger.Error("error-generating-presigned-url", err.Error())
		return ""
	}

	// generate random time between 1/3 duration and 2/3 duration
	// so that not every cache expires at the same time
	maxDuration := int(duration.Nanoseconds() / 3)
	cacheDuration := time.Duration(rand.Intn(maxDuration) + maxDuration)
	if err := s.cache.SetWithExpire(key, resp.URL, cacheDuration); err != nil {
		s.logger.Error("error-setting-cache", err.Error())
	}

	return resp.URL
}

func (s *S3BucketService) CopyToFolder(objectKey, folderName string) (string, error) {
	bucketName := s.env.S3BucketName
	src := fmt.Sprintf("%v/%v/%v", bucketName, s.env.S3TempFolder, objectKey)
	dst := fmt.Sprintf("%v/%v", folderName, objectKey)

	_, err := s.Client.CopyObject(context.TODO(), &s3.CopyObjectInput{
		Bucket:     aws.String(bucketName),
		CopySource: aws.String(src),
		Key:        aws.String(dst),
	})

	if err != nil {
		s.logger.Error(fmt.Sprintf("Couldn't copy object from %v to %v. Here's why: %v\n",
			src, dst, err))
		awsError := utils.MapAWSError(s.logger, err)
		if awsError != nil {
			return "", awsError
		}
		return "", err
	}
	return dst, nil
}

func (s *S3BucketService) DeleteFile(fileKey string) error {
	bucketName := s.env.S3BucketName
	src := fmt.Sprintf("%v/%v", s.env.S3TempFolder, fileKey)
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(src),
	}
	_, err := s.DeleteObject(context.TODO(), input)
	if err != nil {
		return err
	}
	return nil
}

func (s *S3BucketService) MoveToDestFolder(fileKey string) (string, error) {
	dst, err := s.CopyToFolder(fileKey, s.env.S3AttachmentFolder)
	if err != nil {
		return "", err
	}

	if err := s.DeleteFile(fileKey); err != nil {
		return "", err
	}
	return dst, nil
}
