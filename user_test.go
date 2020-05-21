package wxm

import (
	"fmt"
	"testing"
)

func TestClient_GetUserOpenIdList(t *testing.T) {
	rsp, err := client.GetUserOpenIdList("")
	if err != nil {
		t.Error(err)
	}

	t.Log(len(rsp.Data.OpenId))
	t.Log(rsp.NextOpenId)
}

func TestClient_GetUserBasicInfo(t *testing.T) {
	rsp, err := client.GetUserBasicInfo("ocGXK1P38zSHpd_jG1A0hRDNgQo8", "")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(rsp)
}
