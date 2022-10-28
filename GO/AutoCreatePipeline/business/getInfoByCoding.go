package business

import (
	sdk "AutoCreatePipeline/sdk"
	"fmt"
	"log"
	"strconv"
)

/*
https://mojotv.cn/go/bad-go-pointer-returns
通常情况下我们最好还是return value
生命时间短的struct return pointer是不推荐的
大size的struct推荐使用pointer返回
*/
//通过模板id获取环境变量
func GetEnvKey(templateID, token string) []string {
	var envNameArray []string
	templateResp := sdk.GetTemplateInfo(templateID, token)
	pipelineConfigID := templateResp.PipelineConfigID
	log.Println(fmt.Sprintln("pipelineConfigID=%s", pipelineConfigID))
	pipelineConfigResp := sdk.GetPipelineConfigInfo(strconv.Itoa(pipelineConfigID), token)
	envMap := pipelineConfigResp.PipelineJSON.Env
	for envName, envOtherInfo := range envMap {
		log.Println(envName, envOtherInfo.Desc)
		envNameArray = append(envNameArray, envName)
	}
	return envNameArray
}
