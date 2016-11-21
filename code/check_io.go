package main

import (
	"fmt"

	"os"

	"strconv"
	"time"

	"github.com/shirou/gopsutil/disk"
)

func main() {
	argv := os.Args
	if len(argv) != 6 {
		fmt.Println("Please InPut Args")
		os.Exit(0)
	}
	d := argv[1]
	rw, _ := strconv.Atoi(argv[2])
	rc, _ := strconv.Atoi(argv[3])
	ww, _ := strconv.Atoi(argv[4])
	wc, _ := strconv.Atoi(argv[5])
	v, _ := disk.IOCounters()
	time.Sleep(1000)
	v2, _ := disk.IOCounters()
	rs := (v2[d].ReadBytes - v[d].ReadBytes) >> 10
	ws := (v2[d].WriteBytes - v[d].WriteBytes) >> 10
	var sta string
	var exit_code int
	if (rs > rw && rs < rc) || (ws > ww && ws < wc) {
		sta = "w"
		exit_code = 1
	}

}
