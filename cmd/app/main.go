package main

import (
	"fmt"
	"runtime/metrics"
	"strings"
	"unsafe"
)

func sayHello() {
	var sl []int
	sl = append(sl, 1)
	sl = append(sl, 2, 3, 4, 5)

	fmt.Println(sl)
}

func getSizes[T any](a T) {
	fmt.Printf("Размер на стеке: %d байт. ", unsafe.Sizeof(a))
	samples := make([]metrics.Sample, 1)
	samples[0].Name = "/memory/classes/heap/objects:bytes"
	metrics.Read(samples)
	heapObjects := samples[0].Value.Uint64()
	fmt.Printf("Размер объектов в куче: %d байт\n", heapObjects)
}

func pointers() {
	p := func(x int) int { return x + x }
	i := 1
	for i < 1000000 {
		fmt.Println(p(i))
		i++
		fmt.Printf("Размер на стеке: %d байт. ", unsafe.Sizeof(i))
		samples := make([]metrics.Sample, 1)
		samples[0].Name = "/memory/classes/heap/objects:bytes"
		metrics.Read(samples)
		heapObjects := samples[0].Value.Uint64()
		fmt.Printf("Размер объектов в куче: %d байт\n", heapObjects)
	}
}

func pointers2() {
	p := "some"
	a := &p
	//fmt.Println(*a, **b, ***c)
	i := 1
	for i < 1000 {
		fmt.Println(*a)
		i++
	}
	fmt.Printf("Размер на стеке: %d байт. ", unsafe.Sizeof(a))
	samples := make([]metrics.Sample, 1)
	samples[0].Name = "/memory/classes/heap/objects:bytes"
	metrics.Read(samples)
	heapObjects := samples[0].Value.Uint64()
	fmt.Printf("Размер объектов в куче: %d байт\n", heapObjects)
}

func slices() {
	nums := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	nums_slice := nums[2:]
	fmt.Println("nums:", nums)
	fmt.Println("nums_slice:", nums_slice)
	fmt.Println(cap(nums_slice), len(nums_slice))

	nums_slice2 := nums_slice[0:2:4]
	fmt.Println("nums_slice2:", nums_slice2)
	fmt.Println(cap(nums_slice2), len(nums_slice2))

	fmt.Printf("Размер на стеке: %d байт. ", unsafe.Sizeof(nums_slice))
	samples := make([]metrics.Sample, 1)
	samples[0].Name = "/memory/classes/heap/objects:bytes"
	metrics.Read(samples)
	heapObjects := samples[0].Value.Uint64()
	fmt.Printf("Размер объектов в куче: %d байт\n", heapObjects)
}

func someStrings() {
	var str string
	str = "some,string"
	s := strings.Split(str, "")
	getSizes(s)
	fmt.Println(s[2])
}

func main() {
	//studing_goschool.Lesson3()
	//studing_gotour.Testing()
	//pointers2()
	//slices()
	someStrings()
}
