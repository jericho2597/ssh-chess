package app

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gliderlabs/ssh"
)

var SessionManagerService *SessionManager

const cleanupInterval = 1 * time.Minute

type SessionData struct {
	ID        string
	session   ssh.Session
	program   tea.Program
	lastUsed  time.Time
	ctx       context.Context
	cancelCtx context.CancelFunc
}

type SessionManager struct {
	sessions      map[string]*SessionData
	mu            sync.RWMutex
	cleanupTicker *time.Ticker
	done          chan struct{}
}

func NewSessionManager() {
	SessionManagerService = &SessionManager{
		sessions:      make(map[string]*SessionData),
		cleanupTicker: time.NewTicker(cleanupInterval),
		done:          make(chan struct{}),
	}
	go SessionManagerService.cleanupRoutine()
}

func (sm *SessionManager) AddSession(session ssh.Session, program tea.Program) (string, error) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	id := session.Context().SessionID()
	ctx, cancel := context.WithCancel(context.Background())

	sm.sessions[id] = &SessionData{
		ID:        id,
		session:   session,
		program:   program,
		lastUsed:  time.Now(),
		ctx:       ctx,
		cancelCtx: cancel,
	}

	go sm.monitorSession(id, session)

	return id, nil
}

func (sm *SessionManager) GetSession(id string) (*SessionData, error) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	session, exists := sm.sessions[id]
	if !exists {
		return nil, errors.New("session not found")
	}

	session.lastUsed = time.Now()
	return session, nil
}

func (sm *SessionManager) Send(id string, msg tea.Msg) {

	sessionData, err := sm.GetSession(id)
	if err != nil {
		log.Println(err)
		return
	}

	sessionData.program.Send(msg)
}

func (sm *SessionManager) RemoveSession(id string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	if session, exists := sm.sessions[id]; exists {
		session.cancelCtx()
		delete(sm.sessions, id)
	}
}

func (sm *SessionManager) monitorSession(id string, session ssh.Session) {
	<-session.Context().Done()
	sm.RemoveSession(id)
}

func (sm *SessionManager) cleanupRoutine() {
	for {
		select {
		case <-sm.cleanupTicker.C:
			sm.cleanupInactiveSessions()
		case <-sm.done:
			return
		}
	}
}

func (sm *SessionManager) cleanupInactiveSessions() {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	now := time.Now()
	for id, session := range sm.sessions {
		if now.Sub(session.lastUsed) > 30*time.Minute {
			session.cancelCtx()
			delete(sm.sessions, id)
		}
	}
}

func (sm *SessionManager) Shutdown() {
	close(sm.done)
	sm.cleanupTicker.Stop()

	sm.mu.Lock()
	defer sm.mu.Unlock()

	for _, session := range sm.sessions {
		session.cancelCtx()
	}
	sm.sessions = make(map[string]*SessionData)
}

func (sm *SessionManager) GetAllSessions() []*SessionData {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	sessions := make([]*SessionData, 0, len(sm.sessions))
	for _, session := range sm.sessions {
		sessions = append(sessions, session)
	}
	return sessions
}

func (sm *SessionManager) SessionCount() int {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return len(sm.sessions)
}
