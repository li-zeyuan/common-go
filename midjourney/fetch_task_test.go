package midjourney

import (
	"context"
	"testing"

	"github.com/li-zeyuan/common-go/model"
	"github.com/stretchr/testify/assert"
)

func TestFetchTasks(t *testing.T) {
	cli := NewCli()

	req := new(model.FetchTasksReq)
	//req.Ids = []string{"1719395287129922"}
	req.Ids = []string{"1719922852934765"}
	resp, err := cli.FetchTasks(context.Background(), req)
	assert.Nil(t, err)

	t.Log(resp)
}
