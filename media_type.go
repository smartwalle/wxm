package wxm

type MediaType string

const (
	MediaTypeOfImage MediaType = "image"
)

type UploadMediaResponse struct {
	Error
	Type      MediaType `json:"type"`
	MediaId   string    `json:"media_id"`
	CreatedAt int64     `json:"created_at"`
}
