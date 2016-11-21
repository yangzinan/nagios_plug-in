package main

import "github.com/shirou/gopsutil/load"
import "os"
import "fmt"
import "strconv"

const MSG = "Load-Status: %s - load1:%f, load5:%f, load15:%f | load1=%f;%f;0;0; load5=%f;%f;0;0; load15=%f;%f;0;0; \n"

func main() {
	argv := os.Args
	if len(argv) != 7 {
		fmt.Println("Please InPut Args")
		os.Exit(0)
	}
	w1, _ := strconv.ParseFloat(argv[1], 64)
	c1, _ := strconv.ParseFloat(argv[2], 64)
	w5, _ := strconv.ParseFloat(argv[3], 64)
	c5, _ := strconv.ParseFloat(argv[4], 64)
	w15, _ := strconv.ParseFloat(argv[5], 64)
	c15, _ := strconv.ParseFloat(argv[6], 64)
	v, _ := load.Avg()
	var sta string
	var exit_code int
	if (v.Load1 > w1 && v.Load1 < c1) || (v.Load5 > w5 && v.Load5 < c5) || (v.Load15 > w15 && v.Load15 < c15) {
		sta = "WARNING"
		exit_code = 1
	} else if v.Load1 > c1 || v.Load5 > c5 || v.Load15 > c15 {
		sta = "Critical"
		exit_code = 2
	} else {
		sta = "OK"
		exit_code = 0
	}
	fmt.Printf(MSG, sta, v.Load1, v.Load5, v.Load15, w1, c1, w5, c5, w15, c15)
	os.Exit(exit_code)
}
