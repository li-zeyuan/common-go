package minioclient

import (
	"bytes"
	"context"
	"io"

	"github.com/li-zeyuan/common-go/mylogger"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
)

type Client struct {
	conf   *Config
	client *minio.Client
}

func New(ctx context.Context, conf *Config) (*Client, error) {
	minioClient, err := minio.New(conf.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(conf.AccessKeyID, conf.SecretAccessKey, ""),
	})
	if err != nil {
		mylogger.Error(ctx, "new minio client fail", zap.Error(err))
		return nil, err
	}

	return &Client{
		conf:   conf,
		client: minioClient,
	}, nil
}

func (c *Client) GetConfig() *Config {
	return c.conf
}

func (c *Client) CreateBucketIfNotExist(ctx context.Context) error {
	found, err := c.client.BucketExists(ctx, c.conf.Bucket)
	if err != nil {
		zap.L().Error("check bucket if exist fail", zap.Error(err))
		return err
	}
	if found {
		return nil
	}

	err = c.client.MakeBucket(ctx, c.conf.Bucket, minio.MakeBucketOptions{Region: "us-east-1", ObjectLocking: false})
	if err != nil {
		zap.L().Error("create bucket fail", zap.Error(err))
		return err
	}

	return nil
}

func (c *Client) PresignedPutObject(ctx context.Context, objectKey string) (string, error) {
	url, err := c.client.PresignedPutObject(ctx, c.conf.Bucket, objectKey, c.conf.PresignedPutExpiry)
	if err != nil {
		zap.L().Error("presigned put object fail", zap.Error(err))
		return "", err
	}

	return url.String(), nil
}

func (c *Client) PresignedGetObject(ctx context.Context, objectKey string) (string, error) {
	if len(objectKey) == 0 {
		return "", nil
	}

	url, err := c.client.PresignedGetObject(ctx, c.conf.Bucket, objectKey, c.conf.PresignedGetExpiry, nil)
	if err != nil {
		zap.L().Error("presigned get object fail", zap.Error(err))
		return "", err
	}

	return url.String(), nil
}

func (c *Client) PutObject(ctx context.Context, objectKey string, buf []byte) error {
	if len(buf) == 0 {
		return nil
	}

	_, err := c.client.PutObject(ctx, c.conf.Bucket, objectKey, bytes.NewReader(buf), int64(len(buf)), minio.PutObjectOptions{})
	if err != nil {
		zap.L().Error("put object fail", zap.Error(err))
		return err
	}

	return nil
}

func (c *Client) GetObject(ctx context.Context, objectKey string) ([]byte, error) {
	obj, err := c.client.GetObject(ctx, c.conf.Bucket, objectKey, minio.GetObjectOptions{})
	if err != nil {
		zap.L().Error("get object fail", zap.Error(err))
		return nil, err
	}

	b, err := io.ReadAll(obj)
	if err != nil {
		zap.L().Error("read object fail", zap.Error(err))
		return nil, err
	}

	return b, nil
}

func (c *Client) Close() {
	return
}
