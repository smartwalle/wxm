package wxm_test

import (
	"github.com/smartwalle/wxm"
	"testing"
)

var officialAccount = wxm.NewOfficialAccount("xxx", "xxx")

func TestOfficialAccount_GetToken(t *testing.T) {
	t.Log(miniProgram.GetToken())
}
