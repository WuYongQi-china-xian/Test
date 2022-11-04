package business

import (
	sdk "AutoCreatePipeline/sdk"
	"AutoCreatePipeline/sdk/codingInterface"
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

/*
根据工蜂url获取repo_id
repoName: "http://git.woa.com/wolfwwu/WorkerBeeToGithub.git"
projectId: coding项目id 2197
token: coding令牌
*/
func GetRepoId(projectId, repoName, token string) int {
	var repoReq codingInterface.ReposReq
	repoReq.RepoName = repoName
	repoReq.Page = 1
	repoReq.PageSize = 200
	ReposRespInfo := sdk.GetReposInfo(repoReq, projectId, token)
	repoId := ReposRespInfo.Repos[0].ID
	log.Println(fmt.Sprintf("RepoID:%v", repoId))
	return repoId
}
