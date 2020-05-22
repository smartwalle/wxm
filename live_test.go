package wxm_test

import (
	"encoding/json"
	"github.com/smartwalle/wxm"
	"testing"
)

func TestMiniProgram_GetLiveInf(t *testing.T) {
	var p = wxm.GetLiveInfoParam{}
	p.Start = 0
	p.Limit = 100

	rsp, err := miniProgram.GetLiveInfo(p)
	if err != nil {
		t.Error(err)
	}

	rspBytes, _ := json.Marshal(rsp)
	t.Log(string(rspBytes))
}
