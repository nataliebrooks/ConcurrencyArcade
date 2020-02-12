# Concurrency Arcade and Games

This is an arcade project that I completed in my second CS class. It was the culmination of our study of the Go Programming Language and concurrency, and takes advantage of concurrency to create an arcade server where different processes may register their games and remove their games concurrently.

It also includes my implementation of a two-player game version of Egyptian Ratscrew (ers_server.go), which heavily relies on close timing between two people hitting specific keys upon certain cards appearing.

Environment:

You must have GO installed.

This program relies on a CS221 package that helps with creating and connecting to the server. It is supplied in the main folder of the repository, but must be copied to this location (or the location that GO is installed on your machine):

/usr/local/Cellar/go/1.7.3/libexec/src

Create a new folder titled "cs221" in the src folder, then copy the cs221.go file from this repository into it.

  
  Game (Egyptian Ratscrew):
  Explanations for how the game works are in the headers of both
  the ers_client_arc.go and ers_server_arc.go.
  
  Game (Echo):
  This is not really a game, just an echo that exhibits concurrent programming.



Running the Program:

Arcade:

  You must first start the arcade running so that it can host the games.
  To do this, run:
  
     "go run arcade.go"

  In order to get games to register you just run their code:
  (you will be prompted for machine (localhost) and port (8888))
     "go run echoservice.go" 
     "go run ers_server_arc.go"
     
  Once you have registered the game to the arcade, 
  all you have to do is run the client code:
     "go run echoclient.go"
     "go run ers_client_.go"
  It will automatically begin running the game so you can
  play!
