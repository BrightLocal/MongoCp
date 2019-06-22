package dsn

import (
	"strings"
)

type DSN struct {
	UserName string
	Password string
	HostName string
	Port     string
	Database string
	Extra    string
}

func Parse(dsn string) DSN {
	var out DSN
	parts := make([]string, 2)
	// split by @ into "user:pass" and "host:port/db"
	if strings.Contains(dsn, "@") {
		parts = strings.SplitN(dsn, "@", 2)
		out.UserName, out.HostName = parts[0], parts[1]
	} else {
		out.HostName = dsn
	}
	// split "user:pass" into "user" and "pass"
	if strings.Contains(out.UserName, ":") {
		parts = strings.SplitN(out.UserName, ":", 2)
		out.UserName, out.Password = parts[0], parts[1]
	}
	// split "host:port/db?extra" into "host:port" and "db?extra"
	if strings.Contains(out.HostName, "/") {
		parts = strings.SplitN(out.HostName, "/", 2)
		out.HostName, out.Database = parts[0], parts[1]
	}
	// split "host:port"
	if strings.Contains(out.HostName, ":") {
		parts = strings.SplitN(out.HostName, ":", 2)
		out.HostName, out.Port = parts[0], parts[1]
	}
	// split "db?extra"
	if strings.Contains(out.Database, "?") {
		parts = strings.SplitN(out.Database, "?", 2)
		out.Database, out.Extra = parts[0], parts[1]
	}
	return out
}

func (d DSN) GetExtra(key string) string {
	if d.Extra == "" {
		return ""
	}
	pairs := strings.Split(d.Extra, "&")
	for _, pair := range pairs {
		p := strings.SplitN(pair, "=", 2)
		if p[0] == key {
			return p[1]
		}
	}
	return ""
}
