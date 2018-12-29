package day07

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var digits = regexp.MustCompile("^\\d+$")

type and struct {
	first  string
	second string
}

type andNum struct {
	val uint16
	cmd string
}

type or struct {
	first  string
	second string
}

type lshift struct {
	cmd string
	val uint16
}

type rshift struct {
	cmd string
	val uint16
}

type not string

func reduce(lines []string, wireToCheck string) (uint16, uint16) {
	commands := make(map[string]interface{}, len(lines))
	origNums := map[string]uint16{}
	nums := map[string]uint16{}
	sort.Strings(lines)
	for _, line := range lines {
		cmd := strings.Split(line, " -> ")
		wire := cmd[1]
		if digits.MatchString(cmd[0]) {
			value, _ := strconv.ParseUint(cmd[0], 10, 16)
			nums[wire] = uint16(value)
			origNums[wire] = uint16(value)
		} else {
			parts := strings.Split(cmd[0], " ")
			switch len(parts) {
			case 1:
				commands[wire] = parts[0]
			case 2:
				commands[wire] = not(parts[1])
			case 3:
				switch parts[1] {
				case "AND":
					if digits.MatchString(parts[0]) {
						value, _ := strconv.ParseUint(parts[0], 10, 16)
						commands[wire] = andNum{uint16(value), parts[2]}
					} else {
						commands[wire] = and{parts[0], parts[2]}
					}
				case "OR":
					commands[wire] = or{parts[0], parts[2]}
				case "LSHIFT":
					value, _ := strconv.ParseUint(parts[2], 10, 16)
					commands[wire] = lshift{parts[0], uint16(value)}
				case "RSHIFT":
					value, _ := strconv.ParseUint(parts[2], 10, 16)
					commands[wire] = rshift{parts[0], uint16(value)}
				}
				break
			}
		}
	}

	// Refine until there's no more changes that can be done
	for len(nums) != len(lines) {
		for key, cmd := range commands {
			if _, ok := nums[key]; !ok {
				if num, ok := tryReduce(cmd, nums); ok {
					nums[key] = num
				}
			}
		}
	}

	wire := nums[wireToCheck]
	nums = origNums
	nums["b"] = wire
	delete(commands, "b")

	// Refine until there's no more changes that can be done
	for changes := true; changes; {
		changes = false
		for key, cmd := range commands {
			if _, ok := nums[key]; !ok {
				if num, ok := tryReduce(cmd, nums); ok {
					if key == wireToCheck {
						return wire, num
					}
					nums[key] = num
					changes = true
				}
			}
		}
	}

	return wire, nums[wireToCheck]
}

func tryReduce(cmd interface{}, nums map[string]uint16) (uint16, bool) {
	switch cmd.(type) {
	case string:
		num, ok := nums[cmd.(string)]
		return num, ok
	case not:
		if num, ok := nums[string(cmd.(not))]; ok {
			return ^num, true
		}
	case lshift:
		c := cmd.(lshift)
		if num, ok := nums[c.cmd]; ok {
			return num << c.val, true
		}
	case rshift:
		c := cmd.(rshift)
		if num, ok := nums[c.cmd]; ok {
			return num >> c.val, true
		}
	case or:
		c := cmd.(or)
		first, firstOk := nums[c.first]
		second, secondOk := nums[c.second]
		if firstOk && secondOk {
			return first | second, true
		}
	case and:
		c := cmd.(and)
		first, firstOk := nums[c.first]
		second, secondOk := nums[c.second]
		if firstOk && secondOk {
			return first & second, true
		}
	case andNum:
		c := cmd.(andNum)
		if num, ok := nums[c.cmd]; ok {
			return num & c.val, true
		}
	}
	return 0, false
}
