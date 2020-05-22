package wxm_test

import (
	"fmt"
	"github.com/smartwalle/wxm"
	"testing"
)

var officialAccount = wxm.NewOfficialAccount("xxx", "xxx")

func TestOfficial_GetToken(t *testing.T) {
	fmt.Println(miniProgram.GetToken())
}
