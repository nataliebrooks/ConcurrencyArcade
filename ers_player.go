// natalie brooks
// cs final project
//
// MATH 221 FALL 2016
//
// EGYPTIAN RATSCREW PLAYING SERVICE PLAYER
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
)

func scan(c chan string) {
	var message string
	fmt.Scanln(&message)
	c <- message
}

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Error: too few command-line arguments.\n")
		fmt.Printf("usage: go run player.go <host> <port>\n")
		os.Exit(0)
	}

	// Get the hostname from the command line.
	hostname := os.Args[1]
	// Get the port number.
	port,_ := strconv.Atoi(os.Args[2])

	// Make a connection, getting "proxy" channels cout/cin.
	cout, cin, e := cs221.MakeConnection(hostname, port, "player")
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(1)
	}

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