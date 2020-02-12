# Concurrency Games and Arcade

This is an arcade project that I completed in my second CS class. It was the culmination of our study of the Go Programming Language and concurrency, and takes advantage of concurrency to create an arcade server where different processes may register their games and remove their games concurrently.

It also includes my implementation of a two-player game version of Egyptian Ratscrew (ers_server.go), which heavily relies on close timing between two people hitting specific keys upon certain cards appearing.

Game:
  For my game I decided to do a simplified Egyptian Ratscrew.
  Explanations for how that works are in the headers of both
  the ers_client_arc.go and ers_server_arc.go.

Arcade:

  All you need to do to get arcade running is to run its code:
     "go run arcade.go"

  In order to get games to register you just run their code:
     "go run echoservice.go" 

  They will prompt you to input a machine and port code to 
  register with. once you have registered the game to the arcade, 
  all you have to do is run the client code:
     "go run echoclient.go"
  It will automatically begin running the game so you can
  play!
