package main

import "go-interfaces-punteros/lib/logger"

func main() {
	var lg logger.Logger
	lg = logger.NewTextFile("log.txt")

	lg.Info("Hello World")
	lg = logger.NewLocal("red")
	lg.Info("Hello World 2!")
}
