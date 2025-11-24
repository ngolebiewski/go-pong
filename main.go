package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 320
	screenHeight = 240
	PaddleWidth  = 5.0
	PaddleHeight = 40.0
	PaddleSpeed  = 4.0
	ballWidth    = 5.0
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
}

// Init entitites in game
var p1 = Paddle{x: 10.0, y: screenHeight/2 - PaddleHeight/2, Width: PaddleWidth, Height: PaddleHeight}
var p2 = Paddle{x: 305.0, y: screenHeight/2 - PaddleHeight/2, Width: PaddleWidth, Height: PaddleHeight}
var ball = Ball{x: screenWidth / 2, y: screenHeight / 2, vx: 0, vy: 0, speed: 5.0}
var player1 = Player{score: 0}
var player2 = Player{score: 0}
var state = State{start: true}

func reset() {
	p1 = Paddle{x: 10.0, y: screenHeight/2 - PaddleHeight/2, Width: PaddleWidth, Height: PaddleHeight}
	p2 = Paddle{x: 305.0, y: screenHeight/2 - PaddleHeight/2, Width: PaddleWidth, Height: PaddleHeight}
	ball = Ball{x: screenWidth / 2, y: screenHeight / 2, vx: 0, vy: 0, speed: 1.0}
	player1 = Player{score: 0}
	player2 = Player{score: 0}
	state = State{start: true}
}

func resetListener() {
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		reset()
	}
}

// Paddle Methods
func (d Paddle) drawPaddle(p *Paddle, screen *ebiten.Image) {
	vector.FillRect(screen, p.x, p.y, p.Width, p.Height, color.White, false)
}

func (d *Paddle) updatePaddle1() {
	if ebiten.IsKeyPressed(ebiten.KeyW) && d.y >= 0 {
		d.y += -PaddleSpeed
	} else if ebiten.IsKeyPressed(ebiten.KeyS) && d.y <= screenHeight-PaddleHeight {
		d.y += PaddleSpeed
		fmt.Println(screenHeight, d.y)
	}
}
func (d *Paddle) updatePaddle2() {
	if (ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyArrowUp)) && d.y >= 0 {
		d.y += -4.0
	} else if (ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyArrowDown)) && d.y <= screenHeight-PaddleHeight {
		d.y += 4.0
	}
}

// Ball Methods
func (b *Ball) drawBall(screen *ebiten.Image) {
	vector.FillRect(screen, b.x, b.y, ballWidth, ballWidth, color.White, false)
}

func (d *Ball) update() {
	if state.start {
		if ebiten.IsKeyPressed(ebiten.KeySpace) || ebiten.IsKeyPressed(ebiten.KeyEnter) {
			d.vx += (rand.Float32() - .5) * 2 * d.speed
			d.vy += (rand.Float32() / 2) * d.speed
			state.start = !state.start
		}
	}
	//move ball
	d.x += d.vx
	d.y += d.vy
	//collissions?

	// d.x += 1 * d.speed // test to move it!
}

func (d *Ball) reset() {
	d.x = screenWidth / 2
	d.y = screenHeight / 2
	d.vx, d.vy, d.speed = 0, 0, 1.0
}

func (g *Game) Update() error {
	p1.updatePaddle1()
	p2.updatePaddle2()
	ball.update()
	// Hit space to start ball -- start state
	// -> make function to restart ball
	// Ball updates with some sort of direction!!!! should not be straight up down! How do you do that, radians?
	// Ball picks up velocity each hit, Direction opposite of incoming force? Based on where it hits paddle? --> ALL ABOUT THE velocityx and vy for the direction!
	// if player misses (goes over The 'X' borders) other player gets a point
	//
	// make an AI opponent, and have it shout insults! LOL.
	resetListener() // Hit R to RESET GAME
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
