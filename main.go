package main

import (
	"log"
	"time"
)

var arr1 = []int{1, 2, 3, 4, 3, 8, 9, 0, 9, 8, 8, 8, 8, 8, 7, 7, 6, 6, 6, 5, 5}
var arr2 = []int{9, 5, 3}

func main() {
	start := time.Now()
	solution1(arr1, arr2)
	log.Print(time.Since(start))

  
}

func solution1(arr1 []int, arr2 []int) {
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if arr2[j] == arr1[i] {
				log.Println("success")
			}
		}
	}
}

func sol2(arr1 []int, arr2 []int){
  //
}
