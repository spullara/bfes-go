package bfes

import (
	"errors"
	"math"
	"sort"
	"sync"
)

type BFES struct {
	index [][]float32
	dim   int
}

func New(dim int) *BFES {
	return &BFES{
		dim: dim,
	}
}

func (b *BFES) Add(vec []float32) {
	if len(vec) != b.dim {
		panic("vector dimension mismatch")
	}
	unitFactor := float32(math.Sqrt(float64(b.dot(vec, vec))))
	clone := make([]float32, len(vec))
	for i := 0; i < b.dim; i++ {
		clone[i] = vec[i] / unitFactor
	}
	b.index = append(b.index, clone)
}

func (b *BFES) Search(query []float32, k int) []Score {
	if len(query) != b.dim {
		panic("query dimension mismatch")
	}
	unitFactor := float32(1.0 / math.Sqrt(float64(b.dot(query, query))))

	scores := make([]Score, len(b.index))
	var wg sync.WaitGroup
	wg.Add(len(b.index))
	for i := 0; i < len(b.index); i++ {
		go func(i int) {
			defer wg.Done()
			scores[i] = Score{ID: i, Value: b.dot(query, b.index[i])}
		}(i)
	}
	wg.Wait()

	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Value > scores[j].Value
	})

	if len(scores) > k {
		scores = scores[:k]
	}

	for i := range scores {
		scores[i].Value *= unitFactor
	}

	return scores
}

func (bfes *BFES) dot(a, b []float32) float32 {
	if len(a) != len(b) {
		panic(errors.New("arrays must have the same length"))
	}

	var result float32
	for i := 0; i < len(a); i++ {
		result += a[i] * b[i]
	}

	return result
}

type Score struct {
	ID    int
	Value float32
}
