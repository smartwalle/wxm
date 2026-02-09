package miniprogram

import (
	"context"
	"net/url"
	"path/filepath"
)

const (
	kUploadMedia = "https://api.weixin.qq.com/cgi-bin/media/upload"
)

// UploadTempMedia 上传媒体文件到微信服务器
//
//	接口文档：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.uploadTempMedia.html
func (m *MiniProgram) UploadTempMedia(ctx context.Context, accessToken string, mediaType MediaType, filename string) (response *UploadMediaResponse, err error) {
	if mediaType == "" {
		mediaType = MediaTypeImage
	}

	var query = url.Values{}
	query.Add("type", string(mediaType))

	if err = m.Upload(ctx, accessToken, kUploadMedia, "media", filepath.Base(filename), filename, query, &response); err != nil {
		return nil, err
	}
	return response, nil
}
