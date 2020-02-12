package main

import (
	"cs221"
	"fmt"
	"strconv"
	"os"
	"strings"
)

func search() (out chan<- string, in <-chan string) {
	success := false
	for success == false {
		out, in, e := cs221.MakeConnection("localhost",6666,"Shout")
		out <- fmt.Sprintln("LOOKUP Echo")
		out <- "\n"
		reply := cs221.HeadLine(<-in)
		parts := strings.Split(reply," ")

		if parts[0] == "No" {
			fmt.Println(reply)
			os.Exit(1)
		}
		port, _ := strconv.Atoi(parts[1])
		host := parts[0]
		out, in, e = cs221.MakeConnection(host, port, "Shout")
		if e != nil {
			fmt.Println("Removing.")
			out, in, e = cs221.MakeConnection("localhost",6666,"Shout")
			out <- "REMOVE "+parts[1]+" \n"
			out <- "\n"
			ok := (<-in)
			fmt.Println(ok)
			out <- ""
		} else {
			success = true
			return out, in
		}
		if e != nil {
		fmt.Println(e.Error())
		os.Exit(1)
		}
	}
return out, in
}

func main() {
	out, in := search()

	fmt.Println("Enter lines of text below:")
	for {
		var message string	
		fmt.Scanln(&message)	
		out <- message + "\n"
		out <- "\n"
		reply := <-in
		fmt.Println(cs221.HeadLine(reply))
	}
}
