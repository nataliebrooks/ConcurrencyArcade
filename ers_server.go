// natalie brooks
// cs final project
//
// MATH 221 FALL 2016
//
// EGYPTIAN RATSCREW PLAYING SERVICE HOST
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
    "fmt"
    "math/rand"
    "strconv"
    "time"
    "os"
    "cs221"
)

// Global channel shared by all connection handlers.
// Used to report game results.
type report struct {
    player string
    result string
}
var creport chan report = make(chan report)

//
// This flag should be true if the Terminal background
// is bright; false if it is dark.
//
var BACKGROUND_IS_BRIGHT bool = false

//
// Converts a card's code (0-51) into it's number,
// from 1 (for Ace) to 13 (for King).
//
func number(c int) int {
    return (c%13 + 1)
}

//
// Converts a card's code (0-51) into it's rank,
// a string that is A, 2, etc., 10, J, Q, K.
//
func rank(c int) string {
    var face = map[int]string{
        1:  "A",
        11: "J",
        12: "Q",
        13: "K",
    }
    var r int = number(c)
    if r >= 2 && r <= 10 {
        return strconv.Itoa(r)
    } else {
        return face[r]
    }
}

//
// Converts a card's code (0-51) into it's suit,
// a unicode character for spades, clubs, hearts,
// or diamonds.
//
func suit(c int) string {

    var suits map[int]string

    if BACKGROUND_IS_BRIGHT {
        suits = map[int]string{
            0: "\u2660", // spades
            1: "\u2663", // clubs
            2: "\u2661", // hearts
            3: "\u2662", // diamonds
        }
    } else {
        suits = map[int]string{
            0: "\u2664", // spades
            1: "\u2667", // clubs
            2: "\u2665", // hearts
            3: "\u2666", // diamonds
        }
    }
    var s int = c % 4
    return suits[s]
}

//
// Given a playing card's code (0-51), returns a
// string depicting that card.
//
func card(c int) string {
    return rank(c) + suit(c)
}

//
// Creates a 52 card deck, a random permutation of
// the card code values from 0 to 51.
//
//    0-12  spades
//    13-25 clubs
//    26-38 hearts
//    39-51 diamonds
//
func makeDeck() []int {
    var deck []int = make([]int, 52)
    for position, _ := range deck {
        deck[position] = position
    }
    rand.Seed(time.Now().UnixNano())
    for position, _ := range deck {
        var swapWith int = rand.Intn(52 - position) + position
        if swapWith != position {
            deck[position], deck[swapWith] = deck[swapWith], deck[position]
        }
        position = position + 1
    }
    return deck
}

func deal_card(p []int, h []int) ([]int, []int){
	p = append([]int{h[0]}, p...)
	h = h[1:]
	return p, h
}

//
// determines whether it is a slap situation based on the
// contents of the pile.
//
func is_slap(pile []int) bool {
	if len(pile) == 1 {
		return false
	} else if rank(pile[0]) == rank(pile[1]) {
		return true
	} else if len(pile) == 2 {
		return false
	} else if rank(pile[0]) == rank(pile[2]) {
		return true
	} else {
		return false
	}
}

//
// builds primary hands to begin the game
//
func make_hands(deck []int) ([]int, []int) {
	var host_hand []int = make([]int, 26)
	var player_hand []int = make([]int, 26)
	host_hand = deck[:26]
	player_hand = deck[26:]
	return host_hand, player_hand
} 

//
// adds the cards in the pile to the cards in a player's hand
//
func add_pile(hand []int, pile []int) []int {
	var l int = len(pile) + len(hand)
	var index int = 0
	var a []int = make([]int, l)
	for index < l {
		if index < len(hand) {
			a[index] = hand[index]
		} else {
			a[index] = pile[(index - len(hand))]
		}
		index += 1
	}
	return a
}

//
// given the two players hands, returns bool of whether game
// is over.
//
func is_over(player_hand []int, host_hand []int) bool {
	if len(player_hand) == 0 {
		return true
	} else if len(host_hand) == 0 {
		return true
	} else {
		return false
	}
}

