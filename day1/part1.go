package main

import (
	utils "adventofcode/2015/go"
	"io"
	"os"
)

func Part1(filename string) int {
	return Part1_Seek(filename)
}

func Part1_ReadFile(filename string) int {
	dat, err := os.ReadFile(filename)
	utils.Check(err)

	str := string(dat)

	floor := 0
	for _, char := range str {
		switch char {
		case '(':
			floor++
		case ')':
			floor--
		}
	}

	return floor
}

func Part1_Seek(filename string) int {
	f, open_err := os.Open(filename)
	utils.Check(open_err)
	defer f.Close()

	floor := 0
	const read_size = 100
	var seek_offset int64 = 0
	var seek_err error
	var read_err error

	b := make([]byte, read_size)

	for {
		_, seek_err = f.Seek(seek_offset, 0)
		utils.Check(seek_err)

		_, read_err = f.Read(b)
		if read_err == io.EOF {
			break
		}
		utils.Check(read_err)

		floor += findFloor(b)
		seek_offset += read_size
	}

	return floor
}

func findFloor(input []byte) int {
	floor := 0

	for _, b := range input {
		switch b {
		case '(':
			floor++
		case ')':
			floor--
		}
	}

	return floor
}
