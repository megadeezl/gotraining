// Example shows how to takes slices of slices to create different
// views of and make changes to the underlying array.
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	slice1 := make([]string, 5, 8)
	slice1[0] = "Apple"
	slice1[1] = "Orange"
	slice1[2] = "Banana"
	slice1[3] = "Grape"
	slice1[4] = "Plum"

	InspectSlice(slice1)

	// For slice[i:j] with an underlying array of capacity k
	// Length: j - i
	// Capacity: k - i

	// Slice slice1 from element two for length 4 - 2
	// With a capacity of 8 - 2
	slice2 := slice1[2:4]
	InspectSlice(slice2)

	// Slice slice2 from element one for length 6 - 1
	// With a capacity 6 - 1
	slice3 := slice2[1:cap(slice2)]
	InspectSlice(slice3)

	slice3[0] = "CHANGED"
	fmt.Println("slice1[3]", slice1[3])
	fmt.Println("slice2[1]", slice2[1])
	fmt.Println("slice3[0]", slice3[0])
}

func InspectSlice(slice []string) {
	// Capture the address to the slice structure
	address := unsafe.Pointer(&slice)

	// Capture the address where the length and cap size is stored
	lenAddr := uintptr(address) + uintptr(8)
	capAddr := uintptr(address) + uintptr(16)

	// Create pointers to the length and cap size
	lenPtr := (*int)(unsafe.Pointer(lenAddr))
	capPtr := (*int)(unsafe.Pointer(capAddr))

	// Create a pointer to the underlying array
	addPtr := (*[8]string)(unsafe.Pointer(*(*uintptr)(address)))

	fmt.Printf("Slice Addr[%p] Len Addr[0x%x] Cap Addr[0x%x]\n",
		address,
		lenAddr,
		capAddr)

	fmt.Printf("Slice Length[%d] Cap[%d]\n",
		*lenPtr,
		*capPtr)

	for index := 0; index < *lenPtr; index++ {
		fmt.Printf("[%d] %p %s\n",
			index,
			&(*addPtr)[index],
			(*addPtr)[index])
	}

	fmt.Printf("\n\n")
}