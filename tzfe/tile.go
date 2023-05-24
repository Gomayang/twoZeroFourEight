package tzfe

type tile struct {
	x        int
	y        int
	num      int
	isMerged bool
}

func (t *tile) move(dir direction) {
	for {
		if t.x+dir[0] < 0 || t.x+dir[0] >= mapSize || t.y+dir[1] < 0 || t.y+dir[1] >= mapSize {
			print("out of map", t.x, t.y, "\n")
			return
		}
		dest := checkTile(t.x+dir[0], t.y+dir[1])
		if *dest == *emptyTile() {
			t.x += dir[0]
			t.y += dir[1]
			continue
		}
		if dest.num == t.num && !dest.isMerged {
			t.num *= 2
			t.isMerged = true
			deleteTile(dest.x, dest.y)
			t.x += dir[0]
			t.y += dir[1]
		}
		return
	}
}

func emptyTile() *tile {
	return &tile{}
}

func deleteTile(x, y int) {
	for i, t := range tiles {
		if t.x == x && t.y == y {
			tiles = append(tiles[:i], tiles[i+1:]...)
			return
		}
	}
}

func checkTile(x, y int) *tile {
	for _, t := range tiles {
		if t.x == x && t.y == y {
			return t
		}
	}
	return &tile{}
}

func initMergeTiles() {
	for _, t := range tiles {
		t.isMerged = false
	}
}
