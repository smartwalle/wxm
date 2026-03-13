package wxm

import "context"

const (
	APIGetAccountBasicInfo = "/cgi-bin/account/getaccountbasicinfo"
)

// GetAccountBasicInfo 获取基本信息
//
//	接口文档：https://developers.weixin.qq.com/doc/oplatform/openApi/miniprogram-management/basic-info-management/api_getaccountbasicinfo.html
//	说明：调用本 API 可以获取小程序的基本信息，该接口同适用于获取公众号基本信息
func (c *Client) GetAccountBasicInfo(ctx context.Context, accessToken string) (response *GetAccountBasicInfoResponse, err error) {
	var request = struct{}{}
	if err = c.Post(ctx, APIGetAccountBasicInfo, accessToken, request, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}
