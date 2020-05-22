package wxm_test

import (
	"fmt"
	"github.com/smartwalle/wxm"
	"testing"
)

var miniProgram = wxm.NewMiniProgram("wx6149efb9af013077", "871f204df0dfedc51ec57bef56eca353")

func TestMiniProgram_GetToken(t *testing.T) {
	fmt.Println(miniProgram.GetToken())
}
