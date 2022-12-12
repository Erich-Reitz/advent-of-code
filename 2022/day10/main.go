package main

import (
	"fmt"
	"github.com/Erich-Reitz/commonGo"
	"strconv"
	"strings"
)

type CPU_INSTRUCTION int

const (
	NOOP    CPU_INSTRUCTION = iota
	ADDX    CPU_INSTRUCTION = iota
	UNKNOWN CPU_INSTRUCTION = iota
)

type PendingCpuOperation struct {
	operation        CPU_OPERATION
	cyclesToComplete int
	register         int
}

type CPU_OPERATION struct {
	instruction CPU_INSTRUCTION
	value       int
}

type Register struct {
	Value int
}

type CPU struct {
	Registers         []Register
	PendingOperations []PendingCpuOperation
}

func convertLineToCpuOperation(line string) CPU_OPERATION {
	if line == "noop" {
		return CPU_OPERATION{NOOP, 0}
	} else if strings.HasPrefix(line, "addx ") {
		valueAsString := strings.Split(line, " ")[1]
		valueAsInt, err := strconv.Atoi(valueAsString)
		if err != nil {
			panic(err)
		}
		return CPU_OPERATION{ADDX, valueAsInt}
	}
	return CPU_OPERATION{UNKNOWN, 0}
}

func (cpu *CPU) addOperation(operation CPU_OPERATION) {
	var pendingOperation PendingCpuOperation
	if operation.instruction == NOOP {
		pendingOperation = PendingCpuOperation{operation: operation, cyclesToComplete: 1, register: 0}
	} else if operation.instruction == ADDX {
		pendingOperation = PendingCpuOperation{operation: operation, cyclesToComplete: 2, register: 0}
	}
	cpu.PendingOperations = append(cpu.PendingOperations, pendingOperation)
}

func (cpu *CPU) initalizeCpu(registerCount uint) {
	cpu.Registers = make([]Register, registerCount)
	for i := 0; i < int(registerCount); i++ {
		cpu.Registers[i] = Register{1}
	}
}

func (cpu *CPU) pendingOperations() []PendingCpuOperation {
	return cpu.PendingOperations
}
func (cpu *CPU) currentOperation() *PendingCpuOperation {
	return &cpu.PendingOperations[0]
}

func (cpu *CPU) addValueToRegister(register, value int) {
	cpu.Registers[register].Value += value
}

func (cpu *CPU) performOperation() {
	currentOperation := cpu.currentOperation()
	currentOperation.cyclesToComplete -= 1
	if currentOperation.cyclesToComplete == 0 {
		cpu.addValueToRegister(currentOperation.register, currentOperation.operation.value)
		cpu.PendingOperations = cpu.PendingOperations[1:]
	}
}

func (cpu *CPU) CalculateSignalStrength(cpuCycle int) int {
	return cpuCycle * cpu.Registers[0].Value
}

func part1(contents string) {
	lines := strings.Split(contents, "\n")
	cpu := CPU{}
	cpu.initalizeCpu(1)
	operations := make([]CPU_OPERATION, 0)
	for _, line := range lines {
		operations = append(operations, convertLineToCpuOperation(line))
	}

	cpuCycles := 1
	sumOfSignalStrengthCalculations := 0
	for {
		if cpuCycles == 20 || (cpuCycles-20)%40 == 0 {
			cpuStrength := cpu.CalculateSignalStrength(cpuCycles)
			sumOfSignalStrengthCalculations += cpuStrength
		}

		if len(operations) == 0 {
			break
		}
		if len(cpu.pendingOperations()) == 0 {
			var op CPU_OPERATION
			op, operations = operations[0], operations[1:]
			cpu.addOperation(op)
		}
		cpu.performOperation()
		cpuCycles += 1
	}
	fmt.Println(sumOfSignalStrengthCalculations)
}

func draw(crtScreen [][]string, crtGunRow, crtGunCol, spriteMiddlePosition int) {
	if advent.AbsDiffInt(crtGunCol, spriteMiddlePosition) <= 1 {
		crtScreen[crtGunRow][crtGunCol] = "\u2588"
	} else {
		crtScreen[crtGunRow][crtGunCol] = " "
	}
}

func part2(contents string) {
	lines := strings.Split(contents, "\n")
	cpu := CPU{}
	cpu.initalizeCpu(1)
	operations := make([]CPU_OPERATION, 0)
	for _, line := range lines {
		operations = append(operations, convertLineToCpuOperation(line))
	}

	cpuCycles := 1
	crtScreen := make([][]string, 6)
	crtGunPosition := 0
	for i := range crtScreen {
		crtScreen[i] = make([]string, 40)
	}
	for {
		if len(operations) == 0 {
			break
		}
		crtGunRow := crtGunPosition / 40 % 6
		crtGunCol := crtGunPosition % 40
		fmt.Println(crtGunPosition, crtGunRow, crtGunCol)
		draw(crtScreen, crtGunRow, crtGunCol, cpu.Registers[0].Value)

		if len(cpu.pendingOperations()) == 0 {
			var op CPU_OPERATION
			op, operations = operations[0], operations[1:]
			cpu.addOperation(op)
		}
		cpu.performOperation()
		cpuCycles += 1
		crtGunPosition += 1
	}
	for _, line := range crtScreen {
		fmt.Println(line)
	}
}

func main() {
	data := advent.ReadFileAsString("input.txt")
	// part1(data)
	part2(data)
}
