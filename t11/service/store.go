package service

import (
	"sync"
)

type Store struct {
	sync.Mutex
	store map[int]EventCalendar
}

func NewStore() *Store {
	return &Store{Mutex: sync.Mutex{}, store: make(map[int]EventCalendar)}
}
