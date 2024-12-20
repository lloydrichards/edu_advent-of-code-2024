package day17

import (
	U "advent-of-code-2024/internal/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Program struct {
	a, b, c      int
	instructions []int
}

func (p *Program) combo(oprand int) int {
	switch oprand {
	case 0, 1, 2, 3:
		return oprand
	case 4:
		return p.a
	case 5:
		return p.b
	case 6:
		return p.c
	default:
		return -1
	}
}

func (p *Program) run() string {
	pointer := 0
	output := []int{}

	for pointer < len(p.instructions) {
		instruction := p.instructions[pointer]
		opcode := p.instructions[pointer+1]

		switch instruction {
		case 0: // adv
			p.a = p.a >> p.combo(opcode)
		case 1: // bxl
			p.b = p.b ^ opcode
		case 2: // bst
			p.b = p.combo(opcode) % 8
		case 3: // jnz
			if p.a != 0 {
				pointer = opcode - 2
			}
		case 4: // bxc
			p.b = p.b ^ p.c
		case 5: // out
			output = append(output, p.combo(opcode)%8)
		case 6: //bdv
			p.b = p.a >> p.combo(opcode)
		case 7: //bdv
			p.c = p.a >> p.combo(opcode)
		}

		pointer += 2
	}

	outputStr := make([]string, len(output))
	for i, o := range output {
		outputStr[i] = strconv.Itoa(o)
	}

	return strings.Join(outputStr, ",")

}

func parseProgram(input string) Program {
	parts := strings.Split(input, "\n\n")

	re := regexp.MustCompile(`\d+`)
	registrar := re.FindAllString(parts[0], -1)
	instrument := re.FindAllString(parts[1], -1)

	a, _ := strconv.Atoi(registrar[0])
	b, _ := strconv.Atoi(registrar[1])
	c, _ := strconv.Atoi(registrar[2])

	instructions := make([]int, 0)
	for _, i := range instrument {
		j, _ := strconv.Atoi(i)
		instructions = append(instructions, j)
	}

	return Program{a, b, c, instructions}
}

func Part1(dir string) (string, error) {
	input, err := U.LoadInputFile(dir)
	if err != nil {
		return "", err
	}

	program := parseProgram(input)

	output := program.run()

	return output, nil
}

func (p *Program) reverseEngineer(program []int, ans int) (int, error) {

	if len(program) == 0 {
		return ans, nil
	}

	list := make([]int, 8)

	// b = a % 8       <- 2,4
	// b = b ^ 1       <- 1,1
	// c = a >> b      <- 7,5
	// b = b ^ 5       <- 1,5
	// b = b ^ c       <- 4,3

	// out = b % 8     <- 5,5
	// a = a << 3      <- 0,3
	// if a == 0: loop <- 3,0

	for bit := range list {
		p.a = (ans << 3) + bit

		p.b = p.a % 8    // <- 2,4
		p.b = p.b ^ 1    // <- 1,1
		p.c = p.a >> p.b // <- 7,5
		p.b = p.b ^ 5    // <- 1,5
		p.b = p.b ^ p.c  // <- 4,3

		if p.b%8 == program[len(program)-1] {
			valid, err := p.reverseEngineer(program[:len(program)-1], p.a)
			if err != nil {
				continue
			}
			return valid, nil
		}

	}
	return -1, fmt.Errorf("no valid answer")

}

func Part2(dir string) (int, error) {
	input, err := U.LoadInputFile(dir)
	if err != nil {
		return -1, err
	}
	program := parseProgram(input)

	ans, err := program.reverseEngineer(program.instructions, 0)
	if err != nil {
		return -1, err
	}

	return ans, nil
}
