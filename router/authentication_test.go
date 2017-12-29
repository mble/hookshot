package router_test

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/go-chi/chi"
	"github.com/mble/hookshot/router"
)

const testSecret = "foobar"

func testRequest(router *chi.Mux, path string, method string, signature string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, strings.NewReader("{}"))
	req.Header.Set("X-Hub-Signature", signature)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	return w
}

func testRouter() *chi.Mux {
	os.Setenv("HUB_SECRET", testSecret)
	r := chi.NewRouter()
	r.Use(router.Authenticate)
	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
	})
	return r
}

func signature(body string) string {
	dst := make([]byte, 40)
	computed := hmac.New(sha1.New, []byte(testSecret))
	computed.Write([]byte(body))
	hex.Encode(dst, computed.Sum(nil))
	return "sha1=" + string(dst)
}

func testAuthentication(t *testing.T, expectedCode int, signature string) {
	r := testRouter()
	resp := testRequest(r, "/", "GET", signature)

	if resp.Code != expectedCode {
		t.Errorf("expected: %d, got: %d", expectedCode, resp.Code)
	}
}

func TestEmptySignature(t *testing.T) {
	expectedCode := 404
	signature := ""

	testAuthentication(t, expectedCode, signature)
}

func TestBadSignature(t *testing.T) {
	expectedCode := 404
	signature := "sha1=25af6174a0fcecc4d346680a72b7ce644b9a88e8"

	testAuthentication(t, expectedCode, signature)
}

func TestGoodSignature(t *testing.T) {
	expectedCode := 200
	signature := signature("{}")

	testAuthentication(t, expectedCode, signature)
}
