package clients

import (
	"bytes"
	"fmt"
	"text/template"
)

var dsn = map[string]string{
	"mysql":    "{{.User}}:{{.Password}}@tcp({{.Host}}:{{.Port}})/{{.Name}}?charset=utf8mb4&parseTime=True&loc=UTC",
	"postgres": "user={{.User}} password={{.Password}} dbname={{.Name}} host={{.Host}} port={{.Port}} sslmode=disable TimeZone=Etc/UTC",
	"sqlite":   "{{.File}}?cache=shared&mode=rwc&parseTime=True",
}

// Config for clients connections.
type Config struct {
	DB       int
	File     string
	Host     string
	Name     string
	Password string
	Port     string
	Type     string
	User     string
}

// Addr returns the host and port address.
func (cfg Config) Addr() (s string) {
	s = fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	return
}

// DSN will return the string dsn template with the configuration data
// already connected to it.
func (cfg Config) DSN() (s string, err error) {
	var ok bool
	if s, ok = dsn[cfg.Type]; !ok {
		err = fmt.Errorf("unable to find DSN type for %s", cfg.Type)
		return
	}
	var tmpl *template.Template
	if tmpl, err = template.New(s).Parse(s); err != nil {
		return
	}
	buf := new(bytes.Buffer)
	if err = tmpl.Execute(buf, cfg); err != nil {
		return
	}
	s = buf.String()
	return
}
