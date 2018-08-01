package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/url"
	"sort"
	"strings"
)

var skips = map[string]bool{
	"sig": true}

func CreateSig(method string, uri string, secretkey string, param map[string]string) string {
	var sig = strings.ToUpper(method) + "&" + UrlEncode(uri) + "&"
	var keys []string
	for k := range param {
		if skips[k] {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var queryStr string
	for _, key := range keys {
		queryStr += key + "=" + param[key] + "&"
	}
	sig += UrlEncode(queryStr[:len(queryStr)-1])
	h := hmac.New(sha1.New, []byte(secretkey))
	h.Write([]byte(sig))
	sig = base64.StdEncoding.EncodeToString(h.Sum(nil))
	return UrlEncode(sig)
}

func UrlEncode(s string) string {
	var builder string
	for i := 0; i < len(s); i++ {
		switch s[i : i+1] {
		case "'":
			builder += "%27"
		case " ":
			builder += "%20"
		case "(":
			builder += "%28"
		case ")":
			builder += "%29"
		case "!":
			builder += "%21"
		case "*":
			builder += "%2A"
		default:
			builder += url.QueryEscape(s[i : i+1])
		}
	}
	return builder
}
