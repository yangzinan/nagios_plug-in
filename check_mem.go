package main

import (
	"fmt"
	"strconv"

	"os"

	"github.com/shirou/gopsutil/mem"
)

func main() {
	argv := os.Args
	if len(argv) != 3 {
		fmt.Println("Please InPut Args")
		os.Exit(0)
	}
	w, _ := strconv.ParseFloat(argv[1], 64)
	c, _ := strconv.ParseFloat(argv[2], 64)
	v, _ := mem.VirtualMemory()
	total := v.Total >> 30
	free := v.Free >> 30
	if v.UsedPercent > w && v.UsedPercent < c {
		fmt.Printf("Total:%vG, Free:%vG, UsedPercent:%f%% | UsedPercent=%f;%f;0;100", total, free, v.UsedPercent, w, c)
		os.Exit(1)
	} else if v.UsedPercent > c {
		fmt.Printf("Total:%vG, Free:%vG, UsedPercent:%f%% | UsedPercent=%f;%f;0;100", total, free, v.UsedPercent, w, c)
		os.Exit(2)
	} else {
		fmt.Printf("Total:%vG, Free:%vG, UsedPercent:%f%% | UsedPercent=%f;%f;0;100", total, free, v.UsedPercent, w, c)
		os.Exit(0)
	}
}
