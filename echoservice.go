package main

import (
	"cs221"
	"fmt"
	"strconv"
	"os"
	"math/rand"
)

func handleEcho(out chan<- string, in <-chan string, info interface{}) {
	for {
		message := <-in
		if rand.Intn(10) == 0 {
			out <- fmt.Sprintf("Come again, eh?!??\n")
			out <- "\n"
		} else {
			msg := cs221.HeadLine(message)
			out <- fmt.Sprintf("%s, eh!\n",msg)
			out <- "\n"
		}
	}
}

func main() {
	var hostname string = ""
	fmt.Println("Enter your machine: ")
	fmt.Scanln(&hostname)
	var port string = ""
	fmt.Println("Enter your port: ")
	fmt.Scanln(&port)

	out, in, f := cs221.MakeConnection("localhost", 6666, "Echo")
	out <- "ADD Echo "+hostname+" "+port+" \n"
	out <- "\n"
	okay := (<-in)
	fmt.Println(okay)

	var p int
	p, _ = strconv.Atoi(port)
	f = cs221.HandleConnections(hostname, p, handleEcho, "Echo", nil)

	if f != nil {
		fmt.Println(f.Error())
		os.Exit(0)

	}
}




