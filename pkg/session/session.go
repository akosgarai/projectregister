package session

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/akosgarai/projectregister/pkg/config"
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

// GetUser returns the user from the session
func (s *Session) GetUser() *model.User {
	return s.user
}

// Store type a simple in memory session store
type Store struct {
	sessions  map[string]*Session
	length    time.Duration
	keyLength int
	alphabet  string
}

// NewStore creates a new session store
func NewStore(c *config.Environment) *Store {
	return &Store{
		sessions:  make(map[string]*Session),
		length:    time.Duration(c.GetSessionLength()) * time.Minute,
		keyLength: c.GetSessionNameLength(),
		alphabet:  c.GetSessionNameAlphabet(),
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

// GenerateSessionKey returns a securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
// https://gist.github.com/dopey/c69559607800d2f2f90b1b1ed4e550fb
func (s *Store) GenerateSessionKey() (string, error) {
	ret := make([]byte, s.keyLength)
	for i := 0; i < s.keyLength; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(s.alphabet))))
		if err != nil {
			return "", err
		}
		ret[i] = s.alphabet[num.Int64()]
	}

	return string(ret), nil
}
