package minioclient

import (
	"context"
	"testing"

	"github.com/li-zeyuan/common-go/utils"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	conf := &Config{}
	err := utils.DecodeConfigFile("/Users/zeyuan.li/Desktop/workspace/ggo/src/github.com/li-zeyuan/sun/testdata/dev_config.yaml", conf)
	if err != nil {
		t.Fatal(err)
	}

	New(context.Background(), conf)
}

func TestCreateBucketIfNotExist(t *testing.T) {
	conf := &Config{}
	err := utils.DecodeConfigFile("/Users/zeyuan.li/Desktop/workspace/ggo/src/github.com/li-zeyuan/sun/testdata/dev_config.yaml", conf)
	if err != nil {
		t.Fatal(err)
	}

	//conf.Bucket = "test"

	cli, err := New(context.Background(), conf)
	err = cli.CreateBucketIfNotExist(context.Background())
	assert.Nil(t, err)
}

func TestPresignedGetObject(t *testing.T) {
	conf := &Config{}
	err := utils.DecodeConfigFile("/Users/zeyuan.li/Desktop/workspace/ggo/src/github.com/li-zeyuan/sun/testdata/dev_config.yaml", conf)
	if err != nil {
		t.Fatal(err)
	}

	cli, err := New(context.Background(), conf)
	url, err := cli.PresignedGetObject(context.Background(), "ai_logo (1).jpg")
	assert.Nil(t, err)
	t.Log(url)
}

func TestPublicGetObject(t *testing.T) {
	conf := &Config{}
	err := utils.DecodeConfigFile("/Users/zeyuan.li/Desktop/workspace/ggo/src/github.com/li-zeyuan/sun/testdata/dev_config.yaml", conf)
	if err != nil {
		t.Fatal(err)
	}

	cli := &Client{conf: conf}
	url, err := cli.PublicGetObject(context.Background(), "ai_logo (1).jpg")
	assert.Nil(t, err)
	t.Log(url)
}
