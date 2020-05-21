package wxm

import (
	"encoding/json"
	"testing"
)

func TestClient_SendSubscribeMessage(t *testing.T) {
	var p = SendSubscribeMessageParam{}
	p.ToUser = "o45lH49xBSpfFndnPY5g6dM9cgvE"
	p.TemplateId = "eyVwkflimGWuuO0n_jJ5QQaiiJpOfcsg53NqPPA_D6k"
	p.AddData("thing1", "hah")
	p.AddData("thing2", "hhh")
	p.AddData("date3", "2015年01月05日")
	p.AddData("thing4", "aaa")

	rsp, err := client.SendSubscribeMessage(p)
	if err != nil {
		t.Error(err)
	}

	rspBytes, _ := json.Marshal(rsp)
	t.Log(string(rspBytes))
}

func TestClient_SendTemplateMessage(t *testing.T) {
	var p = SendTemplateMessageParam{}
	p.ToUser = "ocGXK1H6qbqcJ84MmES8Z5y5ItaE"
	p.TemplateId = "WITmIZMRP_GkiDmGwdMmeMM2Qlr6dZ8EqtiLVapf67Q"
	p.AddData("first", "hello first message", "#173177")
	p.AddData("Day", "2020-05-21", "#173177")
	p.AddData("orderId", "test-order-id", "#173177")
	p.AddData("orderType", "order-type", "#173177")
	p.AddData("customerName", "customer", "#173177")
	p.AddData("customerPhone", "18180103029", "#173177")
	p.AddData("remark", "something else", "#173177")

	rsp, err := client.SendTemplateMessage(p)
	if err != nil {
		t.Error(err)
	}

	rspBytes, _ := json.Marshal(rsp)
	t.Log(string(rspBytes))
}
