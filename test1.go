package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func program1(x int) {
	// Array
	// array_start := time.Now()
	// var a [101]int
	// for i := 0; i < x+1; i++ {
	// 	a[i] = i
	// }
	// array_elapsed := time.Since(array_start)
	// log.Printf("Initialize array took %s", array_elapsed)

	// Slice
	slice_start := time.Now()
	var s []int
	for i := 0; i < x+1; i++ {
		s = append(s, i)
	}
	slice_elapsed := time.Since(slice_start)
	log.Printf("Initialize slice took %s", slice_elapsed)

	// Map
	map_start := time.Now()
	m := make(map[int]int)
	for i := 0; i < x+1; i++ {
		m[i] = i
	}
	map_elapsed := time.Since(map_start)
	log.Printf("Initialize map took %s", map_elapsed)

	// Increment Slice
	slice_start = time.Now()
	for index, value := range s {
		s[index] = value + 1
	}
	slice_elapsed = time.Since(slice_start)
	log.Printf("Increment slice took %s", slice_elapsed)

	// Increment Map
	map_start = time.Now()
	for i := 0; i < x+1; i++ {
		m[i]++
	}
	map_elapsed = time.Since(map_start)
	log.Printf("Increment map took %s", map_elapsed)
}

func program2(x int) {
	// Array

	// Slice
	var s []int
	for i := 0; i < x+1; i++ {
		s = append(s, rand.Int())
	}

	// Sort Slice
	slice_start := time.Now()
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	slice_elapsed := time.Since(slice_start)
	log.Printf("Sorting slice took %s", slice_elapsed)

	// Sort Stable
	slice_start = time.Now()
	slice_elapsed = time.Since(slice_start)
	log.Printf("Stable sorting slice took %s", slice_elapsed)
}

func main() {
	arg := os.Args[1]
	x, _ := strconv.Atoi(arg)
	fmt.Println(x)
	program1(x)
	program2(x)
}
