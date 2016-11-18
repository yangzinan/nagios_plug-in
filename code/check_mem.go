package main

import (
	"fmt"     
	"strconv"  //引入字符串转floadt64库
	"os"
	"github.com/shirou/gopsutil/mem" //引入内存操作库
)

const MSG = "Memory_Status: %s - Total:%vM, Free:%vM, UsedPercent:%f%% | UsedPercent=%f;%f;0;100\n" //定义输出常量

func main() {
	argv := os.Args
	if len(argv) != 3 {
		fmt.Println("Please InPut Args")
		os.Exit(0)
	}
	
	w, _ := strconv.ParseFloat(argv[1], 64) //获取参数1为内存使用率的警告阀值
	c, _ := strconv.ParseFloat(argv[2], 64) //获取参数2为内存使用率的灾难阀值
	v, _ := mem.VirtualMemory()  //获取内存信息
	total := v.Total >> 20  //获取内存总量（MB）
	free := v.Free >> 20    //获取剩余内存（MB）
	var sta string      //定义状态
	var exit_code int   //定义退出码

	if v.UsedPercent > w && v.UsedPercent < c {
		sta = "WARNING"
		exit_code = 1
	} else if v.UsedPercent > c {
		sta = "Critical"
		exit_code = 2
	} else {
		sta = "OK"
		exit_code = 0
	}
	fmt.Printf(MSG, sta, total, free, v.UsedPercent, w, c)
	os.Exit(exit_code)
}
