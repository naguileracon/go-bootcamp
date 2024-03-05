package main

import "os"

func main() {
	file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_RDONLY|os.O_RDONLY, 0644)
	if err != nil {
		println(err.Error())
		return
	}
	defer file.Close()

	data := []byte("Hello, World!")

	_, err = file.Write(data)
	if err != nil {
		println(err.Error())
		return
	}

	println("Write data to file success")

}
