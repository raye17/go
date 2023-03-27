package main

type Computer struct {
	CPU      string
	Memory   string
	HardDisk string
}

func (c *Computer) SetCPU(cpu string) {
	c.CPU = cpu
}
func (c *Computer) GetCPU() string {
	return c.CPU
}

func (c *Computer) SetMemory(memory string) {
	c.Memory = memory
}

func (c *Computer) GetMemory() string {
	return c.Memory
}

func (c *Computer) SetHardDisk(hardDisk string) {
	c.HardDisk = hardDisk
}

func (c *Computer) GetHardDisk() string {
	return c.HardDisk
}

type Builder interface {
	SetCPU(cpu string) Builder
	SetMemory(memory string) Builder
	SetHardDisk(hardDisk string) Builder
	Build() *Computer
}
type ComputerBuilder struct {
	computer *Computer
}

func (c *ComputerBuilder) SetCPU(cpu string) Builder {
	if c.computer == nil {
		c.computer = new(Computer)
	}
	c.computer.SetCPU(cpu)
	return c
}
func (c *ComputerBuilder) SetMemory(memory string) Builder {
	if c.computer == nil {
		c.computer = new(Computer)
	}
	c.computer.SetMemory(memory)
	return c
}

func (c *ComputerBuilder) SetHardDisk(hardDisk string) Builder {
	if c.computer == nil {
		c.computer = new(Computer)
	}
	c.computer.SetHardDisk(hardDisk)
	return c
}

func (c *ComputerBuilder) Build() *Computer {
	return c.computer
}

type Director struct {
	Builder Builder
}

func (d Director) Create(cpu, memory, hardDisk string) *Computer {
	return d.Builder.SetHardDisk(hardDisk).SetCPU(cpu).SetMemory(memory).Build()
}

func main() {

}
