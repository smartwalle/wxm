package wxm_test

import (
	"context"
	"github.com/smartwalle/wxm"
	"testing"
)

var officialAccount = wxm.NewOfficialAccount("xxx", "xxx")

func TestOfficialAccount_GetToken(t *testing.T) {
	t.Log(officialAccount.GetToken(context.Background()))
}
