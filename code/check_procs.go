package main

import "github.com/shirou/gopsutil/load"
import "os"
import "fmt"
import "strconv"

const MSG = "Procs-Status: %s - ProcsRunning:%d, ProcsBlocked:%d | procsRunning=%d;%d;0;0 ProcsBlocked=%d;%d;0;0\n"

func main() {
	argv := os.Args
	if len(argv) != 5 {
		fmt.Println("Please InPut Args")
		os.Exit(0)
	}
	rw, _ := strconv.Atoi(argv[1])
	rc, _ := strconv.Atoi(argv[2])
	bw, _ := strconv.Atoi(argv[3])
	bc, _ := strconv.Atoi(argv[4])
	v, _ := load.Misc()
	procsRunning := v.ProcsRunning
	ProcsBlocked := v.ProcsBlocked
	var sta string
	var exit_code int
	if (procsRunning > rw && procsRunning < rc) || (ProcsBlocked > bw && ProcsBlocked < bc) {
		sta = "WARNING"
		exit_code = 1
	} else if procsRunning > rc || ProcsBlocked > bc {
		sta = "Critical"
		exit_code = 2
	} else {
		sta = "OK"
		exit_code = 0
	}
	fmt.Printf(MSG, sta, procsRunning, ProcsBlocked, rw, rc, bw, bc)
	os.Exit(exit_code)
}
