package wxm_test

import (
	"fmt"
	"github.com/smartwalle/wxm"
	"testing"
)

var miniProgram = wxm.NewMiniProgram("xxx", "xxx")

func TestMiniProgram_GetToken(t *testing.T) {
	fmt.Println(miniProgram.GetToken())
}
