package tzfe

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var tiles = []*tile{}
var mapSize = 4
var isDebug = true

func generateNewTile() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		x := rand.Intn(mapSize)
		y := rand.Intn(mapSize)

		if *checkTile(x, y) == *emptyTile() {
			tiles = append(tiles, &tile{
				x:        x,
				y:        y,
				num:      2,
				isMerged: false,
			})
			break
		}
	}
}

func printTiles() {
	for i := mapSize - 1; i >= 0; i-- {
		for j := 0; j < mapSize; j++ {
			fmt.Printf("┌──────┐")
		}
		fmt.Println("")
		for j := 0; j < mapSize; j++ {
			tile := checkTile(j, i)
			if *tile == *emptyTile() {
				fmt.Printf("│      │")
				continue
			}
			fmt.Printf("│%5d │", tile.num)
		}
		fmt.Println("")
		for j := 0; j < mapSize; j++ {
			fmt.Printf("└──────┘")
		}
		fmt.Println("")
	}
}

func selectDirection() direction {
	fmt.Print(`
===================
2048

1. left
2. right
3. up
4. down
5. quit
===================

숫자를 입력해주세요 : `)
	input := ""
	num, err := fmt.Scan(&input)
	if err != nil {
		print(err)
		return selectDirection()
	}
	if num != 1 {
		return selectDirection()
	}
	switch input {
	case "1":
		return leftDir()
	case "2":
		return rightDir()
	case "3":
		return upDir()
	case "4":
		return downDir()
	case "5":
		return direction{}
	default:
		return selectDirection()
	}
}

func moveTiles(dir direction) {
	switch dir {
	case leftDir():
		for i := 0; i < mapSize; i++ {
			for j := 0; j < mapSize; j++ {
				tile := checkTile(i, j)
				if *tile == *emptyTile() {
					continue
				}
				tile.move(dir)
			}
		}
	case rightDir():
		for i := mapSize - 1; i >= 0; i-- {
			for j := 0; j < mapSize; j++ {
				tile := checkTile(i, j)
				if *tile == *emptyTile() {
					continue
				}
				tile.move(dir)
			}
		}
	case upDir():
		for i := 0; i < mapSize; i++ {
			for j := 0; j < mapSize; j++ {
				tile := checkTile(i, j)
				if *tile == *emptyTile() {
					continue
				}
				tile.move(dir)
			}
		}
	case downDir():
		for i := 0; i < mapSize; i++ {
			for j := mapSize - 1; j >= 0; j-- {
				tile := checkTile(i, j)
				if *tile == *emptyTile() {
					continue
				}
				tile.move(dir)
			}
		}
	}
}

func checkHighestTile() int {
	highest := 0
	for _, t := range tiles {
		if t.num > highest {
			highest = t.num
		}
	}
	return highest
}

func StartGame() {
	print("select map size")
	mapSize = selectMapSize()
	if mapSize == 0 {
		return
	}

	print("Generating new tile...")
	generateNewTile()

	print("print tiles")
	printTiles()

	for {
		print("init merge tiles")
		initMergeTiles()
		print("select direction")
		dir := selectDirection()
		if dir == (direction{}) {
			StartGame()
			return
		}
		print("move tiles")
		moveTiles(dir)
		print("generate new tile")
		generateNewTile()
		print("print tiles")
		printTiles()

		if len(tiles) == mapSize*mapSize {
			fmt.Println("Game Over")
			StartGame()
			return
		}
		if checkHighestTile() == 2048 {
			fmt.Println("You Win")
			StartGame()
			return
		}
	}

}

func selectMapSize() int {
	fmt.Print(`
===================
2048
	
1. 4x4
2. 5x5
3. quit
===================

숫자를 입력해주세요 : `)
	input := ""
	for {
		num, err := fmt.Scan(&input)
		if err != nil {
			print(err)
			continue
		}
		if num != 1 {
			continue
		}
		intInput, err := strconv.Atoi(input)
		if err != nil {
			print(err)
			continue
		}
		switch intInput {
		case 1:
			return 4
		case 2:
			return 5
		case 3:
			return 0
		}
	}

}

func print(a ...any) {
	if isDebug {
		fmt.Println(a...)
	}
}
