package main

import "fmt"

func variableTwo() {
	var (
		i   int
		u   uint // positive number
		u8  uint8
		i8  int8
		i16 int16
		u16 uint16
		f   float32
		s   string
		b   bool
	)

	i = -15      // int32 (32 bits) | int64 (64 bits)
	u = 15       // uint32 (32 bits) | uint64 (64 bits)
	i8 = -128    // -128 - 127
	u8 = 254     // 0 - 255
	i16 = -32768 // -32768 - 32767
	u16 = 65535  // 0 - 65535
	f = 3.14     // float32
	s = "hello"
	s = "hello"
	b = true
	s = "Djood"
	b = true

	fmt.Println(i, u, i8, u8, i16, u16, f, s, b)
}
