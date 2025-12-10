package matrix

import (
	"errors"
	"fmt"
)

type SparseMatrix struct {
	values map[string]int
	m      int
	n      int
}

func positionToKey(i, j int) string {
	return fmt.Sprintf("%d,%d", i, j)
}

func NewSparseMatrix(m, n int) SparseMatrix {
	return SparseMatrix{map[string]int{}, m, n}
}

func (s *SparseMatrix) Size() (int, int) {
	return s.m, s.n
}

func (s *SparseMatrix) Set(i, j, value int) {
	key := positionToKey(i, j)
	s.values[key] = value
}

func (s *SparseMatrix) Get(i, j int) (int, error) {
	if i < 0 || i >= s.m || j < 0 || j >= s.n {
		return -1, errors.New("Outside of bounds")
	}
	key := positionToKey(i, j)
	value, exists := s.values[key]
	if !exists {
		return 0, nil
	}
	return value, nil
}

func (s *SparseMatrix) PrintValues() {
	m, n := s.Size()
	for i := range m {
		for j := range n {
			v, err := s.Get(i, j)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%d ", v)
		}
		fmt.Print("\n")
	}
}
