package session

import "github.com/go-session/session/v3"

const (
	CookieKey = "X-Access-token"
)

var Middleware = New(session.SetCookieName(CookieKey), session.SetExpired(3600*24), session.SetEnableSIDInURLQuery(false))
