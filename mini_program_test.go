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

func TestMiniProgram_CheckMessagePushServer(t *testing.T) {
	var r = miniProgram.CheckMessagePushServer("feb1e4eb9a304c3eaae53ee6bf35071c", "1594265641", "1812801742", "a70f5cb26cd3c2598fb3c504b934a96659951791")
	t.Log(r)
}
