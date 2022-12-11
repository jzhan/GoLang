package main

import (
	"fmt"
	"main/winapi"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const (
	BLACK   = 0
	BLUE    = 1
	GREEN   = 2
	AQUA    = 3
	RED     = 4
	PURPLE  = 5
	YELLOW  = 6
	WHITE   = 7
	GRAY    = 8
	LBLUE   = 9
	LGREEN  = 10
	LAQUA   = 11
	LRED    = 12
	LPURPLE = 13
	LYELLOW = 14
	BWHITE  = 15

	VK_ESCAPE = 0x1B

	VK_CONVERT    = 0x1C
	VK_NONCONVERT = 0x1D
	VK_ACCEPT     = 0x1E
	VK_MODECHANGE = 0x1F

	VK_SPACE    = 0x20
	VK_PRIOR    = 0x21
	VK_NEXT     = 0x22
	VK_END      = 0x23
	VK_HOME     = 0x24
	VK_LEFT     = 0x25
	VK_UP       = 0x26
	VK_RIGHT    = 0x27
	VK_DOWN     = 0x28
	VK_SELECT   = 0x29
	VK_PRINT    = 0x2A
	VK_EXECUTE  = 0x2B
	VK_SNAPSHOT = 0x2C
	VK_INSERT   = 0x2D
	VK_DELETE   = 0x2E
	VK_HELP     = 0x2F

	BORDER_WIDTH  = 39
	BORDER_HEIGHT = 15
)

type Coord struct {
	X int32
	Y int32
}

type Body struct {
	dir rune
	pos Coord
}

type Snake struct {
	len  int
	body [BORDER_WIDTH * BORDER_HEIGHT]Body
}

type Food struct {
	pos Coord
}

var snake Snake
var food Food
var prev_snake_coord Coord = Coord{-1, -1}
var isEaten bool = true
var score = -1

func main() {
	winapi.ShowCursor(0)

	for {
		Render()

		if winapi.GetAsyncKeyState(VK_ESCAPE)&0x800 != 0 {
			break
		} else if winapi.GetAsyncKeyState(VK_SPACE)&0x8000 != 0 {

		} else if winapi.GetAsyncKeyState(VK_LEFT)&0x8001 != 0 && snake.body[0].dir != 'R' {
			snake.body[0].dir = 'L'
		} else if winapi.GetAsyncKeyState(VK_UP)&0x8001 != 0 && snake.body[0].dir != 'D' {
			snake.body[0].dir = 'U'
		} else if winapi.GetAsyncKeyState(VK_RIGHT)&0x8001 != 0 && snake.body[0].dir != 'L' {
			snake.body[0].dir = 'R'
		} else if winapi.GetAsyncKeyState(VK_DOWN)&0x8001 != 0 && snake.body[0].dir != 'U' {
			snake.body[0].dir = 'D'
		}

		winapi.Sleep(150)
		Move()

		if CollisionCheck() == -1 {
			winapi.Gotoxy(BORDER_WIDTH/2-5, BORDER_HEIGHT/2)
			fmt.Print("GAME OVER")

			break
		}

	}

	winapi.Gotoxy(0, BORDER_HEIGHT)
}

// Start automatically before main function
func init() {
	snake.body[0].dir = 'L'
	snake.len = 5
	snake.body[0].pos.X = BORDER_WIDTH / 2
	snake.body[0].pos.Y = BORDER_HEIGHT / 2

	for i := 1; i < snake.len; i++ {
		snake.body[i].dir = 'L'
		snake.body[i].pos.X = snake.body[i-1].pos.X + 1
		snake.body[i].pos.Y = snake.body[i-1].pos.Y
	}

	cls := exec.Command("cmd", "/C", "cls")
	cls.Stdout = os.Stdout
	cls.Run()

	// Draw border
	for i := 0; i < BORDER_HEIGHT; i++ {
		for j := 0; j < BORDER_WIDTH; j++ {
			if i >= 0 && i < BORDER_HEIGHT-1 && (j == 0 || j == BORDER_WIDTH-1) {
				fmt.Print(string(rune(0x2588)))
			} else if i == 0 || i == BORDER_HEIGHT-1 || j == 0 || j == BORDER_WIDTH-1 {
				fmt.Print(string(rune(0x2580)))
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}
}

func GenerateFoodNewposition() {
	rand.Seed(time.Now().UTC().UnixNano())

	for {
		flag := false

		food.pos.X = (rand.Int31() % (BORDER_WIDTH - 2)) + 1
		food.pos.Y = (rand.Int31() % (BORDER_HEIGHT - 2)) + 1

		for i := 0; i < snake.len; i++ {
			if food.pos.X == snake.body[i].pos.X && food.pos.Y == snake.body[i].pos.Y {
				flag = true

				break
			}
		}

		if !flag {
			break
		}
	}
}

func CollisionCheck() int32 {
	if snake.body[0].pos.X == food.pos.X && snake.body[0].pos.Y == food.pos.Y {
		isEaten = true
		snake.len += 1

		switch snake.body[snake.len-2].dir {
		case 'L':
			new_x := snake.body[snake.len-2].pos.X + 1
			new_y := snake.body[snake.len-2].pos.Y

			if new_x == BORDER_WIDTH-1 {
				new_x -= 1

				if new_y+1 != BORDER_WIDTH-1 {
					new_y += 1
				} else {
					new_y -= 1
				}
			}

			snake.body[snake.len-1].pos = Coord{new_x, new_y}
		case 'U':
			new_x := snake.body[snake.len-2].pos.X
			new_y := snake.body[snake.len-2].pos.Y + 1

			if new_y == BORDER_HEIGHT-1 {
				new_y -= 1

				if new_x+1 != BORDER_WIDTH-1 {
					new_x += 1
				} else {
					new_x -= 1
				}
			}

			snake.body[snake.len-1].pos = Coord{new_x, new_y}
		case 'R':
			new_x := snake.body[snake.len-2].pos.X - 1
			new_y := snake.body[snake.len-2].pos.Y

			if new_x == 0 {
				new_x += 1

				if new_y+1 != BORDER_HEIGHT-1 {
					new_y += 1
				} else {
					new_y -= 1
				}
			}

			snake.body[snake.len-1].pos = Coord{new_x, new_y}
		case 'D':
			new_x := snake.body[snake.len-2].pos.X
			new_y := snake.body[snake.len-2].pos.Y - 1

			if new_y == 0 {
				new_y += 1

				if new_x+1 != BORDER_WIDTH-1 {
					new_x += 1
				} else {
					new_x -= 1
				}
			}

			snake.body[snake.len-1].pos = Coord{new_x, new_y}
		}
	}

	// check if head hit the border
	if snake.body[0].pos.X == 0 || snake.body[0].pos.Y == 0 ||
		snake.body[0].pos.X == BORDER_WIDTH-1 || snake.body[0].pos.Y == BORDER_HEIGHT-1 {
		return -1
	}

	// check if head hit body
	for i := 1; i < snake.len; i++ {
		if snake.body[0].pos.X == snake.body[i].pos.X && snake.body[0].pos.Y == snake.body[i].pos.Y {
			return -1
		}
	}

	return 0
}

func Render() {
	if isEaten {
		isEaten = false
		score += 1

		GenerateFoodNewposition()
		winapi.Gotoxy(food.pos.X, food.pos.Y)
		winapi.SetTextColor(LRED)

		print("@")

		winapi.Gotoxy(0, BORDER_HEIGHT)
		winapi.SetTextColor(WHITE)

		fmt.Print("Score: ", score)

	}

	winapi.SetTextColor(LBLUE)

	for i := 0; i < snake.len; i++ {
		winapi.Gotoxy(snake.body[i].pos.X, snake.body[i].pos.Y)

		print("*")

		if i == 0 {
			winapi.SetTextColor(LAQUA)
		} else if i == snake.len-2 {
			winapi.SetTextColor(LYELLOW)
		}
	}
	winapi.SetTextColor(WHITE)
	winapi.Gotoxy(prev_snake_coord.X, prev_snake_coord.Y)
	print(" ")
}

func Move() {
	prev_snake_coord.X = snake.body[0].pos.X
	prev_snake_coord.Y = snake.body[0].pos.Y

	switch snake.body[0].dir {
	case 'L':
		snake.body[0].pos.X -= 1
	case 'U':
		snake.body[0].pos.Y -= 1
	case 'R':
		snake.body[0].pos.X += 1
	case 'D':
		snake.body[0].pos.Y += 1
	}

	for i := 1; i < snake.len; i++ {
		if prev_snake_coord.X != snake.body[i].pos.X {
			if prev_snake_coord.X > snake.body[i].pos.X {
				snake.body[i].dir = 'R'
			} else {
				snake.body[i].dir = 'L'
			}
		} else {
			if prev_snake_coord.Y > snake.body[i].pos.Y {
				snake.body[i].dir = 'U'
			} else {
				snake.body[i].dir = 'D'
			}
		}

		temp_x := prev_snake_coord.X
		temp_y := prev_snake_coord.Y

		prev_snake_coord.X = snake.body[i].pos.X
		prev_snake_coord.Y = snake.body[i].pos.Y

		snake.body[i].pos.X = temp_x
		snake.body[i].pos.Y = temp_y
	}
}
