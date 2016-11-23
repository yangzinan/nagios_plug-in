package main

import (
	"fmt"

	"os"

	"strconv"

	"github.com/shirou/gopsutil/cpu"
)

const MSG = "CPU_status: %s - UsedPercent=%v%% | UsedPercent=%f;%f;0;100;\n"

func main() {
	argv := os.Args
	if len(argv) != 3 {
		fmt.Println("Please InPut Args")
		os.Exit(0)
	}

	w, _ := strconv.ParseFloat(argv[1], 64)
	c, _ := strconv.ParseFloat(argv[2], 64)
	v, _ := cpu.Percent(1000000000, true)
	var sta string
	var exit_code int
	if v[0] > w && v[0] < c {
		sta = "WARNING"
		exit_code = 1
	} else if v[0] > c {
		sta = "Critical"
		exit_code = 2
	} else {
		sta = "OK"
		exit_code = 0
	}
	fmt.Printf(MSG, sta, v[0], w, c)
	os.Exit(exit_code)
}
