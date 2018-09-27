package main

import (
	"fmt"
	"github.com/teejays/clog"
)

var ERR_KEY_NOT_EXIST error = fmt.Errorf("Key does not exist")

type Store struct {
	Kv        map[string]string
	CountDiff map[string]int
	Parent    *Store
}

func NewStore(parent *Store) *Store {
	s := new(Store)
	s.Kv = make(map[string]string)
	s.CountDiff = make(map[string]int)
	s.Parent = parent
	return s
}

func (s *Store) Set(key, value string) error {
	oldValue, err := s.Get(key)
	if err != nil && err != ERR_KEY_NOT_EXIST {
		return err
	}
	if err == nil {
		s.CountDiff[oldValue] -= 1
	}

	s.Kv[key] = value
	if value != "" {
		s.CountDiff[value] += 1
	}

	return nil
}

func (s *Store) Get(key string) (string, error) {
	value, exists := s.Kv[key]

	// Base Condition
	if exists {
		return value, nil
	}
	if s.Parent == nil {
		return value, ERR_KEY_NOT_EXIST
	}
	return s.Parent.Get(key)
}

func (s *Store) Delete(key string) error {
	value, err := s.Get(key)
	if err != nil {
		return err
	}
	clog.Debugf("[DELETE] Key: %s | Original Value %s ", key, value)

	// Let's not delete the key for now, so when we're commiting a transaction, we know that a key has been deleted
	// delete(s.Kv, key)
	s.Set(key, "")
	return nil
}

func (s *Store) Count(value string) (int, error) {
	var cnt int = s.CountDiff[value]
	if s.Parent == nil {
		return cnt, nil
	}
	pCnt, err := s.Parent.Count(value)
	if err != nil {
		return -1, err
	}
	return cnt + pCnt, nil
}
