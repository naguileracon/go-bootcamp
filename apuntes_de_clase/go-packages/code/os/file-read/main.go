package main

import (
	"bufio"
	"io"
	"os"
)

func main() {
	file, err := os.OpenFile("./code/os/file-read/test.txt", os.O_RDONLY, 0644)
	if err != nil {
		println(err.Error())
		return
	}
	defer file.Close()

	// read data step by step
	buffer := make([]byte, 100)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				println(err.Error())
				break
			}
			println(err.Error())
			return
		}
		println(string(buffer[:n]))

	}
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		println(err.Error())
		return
	}

	rd := bufio.NewReader(file)

	for {
		data, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				println(err.Error())
				break
			}
		}
		println(data)
	}

}
