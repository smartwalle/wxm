package wxm_test

import (
	"github.com/smartwalle/wxm"
	"testing"
)

func TestMiniProgram_GetLiveInfo(t *testing.T) {
	var p = wxm.GetLiveInfoParam{}
	p.Start = 0
	p.Limit = 100

	rsp, err := miniProgram.GetLiveInfo(p)
	if err != nil {
		t.Fatal(err)
	}

	if rsp.IsFailure() {
		t.Fatal(rsp.Msg)
	}
	t.Logf("%v", rsp)
}
