package main

import (
	"strconv"
	"strings"
)

type d5stack []byte

func (s d5stack) Push(v ...byte) d5stack {
	return append(s, v...)
}

func (s d5stack) Pop() (d5stack, byte) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s d5stack) PopM(qty int) (d5stack, []byte) {
	l := len(s)
	return s[:l-qty], s[l-qty:]
}

func (*methods) D5P1(input string) string {
	stacks := make(map[int]d5stack)

	parts := strings.Split(input, "\n\n")

	// parse and populate stacks
	stackLines := strings.Split(parts[0], "\n")
	for o := len(stackLines) - 2; o >= 0; o-- {
		line := stackLines[o]
		var c int
		for i := 1; i < len(line)-1; i += 4 {
			if line[i] != 0x20 {
				stacks[c] = stacks[c].Push(line[i])
			}
			c++
		}
		c = 0
	}

	// parse and execute instructions
	cmds := strings.Split(parts[1], "\n")
	for _, cmd := range cmds {
		parts := strings.Split(cmd, " ")
		qty, _ := strconv.Atoi(parts[1])
		src, _ := strconv.Atoi(parts[3])
		src--
		dest, _ := strconv.Atoi(parts[5])
		dest--

		for i := 0; i < qty; i++ {
			stack, block := stacks[src].Pop()
			stacks[src] = stack
			stacks[dest] = stacks[dest].Push(block)
		}
	}

	var result []byte
	var i int
	for {
		stack, ok := stacks[i]
		if !ok {
			break
		}
		result = append(result, stack[len(stack)-1])
		i++
	}

	return string(result)
}

func (*methods) D5P2(input string) string {
	stacks := make(map[int]d5stack)

	parts := strings.Split(input, "\n\n")

	// parse and populate stacks
	stackLines := strings.Split(parts[0], "\n")
	for o := len(stackLines) - 2; o >= 0; o-- {
		line := stackLines[o]
		var c int
		for i := 1; i < len(line)-1; i += 4 {
			if line[i] != 0x20 {
				stacks[c] = stacks[c].Push(line[i])
			}
			c++
		}
		c = 0
	}

	// parse and execute instructions
	cmds := strings.Split(parts[1], "\n")
	for _, cmd := range cmds {
		parts := strings.Split(cmd, " ")
		qty, _ := strconv.Atoi(parts[1])
		src, _ := strconv.Atoi(parts[3])
		src--
		dest, _ := strconv.Atoi(parts[5])
		dest--

		stack, blocks := stacks[src].PopM(qty)
		stacks[src] = stack
		stacks[dest] = stacks[dest].Push(blocks...)
	}

	var result []byte
	var i int
	for {
		stack, ok := stacks[i]
		if !ok {
			break
		}
		result = append(result, stack[len(stack)-1])
		i++
	}

	return string(result)
}
