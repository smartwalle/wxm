package wxm

type MediaType string

const (
	MediaTypeOfImage MediaType = "image"
)

type UploadMediaRsp struct {
	ErrCode   ErrCode   `json:"errcode"`
	ErrMsg    string    `json:"errmsg"`
	Type      MediaType `json:"type"`
	MediaId   string    `json:"media_id"`
	CreatedAt int64     `json:"created_at"`
}
