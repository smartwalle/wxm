package wxm_test

import (
	"fmt"
	"github.com/smartwalle/wxm"
	"testing"
)

func TestMiniProgram_GetUnlimited(t *testing.T) {
	var p = wxm.GetUnlimitedParam{}
	p.Scene = "1"
	p.Page = "pages/assist/gooddetail/gooddetail"

	fmt.Println(miniProgram.GetUnlimited(p))
}
