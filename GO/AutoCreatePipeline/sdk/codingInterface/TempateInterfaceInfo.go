package codingInterface

/*
Url="http://tencent.coding.woa.com/api/qci/rest-api/template/22868"
Method = "GET"
headers = {"Authorization": "token xxxxx"}
 */

type TemplateReq struct {}
type TemplateResp struct {
	ID                 int           `json:"id"`
	Name               string        `json:"name"`
	Desc               interface{}   `json:"desc"`
	PipelineConfigID   int           `json:"pipeline_config_id"`
	Admin              string        `json:"admin"`
	Member             interface{}   `json:"member"`
	Icon               interface{}   `json:"icon"`
	Category           []interface{} `json:"category"`
	Usage              interface{}   `json:"usage"`
	IsRecommended      bool          `json:"is_recommended"`
	State              int           `json:"state"`
	ProjectID          string        `json:"project_id"`
	Scope              int           `json:"scope"`
	PipelineCnt        int           `json:"pipeline_cnt"`
	Power              int           `json:"power"`
	PipelineConfigJSON struct {
		Name                       string        `json:"name"`
		Desc                       interface{}   `json:"desc"`
		CodeSource                 string        `json:"code_source"`
		YamlLocation               interface{}   `json:"yaml_location"`
		CurRevision                interface{}   `json:"cur_revision"`
		CodeBranch                 []interface{} `json:"code_branch"`
		CodeSourceType             string        `json:"code_source_type"`
		DefaultCodeBranch          interface{}   `json:"default_code_branch"`
		ID                         int           `json:"id"`
		ProductID                  interface{}   `json:"product_id"`
		Admin                      string        `json:"admin"`
		State                      int           `json:"state"`
		Creator                    string        `json:"creator"`
		Created                    string        `json:"created"`
		Timeout                    interface{}   `json:"timeout"`
		PipelineJSON               string        `json:"pipeline_json"`
		Modifier                   string        `json:"modifier"`
		Updated                    string        `json:"updated"`
		UserEnv                    string        `json:"user_env"`
		CurrentRunResult           interface{}   `json:"current_run_result"`
		Members                    string        `json:"members"`
		TriggerType                string        `json:"trigger_type"`
		RepoUsername               interface{}   `json:"repo_username"`
		RepoPassword               interface{}   `json:"repo_password"`
		AuthType                   int           `json:"auth_type"`
		Power                      int           `json:"power"`
		ProductName                interface{}   `json:"product_name"`
		StateName                  string        `json:"state_name"`
		IsNochangeBuild            bool          `json:"is_nochange_build"`
		IsBlockMr                  bool          `json:"is_block_mr"`
		CredentialID               interface{}   `json:"credential_id"`
		Visibility                 int           `json:"visibility"`
		MaximumBuilds              int           `json:"maximum_builds"`
		IsLocalMr                  bool          `json:"is_local_mr"`
		YamlSourceType             int           `json:"yaml_source_type"`
		BuildVersion               string        `json:"build_version"`
		Webhook                    interface{}   `json:"webhook"`
		IsAutoCancel               bool          `json:"is_auto_cancel"`
		AutoCancelSameMergeRequest bool          `json:"auto_cancel_same_merge_request"`
		ProjectID                  int           `json:"project_id"`
		TemplateID                 int           `json:"template_id"`
		ConfigType                 int           `json:"config_type"`
		IsNoRevert                 bool          `json:"is_no_revert"`
		RepoID                     interface{}   `json:"repo_id"`
		OpTypeData                 struct {
			ID     interface{} `json:"id"`
			OpType interface{} `json:"op_type"`
		} `json:"op_type_data"`
		ProjectInfo struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			DisplayName string `json:"display_name"`
			URL         string `json:"url"`
			Icon        string `json:"icon"`
			PlmCatalog  struct {
				ID         int    `json:"id"`
				CatalogID  int    `json:"catalogId"`
				Code       string `json:"code"`
				Level      int    `json:"level"`
				NameZh     string `json:"nameZh"`
				NameEn     string `json:"nameEn"`
				PStatus    int    `json:"pStatus"`
				Admin      string `json:"admin"`
				ParentCode string `json:"parentCode"`
				Dept1ID    int    `json:"dept1Id"`
				Dept1Name  string `json:"dept1Name"`
				Dept2ID    int    `json:"dept2Id"`
				Dept2Name  string `json:"dept2Name"`
				Intro      string `json:"intro"`
				Version    int    `json:"version"`
				BrandName  string `json:"brandName"`
				CreatedAt  string `json:"createdAt"`
				UpdatedAt  string `json:"updatedAt"`
				DeletedAt  string `json:"deletedAt"`
			} `json:"plm_catalog"`
			Belonging int `json:"belonging"`
		} `json:"project_info"`
	} `json:"pipeline_config_json"`
	TeamID             interface{}   `json:"team_id"`
	IsForceSync        bool          `json:"is_force_sync"`
	IsConstrained      bool          `json:"is_constrained"`
	Version            interface{}   `json:"version"`
	VersionCount       int           `json:"version_count"`
	Creator            string        `json:"creator"`
	Created            string        `json:"created"`
	Updated            string        `json:"updated"`
	Modifier           string        `json:"modifier"`
	AuthorizedProjects []interface{} `json:"authorized_projects"`
}

