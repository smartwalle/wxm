package wxm

import (
	"fmt"
	"os"
	"testing"
)

var client *Client

func TestMain(m *testing.M) {
	//client = New("wx7262a2f023e9aef8", "133850581b156b304ed23b30766aee90")
	client = New("wx6149efb9af013077", "871f204df0dfedc51ec57bef56eca353")
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
	fmt.Println(client.JSCode2Session("033QiQle1WQX1y0QYKje1j1Wle1QiQlk"))
}
