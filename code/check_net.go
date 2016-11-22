package main

import (
	"fmt"

	"time"

	"os"
	"strconv"

	nnet "github.com/shirou/gopsutil/net"
)

const MSG = "[%s]_Status: %s - RecvBytes/s=%dKB/s, SendBytes/s=%dKB/s |　RecvBytes=%d;%d;0;0; SendBytes=%d;%d;0;0;"

func main() {
	argv := os.Args
	if len(argv) != 6 {
		fmt.Println("Please InPut Args")
		os.Exit(0)
	}
	net_car := argv[1]
	rw, _ := strconv.Atoi(argv[2])
	rc, _ := strconv.Atoi(argv[3])
	sw, _ := strconv.Atoi(argv[4])
	sc, _ := strconv.Atoi(argv[5])
	v, _ := nnet.IOCounters(true)
	time.Sleep(1000000000)
	v2, _ := nnet.IOCounters(true)
	var num int
	for i := 0; i < len(v); i++ {
		if v[i].Name == net_car {
			num = i
			break
		}
	}
	r := int((v2[num].BytesRecv - v[num].BytesRecv) >> 10)
	s := int((v2[num].BytesSent - v[num].BytesSent) >> 10)
	var sta string
	var exit_code int
	if (r > rw && r < rc) || (s > sw && s < sc) {
		sta = "WARNING"
		exit_code = 1
	} else if r > rc || s > sc {
		sta = "Critical"
		exit_code = 2
	} else {
		sta = "OK"
		exit_code = 0
	}
	fmt.Printf(MSG, net_car, sta, r, s, rw, rc, sw, sc)
	os.Exit(exit_code)
}
