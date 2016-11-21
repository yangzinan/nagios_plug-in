# nagios_plug-in
## check_mem
1. 使用方法： check_mem 80 90
2. 参数说明： 第一个参数是警告阀值，第二个参数是灾难阀值

## chech_disk
1. 使用方法： check_disk / 80 90
2. 参数说明： 参数1是要监控的分区挂载点，参数2是警告阀值，参数3是灾难阀值

## check_procs
1. 使用方法： check_procs 100 200 10 30
2. 参数说明   参数1，2是ProcsRunning（正在运行的任务数的警告个灾难的阀值）参数3，4是ProcsBlocked（当前被阻塞的任务数的警告和灾难的阀值）