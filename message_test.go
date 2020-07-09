package wxm_test

import (
	"encoding/json"
	"github.com/smartwalle/wxm"
	"testing"
)

func TestMiniProgram_SendSubscribeMessage(t *testing.T) {
	var p = wxm.SendSubscribeMessageParam{}
	p.ToUser = "o45lH49xBSpfFndnPY5g6dM9cgvE"
	p.TemplateId = "eyVwkflimGWuuO0n_jJ5QQaiiJpOfcsg53NqPPA_D6k"
	p.AddData("thing1", "hah")
	p.AddData("thing2", "hhh")
	p.AddData("date3", "2015年01月05日")
	p.AddData("thing4", "aaa")

	rsp, err := miniProgram.SendSubscribeMessage(p)
	if err != nil {
		t.Error(err)
	}

	rspBytes, _ := json.Marshal(rsp)
	t.Log(string(rspBytes))
}

func TestMiniProgram_SendUniformMessage(t *testing.T) {
	var p = wxm.SendUniformMessageParam{}
	p.ToUser = "o-wmv4nnpvU_yj0fvwu_jAX2s38w"
	p.MPTemplateMsg = &wxm.MPTemplateMsg{}
	p.MPTemplateMsg.AppId = "wx7262a2f023e9aef8"
	p.MPTemplateMsg.TemplateId = "zbTRULuZeGLJgjQbXhlp6DkhwF5Gb6uXMt-VAV_APrg"
	p.MPTemplateMsg.MiniProgram = wxm.NewMiniProgramInfo("wx6149efb9af013077", "/packageC/pages/cindex/cindex")
	p.MPTemplateMsg.AddData("first", "您好，您有新的订单！", "#173177")
	p.MPTemplateMsg.AddData("keyword1", "晚安服务匹配订单", "#173177")
	p.MPTemplateMsg.AddData("keyword2", "nickname", "#173177")
	p.MPTemplateMsg.AddData("keyword3", "2006-01-02 15:04:06", "#173177")
	p.MPTemplateMsg.AddData("keyword4", "10.19元", "#173177")
	p.MPTemplateMsg.AddData("remark", "请及时查看处理！", "#173177")

	rsp, err := miniProgram.SendUniformMessage(p)
	if err != nil {
		t.Error(err)
	}

	rspBytes, _ := json.Marshal(rsp)
	t.Log(string(rspBytes))
}

func TestMiniProgram_SendCustomerServiceMessage(t *testing.T) {
	var p = wxm.SendCustomerServiceMessageParam{}
	p.ToUser = "o45lH49xBSpfFndnPY5g6dM9cgvE"
	p.MsgType = wxm.MsgTypeOfImage
	p.Image = &wxm.MsgImage{}
	p.Image.MediaId = "Ge2GGIhdVoRBX2-R9wYrdWjjENa_KEp4Ag9VDy3VhALgKQBMA2EFdcRQjKKPl-J4"

	rsp, err := miniProgram.SendCustomerServiceMessage(p)
	if err != nil {
		t.Error(err)
	}
	rspBytes, _ := json.Marshal(rsp)
	t.Log(string(rspBytes))
}

func TestOfficialAccount_SendTemplateMessage(t *testing.T) {
	var p = wxm.SendTemplateMessageParam{}
	p.ToUser = "ocGXK1H6qbqcJ84MmES8Z5y5ItaE"
	p.TemplateId = "WITmIZMRP_GkiDmGwdMmeMM2Qlr6dZ8EqtiLVapf67Q"
	p.AddData("first", "hello first message", "#173177")
	p.AddData("Day", "2020-05-21", "#173177")
	p.AddData("orderId", "test-order-id", "#173177")
	p.AddData("orderType", "order-type", "#173177")
	p.AddData("customerName", "customer", "#173177")
	p.AddData("customerPhone", "18180103029", "#173177")
	p.AddData("remark", "something else", "#173177")

	rsp, err := officialAccount.SendTemplateMessage(p)
	if err != nil {
		t.Error(err)
	}

	rspBytes, _ := json.Marshal(rsp)
	t.Log(string(rspBytes))
}
