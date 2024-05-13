package main

import (
	utils "adventofcode/2015/go"
	"errors"
	"io"
	"os"
)

func Part2(filename string) (int64, error) {
	f, open_err := os.Open(filename)
	utils.Check(open_err)
	defer f.Close()

	floor := 0
	const read_size = 100
	var seek_offset int64 = 0
	var seek_err error
	var read_err error

	bytes := make([]byte, read_size)

	for {
		_, seek_err = f.Seek(seek_offset, 0)
		utils.Check(seek_err)

		_, read_err = f.Read(bytes)
		if read_err == io.EOF {
			break
		}
		utils.Check(read_err)

		for i, b := range bytes {
			switch b {
			case '(':
				floor++
			case ')':
				floor--
			}

			if floor < 0 {
				return seek_offset + int64(i) + 1, nil
			}
		}

		seek_offset += read_size
	}

	return 0, errors.New("has never been to the basement")
}
