package gitInterface

/*
获取工蜂项目信息
Url="http://git.woa.com/api/v3/projects/:id"
Method = "GET"
headers = {"PRIVATE-TOKEN": "xxxxx"}
*/
type ProjectInfoResp struct {
	ID               int         `json:"id"`
	Description      interface{} `json:"description"`
	Public           bool        `json:"public"`
	Archived         bool        `json:"archived"`
	VisibilityLevel  int         `json:"visibility_level"`
	PublicVisibility int         `json:"public_visibility"`
	Namespace        struct {
		CreatedAt   string `json:"created_at"`
		Description string `json:"description"`
		ID          int    `json:"id"`
		Name        string `json:"name"`
		OwnerID     int    `json:"owner_id"`
		Path        string `json:"path"`
		UpdatedAt   string `json:"updated_at"`
	} `json:"namespace"`
	Owner struct {
		ID        int    `json:"id"`
		Username  string `json:"username"`
		WebURL    string `json:"web_url"`
		Name      string `json:"name"`
		State     string `json:"state"`
		AvatarURL string `json:"avatar_url"`
	} `json:"owner"`
	Name                 string        `json:"name"`
	NameWithNamespace    string        `json:"name_with_namespace"`
	Path                 string        `json:"path"`
	PathWithNamespace    string        `json:"path_with_namespace"`
	DefaultBranch        string        `json:"default_branch"`
	SSHURLToRepo         string        `json:"ssh_url_to_repo"`
	HTTPURLToRepo        string        `json:"http_url_to_repo"`
	HTTPSURLToRepo       string        `json:"https_url_to_repo"`
	WebURL               string        `json:"web_url"`
	TagList              []interface{} `json:"tag_list"`
	IssuesEnabled        bool          `json:"issues_enabled"`
	MergeRequestsEnabled bool          `json:"merge_requests_enabled"`
	WikiEnabled          bool          `json:"wiki_enabled"`
	SnippetsEnabled      bool          `json:"snippets_enabled"`
	ReviewEnabled        bool          `json:"review_enabled"`
	ForkEnabled          bool          `json:"fork_enabled"`
	TagNameRegex         interface{}   `json:"tag_name_regex"`
	TagCreatePushLevel   int           `json:"tag_create_push_level"`
	CreatedAt            string        `json:"created_at"`
	LastActivityAt       string        `json:"last_activity_at"`
	CreatorID            int           `json:"creator_id"`
	AvatarURL            string        `json:"avatar_url"`
	WatchsCount          int           `json:"watchs_count"`
	StarsCount           int           `json:"stars_count"`
	ForksCount           int           `json:"forks_count"`
	ConfigStorage        struct {
		LimitLfsFileSize int `json:"limit_lfs_file_size"`
		LimitSize        int `json:"limit_size"`
		LimitFileSize    int `json:"limit_file_size"`
		LimitLfsSize     int `json:"limit_lfs_size"`
	} `json:"config_storage"`
	ForkedFromProject string `json:"forked_from_project"`
	Statistics        struct {
		CommitCount       int `json:"commit_count"`
		RepositorySize    int `json:"repository_size"`
		LfsRepositorySize int `json:"lfs_repository_size"`
	} `json:"statistics"`
	CreatedFromID       interface{} `json:"created_from_id"`
	TemplateRepository  bool        `json:"template_repository"`
	SuggestionReviewers []struct {
		ID        int    `json:"id"`
		Username  string `json:"username"`
		WebURL    string `json:"web_url"`
		Name      string `json:"name"`
		State     string `json:"state"`
		AvatarURL string `json:"avatar_url"`
	} `json:"suggestion_reviewers"`
	NecessaryReviewers        []interface{} `json:"necessary_reviewers"`
	PathReviewerRules         string        `json:"path_reviewer_rules"`
	ApproverRule              int           `json:"approver_rule"`
	NecessaryApproverRule     int           `json:"necessary_approver_rule"`
	CommitMrCheck             bool          `json:"commit_mr_check"`
	ReviewCheck               bool          `json:"review_check"`
	CanApproveByCreator       bool          `json:"can_approve_by_creator"`
	AutoCreateReviewAfterPush bool          `json:"auto_create_review_after_push"`
	ForbiddenModifyRule       bool          `json:"forbidden_modify_rule"`
	ForceAddLabelsInNote      bool          `json:"force_add_labels_in_note"`
	ResolvedCheck             bool          `json:"resolved_check"`
	PushResetEnabled          bool          `json:"push_reset_enabled"`
	MergeRequestTemplate      interface{}   `json:"merge_request_template"`
	FileOwnerPathRules        string        `json:"file_owner_path_rules"`
}
