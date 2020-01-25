package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type gameState int

const (
	start gameState = iota
	play
)

var state = start

const (
	screenWidth  = 800
	screenHeight = 600
)

var nums = [][]byte{
	{
		1, 1, 1,
		1, 0, 1,
		1, 0, 1,
		1, 0, 1,
		1, 1, 1,
	},
	{
		1, 1, 0,
		0, 1, 0,
		0, 1, 0,
		0, 1, 0,
		1, 1, 1,
	},
	{
		1, 1, 1,
		0, 0, 1,
		1, 1, 1,
		1, 0, 0,
		1, 1, 1,
	},
	{
		1, 1, 1,
		0, 0, 1,
		0, 1, 1,
		0, 0, 1,
		1, 1, 1,
	},
}

type color struct {
	r, g, b byte
}

type pos struct {
	x, y float32
}
type ball struct {
	pos
	xVel, yVel, radius float32
	color              color
}

func drawNum(pos pos, color color, size int, num int, pixels []byte) {
	startX := int(pos.x) - (size*3)/2
	startY := int(pos.y) - (size-5)/2

	for i, v := range nums[num] {
		if v == 1 {
			for y := startY; y < startY+size; y++ {
				for x := startX; x < startX+size; x++ {
					setPixel(x, y, color, pixels)
				}
			}
		}
		startX += size
		if (i+1)%3 == 0 {
			startY += size
			startX -= size * 3
		}
	}
}

func (ball *ball) draw(pixels []byte) {
	for y := -ball.radius; y < ball.radius; y++ {
		for x := -ball.radius; x < ball.radius; x++ {
			if x*x+y*y < ball.radius*ball.radius {
				setPixel(int(ball.x+x), int(ball.y+y), ball.color, pixels)
			}
		}
	}
}

func center() pos {
	return pos{screenHeight / 2, screenWidth / 2}
}

func (ball *ball) update(Lpaddle *paddle, Rpaddle *paddle, elapsedTime float32) {
	ball.x += ball.xVel * elapsedTime
	ball.y += ball.yVel * elapsedTime

	if ball.y-ball.radius < 0 || ball.y+ball.radius > screenHeight {
		ball.yVel = -ball.yVel
	}
	if ball.x < 0 {
		Rpaddle.score++
		ball.x, ball.y = screenWidth/2, screenHeight/2
		state = start
	} else if ball.x > screenWidth {
		Lpaddle.score++
		ball.x, ball.y = screenWidth/2, screenHeight/2
		state = start
	}

	if ball.x-ball.radius < Lpaddle.x+Lpaddle.w/2 {
		if ball.y > Lpaddle.y-Lpaddle.h/2 && ball.y < Lpaddle.y+Lpaddle.h/2 {
			ball.xVel = -ball.xVel
			ball.x = Lpaddle.x + Lpaddle.w/2.0 + ball.radius
		}
	}
	if ball.x+ball.radius > Rpaddle.x-Rpaddle.w/2 {
		if ball.y > Rpaddle.y-Rpaddle.h/2 && ball.y < Rpaddle.y+Rpaddle.h/2 {
			ball.xVel = -ball.xVel
			ball.x = Rpaddle.x - Rpaddle.w/2.0 - ball.radius
		}
	}
}

type paddle struct {
	pos
	h, w, speed float32
	color       color
	score       int
}

func lerp(a, b, percent float32) float32 {
	return a + percent*(b-a)
}

func (paddle *paddle) draw(pixels []byte) {
	startX := int(paddle.x - paddle.w/2)
	startY := int(paddle.y - paddle.h/2)

	for y := 0; y < int(paddle.h); y++ {
		for x := 0; x < int(paddle.w); x++ {
			setPixel(startX+x, startY+y, paddle.color, pixels)
		}
	}

	numX := lerp(paddle.x, center().x, 0.2)
	drawNum(pos{numX, 35}, paddle.color, 10, paddle.score, pixels)
}

func (paddle *paddle) update(keyState []uint8, elapsedTime float32) {
	if keyState[sdl.SCANCODE_W] != 0 {
		paddle.y -= paddle.speed * elapsedTime
	}
	if keyState[sdl.SCANCODE_S] != 0 {
		paddle.y += paddle.speed * elapsedTime
	}
}

func (paddle *paddle) aiUpdate(ball *ball, elapsedTime float32) {
	paddle.y = ball.y
}

func clear(pixels []byte) {
	for i := range pixels {
		pixels[i] = 0
	}
}

func setPixel(x, y int, c color, pixels []byte) {
	index := (y*screenWidth + x) * 4

	if index < len(pixels)-4 && index >= 0 {
		pixels[index] = c.r
		pixels[index+1] = c.g
		pixels[index+2] = c.b
	}
}

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("Init SDL:", err)
		return
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Pong", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println("Init Window:", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Init Renderer:", err)
		return
	}
	defer renderer.Destroy()

	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, screenWidth, screenHeight)
	if err != nil {
		fmt.Println("Init tex:", err)
		return
	}
	defer tex.Destroy()

	pixels := make([]byte, screenWidth*screenHeight*4)

	player1 := paddle{pos{50, screenHeight / 2}, 50, 10, 300, color{255, 255, 255}, 0}
	player2 := paddle{pos{screenWidth - 50, screenHeight / 2}, 50, 10, 300, color{255, 255, 255}, 0}
	ball := ball{pos{screenWidth / 2, screenHeight / 2}, 200, 200, 4, color{255, 255, 255}}

	keyState := sdl.GetKeyboardState()

	var frameStart time.Time
	var elapsedTime float32

	//game loop
	for {
		frameStart = time.Now()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		if state == play {
			player1.update(keyState, elapsedTime)
			player2.aiUpdate(&ball, elapsedTime)
			ball.update(&player1, &player2, elapsedTime)
		} else if state == start {
			if keyState[sdl.SCANCODE_SPACE] != 0 {
				if player1.score == 3 || player2.score == 3 {
					player1.score, player2.score = 0, 0
					player1.x, player1.y = 50, screenHeight/2
					player2.x, player2.y = screenWidth-50, screenHeight/2
					ball.x, ball.y = screenWidth/2, screenHeight/2
				}
				state = play
			}
		}

		clear(pixels)
		player1.draw(pixels)
		player2.draw(pixels)
		ball.draw(pixels)

		tex.Update(nil, pixels, screenWidth*4)
		renderer.Copy(tex, nil, nil)
		renderer.Present()

		elapsedTime = float32(time.Since(frameStart).Seconds())
		if elapsedTime < .005 {
			sdl.Delay(5 - uint32(elapsedTime*1000.0))
			elapsedTime = float32(time.Since(frameStart).Seconds())
		}
	}
}
