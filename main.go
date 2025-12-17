package main

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"math/rand"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth   = 320
	screenHeight  = 240
	PaddleWidth   = 5.0
	PaddleHeight  = 40.0
	PaddleSpeed   = 4.0
	ballWidth     = 5.0
	ballInitSpeed = 2.0
)

type Game struct{}

type Paddle struct {
	x      float32
	y      float32
	Width  float32
	Height float32
}

type Ball struct {
	x     float32
	y     float32
	vx    float32
	vy    float32
	speed float32
}

type Player struct {
	score int
}

type State struct {
	start bool
	AI    bool
}

// Init entitites in game
var p1 = Paddle{x: 10.0, y: screenHeight/2 - PaddleHeight/2, Width: PaddleWidth, Height: PaddleHeight}
var p2 = Paddle{x: 305.0, y: screenHeight/2 - PaddleHeight/2, Width: PaddleWidth, Height: PaddleHeight}
var ball = Ball{x: screenWidth / 2, y: screenHeight / 2, vx: 0, vy: 0, speed: ballInitSpeed}
var player1 = Player{score: 0}
var player2 = Player{score: 0}
var state = State{start: true, AI: false}

// Seems too similar to the init!!! Starts game over.
func reset() {
	p1 = Paddle{x: 10.0, y: screenHeight/2 - PaddleHeight/2, Width: PaddleWidth, Height: PaddleHeight}
	p2 = Paddle{x: 305.0, y: screenHeight/2 - PaddleHeight/2, Width: PaddleWidth, Height: PaddleHeight}
	ball = Ball{x: screenWidth / 2, y: screenHeight / 2, vx: 0, vy: 0, speed: ballInitSpeed}
	player1 = Player{score: 0}
	player2 = Player{score: 0}
	state = State{start: true}
}

// Resets the Game to a new game state
func keyListener() error {
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		reset()
	}
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		fmt.Println("Q is pressed, but are we terminated?")
		return ebiten.Termination
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		state.AI = !state.AI
		fmt.Println("AI mode is turned", map[bool]string{true: "on", false: "off"}[state.AI])
	}
	return nil
}

// Paddle Methods
func (d Paddle) drawPaddle(p *Paddle, screen *ebiten.Image) {
	vector.FillRect(screen, p.x, p.y, p.Width, p.Height, color.White, false)
}

func (d *Paddle) updatePaddle1() {
	if (ebiten.IsKeyPressed(ebiten.KeyW) || gamepadUp(0)) && d.y >= 0 {
		d.y += -PaddleSpeed
	} else if (ebiten.IsKeyPressed(ebiten.KeyS) || gamepadDown(0)) && d.y <= screenHeight-PaddleHeight {
		d.y += PaddleSpeed
	}
}
func (d *Paddle) updatePaddle2() {
	if (ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) || gamepadUp(1)) && d.y >= 0 {
		d.y += -4.0
	} else if (ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) || gamepadDown(1)) && d.y <= screenHeight-PaddleHeight {
		d.y += 4.0
	}
}

// Ball Methods
func (b *Ball) drawBall(screen *ebiten.Image) {
	vector.FillRect(screen, b.x, b.y, ballWidth, ballWidth, color.White, false)
}

func (d *Ball) update() {
	if state.start {
		startPressed :=
			ebiten.IsKeyPressed(ebiten.KeySpace) ||
				ebiten.IsKeyPressed(ebiten.KeyEnter)

		// Gamepads 0 and 1
		for _, gp := range []ebiten.GamepadID{0, 1, 2, 3, 4} {
			if startButtonPressed(gp) {
				startPressed = true
			}
		}

		if startPressed {
			playBounce("start")
			dirX := float32(1.0)
			if rand.Float32() < .5 {
				dirX = -dirX
			}
			d.vx += dirX
			d.vy += float32(math.Min((rand.Float64()-.5)*4, 1.8))
			fmt.Println("Random start velocity -- vx:", d.vx, "vy:", d.vy)
			state.start = !state.start // Still needed?
			startPressed = false
		}
	}

	//move ball
	d.x += d.vx * d.speed
	d.y += d.vy * d.speed
}

