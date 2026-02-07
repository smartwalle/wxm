package officialaccount

type AuthScope string

const (
	AuthScopeBase     AuthScope = "snsapi_base"
	AuthScopeUserInfo AuthScope = "snsapi_userinfo"
)
