package wxm_test

import (
	"fmt"
	"github.com/smartwalle/wxm"
	"testing"
)

var officialAccount = wxm.NewOfficialAccount("wx7262a2f023e9aef8", "133850581b156b304ed23b30766aee90")

func TestOfficialAccount_GetToken(t *testing.T) {
	fmt.Println(miniProgram.GetToken())
}
