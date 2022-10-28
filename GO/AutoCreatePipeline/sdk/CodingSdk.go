package sdk

import (
	codinginterface "AutoCreatePipeline/sdk/codingInterface"
	"encoding/json"
	"fmt"
)

var rootApi = "http://tencent.coding.woa.com"
var headers = make(map[string]string)

func GetTemplateInfo(templateId, token string) codinginterface.TemplateResp {
	var template codinginterface.TemplateResp
	url := fmt.Sprintf("%s/api/qci/rest-api/template/%s", rootApi, templateId)
	headers["Authorization"] = fmt.Sprintf("token %s", token)
	resp := SendReqJson("GET", url, "", headers)
	err := json.Unmarshal(resp, &template)
	if err != nil {
		fmt.Println(err)
	}
	return template
}

func GetPipelineConfigInfo(piplineConfigID, token string) codinginterface.PipelineConfigResp {
	var pipelineConfig codinginterface.PipelineConfigResp
	url := fmt.Sprintf("%s/api/qci/rest-api/pipeline/%s", rootApi, piplineConfigID)
	headers["Authorization"] = fmt.Sprintf("token %s", token)
	resp := SendReqJson("GET", url, "", headers)
	err := json.Unmarshal(resp, &pipelineConfig)
	if err != nil {
		fmt.Println(err)
	}
	return pipelineConfig
}