/*
Url="http://tencent.coding.woa.com/api/qci/rest-api/pipeline/798430"
Method = "GET"
headers = {"Authorization": "token xxxxx"}
*/

type PipelineConfigResp struct {
	Name              string        `json:"name"`
	Desc              interface{}   `json:"desc"`
	CodeSource        string        `json:"code_source"`
	YamlLocation      interface{}   `json:"yaml_location"`
	CurRevision       interface{}   `json:"cur_revision"`
	CodeBranch        []interface{} `json:"code_branch"`
	CodeSourceType    string        `json:"code_source_type"`
	DefaultCodeBranch interface{}   `json:"default_code_branch"`
	ID                int           `json:"id"`
	ProductID         interface{}   `json:"product_id"`
	Admin             string        `json:"admin"`
	State             int           `json:"state"`
	Creator           string        `json:"creator"`
	Created           string        `json:"created"`
	Timeout           interface{}   `json:"timeout"`
	PipelineJSON      struct {
		Stages []struct {
			Stage string `json:"stage,omitempty"`
			Tasks []struct {
				Task string `json:"task"`
				Cmds []struct {
					Plugin string `json:"plugin"`
					Params struct {
						Cmds []string `json:"cmds"`
					} `json:"params"`
					Label string `json:"label,omitempty"`
				} `json:"cmds"`
				If    string   `json:"if"`
				Temps []string `json:"temps"`
			} `json:"tasks,omitempty"`
			Prompt []struct {
				Msg           string   `json:"msg"`
				To            []string `json:"to"`
				Detail        string   `json:"detail"`
				TimeoutStatus int      `json:"timeout_status"`
				Env           struct {
					Configuration struct {
						Value  string `json:"value"`
						Type   string `json:"type"`
						Label  string `json:"label"`
						Desc   string `json:"desc"`
						Option string `json:"option"`
					} `json:"Configuration"`
				} `json:"env"`
				Notifications struct {
					Channel       string `json:"channel"`
					EnwechatGroup string `json:"enwechat_group"`
					Mentions      string `json:"mentions"`
				} `json:"notifications"`
			} `json:"prompt,omitempty"`
			If string `json:"if,omitempty"`
		} `json:"stages"`
		Version string `json:"version"`
		Worker  struct {
			Label  string        `json:"label"`
			Tools  []interface{} `json:"tools"`
			Docker struct {
				Image string `json:"image"`
			} `json:"docker"`
		} `json:"worker"`
		Env map[string] struct{
				Value  string `json:"value"`
				Type   string `json:"type"`
				Desc   string `json:"desc"`
				Option string `json:"option"`
		} `json:"env"`
		Trigger struct {
			Tags struct {
				Include []string `json:"include"`
			} `json:"tags"`
		} `json:"trigger"`
		Finally struct {
			Success struct {
				Cmds []struct {
					Plugin string `json:"plugin"`
					Params struct {
						Cmds []string `json:"cmds"`
					} `json:"params"`
					Label string `json:"label,omitempty"`
				} `json:"cmds"`
			} `json:"success"`
			Failure struct {
				Cmds []struct {
					Plugin string `json:"plugin"`
					Params struct {
						FailCheckResult string `json:"fail_check_result"`
						FailClassResult string `json:"fail_class_result"`
						CheckBody       string `json:"check_body"`
					} `json:"params"`
					Version string `json:"version,omitempty"`
					Label   string `json:"label,omitempty"`
				} `json:"cmds"`
			} `json:"failure"`
		} `json:"finally"`
	} `json:"pipeline_json"`
	TriggerTypeList []int  `json:"trigger_type_list"`
	Modifier        string `json:"modifier"`
	Updated         string `json:"updated"`
	UserEnv         struct {
	} `json:"user_env"`
	CurrentRunResult interface{}   `json:"current_run_result"`
	Members          string        `json:"members"`
	TriggerType      string        `json:"trigger_type"`
	RepoUsername     string        `json:"repo_username"`
	RepoPassword     string        `json:"repo_password"`
	AuthType         int           `json:"auth_type"`
	Power            int           `json:"power"`
	ProductName      interface{}   `json:"product_name"`
	StateName        string        `json:"state_name"`
	Oauth2User       string        `json:"oauth2_user"`
	TriggerContent   []interface{} `json:"trigger_content"`
	IsNochangeBuild  bool          `json:"is_nochange_build"`
	IsBlockMr        bool          `json:"is_block_mr"`
	CredentialID     interface{}   `json:"credential_id"`
	CredentialData   interface{}   `json:"credential_data"`
	Visibility       int           `json:"visibility"`
	MaximumBuilds    int           `json:"maximum_builds"`
	Certificate      interface{}   `json:"certificate"`
	IsLocalMr        bool          `json:"is_local_mr"`
	YamlSourceType   int           `json:"yaml_source_type"`
	BuildVersion     struct {
		EnableBuildVersion   bool   `json:"EnableBuildVersion"`
		FixVersion           int    `json:"FixVersion"`
		MajorVersion         int    `json:"MajorVersion"`
		FeatureVersion       int    `json:"FeatureVersion"`
		BuildVersionStrategy string `json:"BuildVersionStrategy"`
		BuildVersion         string `json:"BuildVersion"`
	} `json:"build_version"`
	RepoURL    string `json:"repo_url"`
	BranchList struct {
		Ret     bool          `json:"ret"`
		Data    []interface{} `json:"data"`
		Details string        `json:"details"`
	} `json:"branch_list"`
	Webhook                    interface{} `json:"webhook"`
	IsAutoCancel               bool        `json:"is_auto_cancel"`
	AutoCancelSameMergeRequest bool        `json:"auto_cancel_same_merge_request"`
	ProjectID                  int         `json:"project_id"`
	ProjectInfo                struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		DisplayName string `json:"display_name"`
		URL         string `json:"url"`
		Icon        string `json:"icon"`
		PlmCatalog  struct {
			ID         int    `json:"id"`
			CatalogID  int    `json:"catalogId"`
			Code       string `json:"code"`
			Level      int    `json:"level"`
			NameZh     string `json:"nameZh"`
			NameEn     string `json:"nameEn"`
			PStatus    int    `json:"pStatus"`
			Admin      string `json:"admin"`
			ParentCode string `json:"parentCode"`
			Dept1ID    int    `json:"dept1Id"`
			Dept1Name  string `json:"dept1Name"`
			Dept2ID    int    `json:"dept2Id"`
			Dept2Name  string `json:"dept2Name"`
			Intro      string `json:"intro"`
			Version    int    `json:"version"`
			BrandName  string `json:"brandName"`
			CreatedAt  string `json:"createdAt"`
			UpdatedAt  string `json:"updatedAt"`
			DeletedAt  string `json:"deletedAt"`
		} `json:"plm_catalog"`
		Belonging int `json:"belonging"`
	} `json:"project_info"`
	TemplateID int  `json:"template_id"`
	ConfigType int  `json:"config_type"`
	IsNoRevert bool `json:"is_no_revert"`
	Tags       []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"tags"`
	RepoID       interface{} `json:"repo_id"`
	TemplateJSON struct {
		ID               int         `json:"id"`
		Name             string      `json:"name"`
		Desc             interface{} `json:"desc"`
		PipelineConfigID int         `json:"pipeline_config_id"`
		Creator          string      `json:"creator"`
		Created          string      `json:"created"`
		Updated          string      `json:"updated"`
		Admin            string      `json:"admin"`
		Member           interface{} `json:"member"`
		Modifier         string      `json:"modifier"`
		State            int         `json:"state"`
		ProjectID        string      `json:"project_id"`
		Scope            int         `json:"scope"`
		IsForceSync      bool        `json:"is_force_sync"`
		IsConstrained    bool        `json:"is_constrained"`
		Version          interface{} `json:"version"`
		VersionDesc      interface{} `json:"version_desc"`
	} `json:"template_json"`
	TemplateInfoBinded struct {
	} `json:"template_info_binded"`
	OpTypeData struct {
		ID     interface{} `json:"id"`
		OpType interface{} `json:"op_type"`
	} `json:"op_type_data"`
	TriggerEnvTemplate []interface{} `json:"trigger_env_template"`
	Automaint          struct {
	} `json:"automaint"`
	TemplateEnvs struct {
	} `json:"template_envs"`
	TeamID    int `json:"team_id"`
	Customize struct {
	} `json:"customize"`
}
