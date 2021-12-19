package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
)

//go:embed input.txt
var puzzleInput string

func main() {
	fmt.Println(task1(puzzleInput))
	fmt.Println(task2(puzzleInput))
}

func task1(input string) string {
	bits := NewBitMachine(strings.TrimSpace(input))
	pkg := parsePacket(bits)
	return fmt.Sprint(pkg.VersionCount())
}

func task2(input string) string {
	bits := NewBitMachine(strings.TrimSpace(input))
	pkg := parsePacket(bits)
	return fmt.Sprint(pkg.Calculate())
}

// BitMachine parses and holdes the input and returns the bit stream.
type BitMachine struct {
	input  []byte
	buffer []bool
}

// NewBitMachine initializes the BitMachine.
func NewBitMachine(input string) *BitMachine {
	return &BitMachine{
		input: []byte(input),
	}
}

// NextBits returns the next `count` bits from the BitMachine as uint64.
//
// count has to be a number between 1 and 64.
func (b *BitMachine) NextBits(count int) uint64 {
	var out uint64
	for i := 0; i < count; i++ {
		out <<= 1

		if len(b.buffer) == 0 {
			b.buffer = letterToBuffer(b.input[0])
			b.input = b.input[1:]
		}

		if b.buffer[0] {
			out |= 1
		}
		b.buffer = b.buffer[1:]
	}
	return out
}

// subBits implements the NextBitser interface for sub packages.
type subBits struct {
	parent NextBitser
	len    int
}

func (b *subBits) NextBits(count int) uint64 {
	out := b.parent.NextBits(count)
	b.len -= count
	return out
}

// NextBitser returns the next bit from some kind of source.
//
// It is implemented by the BitMachine and for subBits.
type NextBitser interface {
	NextBits(count int) uint64
}

// parseLiteral uses a NextBitser to parse a literal value.
func parseLiteral(bits NextBitser) uint64 {
	var out uint64
	for {
		out <<= 4
		prefix := bits.NextBits(1)
		number := bits.NextBits(4)
		out |= number
		if prefix == 0 {
			break
		}
	}
	return out
}

// Packet is a day16 packat
//
// Either value or sub is set
type Packet struct {
	version int
	typeID  int
	value   uint64
	sub     []Packet
}

// parsePacket uses a NextBitser to parse one package with all sub packages.
func parsePacket(bits NextBitser) Packet {
	var p Packet
	p.version = int(bits.NextBits(3))
	p.typeID = int(bits.NextBits(3))

	if p.typeID == 4 {
		p.value = parseLiteral(bits)
		return p
	}

	lengthType := bits.NextBits(1)
	if lengthType == 1 {
		subPackageCount := int(bits.NextBits(11))
		for i := 0; i < subPackageCount; i++ {
			p.sub = append(p.sub, parsePacket(bits))
		}
		return p
	}

	length := int(bits.NextBits(15))

	sbits := &subBits{
		parent: bits,
		len:    length,
	}

	for sbits.len > 0 {
		p.sub = append(p.sub, parsePacket(sbits))
	}
	return p
}

// VersionCount returns the sum of the version of the packet and all sub
// packages.
//
// Used for task1.
func (p Packet) VersionCount() int {
	s := p.version
	for _, subPkg := range p.sub {
		s += subPkg.VersionCount()
	}
	return s
}

// Calculate executes all packages for task2.
func (p Packet) Calculate() int {
	var operate func(a, b int) int
	var v int
	switch p.typeID {
	case 0:
		operate = func(a, b int) int {
			return a + b
		}

	case 1:
		v = 1
		operate = func(a, b int) int {
			return a * b
		}

	case 2:
		v = math.MaxInt
		operate = func(a, b int) int {
			if a > b {
				return b
			}
			return a
		}
	case 3:
		operate = func(a, b int) int {
			if a < b {
				return b
			}
			return a
		}

	case 4:
		return int(p.value)

	case 5:
		if p.sub[0].Calculate() > p.sub[1].Calculate() {
			return 1
		}
		return 0

	case 6:
		if p.sub[0].Calculate() < p.sub[1].Calculate() {
			return 1
		}
		return 0

	case 7:
		if p.sub[0].Calculate() == p.sub[1].Calculate() {
			return 1
		}
		return 0
	}

	for _, sub := range p.sub {
		v = operate(v, sub.Calculate())
	}
	return v
}

// letterToBuffer parses one hex-byte to a four bit bool-slice.
//
// It uses a table to parse the 16 values.
func letterToBuffer(l byte) []bool {
	switch l {
	case '0':
		return []bool{false, false, false, false}
	case '1':
		return []bool{false, false, false, true}
	case '2':
		return []bool{false, false, true, false}
	case '3':
		return []bool{false, false, true, true}
	case '4':
		return []bool{false, true, false, false}
	case '5':
		return []bool{false, true, false, true}
	case '6':
		return []bool{false, true, true, false}
	case '7':
		return []bool{false, true, true, true}
	case '8':
		return []bool{true, false, false, false}
	case '9':
		return []bool{true, false, false, true}
	case 'A':
		return []bool{true, false, true, false}
	case 'B':
		return []bool{true, false, true, true}
	case 'C':
		return []bool{true, true, false, false}
	case 'D':
		return []bool{true, true, false, true}
	case 'E':
		return []bool{true, true, true, false}
	case 'F':
		return []bool{true, true, true, true}
	}
	return nil
}