func (d *Ball) collide() {
	//bounce off top or bottom
	if d.y <= 0 || d.y >= screenHeight {
		d.vy = -d.vy
		playBounce("wall")
	}
	//out of bounds left/right
	if d.x < 0 {
		player2.score++
		fmt.Println("LEFT Player One: ", player1.score, " Player Two: ", player2.score)
		d.reset()
	}
	if d.x > screenWidth {
		player1.score++
		fmt.Println("RIGHT Player One: ", player1.score, " Player Two: ", player2.score)
		d.reset()
	}

	//TODO: Add in 1/2 ball radius to the X and Y!!!

	if d.vx < 0 {
		//check to see if ball collides with player 1 paddle on left
		if d.x <= p1.x && p1.x <= d.x+PaddleWidth && p1.y <= d.y && d.y <= PaddleHeight+p1.y {
			playBounce("paddle")
			d.vy = -d.vy * (rand.Float32() + 1) // Add some randomness
			d.vx = -d.vx
			//cap the ball speed or else it gets too fast and flies through the paddles without getting detected as a collision!
			if d.speed < 5.0 {
				d.speed += .5
			}
		}
	}
	if d.vx > 0 {
		//check to see if ball collides with player 2 paddle on right
		if d.x <= p2.x && p2.x <= d.x+PaddleWidth && p2.y <= d.y && d.y <= PaddleHeight+p2.y {
			playBounce("paddle")
			d.vy = -d.vy * (rand.Float32() + 1)
			d.vx = -d.vx
			if d.speed < 5.0 {
				d.speed += .5
			}
		}
	}
	//
}

func (d *Ball) reset() {
	d.x = screenWidth / 2
	d.y = screenHeight / 2
	d.vx, d.vy, d.speed = 0, 0, 1.0
	state.start = true
}

// AI MODE
func dumbAI(p *Paddle, b *Ball) {
	//Naive implentation: look where ball is, move paddle towards it.
	//Where is ball y 0 top y=screenHeight is bottom

	// move to center when ball first rebounds off own paddle
	if b.vx < 0 && b.x > screenWidth/7 {
		if b.y < p.y+(PaddleHeight/2) && p.y >= screenHeight/2 {
			p.y -= 4
		}
		//down
		if b.y > p.y-PaddleHeight && p.y+(PaddleHeight/2) <= screenHeight/2 {
			p.y += 4
		}

	} else {
		//move paddle up
		if b.y < p.y+(ballWidth*2) && p.y >= 0 {
			p.y -= 4
		}
		//down
		if b.y > p.y && p.y+(PaddleHeight/2) <= screenHeight {
			p.y += 4
		}
	}
}

func (g *Game) Update() error {
	p1.updatePaddle1()
	if state.AI {
		// dumbAI(&p2, &ball)
		dumbAI(&p2, &ball)
	} else {
		p2.updatePaddle2()
	}
	ball.update()
	ball.collide()
	// MORE GAME FUNCTIONALITY HERE!
	// Hit R to RESET GAME & Q to QUIT on desktop
	if err := keyListener(); err != nil {
		return err // ✨ MUST return this!
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Go Pong!")
	screen.Fill(color.Black)

	//Shows player score. ToDo: Use a bitmap font rather than debug, but works for demo.
	ebitenutil.DebugPrintAt(screen, strconv.Itoa(player1.score), 40, 10)
	ebitenutil.DebugPrintAt(screen, strconv.Itoa(player2.score), screenWidth-40, 10)

	p1.drawPaddle(&p1, screen)
	p2.drawPaddle(&p2, screen)
	ball.drawBall(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Go Pong!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
