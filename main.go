package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spullara/bfes/bfes"
)

const (
	ITERATIONS = 1000
)

var random = rand.New(rand.NewSource(1337))

func main() {
	dim := 512
	k := 10
	b := getBfes(dim)
	vec := getRandomVector(dim)
	start := time.Now()
	for i := 0; i < ITERATIONS; i++ {
		b.Search(vec, k)
	}
	elapsed := time.Since(start)
	fmt.Printf("%.2f ms per search\n", float64(elapsed.Microseconds())/ITERATIONS/1000)
}

func getRandomVector(dim int) []float32 {
	vec := make([]float32, dim)
	for i := 0; i < dim; i++ {
		vec[i] = random.Float32()
	}
	return vec
}

func getBfes(dim int) *bfes.BFES {
	b := bfes.New(dim)
	for i := 0; i < 100000; i++ {
		b.Add(getRandomVector(dim))
	}
	return b
}
