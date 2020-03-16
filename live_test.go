package wxm

import (
	"encoding/json"
	"testing"
)

func TestClient_GetLiveInfo(t *testing.T) {
	var p = GetLiveInfoParam{}
	p.Start = 0
	p.Limit = 100

	rsp, err := client.GetLiveInfo(p)
	if err != nil {
		t.Error(err)
	}

	rspBytes, _ := json.Marshal(rsp)
	t.Log(string(rspBytes))
}
