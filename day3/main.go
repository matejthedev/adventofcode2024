package main

import (
	_ "embed"
	"strconv"
	"unicode"
)

//go:embed data.txt
var dataTxt string

func calculateResult(data string) int {
	sum := 0
	i := 0
	n := len(data)

	for i < n {
		idx := indexOf(data, "mul(", i)
		if idx == -1 {
			break
		}

		pos := idx + 4

		pos = skipSpaces(data, pos)
		xStr, newPos, ok := parseInt(data, pos)
		if !ok {
			i = idx + 4
			continue
		}
		pos = newPos

		pos = skipSpaces(data, pos)
		if pos >= n || data[pos] != ',' {
			i = idx + 4
			continue
		}
		pos++

		pos = skipSpaces(data, pos)
		yStr, newPos, ok := parseInt(data, pos)
		if !ok {
			i = idx + 4
			continue
		}
		pos = newPos

		pos = skipSpaces(data, pos)
		if pos >= n || data[pos] != ')' {
			i = idx + 4
			continue
		}
		pos++

		x, errX := strconv.Atoi(xStr)
		y, errY := strconv.Atoi(yStr)
		if errX == nil && errY == nil {
			sum += x * y
		}

		i = pos
	}

	return sum
}

func indexOf(s, substr string, startIndex int) int {
	if startIndex >= len(s) {
		return -1
	}
	for i := startIndex; i+len(substr) <= len(s); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

func skipSpaces(s string, pos int) int {
	for pos < len(s) && unicode.IsSpace(rune(s[pos])) {
		pos++
	}
	return pos
}

func parseInt(s string, pos int) (string, int, bool) {
	start := pos
	for pos < len(s) && isDigit(s[pos]) {
		pos++
	}
	if pos == start {
		// No digits found
		return "", pos, false
	}
	return s[start:pos], pos, true
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func calculateConditionalResult(data string) int {
	sum := 0
	i := 0
	n := len(data)

	enabled := true

	for i < n {
		if i+4 <= n && data[i:i+3] == "do(" && i+3 < n && data[i+3] == ')' {
			enabled = true
			i += 4
			continue
		}

		if i+7 <= n && data[i:i+6] == "don't(" && data[i+6] == ')' {
			enabled = false
			i += 7
			continue
		}

		if i+4 <= n && data[i:i+4] == "mul(" {
			pos := i + 4

			pos = skipSpaces(data, pos)
			xStr, newPos, ok := parseInt(data, pos)
			if !ok {
				i += 4
				continue
			}
			pos = newPos

			pos = skipSpaces(data, pos)
			if pos >= n || data[pos] != ',' {
				i += 4
				continue
			}
			pos++

			pos = skipSpaces(data, pos)
			yStr, newPos, ok := parseInt(data, pos)
			if !ok {
				i += 4
				continue
			}
			pos = newPos

			pos = skipSpaces(data, pos)
			if pos >= n || data[pos] != ')' {
				i += 4
				continue
			}
			pos++

			x, errX := strconv.Atoi(xStr)
			y, errY := strconv.Atoi(yStr)
			if errX == nil && errY == nil && enabled {
				sum += x * y
			}

			i = pos
			continue
		}

		i++
	}

	return sum
}
