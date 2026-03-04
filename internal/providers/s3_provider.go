package providers

import (
	"context"
	"io"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Provider struct {
	Client     *s3.Client
	BucketName string
	PublicUrl  string
	Presign    *s3.PresignClient
}

type S3Config struct {
	AccountId   string
	AccessKey   string
	SecretKey   string
	BucketName  string
	PublicUrl   string
	ApiEndpoint string
}

func NewS3Config(accountId, accessKey, secretKey, bucketName, publicUrl, apiEndpoint string) *S3Config {
	return &S3Config{
		AccountId:   accountId,
		AccessKey:   accessKey,
		SecretKey:   secretKey,
		BucketName:  bucketName,
		PublicUrl:   publicUrl,
		ApiEndpoint: apiEndpoint,
	}
}

func NewS3Provider(cfg *S3Config) *S3Provider {
	accountId := cfg.AccountId
	accessKey := cfg.AccessKey
	secretKey := cfg.SecretKey
	bucketName := cfg.BucketName
	publicUrl := cfg.PublicUrl
	apiEndpoint := cfg.ApiEndpoint

	if accountId == "" || accessKey == "" || secretKey == "" || bucketName == "" || publicUrl == "" || apiEndpoint == "" {
		log.Fatal("Missing s3 enviroment variables")
	}

	m_cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				accessKey,
				secretKey,
				"",
			),
		),
		config.WithRegion("auto"),
		config.WithBaseEndpoint(apiEndpoint),
	)

	if err != nil {
		log.Fatal("failed to load config", err)
	}

	client := s3.NewFromConfig(m_cfg)
	presign := s3.NewPresignClient(client)

	return &S3Provider{
		Client:     client,
		BucketName: bucketName,
		PublicUrl:  publicUrl,
		Presign:    presign,
	}
}

func (r *S3Provider) Upload(file io.Reader, fileName string) (string, error) {
	_, err := r.Client.PutObject(
		context.TODO(),
		&s3.PutObjectInput{
			Bucket: &r.BucketName,
			Key:    &fileName,
			Body:   file,
		},
	)
	if err != nil {
		return "", err
	}
	return fileName, nil
}

func (r *S3Provider) PresignedUrl(fileName string) (string, error) {
	return fileName, nil
}
