package main

import (
	"encoding/binary"
	"io"
	"os"

	"github.com/whcass/synacor-challenge/computer"
	"github.com/whcass/synacor-challenge/parser"
)

func main() {
	f, err := os.Open("challenges/challenge.bin")
	if err != nil {
		panic(err)
	}
	//fmt.Println(file)
	//temp := strings.Fields(string(file))

	var program []uint16
	for {
		var low, high byte
		var bytes []byte

		err = binary.Read(f, binary.LittleEndian, &low)
		if err == nil {
			err = binary.Read(f, binary.LittleEndian, &high)
		}

		bytes = append(bytes, low)
		bytes = append(bytes, high)

		result := binary.LittleEndian.Uint16(bytes)
		program = append(program, result)

		if err == io.EOF {
			break
		}
	}
	parser.ParseMemory(program)
	vm := computer.NewComputer(program)
	vm.Run()
}
