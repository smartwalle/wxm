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
	p.MPTemplateMsg.TemplateId = "WITmIZMRP_GkiDmGwdMmeMM2Qlr6dZ8EqtiLVapf67Q"
	p.MPTemplateMsg.AddData("first", "ww hello first message", "#173177")
	p.MPTemplateMsg.AddData("Day", "2020-05-21", "#173177")
	p.MPTemplateMsg.AddData("orderId", "test-order-id", "#173177")
	p.MPTemplateMsg.AddData("orderType", "order-type", "#173177")
	p.MPTemplateMsg.AddData("customerName", "customer", "#173177")
	p.MPTemplateMsg.AddData("customerPhone", "18180103029", "#173177")
	p.MPTemplateMsg.AddData("remark", "something else", "#173177")

	rsp, err := miniProgram.SendUniformMessage(p)
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
