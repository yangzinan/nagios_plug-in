package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/shirou/gopsutil/disk"
)

const MSG = "[%s]_Status: %s - Total:%vG, Free:%vG, UsedPercent:%f%% | UsedPercent=%f;%f;0;100\n"

func main() {
	argv := os.Args
	if len(argv) != 4 {
		fmt.Println("Please InPut Args")
		os.Exit(0)
	}
	d := argv[1]
	v, _ := disk.Usage(d)
	total := v.Total >> 30
	free := v.Free >> 30
	UsedPercent := v.UsedPercent
	w, _ := strconv.ParseFloat(argv[2], 64)
	c, _ := strconv.ParseFloat(argv[3], 64)
	var sta string
	var exit_code int
	if UsedPercent > w && UsedPercent < c {
		sta = "WARNING"
		exit_code = 1
	} else if UsedPercent > c {
		sta = "Critical"
		exit_code = 2
	} else {
		sta = "OK"
		exit_code = 0
	}
	fmt.Printf(MSG, d, sta, total, free, UsedPercent, w, c)
	os.Exit(exit_code)
}
