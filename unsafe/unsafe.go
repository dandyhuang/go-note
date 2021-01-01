package main

import (
	"fmt"
	"unsafe"
)

// string like
//typedef struct{
//
//char* buffer;
//
//size_t len;
//
//} string;
type dandy struct {
	sex  uint8
	name string
	f    int
}

func UnsafeTest() {
	d := dandy{66, "dandy", 88}
	fmt.Println(d)
	sex := (*int)(unsafe.Pointer(&d))
	*sex = 2

	f := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&d)) + unsafe.Offsetof(d.f)))
	*f = 3
	fmt.Println(d)
	// equivalent to f
	f = (*int)(unsafe.Pointer(&d.f))
	*f = 6
	fmt.Println(d)
	// because Byte alignment
	fmt.Println("sex:", unsafe.Sizeof(d.sex), "name sizeof:", unsafe.Sizeof(d.name), "len:", len(d.name),
		"sex:", unsafe.Sizeof(d.f))
	// equivalent to f
	f = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&d)) + 8 /*unsafe.Sizeof(d.sex)*/ + unsafe.Sizeof(d.name)))

	*f = 9
	fmt.Println(d)

	fmt.Printf("%p\n", &d.sex)
	fmt.Printf("%p\n", &d.name)
	fmt.Printf("%p\n", &d.f)
	fmt.Printf("x align:%v offset:%v\n", unsafe.Alignof(d.sex), unsafe.Offsetof(d.sex))
	fmt.Printf("x align:%v offset:%v\n", unsafe.Alignof(d.name), unsafe.Offsetof(d.name))
	fmt.Printf("x align:%v offset:%v\n", unsafe.Alignof(d.f), unsafe.Offsetof(d.f))
}

func main() {
	UnsafeTest()
}
