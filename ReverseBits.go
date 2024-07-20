package main

import (
	"fmt"
	"log"
	"strconv"
)

func reverseBits(num uint32) uint32 {
	binary := Uint32To32BitBinary(num)
	log.Printf("binary = " + binary)
	if len(binary) < 32 {
		binary = "0" + binary
	}

	binaryReverse := Reverse(binary)
	log.Printf("binaryReverse = " + binaryReverse)

	result, err := BinaryToUint32(binaryReverse)

	if err != nil {
		return 0
	}
	log.Printf("result = " + strconv.Itoa(int(result)))
	return result
}

// Uint32To32BitBinary converts a uint32 to its 32-bit binary representation as a string
func Uint32To32BitBinary(n uint32) string {
	// Convert the number to binary
	binaryStr := strconv.FormatUint(uint64(n), 2)
	// Ensure the result is 32 bits long by padding with leading zeros if necessary
	return fmt.Sprintf("%032s", binaryStr)
}

// BinaryToUint32 converts a 32-bit binary string to a uint32
func BinaryToUint32(binaryStr string) (uint32, error) {
	// Parse the binary string to a uint64
	value, err := strconv.ParseUint(binaryStr, 2, 32)
	if err != nil {
		return 0, err
	}
	// Cast the uint64 to a uint32
	return uint32(value), nil
}
