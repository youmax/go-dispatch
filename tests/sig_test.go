package tests

import (
	"net/url"
	"testing"
	"upay/utils"
)

type TestCase map[string]string

var testInputs = [...]TestCase{
	{"custid": "001", "method": "get", "url": "/orders/1", "key": "secret", "sig": "4RRcY51t9FPsJvx4bLgDNkdMqiA="}}

func (t TestCase) remove(keys ...string) (vals []string) {
	for _, key := range keys {
		vals = append(vals, t[key])
		delete(t, key)
	}
	return
}
func TestSig(t *testing.T) {
	for _, v := range testInputs {
		ans := url.QueryEscape(v["sig"])
		slice := v.remove("method", "url", "key")
		sig := utils.CreateSig(slice[0], slice[1], slice[2], v)
		if sig != ans {
			t.Errorf("Sig was incorrect, got: %s, want: %s.", sig, ans)
		}
	}
}
