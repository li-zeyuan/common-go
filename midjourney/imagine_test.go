package midjourney

import (
	"context"
	"encoding/base64"
	"os"
	"testing"

	"github.com/li-zeyuan/common-go/model"
	"github.com/li-zeyuan/common-go/mylogger"
	"github.com/stretchr/testify/assert"
)

func NewCli() *Client {
	mylogger.Init(nil)

	conf := NewDefault()
	cli := NewClient(conf)
	return cli
}

func TestImagine(t *testing.T) {
	cli := NewCli()
	imgBytes, err := os.ReadFile("/Users/zeyuan.li/Desktop/workspace/ggo/src/github.com/li-zeyuan/common-go/midjourney/4wbc.png")
	if err != nil {
		t.Fatal(err)
	}

	req := new(model.ImagineReq)
	req.Modes = make([]string, 0)
	req.Remix = true
	req.Base64Array = []string{"data:image/png;base64," + base64.StdEncoding.EncodeToString(imgBytes)}
	req.Prompt = "A newborn Chinese babyeyes closedwhite s kin.rosyfacewhite swaddlingclothessurrounded by dreamy rose flowerswith some light coming infrom the side against a dreamy white background the composition is beautifulrich andbright"
	resp, err := cli.Imagine(context.Background(), req)
	assert.Nil(t, err)

	t.Log(resp)
}
