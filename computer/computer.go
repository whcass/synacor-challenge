package computer

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/whcass/synacor-challenge/parser"
)

const REGISTER_START int = 32768

type Computer struct {
	memory        []uint16
	registers     [8]uint16
	stack         []uint16
	memoryPointer int
	stdin         *bufio.Reader
	logger        *log.Logger
	gameStart     bool
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

func (c Computer) GetRegisterIndex(offset int) int {
	a := c.memory[c.memoryPointer+offset]
	register := c.MapRegister(int(a))
	return register
}

func (c Computer) GetRegisterVal(registerIndex int) uint16 {
	return c.registers[registerIndex]
}

func (c Computer) MapRegister(registerVal int) int {
	register := registerVal % 32768
	return register
}

func (c *Computer) SetRegisterVal(register int, val uint16) {
	c.registers[register] = val
}

func (c *Computer) Push(a uint16) {
	c.stack = append(c.stack, a)
}

func (c *Computer) Pop(register int) {
	val := c.stack[len(c.stack)-1]
	c.stack = c.stack[:len(c.stack)-1]
	c.SetRegisterVal(register, val)
}

func (c Computer) Run() {
	//Grab OpCode
	//Process it
	//Update Memory Pointer

	//game starts at 1798
	for {
		if c.gameStart {
			c.logger.Println("BEGIN REGISTER")
			for _, register := range c.registers {
				c.logger.Println(register)
			}
			c.logger.Println("END REGISTER")
			c.logger.Printf("memoryPointer: %v\r", c.memoryPointer)
			c.logger.Println()
			if c.GetRegisterVal(7) == 0 {
				c.SetRegisterVal(7, 1)
			}
		}
		opCode := parser.Parse(c.GetVar())
		switch opCode {
		case "halt":
			return
		case "set":
			//a := c.memory[c.memoryPointer+1]
			register := c.GetRegisterIndex(1)
			b := c.GetVarOffset(2)
			c.SetRegisterVal(register, b)
			c.memoryPointer += 3
			break
		case "push":
			a := c.GetVarOffset(1)
			c.Push(a)
			c.memoryPointer += 2
			break
		case "pop":
			register := c.GetRegisterIndex(1)
			c.Pop(register)
			c.memoryPointer += 2
			break
		case "eq":
			register := c.GetRegisterIndex(1)
			a := c.GetVarOffset(2)
			b := c.GetVarOffset(3)
			if a == b {
				c.SetRegisterVal(register, 1)
			} else {
				c.SetRegisterVal(register, 0)
			}
			c.memoryPointer += 4
			break
		case "gt":
			register := c.GetRegisterIndex(1)
			a := c.GetVarOffset(2)
			b := c.GetVarOffset(3)
			if a > b {
				c.SetRegisterVal(register, 1)
			} else {
				c.SetRegisterVal(register, 0)
			}
			c.memoryPointer += 4
			break
		case "jmp":
			c.memoryPointer = int(c.GetVarOffset(1))
			break
		case "jt":
			a := c.GetVarOffset(1)
			b := int(c.GetVarOffset(2))
			if a != 0 {
				c.memoryPointer = b
			} else {
				c.memoryPointer += 3
			}
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
		case "add":
			register := c.GetRegisterIndex(1)
			a := c.GetVarOffset(2)
			b := c.GetVarOffset(3)

			ans := int(a+b) % REGISTER_START
			c.SetRegisterVal(register, uint16(ans))
			c.memoryPointer += 4
			break
		case "mult":
			register := c.GetRegisterIndex(1)
			a := c.GetVarOffset(2)
			b := c.GetVarOffset(3)

			ans := int(a*b) % REGISTER_START
			c.SetRegisterVal(register, uint16(ans))
			c.memoryPointer += 4
			break
		case "mod":
			register := c.GetRegisterIndex(1)
			a := c.GetVarOffset(2)
			b := c.GetVarOffset(3)
			ans := int(a%b) % REGISTER_START
			c.SetRegisterVal(register, uint16(ans))
			c.memoryPointer += 4
			break
		case "and":
			register := c.GetRegisterIndex(1)
			a := c.GetVarOffset(2)
			b := c.GetVarOffset(3)

			result := int(a&b) % REGISTER_START
			c.SetRegisterVal(register, uint16(result))
			c.memoryPointer += 4
			break
		case "or":
			register := c.GetRegisterIndex(1)
			a := c.GetVarOffset(2)
			b := c.GetVarOffset(3)
			result := int(a|b) % REGISTER_START
			c.SetRegisterVal(register, uint16(result))
			c.memoryPointer += 4
			break
		case "not":
			register := c.GetRegisterIndex(1)
			a := c.GetVarOffset(2)
			result := int(^a) % REGISTER_START
			c.SetRegisterVal(register, uint16(result))
			c.memoryPointer += 3
			break
		case "rmem":

			register := c.GetRegisterIndex(1)
			a := c.GetVarOffset(2)
			c.SetRegisterVal(register, c.memory[a])
			c.memoryPointer += 3
			break
		case "wmem":
			index := c.GetVarOffset(1)
			val := c.GetVarOffset(2)

			c.memory[index] = val
			c.memoryPointer += 3
			break
		case "call":
			a := c.GetVarOffset(1)
			ret := c.memoryPointer + 2
			c.Push(uint16(ret))
			c.memoryPointer = int(a)
			break
		case "ret":
			val := c.stack[len(c.stack)-1]
			c.stack = c.stack[:len(c.stack)-1]
			c.memoryPointer = int(val)
			break
		case "out":
			out := c.GetVarOffset(1)
			fmt.Print(string(out))
			c.memoryPointer += 2
			break
		case "noop":
			c.memoryPointer++
			break
		case "in":
			c.gameStart = true
			in, err := c.stdin.ReadByte()
			if err != nil {
				panic(err)
			}
			register := c.GetRegisterIndex(1)
			c.SetRegisterVal(register, uint16(in))
			c.memoryPointer += 2
			break
		default:
			fmt.Print("UNEXPECTED OP CODE AT - ")
			fmt.Println(c.memoryPointer)
			os.Exit(1)
		}
	}
}

func NewComputer(program []uint16) *Computer {

	//var registers [8]*uint16
	//for i := 0; i < 8; i++ {
	//	registers[i] = new(uint16)
	//}
	f, err := os.OpenFile("out.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	logger := log.New(f, "[*] ", log.LstdFlags)
	return &Computer{
		memory:        program,
		registers:     [8]uint16{},
		stack:         []uint16{},
		memoryPointer: 0,
		stdin:         bufio.NewReader(os.Stdin),
		logger:        logger,
		gameStart:     false,
	}
}
