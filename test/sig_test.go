package test

import (
	"testing"

	"upay.cash/utils"
)

func TestSig(t *testing.T) {
	var param = make(map[string]string)
	param["custid"] = "001"
	sig := utils.CreateSig("GET", "/orders/1", "secret", param)
	var ans = "4RRcY51t9FPsJvx4bLgDNkdMqiA="
	if sig != ans {
		t.Errorf("Sig was incorrect, got: %s, want: %s.", sig, ans)
	}
}
