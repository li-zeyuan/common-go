package midjourney

import (
	"context"
	"testing"

	"github.com/li-zeyuan/common-go/model"
	"github.com/stretchr/testify/assert"
)

func TestAction(t *testing.T) {
	taskID := "1719375142541391"
	cli := NewCli()

	fReq := new(model.FetchTasksReq)
	fReq.Ids = []string{taskID}
	fMap, err := cli.FetchTasks(context.Background(), fReq)
	if err != nil {
		t.Fatal(err)
	}

	taskResp, ok := fMap[taskID]
	if !ok {
		t.Fatal("no exist task")
	}

	req := new(model.ActionReq)
	req.TaskId = taskID
	req.CustomId = taskResp.Buttons[0].CustomId

	resp, err := cli.Action(context.Background(), req)
	assert.Nil(t, err)

	t.Log(resp)
}
