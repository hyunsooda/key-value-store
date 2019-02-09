package main

import "fmt"

type myElement struct {
	Name string
	Age  uint8
}

type Store map[string]myElement

func newStore() *Store {
	m := make(Store)
	return &m
}

func (s *Store) Add(k string, n myElement) bool {
	if k == "" {
		return false
	}
	if s.Lookup(k) == nil && n.Name != "" && n.Age > 0 {
		(*s)[k] = n
		return true
	}
	return false
}

func (s *Store) Delete(k string) bool {
	if s.Lookup(k) != nil {
		delete(*s, k)
		return true
	}
	return false
}

func (s *Store) Lookup(k string) *myElement {
	_, ok := (*s)[k]
	if ok {
		n := (*s)[k]
		return &n
	}
	return nil
}

func (s *Store) Update(k string, n myElement) bool {
	(*s)[k] = n
	return true
}

func (s *Store) Print() {
	for k, v := range *s {
		fmt.Printf("key: %s value: %v\n", k, v)
	}
}

func main() {
	// test code
	s := newStore()
	s.Add("a", myElement{"hs", 25})
	s.Add("b", myElement{"jw", 21})
	// s.Add("c", myElement{Name: "qw"}) // Ingore due to few members in structure
	s.Print()
	s.Delete("a")
	s.Print()
}
