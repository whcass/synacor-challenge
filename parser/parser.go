package parser

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
	if int(opcode) > len(opcodes) {
		return "UNKNOWN"
	}
	return opcodes[opcode]
}
