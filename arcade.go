// natalie brooks

// MATH 221 FALL 2016
//
// ARCADE SERVICE FOR CLIENTS AND HOSTS
//

package main

import (
	"cs221"
	"fmt"
	"strings"
)


type game struct {
	name string
	hostname string
	port string
}

type register struct {
	games [] game
}

func remove(i int, g [] game) [] game {
	var gg [] game
	if i == 0 && len(g) > 1 {
		gg = g[1:]
	} else if i == 0 {
		gg = nil
	} else {
		first := g[:i - 1]
		last := g[i:]
		gg = append(first, last...)	
	}
return gg
}

type transaction struct {
	request string
	replyChan chan<- string
}

func handleTransaction(request string, r *register) string {

	// Pick apart the request string.
	
	// One of:
	//   "LOOKUP\n<name>
	//   "REMOVE\n<port>"
	//   "ADD\n<name>\n<hostname>\n<port>"
	//
	// NOTE: I wrote my code differently so rather
	// than giving the connecting player a list of
	// all games on the service, it gives the player
	// only the port and hostname of the game that
	// matches them.
	//
	parts := strings.Split(request," ")
	
	if parts[0] == "LOOKUP" {
		var index int = 0
		found := false
		for index < len(r.games) && found == false {
			loc := r.games[index]
			if loc.name == parts[1] {
				fmt.Println(""+parts[1]+" player connected to game "+loc.name+" at "+loc.hostname+" on port "+loc.port+".")
				return ""+loc.hostname+" "+loc.port+""
				found = true
			}
		index += 1
		}
		if found == false {
			return "No games by that name regisitered."
		}	
	} else if parts[0] == "REMOVE" {
		var index int = 0
		found := false
		for index < len(r.games) && found == false {
			loc := r.games[index]
			if loc.port == parts[1] {
				fmt.Println("Game "+loc.name+" at "+loc.hostname+" on port "+loc.port+" removed from Arcade register.")
				r.games = remove(index, r.games)
				fmt.Println(r.games)
				return "Stale game removed."
				found = true
			}
		index += 1
		}	
		return "Stale game removed."	
	} else if parts[0] == "ADD" {
		new_game := game{name:parts[1], hostname:parts[2], port:parts[3]}
		r.games = append(r.games, new_game)
		fmt.Println("Game "+parts[1]+" at "+parts[2]+" on port "+parts[3]+" added to Arcade register.")
		return "Your game has been added to the Arcade register."
	}
	return "Invalid Request"
}

//
// Given the register of hosts, starts off a goroutine that manages
// transactions made to the register.
//
// Gives back a channel `tc` that can be used to send
// transaction requests.
//
func startArcade(r *register) chan transaction {

	tc := make(chan transaction)

	go func () {
		//
		// Handle a sequence of transactions delivered
		// over the channel.
		//
		for t := range tc { // what to do with clients
			reply := handleTransaction(t.request, r)
			t.replyChan <- reply + "\n" // give code of game what to connect
			t.replyChan <- "\n"
		}
	} ()
						
	return tc
}

//
func handleClient(cout chan<- string, cin <-chan string, tc chan transaction) {

	fmt.Println("A client has connected.")

	done := false
	// For each message `r` sent from the ATM...
	for !done {
		msg := (<-cin)
		fmt.Println(msg)
		if msg != "" {
			// ...forward it to the bank over the channel `tc`.
			tc <- transaction{request:cs221.HeadLine(msg),replyChan:cout}
		} else {
			done = true
		}
	}

	fmt.Println("Client disconnected.")
}

func main() {
	clients,_ := cs221.HandleAllConnections("localhost", 6666)

	fmt.Println("Bringing up the arcade.")
	r := &register{games: nil}
	b := startArcade(r)
	
	fmt.Println("Waiting for your client connections.")
	
	for {
		conn := (<-clients) // get a client connection
		go handleClient(conn.Out,conn.In,b) // handle it
	}
}