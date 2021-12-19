package main

import "testing"

func TestTask1(t *testing.T) {
	for _, tt := range []struct {
		input  string
		expect string
	}{
		{"8A004A801A8002F478", "16"},
		{"620080001611562C8802118E34", "12"},
		{"C0015000016115A2E0802F182340", "23"},
		{"A0016C880162017C3686B18A3D4780", "31"},
	} {
		t.Run(tt.input, func(t *testing.T) {
			got := task1(tt.input)
			if got != tt.expect {
				t.Errorf("task1(%q) == %q, expected %q", tt.input, got, tt.expect)
			}
		})
	}
}

func TestTask2(t *testing.T) {
	for _, tt := range []struct {
		input  string
		expect string
	}{
		{"C200B40A82", "3"},
		{"04005AC33890", "54"},
		{"880086C3E88112", "7"},
		{"CE00C43D881120", "9"},
		{"D8005AC2A8F0", "1"},
		{"F600BC2D8F", "0"},
		{"9C005AC2F8F0", "0"},
		{"9C0141080250320F1802104A08", "1"},
	} {
		t.Run(tt.input, func(t *testing.T) {
			got := task2(tt.input)
			if got != tt.expect {
				t.Errorf("task2(%q) == %q, expected %q", tt.input, got, tt.expect)
			}
		})
	}
}

func TestParseLiteral(t *testing.T) {
	bits := &mockNextBitser{[]byte("101111111000101000")}

	got := parseLiteral(bits)
	if got != 2021 {
		t.Errorf("parseLiteral('101111111000101000') == %d, expected 2021", got)
	}
}

type mockNextBitser struct {
	bits []byte
}

func (b *mockNextBitser) NextBits(count int) uint64 {
	var out uint64
	for i := 0; i < count; i++ {
		out <<= 1
		if b.bits[0] == '1' {
			out |= 1
		}
		b.bits = b.bits[1:]
	}
	return out
}
