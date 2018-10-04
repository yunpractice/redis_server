package main

type Set struct {
	len int
}

func (s *Set) GetAll() []interface{} {
	if s.len == 0 {
		return nil
	}
	result := []interface{}{}

	return result
}
