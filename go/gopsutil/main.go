package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

func printMem() {
	vm, _ := mem.VirtualMemory()
	fmt.Println("=== mem ===")
	fmt.Println(vm)
	fmt.Printf("mem.Total: %v	mem.Free: %v	UsedPer: %v\n", vm.Total, vm.Free, vm.UsedPercent)

	fmt.Println()
}

func printCPU() {
	fmt.Println("=== cpu ===")
	per, _ := cpu.Percent(50*time.Millisecond, true)
	fmt.Println(per)
	// utl, _ := cpu.Info()
	// fmt.Println(utl)

	fmt.Println()
}

func printDisk() {
	fmt.Println("=== disk ===")
	root, _ := disk.Usage("/")
	vr, _ := disk.Usage("/var")
	home, _ := disk.Usage("/home")
	fmt.Println(root)
	fmt.Println(vr)
	fmt.Println(home)
	fmt.Println()
}

func main() {
	printMem()
	printCPU()
	printDisk()
}
