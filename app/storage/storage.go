package storage

import (
	"slices"
	"strings"
)

type Storage struct {
	// O(1) by key access
	hash map[string]*Event

	// sorted list
	list []*Event
}

func New() *Storage {
	return &Storage{
		hash: map[string]*Event{},
		list: []*Event{},
	}
}

func (s *Storage) Put(event *Event) bool {
	if _, exists := s.hash[event.ID]; exists {
		return false
	}

	s.hash[event.ID] = event
	s.list = append(s.list, event)

	// order elements by timestamp
	slices.SortFunc(s.list, func(a, b *Event) int {
		return strings.Compare(a.CreatedAt, b.CreatedAt)
	})

	return true
}

func (s *Storage) Get(key string) *Event {
	return s.hash[key]
}

func (s *Storage) GetAll() []*Event {
	return s.list
}
