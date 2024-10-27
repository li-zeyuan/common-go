package robot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/li-zeyuan/common-go/environment"
	"github.com/li-zeyuan/common-go/model"
)

const (
	TitleServerAlter    = "server alter"
	TitleServerNotify   = "server notify"
	TitleBusinessAlter  = "business alter"
	TitleBusinessNotify = "business notify"
)

var _robot *Robot

type Robot struct {
	robotUrl string
}

func Init(url string) {
	if len(url) == 0 {
		return
	}

	_robot = &Robot{
		robotUrl: url,
	}
}

func Send(ct *model.WeComRobotContent) {
	if _robot == nil {
		return
	}

	req := new(model.WeComRobotReq)
	req.MsgType = "text"
	req.Text = new(model.WeComRobotReqText)
	req.Text.Content = fmt.Sprintf("title: %s\nenv: %s\nuid: %s\nrequest_id: %s\nmessage: %s",
		ct.Title, environment.GetEnv(), ct.Uid, ct.RequestId, ct.Message)
	req.Text.MentionedMobileList = []string{"@all"}
	body, err := json.Marshal(req)
	if err != nil {
		log.Printf("marshal robot request fail: %v", err)
		return
	}

	httpReq, err := http.NewRequest("POST", _robot.robotUrl, bytes.NewReader(body))
	if err != nil {
		log.Printf("new robot fail: %v", err)
		return
	}

	cli := http.Client{}
	httpResp, err := cli.Do(httpReq)
	if err != nil {
		log.Printf("do robot request fail: %v", err)
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		respBody, err := io.ReadAll(httpResp.Body)
		if err != nil {
			log.Printf("read body fail: %v", err)
			return
		}

		log.Printf("no 200 return, body: %s, error: %v", string(respBody), err)
		return
	}
}
