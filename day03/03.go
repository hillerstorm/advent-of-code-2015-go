package day03

type pos struct {
	x, y int
}

func deliverPresents(input string, santas int) int {
	curPos := make([]pos, santas)
	for i := 0; i < santas; i++ {
		curPos[i] = pos{0, 0}
	}
	houses := map[pos]bool{
		pos{0, 0}: true,
	}
	idx := 0
	for _, x := range input {
		switch x {
		case '<':
			curPos[idx] = pos{curPos[idx].x - 1, curPos[idx].y}
			break
		case '>':
			curPos[idx] = pos{curPos[idx].x + 1, curPos[idx].y}
			break
		case '^':
			curPos[idx] = pos{curPos[idx].x, curPos[idx].y - 1}
			break
		case 'v':
			curPos[idx] = pos{curPos[idx].x, curPos[idx].y + 1}
			break
		}
		houses[curPos[idx]] = true
		idx = (idx + 1) % santas
	}

	return len(houses)
}
