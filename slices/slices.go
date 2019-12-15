package slices

import "fmt"

func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}
}

func appendOneSliceToAnother() {
	veggies := []string{"potatoes","tomatoes","brinjal"}
	fruits := []string{"oranges","apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:",food)
}

func testingReferencableNatureOfSlice() {
	/**
	Remember that a slice is a struct.
	type slice struct {
	    Length        int
	    Capacity      int
	    ZerothElement *byte
	}
	When a slice is passed to a function, even though it's passed by value, the pointer variable will refer to the same
	underlying array. Hence when a slice is passed to a function as parameter, changes made inside the function are visible
	outside the function too.
	*/
	nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos)
	subtactOne(nos)
	fmt.Println("slice after function call", nos)
}

func multiDimensionalSlices() {
	/**
	Similar to arrays, slices can have multiple dimensions.
	*/
	pls := [][]string {
		{"C", "C++"},
		{"Javascript"},
		{"Go", "Rust"},
	}
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func memoryOptimization() {
	/**
	Slices hold a reference to the underlying array. As long as the slice is in memory, the array cannot be garbage
	collected. This might be of concern when it comes to memory management. Lets assume that we have a very large array
	and we are interested in processing only a small part of it. Henceforth we create a slice from that array and start
	processing the slice. The important thing to be noted here is that the array will still be in memory since the slice
	references it.
	One way to solve this problem is to use the copy function func copy(dst, src []T) int to make a copy of that slice.
	This way we can use the new slice and the original array can be garbage collected.
	*/
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	fmt.Println(countriesCpy)
}

func Run() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4] //creates a slice from a[1] to a[3]
	fmt.Println(b)

	c := []int{6, 7, 8} //creates and array and returns a slice reference
	fmt.Println(c)

	darr := []int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after", darr)

	numa := [3]int{78, 79, 80}
	nums1 := numa[:] //creates a slice which contains all elements of the array
	nums2 := numa[:]
	fmt.Println("array before change 1", numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)
	fmt.Println("slices after all modifications", nums1, nums2)

	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6
	fruitslice = fruitslice[0:cap(fruitslice)] //re-slicing furitslice till its capacity
	fmt.Println("After re-slicing length is",len(fruitslice), "and capacity is",cap(fruitslice))

	makeSlice := make([]int, 5)
	fmt.Println(makeSlice)

	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6

	var names []string //zero value of a slice is nil
	if names == nil {
		fmt.Println("slice has old length", len(names), "and capacity", cap(names)) //capacity of cars is 3
		fmt.Println("slice is nil going to append")
		names = append(names, "Josh", "Sebastian", "Vinay", "asddsaf")
		fmt.Println("slice has new length", len(names), "and capacity", cap(names))
	}

	appendOneSliceToAnother()
	testingReferencableNatureOfSlice()
	multiDimensionalSlices()
	memoryOptimization()
}
