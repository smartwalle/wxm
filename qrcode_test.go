package wxm_test

import (
	"github.com/smartwalle/wxm"
	"testing"
)

func TestMiniProgram_GetUnlimited(t *testing.T) {
	var p = wxm.GetUnlimitedParam{}
	p.Scene = "1"
	p.Page = "pages/assist/gooddetail/gooddetail"

	rsp, err := miniProgram.GetUnlimited(p)
	if err != nil {
		t.Fatal(err)
	}

	if rsp.IsFailure() {
		t.Fatal(rsp.Msg)
	}
	t.Logf("%v", rsp)
}
