package main

import (
	"container/heap"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"utils"
)

const dummy = `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533`

type Puzzle struct {
	grid [][]int
}

type Tuple struct {
	x  int
	y  int
	dx int
	dy int

	Energy int
}

func (t Tuple) Key() string {
	return fmt.Sprintf("%d,%d,%d,%d", t.x, t.y, t.dx, t.dy)
}

// func (p *Puzzle) travelHelper(x, y int, memo map[string]int, seen map[string]bool, dx, dy int) int {
// 	t := Tuple{
// 		x:  x,
// 		y:  y,
// 		dx: dx,
// 		dy: dy,
// 	}
// 	if k, ok := memo[t.Key()]; ok {
// 		return k
// 	}

// 	if seen[t.Key()] {
// 		return largePositive
// 	}

// 	seenCopy := make(map[string]bool)
// 	for k, v := range seen {
// 		seenCopy[k] = v
// 	}
// 	seenCopy[t.Key()] = true

// 	memo[t.Key()] = p.travel(t, memo, seenCopy)
// 	return memo[t.Key()]
// }

// func (p *Puzzle) travel(t Tuple, memo map[string]int, seen map[string]bool) int {
// 	if t.x >= len(p.grid) || t.y >= len(p.grid[0]) || t.y < 0 || t.x < 0 {
// 		return largePositive
// 	}

// 	energy := p.grid[t.x][t.y]
// 	if t.x == len(p.grid)-1 && t.y == len(p.grid[0])-1 {
// 		return energy
// 	}

// 	stack := make([]int, 0)
// 	if t.dx < 3 && t.dx >= 0 {
// 		stack = append(stack, p.travelHelper(t.x+1, t.y, memo, seen, t.dx+1, 0))
// 	}
// 	if t.dx > -3 && t.dx <= 0 {
// 		stack = append(stack, p.travelHelper(t.x-1, t.y, memo, seen, t.dx-1, 0))
// 	}
// 	if t.dy < 3 && t.dy >= 0 {
// 		stack = append(stack, p.travelHelper(t.x, t.y+1, memo, seen, 0, t.dy+1))
// 	}
// 	if t.dy > -3 && t.dy <= 0 {
// 		stack = append(stack, p.travelHelper(t.x, t.y-1, memo, seen, 0, t.dy-1))
// 	}
// 	return energy + utils.GenericMin(stack)
// }

type NodeHeap []*Tuple

func (h NodeHeap) Len() int {
	return len(h)
}

func (h NodeHeap) Less(i, j int) bool {
	return h[i].Energy < h[j].Energy
}

func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NodeHeap) Push(x any) {
	*h = append(*h, x.(*Tuple))
}

func (h *NodeHeap) Pop() any {
	old := *h
	n := len(old)
	*h = old[0 : n-1]
	return old[n-1]
}

func (p *Puzzle) dijkstra() error {
	h := &NodeHeap{}
	heap.Init(h)

	visited := make(map[string]bool)
	dist := make(map[string]int)

	start := &Tuple{0, 0, 0, 0, 0}
	heap.Push(h, start)
	dist[start.Key()] = 0

	for h.Len() > 0 {
		currentNode := heap.Pop(h).(*Tuple)

		// Skip if we've already visited this state
		if visited[currentNode.Key()] {
			continue
		}
		visited[currentNode.Key()] = true

		// Check if we reached the goal
		if currentNode.x == len(p.grid)-1 && currentNode.y == len(p.grid[0])-1 {
			fmt.Println(currentNode.Energy)
			return nil
		}

		// Generate all valid next states
		neighbors := []Tuple{}

		if currentNode.dx >= 0 && currentNode.dx < 3 {
			if currentNode.x+1 < len(p.grid) {
				neighbors = append(neighbors, Tuple{
					x:  currentNode.x + 1,
					y:  currentNode.y,
					dx: currentNode.dx + 1,
					dy: 0,
				})
			}
		}
		if currentNode.dx <= 0 && currentNode.dx > -3 {
			if currentNode.x-1 >= 0 {
				neighbors = append(neighbors, Tuple{
					x:  currentNode.x - 1,
					y:  currentNode.y,
					dx: currentNode.dx - 1,
					dy: 0,
				})
			}
		}
		if currentNode.dy >= 0 && currentNode.dy < 3 {
			if currentNode.y+1 < len(p.grid[0]) {
				neighbors = append(neighbors, Tuple{
					x:  currentNode.x,
					y:  currentNode.y + 1,
					dx: 0,
					dy: currentNode.dy + 1,
				})
			}
		}
		if currentNode.dy <= 0 && currentNode.dy > -3 {
			if currentNode.y-1 >= 0 {
				neighbors = append(neighbors, Tuple{
					x:  currentNode.x,
					y:  currentNode.y - 1,
					dx: 0,
					dy: currentNode.dy - 1,
				})
			}
		}

		// Process each neighbor
		for _, neighbor := range neighbors {
			if visited[neighbor.Key()] {
				continue
			}

			if neighbor.x >= len(p.grid) || neighbor.y >= len(p.grid[0]) || neighbor.y < 0 || neighbor.x < 0 {
				continue
			}

			newEnergy := currentNode.Energy + p.grid[neighbor.x][neighbor.y]

			// Only add to heap if this is a better path
			if oldEnergy, exists := dist[neighbor.Key()]; !exists || newEnergy < oldEnergy {
				dist[neighbor.Key()] = newEnergy
				neighbor.Energy = newEnergy
				heap.Push(h, &neighbor)
			}
		}
	}

	return errors.New("no path found")
}

