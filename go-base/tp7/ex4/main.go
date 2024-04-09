package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func banchmarkUpTo(wg *sync.WaitGroup, values []int) {
	valuesLen := len(values)
	wg.Add(3)
	go func() {
		startTime := time.Now()
		insertionSort(values)
		wg.Done()
		fmt.Println("Insertion Sort: ", valuesLen, time.Since(startTime))
	}()

	go func() {
		startTime := time.Now()
		bubbleSort(values)
		wg.Done()
		fmt.Println("Bubble Sort : ", valuesLen, time.Since(startTime))
	}()

	go func() {
		startTime := time.Now()
		selectionSort(values)
		wg.Done()
		fmt.Println("Selection Sort : ", valuesLen, time.Since(startTime))
	}()
	wg.Wait()
}

func main() {
	variavel1 := rand.Perm(100)
	variavel2 := rand.Perm(1000)
	variavel3 := rand.Perm(10000)
	var wg sync.WaitGroup

	banchmarkUpTo(&wg, variavel1)
	banchmarkUpTo(&wg, variavel2)
	banchmarkUpTo(&wg, variavel3)
}

func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func insertionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return arr
}

func selectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}
