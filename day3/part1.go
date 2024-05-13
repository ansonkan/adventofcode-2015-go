package main

import (
	utils "adventofcode/2015/go"
	"io"
	"os"
)

func Part1(filename string) int {
	f, open_err := os.Open(filename)
	utils.Check(open_err)
	defer f.Close()

	cur := [2]int{0, 0}
	m := make(map[[2]int]bool)
	m[cur] = true
	houses := 1

	const read_size = 10
	var seek_offset int64 = 0
	var seek_err error
	var read_count int
	var read_err error

	bytes := make([]byte, read_size)

	for {
		_, seek_err = f.Seek(seek_offset, 0)
		utils.Check(seek_err)
		seek_offset += read_size

		read_count, read_err = f.Read(bytes)
		if read_err == io.EOF {
			break
		}
		utils.Check(read_err)

	ReadByteLoop:
		for i, b := range bytes {
			if i >= read_count {
				break
			}

			switch b {
			case '^':
				cur[1]++
			case '>':
				cur[0]++
			case 'v':
				cur[1]--
			case '<':
				cur[0]--
			default:
				break ReadByteLoop
			}

			if _, hasKey := m[cur]; !hasKey {
				m[cur] = true
				houses++
			}
		}

	}

	return houses
}
