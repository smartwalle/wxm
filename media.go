package wxm

import (
	"context"
	"net/http"
	"net/url"
)

const (
	kUploadMedia = "https://api.weixin.qq.com/cgi-bin/media/upload"
)

// UploadTempMedia 小程序-上传媒体文件到微信服务器 https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.uploadTempMedia.html
func (m *MiniProgram) UploadTempMedia(ctx context.Context, mediaType MediaType, filename string) (result *UploadMediaResponse, err error) {
	if mediaType == "" {
		mediaType = MediaTypeOfImage
	}

	var values = url.Values{}
	values.Add("type", string(mediaType))

	if err = m.client.upload(ctx, http.MethodPost, kUploadMedia, "media", filename, values, &result); err != nil {
		return nil, err
	}
	return result, nil
}
