package session

import (
	"fmt"
	"time"

	"github.com/akosgarai/projectregister/pkg/model"
)

// Session type
type Session struct {
	user         *model.User
	lastActivity time.Time
}

// New creates a new session
func New(user *model.User) *Session {
	return &Session{
		user:         user,
		lastActivity: time.Now(),
	}
}

// Store type a simple in memory session store
type Store struct {
	sessions map[string]*Session
	length   time.Duration
}

// NewStore creates a new session store
func NewStore(sessionLength time.Duration) *Store {
	return &Store{
		sessions: make(map[string]*Session),
		length:   sessionLength,
	}
}

// Get gets a session from the store
func (s *Store) Get(id string) (*Session, error) {
	session, ok := s.sessions[id]
	if !ok {
		return nil, fmt.Errorf("session not found")
	}
	// if the session it too old, delete it
	if time.Since(session.lastActivity) > (s.length) {
		s.Delete(id)
		return nil, fmt.Errorf("session expired")
	}
	return session, nil
}

// Set sets a session in the store
// if the session already exists, it will be overwritten
// the last activity time will be updated
func (s *Store) Set(id string, session *Session) {
	s.sessions[id] = session
	s.sessions[id].lastActivity = time.Now()
}

// Delete deletes a session from the store
func (s *Store) Delete(id string) {
	delete(s.sessions, id)
}
