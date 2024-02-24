package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Interview question # 1 "What is slice?"
// Slice struct
// https://go.dev/src/runtime/slice.go#L15
//
// func append -> func growslice
// https://go.dev/src/runtime/slice.go#L155
//
// new capacity
// https://go.dev/src/runtime/slice.go#L267

func main() {
	var sliceA []int64 // create slice without values, without initializer, declaration only
	if sliceA == nil {
		// Interview question # 2 "What is default value for slices"
		fmt.Println("slice default value are NIL")
	}
	fmt.Println("default - length is", len(sliceA), "capacity is", cap(sliceA))

	// Interview question # 3 "How append function works?"
	sliceA = append(sliceA, 1, 2, 3, 4)
	fmt.Println("after first appending - length is", len(sliceA), "capacity is", cap(sliceA))

	// Interview question # 4 "Why append inside another function is not good idea?"
	fmt.Println("before not-friend func", sliceA)
	thisFuncIsNotFriendForDeveloper(sliceA)
	fmt.Println("after not-friend func", sliceA)

	// Interview question # 5 "Could you please tell about cases, when you send slice as argument, modify it and everything is okay?"
	// File reader take slice as argument and works fine.
	// Use Q4 (it is okay if func do not change capacity) and this sample to answer Q5.
	file, err := os.Open("_01_slice/hello.txt")
	if err != nil {
		panic(err)
	}
	bufReader := bufio.NewReader(file)
	fileContent := make([]byte, 12)
	_, _ = bufReader.Read(fileContent)
	fmt.Println("File content:", string(fileContent))
}

func thisFuncIsNotFriendForDeveloper(s []int64) {
	ts := time.Now().Unix()
	s = append(s, ts)
	//fmt.Println("inside not-friend func", s)
}
