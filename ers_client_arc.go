// natalie brooks
// cs final project
//
// MATH 221 FALL 2016
//
// EGYPTIAN RATSCREW PLAYING SERVICE PLAYER
//
// NOTE: this is a version that finds a host
//       through joining arcade.go!!!
//
// egyptian ratscrew is a game of speed where 
// cards are separated into two decks and each
// player is given a deck. players take turns
// dealing cards into a pile. 
//
// if two of the same card (e.g. twos or jacks)
// are played with either 0 cards (called a "pair") 
// or 1 card (called a "sandwich") separating
// them, then both players try to slap (press enter)
// before the other. whichever player slaps first
// wins the pile.
//
// if a player slaps the pile without reason they
// must add one card to the bottom of the pile.
//
// the game ends when one of the players ends up
// with the entire deck.
//
// the player must press any key and press enter
// before the host slaps (1 second).

package main

import (
	"cs221"
	"fmt"
	"strconv"
	"os"
	"time"
	"strings"
)

func scan(c chan string) {
	var message string
	fmt.Scanln(&message)
	c <- message
}

func connect() (out chan<- string, in <-chan string) {
	success := false
	for success == false {
		out, in, e := cs221.MakeConnection("localhost",6666,"ers")
		out <- fmt.Sprintln("LOOKUP ers")
		out <- "\n"
		reply := cs221.HeadLine(<-in)
		parts := strings.Split(reply," ")

		if parts[0] == "No" {
			fmt.Println(reply)
			os.Exit(1)
		}
		port, _ := strconv.Atoi(parts[1])
		host := parts[0]
		out, in, e = cs221.MakeConnection(host, port, "ers")
		if e != nil {
			fmt.Println("Removing.")
			out, in, e = cs221.MakeConnection("localhost",6666,"ers")
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

	cout, cin := connect()

	cc := make(chan string)

	// Get the player's information. Send it to the server.
	fmt.Print("Enter your name: ")
	var name string	
	fmt.Scanln(&name)	
	cout <- name + "\n"
	cout <- "\n"

	// Play a hand of blackjack.  Note that this is a
	// "thin" client.  Most of the interaction involves
	// repeating what is sent from the server and 
	// forwarding responses from the user.
	//
	done := false
	for !done {
		//
		// Get a response from the server. Output it.
		reply := <-cin
		fmt.Println(reply)
		
		lines := cs221.Lines(reply)

		if lines[len(lines)-1] == "GAME OVER" {

			//
			// End the proxy's session by sending "".
			cout <- ""
			done = true

		} else {

			// if after one second scan has not happened and
			// sent the value on channel then terminate
			// sends and end line 

			timeout := make(chan bool, 1)
			go func() {
    			time.Sleep(1 * time.Second)
    			timeout <- true
			}()

			go scan(cc)

			select {
				case <-cc:
    				cout <- "s\n"
				case <-timeout:
					cout <- "n\n"			
			}
		cout <- "\n"
		}
	}
}