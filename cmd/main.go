package main

import (
	"fmt"
	"github.com/smartwalle/wxm"
)

func main() {
	var officialAccount = wxm.NewOfficialAccount("wx7262a2f023e9aef8", "133850581b156b304ed23b30766aee90")

	var openIdRsp, _ = officialAccount.GetUserOpenIdList("")
	for _, openId := range openIdRsp.Data.OpenId {
		rsp, _ := officialAccount.GetUserInfo(openId, "")
		fmt.Println(rsp.Nickname, rsp.OpenId, rsp.UnionId)
	}
}
