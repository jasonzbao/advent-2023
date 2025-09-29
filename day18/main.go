package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"utils"
)

var dummy = `R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)`

type Dig struct {
	trench [][]int
}

func (d *Dig) AddColumn(n int, isLeft bool) {
	// adds n columns to the right of the trench
	for i := 0; i < len(d.trench); i++ {
		if isLeft {
			d.trench[i] = append(make([]int, n), d.trench[i]...)
		} else {
			d.trench[i] = append(d.trench[i], make([]int, n)...)
		}
	}
}

func (d *Dig) AddRow(n int, isBottom bool) {
	// adds n rows to the bottom of the trench
	numColumns := len(d.trench[0])
	for i := 0; i < n; i++ {
		if isBottom {
			d.trench = append(d.trench, make([]int, numColumns))
		} else {
			d.trench = append([][]int{make([]int, numColumns)}, d.trench...)
		}
	}
}

func (d *Dig) ProcessInstruction(p Point, instruction Instruction) Point {
	switch instruction.Direction {
	case "R":
		if p.X+instruction.Amount > len(d.trench[0]) {
			d.AddColumn(instruction.Amount-len(d.trench[0])+p.X, false)
		}
		for i := p.X; i < p.X+instruction.Amount; i++ {
			d.trench[p.Y][i] = 1
		}
		p.X += instruction.Amount
		return p
	case "L":
		if p.X-instruction.Amount <= 0 {
			d.AddColumn(1+instruction.Amount-p.X, true)
			p.X = 1 + instruction.Amount
		}
		for i := p.X - 1; i >= p.X-instruction.Amount-1; i-- {
			d.trench[p.Y][i] = 1
		}
		p.X -= instruction.Amount
		return p
	case "U":
		if p.Y-instruction.Amount < 0 {
			d.AddRow(1+instruction.Amount-p.Y, false)
			p.Y = 1 + instruction.Amount
		}
		for i := p.Y - 1; i >= p.Y-instruction.Amount; i-- {
			d.trench[i][p.X-1] = 1
		}
		p.Y -= instruction.Amount
		return p
	case "D":
		if p.Y+instruction.Amount >= len(d.trench) {
			d.AddRow(instruction.Amount-len(d.trench)+p.Y+1, true)
		}
		for i := p.Y + 1; i < p.Y+instruction.Amount+1; i++ {
			d.trench[i][p.X-1] = 1
		}
		p.Y += instruction.Amount
		return p
	}
	return p
}

func (d *Dig) Print() {
	for i := 0; i < len(d.trench); i++ {
		for j := 0; j < len(d.trench[i]); j++ {
			// if i == 100 && j == 100 {
			// 	fmt.Print("@")
			// 	continue
			// }
			if d.trench[i][j] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}

		}
		fmt.Println()
	}
}

func (d *Dig) Area() int {
	area := 0
	for i := 0; i < len(d.trench); i++ {
		for j := 0; j < len(d.trench[i]); j++ {
			if d.trench[i][j] == 1 {
				area += 1
			}
		}
	}
	return area
}

func (d *Dig) FillInArea(x, y int) {
	// x y is a known point inside the trench
	if d.trench[x][y] == 1 {
		return
	}

	d.trench[x][y] = 1
	d.FillInArea(x+1, y)
	d.FillInArea(x-1, y)
	d.FillInArea(x, y+1)
	d.FillInArea(x, y-1)
}

type Instruction struct {
	Direction string
	Amount    int
	Color     string
}

type Point struct {
	X int
	Y int
}

func main() {
	input := utils.FormattedRequest(18)
	// input := utils.FormatInput(dummy)

	var instructions []Instruction
	for _, line := range input {
		parts := strings.Split(line, " ")
		amount, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal("Error parsing amount", err)
		}
		instructions = append(instructions, Instruction{
			Direction: parts[0],
			Amount:    amount,
			Color:     parts[2][1 : len(parts[2])-1],
		})
	}

	trench := make([][]int, 0)
	trench = append(trench, make([]int, 0))
	dig := Dig{
		trench: trench,
	}
	p := Point{X: 0, Y: 0}

	// start with 1
	instructions = append([]Instruction{{
		Direction: "R",
		Amount:    1,
		Color:     "NA",
	}}, instructions...)

	for _, instruction := range instructions {
		p = dig.ProcessInstruction(p, instruction)
	}

	dig.Print()
	dig.FillInArea(100, 100)
	dig.Print()
	fmt.Println(dig.Area())
}
