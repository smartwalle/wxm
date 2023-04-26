package wxm

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	kUploadMediaURL = "https://api.weixin.qq.com/cgi-bin/media/upload"
)

// UploadTempMedia 小程序-上传媒体文件到微信服务器 https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.uploadTempMedia.html
func (this *MiniProgram) UploadTempMedia(mediaType MediaType, filePath string) (result *UploadMediaRsp, err error) {
	if mediaType == "" {
		mediaType = MediaTypeOfImage
	}

	var v = url.Values{}
	v.Add("type", string(mediaType))

	data, err := this.client.upload(http.MethodPost, kUploadMediaURL, "media", filePath, v, true)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
