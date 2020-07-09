package wxm_test

import (
	"fmt"
	"github.com/smartwalle/wxm"
	"testing"
)

var miniProgram = wxm.NewMiniProgram("wx143cd4036f7c65c5", "94b3359451d867b91d0a78eb74030c9c")

func TestMiniProgram_GetToken(t *testing.T) {
	fmt.Println(miniProgram.GetToken())
}
