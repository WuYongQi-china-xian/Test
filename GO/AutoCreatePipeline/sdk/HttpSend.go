package sdk

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"strings"
)

func SendReqJson(method, url, bodyJson string, headers map[string]string) []byte {

	client := &http.Client{}

	req, err := http.NewRequest(method, url, strings.NewReader(bodyJson))
	if err != nil {
		log.Println(err)
		panic(fmt.Sprintf("new req fail, error: %s", err))
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		panic(fmt.Sprintf("sedn req fail, error: %s", err))
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
	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			panic(fmt.Sprintf("read resp fail, error: %s", err))
		}
		log.Println(fmt.Sprintf("resp.Body:%s", string(body)))
		return body
	} else {
		return []byte("")
	}

}

