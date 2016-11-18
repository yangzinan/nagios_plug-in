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
		fmt.Printf("Total:%vG, Free:%vG, UsedPercent:%f%% | Free=%d;%d;0;%v", total, free, v.UsedPercent, w, c, total)
		os.Exit(1)
	} else if v.UsedPercent > c {
		fmt.Printf("Total:%vG, Free:%vG, UsedPercent:%f%% | Free=%f;%f;0;%v", total, free, v.UsedPercent, w, c, total)
		os.Exit(2)
	} else {
		fmt.Printf("Total:%vG, Free:%vG, UsedPercent:%f%% | Free=%f;%f;0;%v", total, free, v.UsedPercent, w, c, total)
		os.Exit(0)
	}
}
