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

	New(conf)
}

func TestCreateBucketIfNotExist(t *testing.T) {
	conf := &Config{}
	err := utils.DecodeConfigFile("/Users/zeyuan.li/Desktop/workspace/ggo/src/github.com/li-zeyuan/sun/testdata/dev_config.yaml", conf)
	if err != nil {
		t.Fatal(err)
	}

	cli, err := New(conf)
	err = cli.CreateBucketIfNotExist(context.Background())
	assert.Nil(t, err)
}
