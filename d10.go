package main

import (
	"strconv"
	"strings"
)

func (*methods) D10P1(input string) string {
	cmds := strings.Split(input, "\n")

	x := 1
	c := 1
	var strength int

	check := 20

	for _, cmd := range cmds {
		parts := strings.Split(cmd, " ")

		c++
		if c == check {
			strength += c * x
			check += 40
		}

		if parts[0] == "noop" {
			continue
		} else {
			c++
			inc, _ := strconv.Atoi(parts[1])
			x += inc
			if c == check {
				strength += c * x
				check += 40
			}
		}
	}

	return strconv.Itoa(strength)
}

func (*methods) D10P2(input string) string {
	cmds := strings.Split(input, "\n")

	x := 1
	var lines []string
	var line string
	for _, cmd := range cmds {
		parts := strings.Split(cmd, " ")

		pos := len(line)
		if pos == x || pos == x-1 || pos == x+1 {
			line += "#"
		} else {
			line += "."
		}
		if parts[0] == "addx" {
			pos := len(line)
			if pos == 40 {
				lines = append(lines, line)
				line = ""
				pos = 0
			}
			if pos == x || pos == x-1 || pos == x+1 {
				line += "#"
			} else {
				line += "."
			}
			inc, _ := strconv.Atoi(parts[1])
			x += inc
		}
		if len(line) == 40 {
			lines = append(lines, line)
			line = ""
		}
	}

	return strings.Join(lines, "\n")
}
