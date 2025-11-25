package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type Storage struct {
	client   *s3.Client
	bucket   string
	endpoint string
}

type KeypairData struct {
	PublicKey       string `json:"publicKey"`
	SecretKey       []byte `json:"secretKey"`
	SecretKeyBase58 string `json:"secretKeyBase58"`
}

func NewStorage(endpoint, region, bucket, accessKey, secretKey string) (*Storage, error) {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:               endpoint,
			SigningRegion:     region,
			HostnameImmutable: true,
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithEndpointResolverWithOptions(customResolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	return &Storage{
		client: s3.NewFromConfig(cfg, func(o *s3.Options) {
			o.UsePathStyle = true
		}),
		bucket:   bucket,
		endpoint: endpoint,
	}, nil
}

func (s *Storage) UploadKeypair(publicKey string, secretKey []byte, secretKeyBase58 string) (string, error) {
	keypairData := KeypairData{
		PublicKey:       publicKey,
		SecretKey:       secretKey,
		SecretKeyBase58: secretKeyBase58,
	}

	jsonData, err := json.MarshalIndent(keypairData, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal keypair data: %w", err)
	}

	fileName := fmt.Sprintf("vanity-keypairs/%s.json", publicKey)
	contentType := "application/json"
	acl := types.ObjectCannedACLPrivate

	_, err = s.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(fileName),
		Body:        bytes.NewReader(jsonData),
		ContentType: &contentType,
		ACL:         acl,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload to S3: %w", err)
	}

	spacesURL := fmt.Sprintf("%s/%s/%s", s.endpoint, s.bucket, fileName)
	return spacesURL, nil
}
