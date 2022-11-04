package codingInterface

/*
获取流水线的信息
Url="http://tencent.coding.woa.com/api/qci/rest-api/pipeline/{流水线id}"
Method = "GET"
headers = {"Authorization": "token xxxxx"}
*/
type PipelineInfoResp struct {
	Name              string        `json:"name"`
	Desc              string        `json:"desc"`
	CodeSource        string        `json:"code_source"`
	YamlLocation      string        `json:"yaml_location"`
	CurRevision       interface{}   `json:"cur_revision"`
	CodeBranch        []interface{} `json:"code_branch"`
	CodeSourceType    string        `json:"code_source_type"`
	DefaultCodeBranch string        `json:"default_code_branch"`
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
						TargetBranch   string `json:"target_branch"`
						GitToken       string `json:"git_token"`
						SourceRepoList string `json:"source_repo_list"`
					} `json:"params"`
				} `json:"cmds"`
			} `json:"tasks,omitempty"`
			If     string `json:"if,omitempty"`
			Ignore bool   `json:"ignore,omitempty"`
			Prompt []struct {
				To            string `json:"to"`
				Detail        string `json:"detail"`
				TimeoutStatus int    `json:"timeout_status"`
				Notifications struct {
					Channel       string `json:"channel"`
					EnwechatGroup string `json:"enwechat_group"`
					Mentions      string `json:"mentions"`
				} `json:"notifications"`
				Env struct {
				} `json:"env"`
				If      string `json:"if"`
				Timeout string `json:"timeout"`
				Msg     string `json:"msg"`
			} `json:"prompt,omitempty"`
		} `json:"stages"`
		Version string `json:"version"`
		Worker  struct {
			Label    string        `json:"label"`
			Language string        `json:"language"`
			Tools    []interface{} `json:"tools"`
		} `json:"worker"`
		Env map[string]struct {
			Value    string `json:"value"`
			Type     string `json:"type"`
			Require  bool   `json:"require"`
			Readonly bool   `json:"readonly"`
			Desc     string `json:"desc"`
			Option   string `json:"option"`
		} `json:"env"`
		Finally struct {
			Success struct {
				Cmds []struct {
					Plugin  string `json:"plugin"`
					Params  any    `json:"params,omitempty"`
					Version string `json:"version,omitempty"`
				} `json:"cmds"`
			} `json:"success"`
			All struct {
				Cmds []struct {
					Plugin string `json:"plugin"`
					Params any    `json:"params,omitempty"`
					Label  string `json:"label,omitempty"`
				} `json:"cmds"`
			} `json:"all"`
			Aborted struct {
				Cmds []struct {
					Plugin string `json:"plugin"`
					Params any    `json:"params,omitempty"`
					Label  string `json:"label"`
				} `json:"cmds"`
			} `json:"aborted"`
			Other struct {
				Cmds []struct {
					Plugin  string `json:"plugin"`
					Params  any    `json:"params,omitempty"`
					Label   string `json:"label,omitempty"`
					Version string `json:"version,omitempty"`
				} `json:"cmds"`
			} `json:"other"`
		} `json:"finally"`
		Mr struct {
			Branches struct {
				Include []string `json:"include"`
			} `json:"branches"`
			Paths struct {
			} `json:"paths"`
			Action []string `json:"action"`
		} `json:"mr"`
		CodeCheckoutType string `json:"code_checkout_type"`
		IsAutoInterrupt  bool   `json:"is_auto_interrupt"`
		Trigger          struct {
			Branches struct {
				Include []string `json:"include"`
			} `json:"branches"`
			Paths struct {
				Include []string `json:"include"`
			} `json:"paths"`
			Tags struct {
			} `json:"tags"`
		} `json:"trigger"`
		Manual any `json:"manual"`
	} `json:"pipeline_json"`
	TriggerTypeList []int  `json:"trigger_type_list"`
	Modifier        string `json:"modifier"`
	Updated         string `json:"updated"`
	UserEnv         struct {
	} `json:"user_env"`
	CurrentRunResult struct {
		ID                   int         `json:"id"`
		Status               int         `json:"status"`
		DetailsInfo          string      `json:"details_info"`
		Starter              string      `json:"starter"`
		StartTime            string      `json:"start_time"`
		FinishTime           string      `json:"finish_time"`
		BlockTime            interface{} `json:"block_time"`
		Comment              string      `json:"comment"`
		Label                string      `json:"label"`
		CurCodeSource        string      `json:"cur_code_source"`
		CurCodeSourceType    string      `json:"cur_code_source_type"`
		CurBranch            string      `json:"cur_branch"`
		CurRevision          string      `json:"cur_revision"`
		TargetCodeSource     string      `json:"target_code_source"`
		SourceBranch         string      `json:"source_branch"`
		TargetBranch         string      `json:"target_branch"`
		RefType              string      `json:"ref_type"`
		CredentialInfo       string      `json:"credential_info"`
		CodeCheckoutType     string      `json:"code_checkout_type"`
		MergeRequestID       string      `json:"merge_request_id"`
		MergeRequestIid      string      `json:"merge_request_iid"`
		IsSimulateMerge      bool        `json:"is_simulate_merge"`
		RunNum               int         `json:"run_num"`
		Created              string      `json:"created"`
		TriggerType          string      `json:"trigger_type"`
		Updated              string      `json:"updated"`
		HookData             string      `json:"hook_data"`
		YamlSourceType       int         `json:"yaml_source_type"`
		MajorVersion         int         `json:"major_version"`
		FeatureVersion       int         `json:"feature_version"`
		FixVersion           int         `json:"fix_version"`
		BuildVersion         int         `json:"build_version"`
		BuildVersionStrategy interface{} `json:"build_version_strategy"`
		Pipeline             int         `json:"pipeline"`
		LastWorker           interface{} `json:"last_worker"`
		RestartFrom          interface{} `json:"restart_from"`
	} `json:"current_run_result"`
	Members         string        `json:"members"`
	TriggerType     string        `json:"trigger_type"`
	RepoUsername    string        `json:"repo_username"`
	RepoPassword    string        `json:"repo_password"`
	AuthType        int           `json:"auth_type"`
	Power           int           `json:"power"`
	ProductName     interface{}   `json:"product_name"`
	StateName       string        `json:"state_name"`
	Oauth2User      string        `json:"oauth2_user"`
	TriggerContent  []interface{} `json:"trigger_content"`
	IsNochangeBuild bool          `json:"is_nochange_build"`
	IsBlockMr       bool          `json:"is_block_mr"`
	CredentialID    interface{}   `json:"credential_id"`
	CredentialData  struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		User  string `json:"user"`
		Type  int    `json:"type"`
		State string `json:"state"`
	} `json:"credential_data"`
	Visibility     int         `json:"visibility"`
	MaximumBuilds  int         `json:"maximum_builds"`
	Certificate    interface{} `json:"certificate"`
	IsLocalMr      bool        `json:"is_local_mr"`
	YamlSourceType int         `json:"yaml_source_type"`
	BuildVersion   struct {
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
	RepoID       int `json:"repo_id"`
	TemplateJSON struct {
	} `json:"template_json"`
	TemplateInfoBinded struct {
		ID               int         `json:"id"`
		Name             string      `json:"name"`
		Desc             string      `json:"desc"`
		PipelineConfigID int         `json:"pipeline_config_id"`
		Creator          string      `json:"creator"`
		Created          string      `json:"created"`
		Updated          string      `json:"updated"`
		Admin            string      `json:"admin"`
		Member           string      `json:"member"`
		Modifier         string      `json:"modifier"`
		State            int         `json:"state"`
		ProjectID        string      `json:"project_id"`
		Scope            int         `json:"scope"`
		IsForceSync      bool        `json:"is_force_sync"`
		IsConstrained    bool        `json:"is_constrained"`
		Version          interface{} `json:"version"`
		VersionDesc      interface{} `json:"version_desc"`
	} `json:"template_info_binded"`
	OpTypeData struct {
		ID     interface{} `json:"id"`
		OpType interface{} `json:"op_type"`
	} `json:"op_type_data"`
	TriggerEnvTemplate []interface{} `json:"trigger_env_template"`
	Automaint          struct {
	} `json:"automaint"`
	TemplateEnvs map[string]struct {
		Value    string `json:"value"`
		Type     string `json:"type"`
		Require  bool   `json:"require"`
		Option   string `json:"option"`
		Desc     string `json:"desc"`
		Readonly bool   `json:"readonly"`
	} `json:"template_envs"`
	TeamID    int `json:"team_id"`
	Customize struct {
		TriggerTemplate []string `json:"trigger_template"`
	} `json:"customize"`
}