//
// main_game loop, returns the outcome "won" | "lost"
//
func main_game(player_hand []int, host_hand []int, cin <-chan string, cout chan<- string) string {
	var turn bool = true
	var pile [] int = nil
	var outcome string = ""

	fmt.Println(len(player_hand))
	fmt.Println(len(host_hand))

	for is_over(player_hand, host_hand) == false {
		fmt.Printf("Player has %d cards left.\n", len(player_hand))
		fmt.Printf("Host has %d cards left.\n", len(host_hand))
		if turn == true {
			pile, host_hand = deal_card(pile, host_hand)
			fmt.Printf("Host deals %s\n", card(pile[0]))
			cout <- fmt.Sprintln(card(pile[0]))
			cout <- "\n"
		} else {
			pile, player_hand = deal_card(pile, player_hand)
			fmt.Printf("Player deals %s\n", card(pile[0]))
			cout <- fmt.Sprintln(card(pile[0]))
			cout <- "\n"
		}
		
		timeout := make(chan bool, 1) // create timeout channel
		go func() {
    	time.Sleep(1 * time.Second) // wait 2 seconds to send true on channel
    	timeout <- true
		}()

		var response string
		var player_slapped bool
		var player_fast bool

		select {
		case response = (<-cin):
			player_fast = true
			cout <- "Slap!\n"
			cout <- "\n"
		case <- timeout:
			player_fast = false
		}

		fmt.Println(response)

		if response == "s" {
			player_slapped = true
		} else {
			player_slapped = false
		}

		if player_slapped == false {
			response = (<-cin)
		}

		var slap bool = is_slap(pile)

		if slap && player_fast {
			fmt.Println("Player slaps.\n")
			fmt.Printf("Player wins pile of %d.\n\n\n", len(pile))
			cout <- fmt.Sprintf("You win pile of %d cards.\n", len(pile))
			player_hand = add_pile(player_hand, pile)
			pile = nil
		} else if slap {
			fmt.Println("Host slaps.\n")
			fmt.Printf("Host wins pile of %d cards.\n", len(pile))
			fmt.Println()
			cout <- fmt.Sprintln("Host slaps.")
			cout <- fmt.Sprintf("Host wins pile of %d cards.\n", len(pile))
			host_hand = add_pile(host_hand, pile) // host gets the pile
			pile = nil // pile is now empty
		} else if !slap && player_fast {
			pile = append(pile, player_hand[0]) // puts a card on the bottom of the pile
			player_hand = player_hand[1:]
			cout <- fmt.Sprintln("Note:\nUnwarranted slap. Lose 1 card.")
		}

		if turn == false {
			turn = true
		} else {
			turn = false
		}
	}
	if len(player_hand) == 0 {
		cout <- "You lost.\n"
		fmt.Println("The host has won.")
		outcome += "lost"
	} else {
		cout <- "You won!\n"
		fmt.Println("The player has won.")
		outcome += "won"
	}
    return outcome 
}

//
// Communication to the user is via a series of messages
// sent over the channel 'cout' and read from the 
// channel 'cin'. 
//
func handleGame(cout chan<- string, cin <-chan string, info interface{}) {

    // Play starts with the player sending their name.
    clientname := cs221.HeadLine(<-cin)

    fmt.Println("Player '"+clientname+"' has connected to play Egyptian Ratscrew.")

    //
    // Set up the deck and the two players' hands.
    var cards []int = makeDeck()
    player_hand, host_hand := make_hands(cards)
    //cout <- fmt.Sprintf("Note:\nHost goes first.\n")
    var outcome string = ""
    outcome = main_game(player_hand, host_hand, cin, cout)
	creport <- report{player:clientname, result:outcome}
	cout <- "GAME OVER\n"
    cout <- "\n"
}

//
// main
//
func main() {    
	if len(os.Args) < 3 {
		fmt.Printf("Error: too few command-line arguments.\n")
		fmt.Printf("usage: go run dealer.go <host> <port>\n")
		os.Exit(0)
	}

	hostname := os.Args[1]
	port,_ := strconv.Atoi(os.Args[2])

    // Run a reporting go routine.
    go func() {
        for {
            r := <-creport 
            fmt.Printf("REPORT: %s %s\n",r.player,r.result) 
        }
    }()

    // Loop, accepting and handling client connections.
    e := cs221.HandleConnections(hostname, port, handleGame, "Host", nil)

    // This code will only be reached when the server
    // is shutting down, or when an error occurs.
    //
    if e != nil {
        fmt.Println(e.Error())
        os.Exit(0)
    }

}






