package parser

import (
	"log"
	"os"
)

var opcodes = []string{
	"halt",
	"set",
	"push",
	"pop",
	"eq",
	"gt",
	"jmp",
	"jt",
	"jf",
	"add",
	"mult",
	"mod",
	"and",
	"or",
	"not",
	"rmem",
	"wmem",
	"call",
	"ret",
	"out",
	"in",
	"noop",
}

func Parse(opcode uint16) string {
	if int(opcode) > len(opcodes)-1 {
		return "UNKNOWN"
	}
	//fmt.Println(opcode)
	//fmt.Println(opcode)
	return opcodes[opcode]
}

func ParseMemory(program []uint16) {
	out, err := os.OpenFile("bytecode.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	out.Truncate(0)
	bytecode := log.New(out, "", 0)
	//bytecode := bufio.NewWriter(out)
	//bytecode
	index := 0
	for {
		if index > len(program)-1 {
			break
		}
		opcode := Parse(program[index])
		switch opcode {
		case "halt":
			bytecode.Printf("%d   \t| %s", index, opcode)
			index++
			break
		case "set":
			bytecode.Printf("%d   \t| %s %d %d\n", index, opcode, program[index+1], program[index+2])
			index += 3
			break
		case "push":
			bytecode.Printf("%d   \t| %s %d", index, opcode, program[index+1])
			index += 2
			break
		case "pop":
			bytecode.Printf("%d   \t| %s %d", index, opcode, program[index+1])
			index += 2
			break
		case "eq":
			bytecode.Printf("%d   \t| %s %d %d %d", index, opcode, program[index+1], program[index+2], program[index+3])
			index += 4
			break
		case "gt":
			bytecode.Printf("%d   \t| %s %d %d %d", index, opcode, program[index+1], program[index+2], program[index+3])
			index += 4
			break
		case "jmp":
			bytecode.Printf("%d   \t| %s %d", index, opcode, program[index+1])
			index += 2
			break
		case "jt":
			bytecode.Printf("%d   \t| %s %d %d\n", index, opcode, program[index+1], program[index+2])
			index += 3
			break
		case "jf":
			bytecode.Printf("%d   \t| %s %d %d\n", index, opcode, program[index+1], program[index+2])
			index += 3
			break
		case "add":
			bytecode.Printf("%d   \t| %s %d %d %d", index, opcode, program[index+1], program[index+2], program[index+3])
			index += 4
			break
		case "mult":
			bytecode.Printf("%d   \t| %s %d %d %d", index, opcode, program[index+1], program[index+2], program[index+3])
			index += 4
			break
		case "mod":
			bytecode.Printf("%d   \t| %s %d %d %d", index, opcode, program[index+1], program[index+2], program[index+3])
			index += 4
			break
		case "and":
			bytecode.Printf("%d   \t| %s %d %d %d", index, opcode, program[index+1], program[index+2], program[index+3])
			index += 4
			break
		case "or":
			bytecode.Printf("%d   \t| %s %d %d %d", index, opcode, program[index+1], program[index+2], program[index+3])
			index += 4
			break
		case "not":
			bytecode.Printf("%d   \t| %s %d %d\n", index, opcode, program[index+1], program[index+2])
			index += 3
			break
		case "wmem":
			bytecode.Printf("%d   \t| %s %d %d\n", index, opcode, program[index+1], program[index+2])
			index += 3
			break
		case "call":
			bytecode.Printf("%d   \t| %s %d", index, opcode, program[index+1])
			index += 2
			break
		case "ret":
			bytecode.Printf("%d   \t| %s", index, opcode)
			index++
			break
		case "out":
			bytecode.Printf("%d   \t| %s %c", index, opcode, program[index+1])
			index += 2
			break
		case "noop":
			bytecode.Printf("%d   \t| %s", index, opcode)
			index++
			break
		case "in":
			bytecode.Printf("%d   \t| %s %d", index, opcode, program[index+1])
			index += 2
			break
		default:
			index++
			break
		}

	}
}
