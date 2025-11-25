# Pong, coded in Go, via Ebitengine

A minimal implementation of Pong in Go. 
Theoretically, to present in a short talk where I will live code Pong.

CONTROLS:
Player 1: up/down with w/s
Player 2: up/down with up arrow/down arrow
Space/Enter: serve ball
Reset: r
Quit: click on X on upper left

Fun, as Go is known as a backend language.

	// Hit space to start ball -- start state
	// -> make function to restart ball
	// Ball updates with some sort of direction!!!! should not be straight up down! How do you do that, radians?
	// Ball picks up velocity each hit, Direction opposite of incoming force? Based on where it hits paddle? --> ALL ABOUT THE velocityx and vy for the direction!
	// if player misses (goes over The 'X' borders) other player gets a point
	//
    TODO
    // Add Q key to quit
	// make an AI opponent, and have it shout insults! LOL.
    // i.e. "I beat Gary Kasparov, Deep Blue, and now I will beat you."

🏓 Add randomness to Y-axis velocity on paddle hits with a random generator. Otherwise you could get in a conidtion where the paddles remain static and just ping and pong back and forth forever. 

![Screenshot of black and white PONG remake](pong_screenshot.png)