package session

import (
	"testing"
	"time"

	"github.com/akosgarai/projectregister/pkg/model"
)

// TestNewSession tests the NewSession function
func TestNewSession(t *testing.T) {
	user := &model.User{ID: 1, Name: "test", Email: "test@email.com", Password: "password", CreatedAt: "2020-01-01", UpdatedAt: "2020-01-01"}
	session := New(user)
	if session == nil {
		t.Errorf("Expected a session, got nil")
	}
	if session.user != user {
		t.Errorf("Expected user, got %v", session.user)
	}
}

// TestNewStore tests the NewStore function
func TestNewStore(t *testing.T) {
	store := NewStore(10 * time.Second)
	if store == nil {
		t.Errorf("Expected a store, got nil")
	}
	if store.length != 10*time.Second {
		t.Errorf("Expected 10 seconds, got %v", store.length)
	}
	if store.sessions == nil {
		t.Errorf("Expected a map, got nil")
	}
	// test the length of the map, it should be 0
	if len(store.sessions) != 0 {
		t.Errorf("Expected 0, got %v", len(store.sessions))
	}
}

// TestStoreSet tests the Store.Set function
func TestStoreSet(t *testing.T) {
	store := NewStore(10 * time.Second)
	user := &model.User{ID: 1, Name: "test", Email: "test@email.com", Password: "password", CreatedAt: "2020-01-01", UpdatedAt: "2020-01-01"}
	session := New(user)
	store.Set("test", session)
	if store.sessions["test"] != session {
		t.Errorf("Expected session, got %v", store.sessions["test"])
	}
}

// TestStoreGet tests the Store.Get function
func TestStoreGet(t *testing.T) {
	store := NewStore(10 * time.Second)
	user := &model.User{ID: 1, Name: "test", Email: "test@email.com", Password: "password", CreatedAt: "2020-01-01", UpdatedAt: "2020-01-01"}
	session := New(user)
	store.Set("test", session)
	// test getting the session
	s, err := store.Get("test")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if s != session {
		t.Errorf("Expected session, got %v", s)
	}
	// test getting a non existing session
	_, err = store.Get("nonexisting")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	// test getting an expired session
	store.Set("expired", session)
	store.sessions["expired"].lastActivity = store.sessions["expired"].lastActivity.Add(-11 * time.Second)
	_, err = store.Get("expired")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

// TestStoreDelete tests the Store.Delete function
func TestStoreDelete(t *testing.T) {
	store := NewStore(10 * time.Second)
	user := &model.User{ID: 1, Name: "test", Email: "test@email.com", Password: "password", CreatedAt: "2020-01-01", UpdatedAt: "2020-01-01"}
	session := New(user)
	store.Set("test", session)
	// test deleting the session
	store.Delete("test")
	if len(store.sessions) != 0 {
		t.Errorf("Expected 0, got %v", len(store.sessions))
	}
	// test deleting a non existing session
	store.Delete("nonexisting")
	if len(store.sessions) != 0 {
		t.Errorf("Expected 0, got %v", len(store.sessions))
	}
}

// TestGenerateSessionKey tests the GenerateSessionKey function
func TestGenerateSessionKey(t *testing.T) {
	store := NewStore(10 * time.Second)
	key, err := store.GenerateSessionKey()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if key == "" {
		t.Errorf("Expected a key, got %v", key)
	}
	if len(key) != 32 {
		t.Errorf("Expected 32, got %v", len(key))
	}
}
