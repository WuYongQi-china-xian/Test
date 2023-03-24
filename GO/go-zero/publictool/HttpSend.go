package publictool

import (
	"bytes"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"io/ioutil"
	"net/http"
)

func SendReqJson(method, url string, bodyJson []byte, headers map[string]string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(bodyJson))
	if err != nil {
		logx.Info(err)
		panic(fmt.Sprintf("new req fail, error: %s", err))
	}
	//req.Close = true // 关闭常链接
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		logx.Info(err)
		panic(fmt.Sprintf("send req fail, error: %s", err))
	}
	defer resp.Body.Close()
	logx.Info("start-------------------------------->")
	logx.Info(fmt.Sprintf("method:%s", method))
	logx.Info(fmt.Sprintf("url:%s", url))
	logx.Info(fmt.Sprintf("headers:%v", req.Header))
	logx.Info(fmt.Sprintf("body:%s", bodyJson))
	logx.Info(fmt.Sprintf("resp:%v", resp))
	logx.Info("<-----------------------------------end")
	defer client.CloseIdleConnections()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logx.Info(err)
		panic(fmt.Sprintf("read resp fail, error: %s", err))
	}
	logx.Info(fmt.Sprintf("resp.Body:%s", string(body)))
	if resp.StatusCode == 200 || resp.StatusCode == 201 {
		return body
	} else {
		return []byte("")
	}
}
