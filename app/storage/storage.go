package storage

import (
	"slices"
	"strings"
)

type Storage struct {
	// ID hash index
	primary map[string]*Event

	// UserID index
	byUserID map[string][]*Event
}

func New() *Storage {
	return &Storage{
		primary:  map[string]*Event{},
		byUserID: map[string][]*Event{},
	}
}

func (s *Storage) Put(event *Event) bool {
	if _, exists := s.primary[event.ID]; exists {
		return false
	}

	s.primary[event.ID] = event

	if _, exists := s.byUserID[event.UserID]; !exists {
		s.byUserID[event.UserID] = []*Event{}
	}
	s.byUserID[event.UserID] = append(s.byUserID[event.UserID], event)

	// order elements by weight
	slices.SortFunc(s.byUserID[event.UserID], func(a, b *Event) int {
		w := int(b.Weight*100) - int(a.Weight*100)
		if w == 0 {
			return strings.Compare(b.CreatedAt, a.CreatedAt)
		}
		return w
	})

	return true
}

func (s *Storage) GetByID(id string) *Event {
	return s.primary[id]
}

func (s *Storage) GetByUserID(userID string) []*Event {
	return s.byUserID[userID]
}
