package session

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
)

type Session struct {
	CookieLifetime string
	CookiePersist  string
	CookieName     string
	CookieDomain   string
	SessionType    string
	CookieSecure   string
	DBPPool        *sql.DB
}

func (c *Session) InitSession() *scs.SessionManager {

	var persist, secure bool

	// how long should sessions last?
	minutes, err := strconv.Atoi(c.CookieLifetime)
	if err != nil {
		minutes = 60
	}

	// should cookies persist
	if strings.ToLower(c.CookiePersist) == "true" {
		persist = true
	}

	// must cookies be secure?
	if (strings.ToLower(c.CookieSecure)) == "true" {
		secure = true
	}

	session := scs.New()
	session.Lifetime = time.Duration(minutes) * time.Minute
	session.Cookie.Persist = persist
	session.Cookie.Name = c.CookieName
	session.Cookie.Secure = secure
	session.Cookie.Domain = c.CookieDomain
	session.Cookie.SameSite = http.SameSiteLaxMode

	// which session store?
	log.Println(c.SessionType)
	switch strings.ToLower(c.SessionType) {
	case "redis":
	case "mysql", "mariadb":
		session.Store = mysqlstore.New(c.DBPPool)
	case "postgres", "postgresql":
		log.Println("Init postgres sessions backend")
		session.Store = postgresstore.New(c.DBPPool)
	default:
		log.Println("Init postgres COOKIES backend")
	}

	return session
}
