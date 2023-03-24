package publictool

import (
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
)

// 接口文档：https://iwiki.woa.com/pages/viewpage.action?pageId=541885776

const PigeonRootApi = "http://dev.ngate.tencent-cloud.com"

var headers = make(map[string]string)

type PigeonCommonRequest struct {
	BotKey string   `json:"botKey"`
	ChatID []string `json:"chatId"`
}

type MarkDownContentRequest struct {
	PigeonCommonRequest
	MsgContent string `json:"msgContent"`
}

type MarkDownButtonRequest struct {
	PigeonCommonRequest
	MsgContent string `json:"msgContent"`
	ExtButton  string `json:"extButton"`
	AutoSplit  string `json:"autoSplit"`
}

/*
发送markdown文字
*/
func SendMarkDownContentRequest(BotKey, MsgContent string, ChatID []string) {
	var mdcr MarkDownContentRequest
	mdcr.MsgContent = MsgContent
	mdcr.BotKey = BotKey
	mdcr.ChatID = ChatID
	url := fmt.Sprintf("%s/pigeon/v1/wecom/bot/markdown", PigeonRootApi)
	headers["Content-Type"] = "application/json"
	body, err := json.Marshal(mdcr)
	logx.Info(fmt.Sprintf("body=%s", body))
	if err != nil {
		log.Fatal(err)
	}
	SendReqJson("POST", url, body, headers)
	//resp := SendReqJson("POST", url, body, headers)
	//err := json.Unmarshal(resp, &template)
	//if err != nil {
	//	log.Println(err)
	//}
	//return template
}
