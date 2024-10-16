package main

import (
	"fmt"
	"log/slog"
)

// операции с битами
func Ex8() {
	var bigint int64
	var bitpos, bit int

	fmt.Printf("int64: ")
	_, err := fmt.Scanf("%d\n", &bigint)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	fmt.Printf("bitpos(1-64): ")
	_, err = fmt.Scanf("%d\n", &bitpos)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	if bitpos < 1 || bitpos > 64 {
		fmt.Println("incorrect input, terminating")
		return
	}
	fmt.Printf("bit(0-1): ")
	_, err = fmt.Scanf("%d\n", &bit)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	if !(bit == 0 || bit == 1) {
		fmt.Println("incorrect input, terminating")
		return
	}

	fmt.Printf("got: %b (%d), ", bigint, bigint)

	switch bit {
	case 0:
		bigint &= ^1 << (bitpos - 1)
	case 1:
		bigint |= 1 << (bitpos - 1)
	}

	fmt.Printf("result: %b (%d)", bigint, bigint)
}
