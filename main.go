package main

import (
	sch "cogocoroutines/scheduler"
	soc "cogocoroutines/socket"
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println("failed to open listener", err.Error())
		return
	}

	fmt.Println("listening on port 8080")

	scheduler := sch.Scheduler{}
	scheduler.AddTask(soc.ConnectionListener(&scheduler, listener))
	scheduler.Run()
}
