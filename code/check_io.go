package main

import (
	"fmt"

	"os"

	"strconv"
	"time"

	"github.com/shirou/gopsutil/disk"
)

const MSG = "[%s]_IOStatus: %s - read(KB)/s=%d, write(KB)/s=%d | read=%d;%d;0;0; write=%d;%d;0;0;\n"

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
	rs := int((v2[d].ReadBytes - v[d].ReadBytes)) >> 10
	ws := int((v2[d].WriteBytes - v[d].WriteBytes)) >> 10
	var sta string
	var exit_code int
	if (rs > rw && rs < rc) || (ws > ww && ws < wc) {
		sta = "WARNING"
		exit_code = 1
	} else if rs > rc || ws > wc {
		sta = "Critical"
		exit_code = 2
	} else {
		sta = "OK"
		exit_code = 0
	}
	fmt.Printf(MSG, d, sta, rs, ws, rw, rc, ww, wc)
	os.Exit(exit_code)

}
