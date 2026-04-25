package internal

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// ── handlePing ────────────────────────────────────────────────────────────────

func TestLeaderHandlePing_OK(t *testing.T) {
	l := NewLeader("127.0.0.1", 0, "v1.2.3", "test-token")

	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	req.Header.Set("Authorization", "Bearer test-token")
	w := httptest.NewRecorder()
	l.handlePing(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", w.Code)
	}

	var body map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("unmarshal body: %v", err)
	}
	if body["status"] != "ok" {
		t.Errorf("status = %q, want ok", body["status"])
	}
	if body["version"] != "v1.2.3" {
		t.Errorf("version = %q, want v1.2.3", body["version"])
	}
}

func TestLeaderHandlePing_MethodNotAllowed(t *testing.T) {
	l := NewLeader("127.0.0.1", 0, "", "test-token")

	for _, method := range []string{http.MethodPost, http.MethodPut, http.MethodDelete} {
		req := httptest.NewRequest(method, "/ping", nil)
		req.Header.Set("Authorization", "Bearer test-token")
		w := httptest.NewRecorder()
		l.handlePing(w, req)

		if w.Code != http.StatusMethodNotAllowed {
			t.Errorf("%s /ping: status = %d, want 405", method, w.Code)
		}
	}
}

// ── handleRPC ─────────────────────────────────────────────────────────────────

func TestLeaderHandleRPC_MethodNotAllowed(t *testing.T) {
	l := NewLeader("127.0.0.1", 0, "", "test-token")

	req := httptest.NewRequest(http.MethodGet, "/rpc", nil)
	req.Header.Set("Authorization", "Bearer test-token")
	w := httptest.NewRecorder()
	l.handleRPC(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("status = %d, want 405", w.Code)
	}
}

func TestLeaderHandleRPC_InvalidJSON(t *testing.T) {
	l := NewLeader("127.0.0.1", 0, "", "test-token")

	req := httptest.NewRequest(http.MethodPost, "/rpc", bytes.NewBufferString("{bad json}"))
	req.Header.Set("Authorization", "Bearer test-token")
	w := httptest.NewRecorder()
	l.handleRPC(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want 400", w.Code)
	}
	var resp RPCResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp.Error == "" {
		t.Error("expected error in response body")
	}
}

func TestLeaderHandleRPC_ValidationError(t *testing.T) {
	l := NewLeader("127.0.0.1", 0, "", "test-token")

	// set_text with nodeId but missing text → validation error
	body, _ := json.Marshal(RPCRequest{
		Tool:    "set_text",
		NodeIDs: []string{"1:1"},
		Params:  map[string]any{},
	})
	req := httptest.NewRequest(http.MethodPost, "/rpc", bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer test-token")
	w := httptest.NewRecorder()
	l.handleRPC(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want 400", w.Code)
	}
	var resp RPCResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp.Error == "" {
		t.Error("expected validation error in response")
	}
}

func TestLeaderHandleRPC_BridgeNotConnected(t *testing.T) {
	l := NewLeader("127.0.0.1", 0, "", "test-token")

	// get_document has no required params — passes validation, hits bridge
	body, _ := json.Marshal(RPCRequest{Tool: "get_document"})
	req := httptest.NewRequest(http.MethodPost, "/rpc", bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer test-token")
	w := httptest.NewRecorder()
	l.handleRPC(w, req)

	// Bridge returns "plugin not connected" error → 200 with error field
	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
	var resp RPCResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp.Error == "" {
		t.Error("expected 'plugin not connected' error in response")
	}
}

// ── Start / Stop ──────────────────────────────────────────────────────────────

func TestLeaderStart_BindsPort(t *testing.T) {
	port := freePort(t)
	l := NewLeader("127.0.0.1", port, "", "")

	if err := l.Start(); err != nil {
		t.Fatalf("Start: %v", err)
	}
	t.Cleanup(l.Stop)

	// Second leader on the same port must fail.
	l2 := NewLeader("127.0.0.1", port, "", "")
	if err := l2.Start(); err == nil {
		l2.Stop()
		t.Error("expected error when binding already-used port")
	}
}

func TestLeaderStop_FreesPort(t *testing.T) {
	port := freePort(t)
	l := NewLeader("127.0.0.1", port, "", "")

	if err := l.Start(); err != nil {
		t.Fatalf("Start: %v", err)
	}
	l.Stop()

	// Allow OS to release the port.
	time.Sleep(20 * time.Millisecond)

	l2 := NewLeader("127.0.0.1", port, "", "")
	if err := l2.Start(); err != nil {
		t.Fatalf("port should be free after Stop: %v", err)
	}
	l2.Stop()
}

func TestLeaderStop_Idempotent(t *testing.T) {
	l := NewLeader("127.0.0.1", 0, "", "")
	// Stop on a never-started leader should not panic.
	l.Stop()
	l.Stop()
}

// ── /ping endpoint (integration via httptest.Server) ─────────────────────────

func TestLeaderPingEndpoint(t *testing.T) {
	port := freePort(t)
	l := NewLeader("127.0.0.1", port, "test-ver", "test-token")
	if err := l.Start(); err != nil {
		t.Fatalf("Start: %v", err)
	}
	t.Cleanup(l.Stop)

	f := NewFollower("http://127.0.0.1:"+itoa(port), "test-token")
	if !f.Ping(t.Context()) {
		t.Error("expected ping to succeed for running leader")
	}
}

// ── Security ──────────────────────────────────────────────────────────────────

func TestLeaderAuth_RPC_Unauthorized(t *testing.T) {
	l := NewLeader("127.0.0.1", 0, "", "secret-token")

	req := httptest.NewRequest(http.MethodPost, "/rpc", bytes.NewReader([]byte("{}")))
	// No Auth header or wrong token
	req.Header.Set("Authorization", "Bearer wrong-token")
	w := httptest.NewRecorder()
	l.handleRPC(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("status = %d, want 401", w.Code)
	}
}

func TestLeaderAuth_Ping_Unauthorized(t *testing.T) {
	l := NewLeader("127.0.0.1", 0, "", "secret-token")

	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	w := httptest.NewRecorder()
	l.handlePing(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("status = %d, want 401", w.Code)
	}
}

func TestLeaderAuth_WS_Unauthorized(t *testing.T) {
	l := NewLeader("127.0.0.1", 0, "", "secret-token")

	req := httptest.NewRequest(http.MethodGet, "/ws", nil)
	// Missing token query param
	w := httptest.NewRecorder()
	l.handleWS(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("status = %d, want 401", w.Code)
	}
}

func TestLeaderAuth_WS_Authorized(t *testing.T) {
	l := NewLeader("127.0.0.1", 0, "", "secret-token")

	req := httptest.NewRequest(http.MethodGet, "/ws?token=secret-token", nil)
	w := httptest.NewRecorder()
	// handleWS will try to upgrade, which fails in httptest.NewRecorder,
	// but we want to see it pass the authentication check.
	// We can check if it gets past authenticate.
	l.handleWS(w, req)

	// Since httptest.NewRecorder doesn't support hijacking, websocket.Accept will fail
	// with "websocket: the client is not using the WebSocket protocol: 'Upgrade' header, or it isn't 'websocket'"
	// but the status code won't be 401.
	if w.Code == http.StatusUnauthorized {
		t.Error("expected not to be unauthorized")
	}
}
