package session

import "sync"

type Session interface {
	Get(cookie string) (Model, bool)
	Set(m Model)
	IsLoggedIn(cookie string) bool
}

type Model struct {
	Token  string
	UserID int64
}

type session struct {
	m sync.Map
}

func (s *session) Get(token string) (Model, bool) {
	value, ok := s.m.Load(token)
	if !ok {
		return Model{}, false
	}
	return value.(Model), true
}

func (s *session) Set(m Model) {
	s.m.Store(m.Token, m)
}

func (s *session) IsLoggedIn(cookie string) bool {
	_, ok := s.Get(cookie)
	return ok
}

func New() Session {
	return &session{}
}

var (
	s    Session
	once sync.Once
)

func GetSessionInstance() Session {
	if s == nil {
		once.Do(func() {
			s = New()
		})
	}
	return s
}
