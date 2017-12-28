package router

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const signaturePrefix = "sha1="
const signatureLength = 45

// Authenticate a request against an `X-Hub-Signature` header
// https://developer.github.com/webhooks/securing/
func Authenticate(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		signature := req.Header.Get("X-Hub-Signature")
		body, _ := ioutil.ReadAll(req.Body)
		if checkSigned(os.Getenv("HUB_SECRET"), signature, body) {
			h.ServeHTTP(w, req)
		} else {
			http.NotFound(w, req)
		}
	}
	return http.HandlerFunc(fn)
}

func checkSigned(secret, signature string, body []byte) bool {
	if len(signature) != signatureLength || !strings.HasPrefix(signature, signaturePrefix) {
		return false
	}
	computed := hmac.New(sha1.New, []byte(secret))
	computed.Write(body)

	actual := make([]byte, 20)
	hex.Decode(actual, []byte(signature[5:]))

	return hmac.Equal(computed.Sum(nil), actual)
}
