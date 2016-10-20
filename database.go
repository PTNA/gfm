package gfm

// Database is the interface of a FileMaker database
type Database interface {
	Server() Server
	Name() string
	Lay(string) Layout
	Lays() []Layout
	// TODO: Find out if this is the proper way of representing script. Maybe
	// they need their own type?
	Scripts() []string
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

func (db database) Lays() []Layout {
	var lays []Layout
	// TODO: implement this using –layoutnames
	return lays
}

func (db database) Scripts() []string {
	var scripts []string
	// TODO: implement this using
	return scripts
}
