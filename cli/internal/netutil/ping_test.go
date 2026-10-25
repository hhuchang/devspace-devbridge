package netutil

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestPingURI_NoServedByOnError 验证：4xx/5xx 且无 X-Served-By 时，结果为
// 「请求未到达隧道服务」而非原始 HTTP 状态。
func TestPingURI_NoServedByOnError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	}))
	defer srv.Close()

	r := PingURI(srv.URL, 5*time.Second)
	if r.Err == nil {
		t.Fatal("expected Err on non-OK without X-Served-By, got nil")
	}
	if r.StatusText == "403 Forbidden" {
		t.Fatalf("should classify as not-served-by-tunnel, got raw status %q", r.StatusText)
	}
}

// TestPingURI_ServedByKeepsStatus 验证：有 X-Served-By 的 4xx/5xx 保留原始状态，
// 不被误判为「未到达」。
func TestPingURI_ServedByKeepsStatus(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Served-By", "tunnels-cn-north-4-bridge")
		w.WriteHeader(http.StatusBadGateway)
	}))
	defer srv.Close()

	r := PingURI(srv.URL, 5*time.Second)
	if r.Err != nil {
		t.Fatalf("expected no Err when X-Served-By present, got %v", r.Err)
	}
	if r.StatusText != "502 Bad Gateway" {
		t.Fatalf("expected raw status 502 Bad Gateway, got %q", r.StatusText)
	}
}

// TestPingURI_DisconnectedOnTransportError 验证：传输层错误归为「隧道连接断开」，
// 不再伪装成 Bad Gateway / Gateway Timeout。
func TestPingURI_DisconnectedOnTransportError(t *testing.T) {
	// 指向一个未监听的本地端口，触发 connection refused（传输层错误）。
	r := PingURI("http://127.0.0.1:1/", 2*time.Second)
	if r.Err == nil {
		t.Fatal("expected transport Err, got nil")
	}
	if r.StatusText == "Bad Gateway" || r.StatusText == "Gateway Timeout" {
		t.Fatalf("must not fake %q for transport error", r.StatusText)
	}
}
