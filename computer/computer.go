package computer

import (
	"fmt"
	"github.com/whcass/synacor-challenge/parser"
)

const REGISTER_START int = 32768

type Computer struct {
	memory        []uint16
	registers     [8]*uint16
	stack         []uint16
	memoryPointer int
}

func (c Computer) GetVar() uint16 {
	return c.memory[c.memoryPointer]
}

func (c Computer) GetVarOffset(offset int) uint16 {
	val := c.memory[c.memoryPointer+offset]
	if int(val) >= REGISTER_START {
		registerIndex := c.MapRegister(int(val))
		return c.GetRegisterVal(registerIndex)
	}
	return c.memory[c.memoryPointer+offset]
}

func (c Computer) GetRegisterVal(registerIndex int) uint16 {
	return *c.registers[registerIndex]
}

func (c Computer) MapRegister(registerVal int) int {
	register := registerVal % 32768
	return register
}

func (c Computer) SetRegisterVal(register int, val uint16) {
	*c.registers[register] = val
}

func (c Computer) Run() {
	//Grab OpCode
	//Process it
	//Update Memory Pointer
	for {
		opCode := parser.Parse(c.GetVar())
		switch opCode {
		case "halt":
			return
		case "set":
			a := c.memory[c.memoryPointer+1]
			register := c.MapRegister(int(a))
			b := c.GetVarOffset(2)
			c.SetRegisterVal(register, b)
			c.memoryPointer += 3
			break
		case "out":
			out := c.GetVarOffset(1)
			fmt.Print(string(out))
			c.memoryPointer += 2
			break
		case "noop":
			c.memoryPointer++
			break
		case "jmp":
			c.memoryPointer = int(c.GetVarOffset(1))
			break
		case "jt":
			a := c.GetVarOffset(1)
			b := int(c.GetVarOffset(2))
			if a != 0 {
				c.memoryPointer = b
				break
			}
			c.memoryPointer += 3
			break
		case "jf":
			a := c.GetVarOffset(1)
			b := int(c.GetVarOffset(2))
			if a == 0 {
				c.memoryPointer = b
				break
			}
			c.memoryPointer += 3
			break
		default:
			fmt.Print(opCode)
			fmt.Println(" - NOT IMPLEMENTED YET")
			c.memoryPointer++
			break
		}
	}
}

func NewComputer(program []uint16) *Computer {

	var registers [8]*uint16
	for i := 0; i < 8; i++ {
		registers[i] = new(uint16)
	}
	return &Computer{
		memory:        program,
		registers:     registers,
		stack:         []uint16{},
		memoryPointer: 0,
	}
}
