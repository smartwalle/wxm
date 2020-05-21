package wxm

type AuthScope string

const (
	AuthScopeBase     AuthScope = "snsapi_base"
	AuthScopeUserInfo AuthScope = "snsapi_userinfo"
)

type AccessToken struct {
	ErrCode      ErrCode   `json:"errcode"`
	ErrMsg       string    `json:"errmsg"`
	AccessToken  string    `json:"access_token"`
	ExpiresIn    int64     `json:"expires_in"`
	RefreshToken string    `json:"refresh_token"`
	OpenId       string    `json:"openid"`
	Scope        AuthScope `json:"scope"`
	UnionId      string    `json"unionid"`
}
