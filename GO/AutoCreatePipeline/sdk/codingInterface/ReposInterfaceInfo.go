package codingInterface

/*
获取项目下已登记的代码库列表
Url="http://tencent.coding.woa.com/api/workflow/rest-api/product/repositories"
Method = "GET"
headers = {"Authorization": "token xxxxx","X-Plat-Id": 1,"X-Project-Id": 项目ID}
*/
type ReposReq struct {
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	RepoName string `json:"repoName"`
}

type ReposResp struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Total    int    `json:"total"`
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	Repos    []struct {
		RepoID      int    `json:"repoId"`
		Type        string `json:"type"`
		Identity    string `json:"identity"`
		Name        string `json:"name"`
		FullName    string `json:"fullName"`
		HTTPURL     string `json:"httpUrl"`
		Creator     string `json:"creator"`
		Description string `json:"description"`
		CreatedAt   string `json:"createdAt"`
		ID          int    `json:"id"`
		Ids         []struct {
			ID         int    `json:"id"`
			Permission string `json:"permission"`
		} `json:"ids"`
		CodeAnalysisResult []interface{} `json:"codeAnalysisResult"`
	} `json:"repos"`
}
