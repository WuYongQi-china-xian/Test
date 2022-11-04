package sdk

import (
	codinginterface "AutoCreatePipeline/sdk/codingInterface"
	"encoding/json"
	"fmt"
	"log"
)

const CodingRootApi = "http://tencent.coding.woa.com"

var headers = make(map[string]string)

/*
根据模板id查询模板信息
*/
func GetTemplateInfo(templateId, token string) codinginterface.TemplateResp {
	var template codinginterface.TemplateResp
	url := fmt.Sprintf("%s/api/qci/rest-api/template/%s", CodingRootApi, templateId)
	headers["Authorization"] = fmt.Sprintf("token %s", token)
	resp := SendReqJson("GET", url, nil, headers)
	err := json.Unmarshal(resp, &template)
	if err != nil {
		log.Println(err)
	}
	return template
}

/*
根据模板信息中流水线配置id查询流水线信息
*/
func GetPipelineConfigInfo(piplineConfigID, token string) codinginterface.PipelineConfigResp {
	var pipelineConfig codinginterface.PipelineConfigResp
	url := fmt.Sprintf("%s/api/qci/rest-api/pipeline/%s", CodingRootApi, piplineConfigID)
	headers["Authorization"] = fmt.Sprintf("token %s", token)
	resp := SendReqJson("GET", url, nil, headers)
	err := json.Unmarshal(resp, &pipelineConfig)
	if err != nil {
		log.Println(err)
	}
	return pipelineConfig
}

/*
根据流水线id查询指定流水线信息
*/
func GetPipelineInfo(pipelineId, token string) codinginterface.PipelineInfoResp {
	var pipeline codinginterface.PipelineInfoResp
	url := fmt.Sprintf("%s/api/qci/rest-api/pipeline/%s", CodingRootApi, pipelineId)
	headers["Authorization"] = fmt.Sprintf("token %s", token)
	resp := SendReqJson("GET", url, nil, headers)
	err := json.Unmarshal(resp, &pipeline)
	if err != nil {
		log.Println(err)
	}
	return pipeline
}

/*
创建流水线
*/
func CreatePipeline(piepelineReq codinginterface.CreatePipelineReq, token string) codinginterface.CreatePipelineResp {
	var pipelineResp codinginterface.CreatePipelineResp
	url := fmt.Sprintf("%s/api/qci/rest-api/v2/pipeline", CodingRootApi)
	headers["Authorization"] = fmt.Sprintf("Token %s", token)
	body, err := json.Marshal(piepelineReq)
	log.Println(fmt.Sprintf("body=%s", body))
	if err != nil {
		log.Fatal(err)
	}
	resp := SendReqJson("POST", url, body, headers)
	err = json.Unmarshal(resp, &pipelineResp)
	if err != nil {
		log.Fatal(fmt.Sprintf("解析创建流水线的响应体失败:%v", err))
	}
	return pipelineResp
}

/*
根据git http url 获取仓库详情
*/
func GetReposInfo(reposReq codinginterface.ReposReq, projectId, token string) codinginterface.ReposResp {
	var reposInfoResp codinginterface.ReposResp
	url := fmt.Sprintf("%s/api/workflow/rest-api/product/repositories", CodingRootApi)
	headers["Authorization"] = fmt.Sprintf("token %s", token)
	headers["X-Plat-Id"] = "1"
	headers["X-Project-Id"] = projectId
	body, err := json.Marshal(reposReq)
	log.Println(fmt.Sprintf("body=%s", body))
	if err != nil {
		log.Fatal(err)
	}
	resp := SendReqJson("GET", url, body, headers)
	err = json.Unmarshal(resp, &reposInfoResp)
	if err != nil {
		log.Println(err)
	}
	return reposInfoResp
}
