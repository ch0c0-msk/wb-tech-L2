package pattern

import "fmt"

// Facade
type Computer struct {
	c *CPU
	m *Memory
}

func NewComputer(freq, aval, mem int) *Computer {
	return &Computer{c: &CPU{freqMHz: freq}, m: &Memory{volumeGb: mem, availableGb: aval}}
}

func (c *Computer) Start() {
	fmt.Println("Computer is starting...")
	c.c.start()
	c.m.start()
	fmt.Println("Computer is ready to work.")
}

// Facade parts
type CPU struct {
	freqMHz int
}

func (c *CPU) start() {
	fmt.Printf("Freq is %d MHZ. Cpu is starting...\n", c.freqMHz)
}

type Memory struct {
	volumeGb    int
	availableGb int
}

func (m *Memory) start() {
	fmt.Printf("Total memory: %d. Available memory: %d.\n", m.availableGb, m.volumeGb)
}
