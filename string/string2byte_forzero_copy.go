package main

import "unsafe"

func string2bytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}
func bytes2string(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// reference
// https://qcrao91.gitbook.io/go/biao-zhun-ku/unsafe/ru-he-shi-xian-zi-fu-chuan-he-byte-qie-pian-de-ling-kao-bei-zhuan-huan
