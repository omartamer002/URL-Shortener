package storage

import "sync"

type Storage struct {
	urls map[string]string
	mu  sync.RWMutex
}

func NewStorage() *Storage {
	return &Storage{
		urls: make(map[string]string),
	}
}

func (s *Storage) Save(shortURL, originalURL string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.urls[shortURL] = originalURL
}

func (s *Storage) Get(shortURL string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	originalURL, exists := s.urls[shortURL]
	return originalURL, exists
}