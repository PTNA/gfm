package gfm

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/http"
)

// Server is the interface of a FileMaker server
type Server interface {
	DB(string) Database
	DBs() ([]Database, error)
	URL() string
}

type server struct {
	host     string
	user     string
	password string
}

// NewServer instantiates a new server type from a host, user and a password
func NewServer(host, user, pass string) Server {
	return server{host: host, user: user, password: pass}
}

func (s server) DB(name string) Database {
	return database{server: s, name: name}
}

func (s server) DBs() ([]Database, error) {
	var dbs []Database

	var b bytes.Buffer
	b.WriteString(s.URL())
	b.WriteString("-dbnames")
	res, err := http.Get(b.String())

	if err != nil {
		// TODO: maybe we need a better error here
		return dbs, err
	}

	var fmrs FMResultSet
	err = xml.NewDecoder(res.Body).Decode(&fmrs)
	if err != nil {
		return dbs, err
	} else if fmrs.Error.Code != 0 {
		return dbs, fmt.Errorf("FMError: %v", fmrs.Error.Code)
	}

	for _, r := range fmrs.ResultSet.Records {
		name := r.Fields[0].Data
		db := database{server: s, name: name}
		dbs = append(dbs, db)
	}

	return dbs, nil
}

func (s server) URL() string {
	var b bytes.Buffer
	b.WriteString("http://")
	b.WriteString(s.user)
	b.WriteString(":")
	b.WriteString(s.password)
	b.WriteString("@")
	b.WriteString(s.host)
	b.WriteString("/fmi/xml/fmresultset.xml?")
	return b.String()
}
