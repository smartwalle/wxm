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

func TestMiniProgram_ParseMessage(t *testing.T) {
	var r, err = miniProgram.ParsePushMessage("feb1e4eb9a304c3eaae53ee6bf35071c", "1594269675", "246935454", "fcaa44ef33b85fb14cb883e993e0ffd7a2253f9e", "d5fwgPK7OLATY96lL45lI5llYggAgKgoyMslvLybNCt", []byte("{\"ToUserName\":\"gh_c12802380770\",\"Encrypt\":\"dUwqwFjp21PyPNRKlBxdAdAJ1EINbgvu5VlPG+RlWe0JjJjibi4+CaCLGtMOVcji7rEqwY9E36i0zyY3a1P4VmYJ0/eAjzhJCLWc32MK+JVHZi2fSUCAOP6nf3s8KB80KffkpL3/8RE3P7+h083B1jcu2/n9Mx6gEW+FVeOyXWk4WSvle5eqcRg8AvWLP6NIkMpGv/JwJM1kJv1ysJEkbEhNmRM1yusKz0USXk/yU+zJBOuB2RBsyGNLYd8HUek/La7twbVVnt5PCEhIJv8LWnxuI3RSkU+dwh2O3F1uL4RBTXS2HIOJiYg8ULa3hLjaDWXPEnL/elTDMEQYp5oenTuPflobzX0MLNFYDkLIyHQXvHcEjixJD8Iq718Wm4xDrhX4HviJin1CgezYGaGxXGW/b6LD69kBJfI95Hr+MJg/xjeZqJ7PnSaekIX5sEvySTy5dxiOUcxwoR+50zvQ/FVILqIozbjaVQnv4rJHRxnZA4pXAMKf/IyT8LPasepkIlzsbgJI8cQW5JgKzWLL1qcRuX0ktq70PFQmHCllx4x5Tklk7A1mwkh85aHMwjPf2ul0SLziUjbnlEHkAB+XIG8gcPoXI423jeInZI8LO1NN5BL/fgiApsvkZicbZJNljwSLPF0F2zdiJXioBlVp1a8uh1oIn3RTIwtCe/lLMd7lDP/dh5opSwjWL/6mDLvOxy1h+Hy/cnZ5fvsOq1U0KFQV9Z55LkEGNBexDEQTscGwhruopqUHe3gNtM/tFxDx\"}"))
	if err != nil {
		t.Error(err)
	}
	t.Log(r)
}
