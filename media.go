package wxm

import (
	"net/http"
	"net/url"
)

const (
	kUploadMedia = "https://api.weixin.qq.com/cgi-bin/media/upload"
)

// UploadTempMedia 小程序-上传媒体文件到微信服务器 https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.uploadTempMedia.html
func (m *MiniProgram) UploadTempMedia(mediaType MediaType, filePath string) (result *UploadMediaRsp, err error) {
	if mediaType == "" {
		mediaType = MediaTypeOfImage
	}

	var v = url.Values{}
	v.Add("type", string(mediaType))

	if err = m.client.uploadWithRetry(http.MethodPost, kUploadMedia, "media", filePath, v, true, &result); err != nil {
		return nil, err
	}
	return result, nil
}
