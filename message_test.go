package wxm

import (
	"encoding/json"
	"testing"
)

func TestClient_SendMessage(t *testing.T) {
	var p = SendMessageParam{}
	p.ToUser = "o45lH49xBSpfFndnPY5g6dM9cgvE"
	p.TemplateId = "eyVwkflimGWuuO0n_jJ5QQaiiJpOfcsg53NqPPA_D6k"
	p.AddData("thing1", "hah")
	p.AddData("thing2", "hhh")
	p.AddData("date3", "2015年01月05日")
	p.AddData("thing4", "aaa")

	rsp, err := client.SendMessage(p)
	if err != nil {
		t.Error(err)
	}

	rspBytes, _ := json.Marshal(rsp)
	t.Log(string(rspBytes))
}