func (p *Puzzle) dijkstra2() error {
	h := &NodeHeap{}
	heap.Init(h)

	visited := make(map[string]bool)
	dist := make(map[string]int)

	start := &Tuple{0, 0, 0, 0, 0}
	heap.Push(h, start)
	dist[start.Key()] = 0

	for h.Len() > 0 {
		currentNode := heap.Pop(h).(*Tuple)

		// Skip if we've already visited this state
		if visited[currentNode.Key()] {
			continue
		}
		visited[currentNode.Key()] = true

		// Check if we reached the goal
		if currentNode.x == len(p.grid)-1 && currentNode.y == len(p.grid[0])-1 {
			fmt.Println(currentNode.Energy)
			return nil
		}

		// Generate all valid next states
		neighbors := []Tuple{}

		if currentNode.dx > 0 && currentNode.dx < 4 {
			neighbors = append(neighbors, Tuple{
				x:  currentNode.x + 1,
				y:  currentNode.y,
				dx: currentNode.dx + 1,
				dy: 0,
			})
		} else if currentNode.dy > 0 && currentNode.dy < 4 {
			neighbors = append(neighbors, Tuple{
				x:  currentNode.x,
				y:  currentNode.y + 1,
				dx: 0,
				dy: currentNode.dy + 1,
			})
		} else if currentNode.dx < 0 && currentNode.dx > -4 {
			neighbors = append(neighbors, Tuple{
				x:  currentNode.x - 1,
				y:  currentNode.y,
				dx: currentNode.dx - 1,
				dy: 0,
			})
		} else if currentNode.dy < 0 && currentNode.dy > -4 {
			neighbors = append(neighbors, Tuple{
				x:  currentNode.x,
				y:  currentNode.y - 1,
				dx: 0,
				dy: currentNode.dy - 1,
			})
		} else {
			// otherwise, we can go in any direction except in the opposite direction
			if currentNode.dx >= 0 && currentNode.dx < 10 {
				neighbors = append(neighbors, Tuple{
					x:  currentNode.x + 1,
					y:  currentNode.y,
					dx: currentNode.dx + 1,
					dy: 0,
				})
			}
			if currentNode.dx <= 0 && currentNode.dx > -10 {
				neighbors = append(neighbors, Tuple{
					x:  currentNode.x - 1,
					y:  currentNode.y,
					dx: currentNode.dx - 1,
					dy: 0,
				})
			}
			if currentNode.dy >= 0 && currentNode.dy < 10 {
				neighbors = append(neighbors, Tuple{
					x:  currentNode.x,
					y:  currentNode.y + 1,
					dx: 0,
					dy: currentNode.dy + 1,
				})
			}
			if currentNode.dy <= 0 && currentNode.dy > -10 {
				neighbors = append(neighbors, Tuple{
					x:  currentNode.x,
					y:  currentNode.y - 1,
					dx: 0,
					dy: currentNode.dy - 1,
				})
			}
		}

		// Process each neighbor
		for _, neighbor := range neighbors {
			if visited[neighbor.Key()] {
				continue
			}

			if neighbor.x >= len(p.grid) || neighbor.y >= len(p.grid[0]) || neighbor.y < 0 || neighbor.x < 0 {
				continue
			}

			newEnergy := currentNode.Energy + p.grid[neighbor.x][neighbor.y]

			// Only add to heap if this is a better path
			if oldEnergy, exists := dist[neighbor.Key()]; !exists || newEnergy < oldEnergy {
				dist[neighbor.Key()] = newEnergy
				neighbor.Energy = newEnergy
				heap.Push(h, &neighbor)
			}
		}
	}

	return errors.New("no path found")
}

func main() {
	input := utils.FormattedRequest(17)
	// input := utils.FormatInput(dummy)

	gridInts := make([][]int, 0)
	for _, line := range input {
		gridIntLine := make([]int, 0)
		for _, char := range strings.Split(line, "") {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				log.Fatal("Error parsing number", err)
			}
			gridIntLine = append(gridIntLine, num)
		}
		gridInts = append(gridInts, gridIntLine)
	}

	puzzle := Puzzle{
		grid: gridInts,
	}
	fmt.Println(puzzle.dijkstra2())

	// // Print nicely formatted grid
	// for i := 0; i < len(gridInts); i++ {
	// 	for j := 0; j < len(gridInts[0]); j++ {
	// 		stack := make([]int, 0)
	// 		for dx := -3; dx <= 3; dx++ {
	// 			for dy := -3; dy <= 3; dy++ {
	// 				if val, ok := seen[Tuple{i, j, dx, dy, 0}.Key()]; ok {
	// 					stack = append(stack, val)
	// 				}
	// 			}
	// 		}

	// 		// Format with 4 digits and right align
	// 		fmt.Printf("%4d", utils.GenericMin(stack))

	// 		if j < len(gridInts[0])-1 {
	// 			fmt.Print(" | ")
	// 		}
	// 	}
	// 	fmt.Println()
	// }

}
