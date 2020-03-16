package wxm

import (
	"fmt"
	"os"
	"testing"
)

var client *Client

func TestMain(m *testing.M) {
	client = New("wx143cd4036f7c65c5", "94b3359451d867b91d0a78eb74030c9c")
	os.Exit(m.Run())
}

func TestClient_GetAccessToken(t *testing.T) {
	fmt.Println(client.GetAccessToken())
}

func TestClient_GetUnlimited(t *testing.T) {
	var p = GetUnlimitedParam{}
	p.Scene = "1"
	p.Page = "pages/assist/gooddetail/gooddetail"

	fmt.Println(client.GetUnlimited(p))
}

func TestClient_JSCode2Session(t *testing.T) {
	fmt.Println(client.JSCode2Session("111"))
}
