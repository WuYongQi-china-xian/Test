package sdk

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func SendReqJson(method, url string, bodyJson []byte, headers map[string]string) []byte {

	client := &http.Client{}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(bodyJson))
	if err != nil {
		log.Println(err)
		panic(fmt.Sprintf("new req fail, error: %s", err))
	}
	//req.Close = true // 关闭常链接
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		panic(fmt.Sprintf("send req fail, error: %s", err))
	}

	defer resp.Body.Close()
	log.Println("start-------------------------------->")
	log.Println(fmt.Sprintf("method:%s", method))
	log.Println(fmt.Sprintf("url:%s", url))
	log.Println(fmt.Sprintf("headers:%v", req.Header))
	log.Println(fmt.Sprintf("body:%s", bodyJson))
	log.Println(fmt.Sprintf("resp:%v", resp))
	log.Println("<-----------------------------------end")
	defer client.CloseIdleConnections()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		panic(fmt.Sprintf("read resp fail, error: %s", err))
	}
	log.Println(fmt.Sprintf("resp.Body:%s", string(body)))
	if resp.StatusCode == 200 || resp.StatusCode == 201 {
		return body
	} else {
		return []byte("")
	}

}
