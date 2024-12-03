package main

import (
	"math"
	"path/filepath"
	"strconv"
	"strings"
)

func (*methods) D7P1(input string) string {
	m := make(map[string]int64)
	dirs := make(map[string]int64)

	var workdir string
	parts := strings.Split(input, "$ ")
	for i, part := range parts {
		if i == 0 {
			continue
		}
		lines := strings.Split(part, "\n")
		cmd := strings.TrimSpace(lines[0])

		if strings.HasPrefix(cmd, "cd ") {
			target := strings.TrimPrefix(cmd, "cd ")
			workdir = filepath.Join(workdir, target)
			continue
		}

		for _, l := range lines[1:] {
			l = strings.TrimSpace(l)
			if l == "" || strings.HasPrefix(l, "dir") {
				continue
			}
			parts := strings.Split(l, " ")
			size, _ := strconv.ParseInt(parts[0], 10, 64)
			m[filepath.Join(workdir, parts[1])] = size
		}
	}

	for file, size := range m {
		dir := filepath.Dir(file)
		for {
			dirs[dir] += size
			if dir == "/" {
				break
			}
			dir = filepath.Dir(dir)
		}
	}

	var total int64
	for _, size := range dirs {
		if size <= 100000 {
			total += size
		}
	}

	return strconv.FormatInt(total, 10)
}

func (*methods) D7P2(input string) string {
	m := make(map[string]int64)
	dirs := make(map[string]int64)
	var totalUsed int64

	var workdir string
	parts := strings.Split(input, "$ ")
	for i, part := range parts {
		if i == 0 {
			continue
		}
		lines := strings.Split(part, "\n")
		cmd := strings.TrimSpace(lines[0])

		if strings.HasPrefix(cmd, "cd ") {
			target := strings.TrimPrefix(cmd, "cd ")
			workdir = filepath.Join(workdir, target)
			continue
		}

		for _, l := range lines[1:] {
			l = strings.TrimSpace(l)
			if l == "" || strings.HasPrefix(l, "dir") {
				continue
			}
			parts := strings.Split(l, " ")
			size, _ := strconv.ParseInt(parts[0], 10, 64)
			m[filepath.Join(workdir, parts[1])] = size
			totalUsed += size
		}
	}

	for file, size := range m {
		dir := filepath.Dir(file)
		for {
			dirs[dir] += size
			if dir == "/" {
				break
			}
			dir = filepath.Dir(dir)
		}
	}

	freeSpace := 70000000 - totalUsed
	needed := 30000000 - freeSpace

	minSize := int64(math.MaxInt64)

	for _, size := range dirs {
		if size >= needed && size < minSize {
			minSize = size
		}
	}

	return strconv.FormatInt(minSize, 10)
}
