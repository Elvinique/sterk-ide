package main

import (
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSafePaths(t *testing.T) {
	root := t.TempDir()
	s := &Server{root: root}
	if _, e := s.safe("../secret"); e == nil {
		t.Fatal("traversal accepted")
	}
	if _, e := s.safe("/etc/passwd"); e == nil {
		t.Fatal("absolute path accepted")
	}
	if _, e := s.safe("normal/file.go"); e != nil {
		t.Fatal(e)
	}
	_ = os.Symlink(t.TempDir(), filepath.Join(root, "link"))
	if _, e := s.safe("link/file"); e == nil {
		t.Fatal("symlink accepted")
	}
}

func TestAPIAuthentication(t *testing.T) {
	t.Setenv("STERK_LOCAL_TOKEN", "test-only-token")
	s := &Server{root: t.TempDir(), goals: []Goal{}}
	h := s.handler()
	unauthorized := httptest.NewRecorder()
	h.ServeHTTP(unauthorized, httptest.NewRequest("GET", "/api/health", nil))
	if unauthorized.Code != 401 {
		t.Fatalf("expected 401, got %d", unauthorized.Code)
	}
	authorized := httptest.NewRequest("GET", "/api/health", nil)
	authorized.Header.Set("X-Sterk-Token", "test-only-token")
	recorded := httptest.NewRecorder()
	h.ServeHTTP(recorded, authorized)
	if recorded.Code != 200 {
		t.Fatalf("expected 200, got %d", recorded.Code)
	}
}

func TestCommandAllowlist(t *testing.T) {
	t.Setenv("STERK_LOCAL_TOKEN", "test-only-token")
	s := &Server{root: t.TempDir(), goals: []Goal{}}
	req := httptest.NewRequest("POST", "/api/verify", strings.NewReader(`{"kind":"rm-rf"}`))
	req.Header.Set("X-Sterk-Token", "test-only-token")
	rec := httptest.NewRecorder()
	s.handler().ServeHTTP(rec, req)
	if rec.Code != 403 {
		t.Fatalf("expected denied command, got %d", rec.Code)
	}
}
