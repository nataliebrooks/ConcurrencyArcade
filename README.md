# arcade

This is an arcade project that I completed in my second CS class. It was the culmination of our study of the Go Programming Language and concurrency, and takes advantage of concurrency to create an arcade server where different processes may register their games and remove their games concurrently.

arcade:
  all you need to do to get arcade running is to run its code:
     "go run arcade.go"

  in order to get games to register you just run their code:
     "go run echoservice.go" 

  and they will prompt you to input a machine and port code to 
  register with. once you have registered the game to the arcade 
  all you have to do is run the client code:
     "go run echoclient.go"
  and it will automatically begin running the game so you can
  play!
