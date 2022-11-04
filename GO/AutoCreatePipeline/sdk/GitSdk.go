package sdk

import (
	"AutoCreatePipeline/sdk/gitInterface"
	"encoding/json"
	"fmt"
	"log"
	url2 "net/url"
	"strings"
)

var gitRootApi = "http://git.woa.com"
var gitHeaders = make(map[string]string)

/*
获取工蜂仓库详细信息
gitRepoUrl:工蜂仓库http url
token:工蜂令牌
*/
func GetProjectID(gitRepoUrl, token string) gitInterface.ProjectInfoResp {
	tmp := strings.Replace(gitRepoUrl, fmt.Sprintf("%s/", gitRootApi), "", 1)
	projectFullPath := strings.Replace(tmp, ".git", "", 1)
	log.Println(fmt.Sprintf("FullName:%s", projectFullPath))
	url := fmt.Sprintf("%s/api/v3/projects/%s", gitRootApi, url2.QueryEscape(projectFullPath))
	gitHeaders["PRIVATE-TOKEN"] = token
	var projectInfo gitInterface.ProjectInfoResp
	resp := SendReqJson("GET", url, nil, gitHeaders)
	err := json.Unmarshal(resp, &projectInfo)
	if err != nil {
		log.Println(err)
	}
	return projectInfo
}
