package store

// Store ...
type Store struct {
	config *Config
}

// New ...
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open ...
func (s *Store) Open() error {
	return nil
}

// Close
func (s *Store) Close() {
	// ...
}
