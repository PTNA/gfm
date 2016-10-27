package gfm

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/http"
)

// Database is the interface of a FileMaker database
type Database interface {
	// Server returns the Server that this database belongs to.
	Server() Server

	// Name returns the name of the database.
	Name() string

	// Lay returns a Layout struct, which can be queried further.
	Lay(string) Layout

	// Lays returns the layouts of the database.
	Lays() ([]Layout, error)

	// Scripts returns an array of the scripts on this database.
	//
	// TODO: Find out if this is the proper way of representing script. Maybe
	// they need their own type?
	Scripts() ([]string, error)

	// URL returns the url for sending requests to this database.
	URL() string
}

type database struct {
	server Server
	name   string
}

func (db database) Server() Server {
	return db.server
}

func (db database) Name() string {
	return db.name
}

func (db database) Lay(name string) Layout {
	return layout{database: db, name: name}
}

func (db database) Lays() ([]Layout, error) {
	var lays []Layout

	var b bytes.Buffer
	b.WriteString(db.URL())
	b.WriteString("-layoutnames")
	res, err := http.Get(b.String())

	if err != nil {
		// TODO: maybe we need a better error here
		return lays, err
	}

	var fmrs FMResultSet
	err = xml.NewDecoder(res.Body).Decode(&fmrs)
	if err != nil {
		return lays, err
	} else if fmrs.Error.Code != 0 {
		return lays, fmt.Errorf("FMError: %v", fmrs.Error.Code)
	}

	for _, r := range fmrs.ResultSet.Records {
		name := r.Fields[0].Data
		lay := layout{database: db, name: name}
		lays = append(lays, lay)
	}

	return lays, nil
}

func (db database) Scripts() ([]string, error) {
	var scripts []string

	var b bytes.Buffer
	b.WriteString(db.URL())
	b.WriteString("-scriptnames")
	res, err := http.Get(b.String())

	if err != nil {
		// TODO: maybe we need a better error here
		return scripts, err
	}
	var fmrs FMResultSet
	err = xml.NewDecoder(res.Body).Decode(&fmrs)
	if err != nil {
		return scripts, err
	} else if fmrs.Error.Code != 0 {
		return scripts, fmt.Errorf("FMError: %v", fmrs.Error.Code)
	}

	for _, r := range fmrs.ResultSet.Records {
		name := r.Fields[0].Data
		scripts = append(scripts, name)
	}

	return scripts, nil
}

func (db database) URL() string {
	var b bytes.Buffer
	b.WriteString(db.Server().URL())
	b.WriteString("-db=")
	b.WriteString(db.Name())
	b.WriteString("&")
	return b.String()
}
