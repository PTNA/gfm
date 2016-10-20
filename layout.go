package gfm

// Layout represents a FileMaker layout
type Layout interface {
	DB() Database
	Name() string
	Delete(id int) error
	Duplicate(id int) (int, error)
	// TODO: Figure out what the edit interface should be. -edit
	// Edit(id int, ...) (..., error)
	// TODO: Figure out what the find interface should be. –find, –findall, or
	// –findany maybe –findquery?
	// Find(...) Query?
	// TODO: Figure out what the new interface should be. -new
	// New(...) int, error
	// TODO: Figure out what the view interface should be. -view
	// View(...) ...
}

type layout struct {
	database Database
	name     string
}

func (lay layout) DB() Database {
	return lay.database
}

func (lay layout) Name() string {
	return lay.name
}

func (lay layout) Delete(id int) error {
	// TODO: implement this using -delete
	return nil
}

func (lay layout) Duplicate(id int) (int, error) {
	// TODO: implement this using -dup
	return 0, nil
}
