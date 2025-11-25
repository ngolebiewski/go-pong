# Pong in Go (with Ebitengine)

A minimal implementation of Pong in Go. 
Theoretically, to present in a short talk where I will live code Pong.

### CONTROLS

| Action | Player 1 Control | Player 2 Control
|---|---|---|
| Move Up | W | Up Arrow |
| Move Down | S | Down Arrow |
| Serve Ball | Space or Enter | - |
| Reset Game | R | - |
| Quit Game | Click on X (Upper Left) | Q |
| Toggle AI Mode | A | - |

Fun, as Go is known as a backend language.

### NOTES
	// Hit space to start ball -- start state
	// -> make function to restart ball
	// Ball updates with some sort of direction!!!! should not be straight up down! How do you do that, radians?
	// Ball picks up velocity each hit, Direction opposite of incoming force? Based on where it hits paddle? --> ALL ABOUT THE velocityx and vy for the direction!
    🏓 Added randomness to Y-axis velocity on paddle hits with a random generator. Otherwise you could get in a conidtion where the paddles remain static and just ping and pong back and forth forever. https://gobyexample.com/random-numbers
	// if player misses (goes over The 'X' borders) other player gets a point
     // Add Q key to quit: NOTE: the Update function in the Game loop needs to return a specific 'Termination' error to cleanly exit: ebiten.Termination
     // Added a very naive implementation of an AI player. It only tracks the Y value of the ball and moves accordingly. A smarter AI would idle after hitting the ball, and do some basic visual calculations with some random timeouts to GUESS where it may go. 

### TODO
   
	// imporve the AI opponent, and have it shout insults! LOL.
    // i.e. "I beat Gary Kasparov, Deep Blue, and now I will destroy you."
    // STATE MACHINE
        - Game Over Player X Wins
        - START Screen P1 P2? or AI -- Option to replace ball with Gopher



![Screenshot of black and white PONG remake](pong_screenshot.png)