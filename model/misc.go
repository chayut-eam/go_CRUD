package model

type Int64Set struct {
	elements map[int64]bool
}

func NewInt64Set() *Int64Set {
	return &Int64Set{
		elements: make(map[int64]bool),
	}
}

func (s *Int64Set) Put(i int64) {
	s.elements[i] = true
}

func (s *Int64Set) Delete(i int64) {
	delete(s.elements, i)
}

func (s *Int64Set) Values() []int64 {
	var values []int64
	for element := range s.elements {
		values = append(values, element)
	}
	return values
}

//---------------------------------------------

type StringSet struct {
	elements map[string]bool
}

func NewStringSet() *StringSet {
	return &StringSet{
		elements: make(map[string]bool),
	}
}

func (s *StringSet) Put(i string) {
	s.elements[i] = true
}

func (s *StringSet) Delete(i string) {
	delete(s.elements, i)
}

func (s *StringSet) Values() []string {
	var values []string
	for element := range s.elements {
		values = append(values, element)
	}
	return values
}

//---------------------------------------------