/*
创建流水线
Url="http://tencent.coding.woa.com/api/qci/rest-api/v2/pipeline"
Method = "POST"
headers = {"Authorization": "token xxxxx"}
*/

type CreatePipelineReq struct {
	CodeSourceType string `json:"code_source_type"`
	RepoID         int    `json:"repo_id"`
	PipelineJSON   struct {
		Version string        `json:"version"`
		Stages  []interface{} `json:"stages"`
	} `json:"pipeline_json"`
	Name       string `json:"name"`
	TemplateID int    `json:"template_id"`
	ConfigType int    `json:"config_type"`
	Customize  struct {
		TriggerTemplate []string `json:"trigger_template"`
	} `json:"customize"`
	Tags         []interface{}     `json:"tags"`
	PipelineTags []interface{}     `json:"pipeline_tags"`
	ProjectId    int               `json:"project_id"`
	UserEnv      map[string]string `json:"user_env"`
}

type CreatePipelineResp struct {
	ID                   int         `json:"id"`
	Name                 string      `json:"name"`
	Desc                 string      `json:"desc"`
	CodeSource           string      `json:"code_source"`
	CodeSourceNoProtocol string      `json:"code_source_no_protocol"`
	YamlLocation         interface{} `json:"yaml_location"`
	CodeSourceType       string      `json:"code_source_type"`
	DefaultCodeBranch    interface{} `json:"default_code_branch"`
	Admin                string      `json:"admin"`
	State                int         `json:"state"`
	Timeout              interface{} `json:"timeout"`
	CurRevision          interface{} `json:"cur_revision"`
	CredentialID         int         `json:"credential_id"`
	MaximumBuilds        int         `json:"maximum_builds"`
	PipelineFile         interface{} `json:"pipeline_file"`
	YamlSourceType       int         `json:"yaml_source_type"`
	TemplateID           int         `json:"template_id"`
	ConfigType           string      `json:"config_type"`
	ProjectID            int         `json:"project_id"`
	Tags                 []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"tags"`
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
	TriggerEnvTemplate []interface{} `json:"trigger_env_template"`
	TeamID             int           `json:"team_id"`
	Customize          struct {
		TriggerTemplate []string `json:"trigger_template"`
	} `json:"customize"`
}
