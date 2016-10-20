package gfm

// Server is the interface of a FileMaker server
type Server interface {
	DB(string) Database
	DBs() []Database
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

func (s server) DBs() []Database {
	var dbs []Database
	// TODO: Implement this using -dbnames
	return dbs
}
