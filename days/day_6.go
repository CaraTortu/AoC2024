package days

import (
	"log"

	"AoC24/utils"
	"github.com/hashicorp/go-set/v3"
)

const (
	NORTH = iota
	SOUTH
	EAST
	WEST
)

const (
	EMPTY = iota
	OBSTACLE
	VISITED
)

type Day6Player struct {
	X         int
	Y         int
	Direction int
}

func (p *Day6Player) Hash() int {
	return p.X * p.Y * (p.Direction + 1)
}

type Day6Board struct {
	cells  []int
	Player Day6Player
	width  int
	height int
}

func day6GetBoard(filename string) Day6Board {
	contents, err := utils.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	boardHeight := len(contents)
	boardWidth := len(contents[0])

	board := make([]int, boardHeight*boardWidth)
	player := Day6Player{}
	for i, row := range contents {
		for j, cell := range row {
			switch cell {
			case '#':
				board[i*boardWidth+j] = OBSTACLE
				break
			case '^':
				player.X = j
				player.Y = i
				player.Direction = NORTH
				board[i*boardWidth+j] = VISITED
				break
			case 'v':
				player.X = j
				player.Y = i
				player.Direction = SOUTH
				board[i*boardWidth+j] = VISITED
				break
			case '>':
				player.X = j
				player.Y = i
				player.Direction = EAST
				board[i*boardWidth+j] = VISITED
				break
			case '<':
				player.X = j
				player.Y = i
				player.Direction = WEST
				board[i*boardWidth+j] = VISITED
				break
			default:
				board[i*boardWidth+j] = EMPTY
			}
		}
	}

	return Day6Board{
		cells:  board,
		Player: player,
		width:  len(contents[0]),
		height: len(contents),
	}
}

func (b *Day6Board) setCell(x, y, value int) {
	b.cells[y*b.width+x] = value
}

func (b *Day6Board) getCell(x, y int) int {
	return b.cells[y*b.width+x]
}

func (b *Day6Board) move() {
	switch b.Player.Direction {
	case NORTH:
		if b.getCell(b.Player.X, b.Player.Y-1) == OBSTACLE {
			b.Player.Direction = EAST
			break
		}
		b.Player.Y--
		break
	case SOUTH:
		if b.getCell(b.Player.X, b.Player.Y+1) == OBSTACLE {
			b.Player.Direction = WEST
			break
		}
		b.Player.Y++
		break
	case EAST:
		if b.getCell(b.Player.X+1, b.Player.Y) == OBSTACLE {
			b.Player.Direction = SOUTH
			break
		}
		b.Player.X++
		break
	case WEST:
		if b.getCell(b.Player.X-1, b.Player.Y) == OBSTACLE {
			b.Player.Direction = NORTH
			break
		}
		b.Player.X--
		break
	}

	b.setCell(b.Player.X, b.Player.Y, VISITED)
}

func (b *Day6Board) isLoop() bool {
	board := b.clone()
	playerHistory := set.New[Day6Player](len(b.cells) / 2)

	for {
		if board.Player.Y <= 0 || board.Player.Y >= board.height-1 || board.Player.X <= 0 || board.Player.X >= board.width-1 {
			return false
		}

		board.move()

		if playerHistory.Contains(board.Player) {
			return true
		}

		playerHistory.Insert(board.Player)
	}
}

func (b *Day6Board) clone() Day6Board {
	newBoard := make([]int, len(b.cells))
	copy(newBoard, b.cells)

	return Day6Board{
		cells:  newBoard,
		Player: b.Player,
		width:  b.width,
		height: b.height,
	}
}

func Day6A(filename string) int {
	board := day6GetBoard(filename)

	for {
		if board.Player.Y <= 0 || board.Player.Y >= board.height-1 || board.Player.X <= 0 || board.Player.X >= board.height-1 {
			break
		}

		board.move()
	}

	count := 0
	for _, cell := range board.cells {
		if cell == VISITED {
			count++
		}
	}

	return count
}

func Day6B(filename string) int {
	board := day6GetBoard(filename)
	goodBoard := board.clone()

	seen := make([]Day6Player, 0)
	for {
		if board.Player.Y <= 0 || board.Player.Y >= board.height-1 || board.Player.X <= 0 || board.Player.X >= board.width-1 {
			break
		}

		board.move()
		seen = append(seen, board.Player)
	}

	// Filter unique moves
	uniqueMoves := make([]Day6Player, 0)

	for _, player := range seen {
		found := false
		for _, unique := range uniqueMoves {
			if player.X == unique.X && player.Y == unique.Y {
				found = true
				break
			}
		}

		if !found {
			uniqueMoves = append(uniqueMoves, player)
		}
	}

	board = goodBoard

	loops := 0
	for _, player := range uniqueMoves {
		if board.getCell(player.X, player.Y) == EMPTY {
			board.setCell(player.X, player.Y, OBSTACLE)

			if board.isLoop() {
				loops++
			}

			board.setCell(player.X, player.Y, EMPTY)
		}
	}

	return loops
}
