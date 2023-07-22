package socket

import (
	"bufio"
	"cogocoroutines/functions"
	s "cogocoroutines/scheduler"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

func HandleConnection(sch *s.Scheduler, conn net.Conn) func() bool {
	return func() bool {
		req, readReqErr := http.ReadRequest(bufio.NewReader(conn))

		if readReqErr != nil {
			closeErr := conn.Close()
			if closeErr != nil {
				return true
			}
			return true
		}

		sleepDurationStr := path.Base(req.URL.Path)
		sleepDuration, parseIntErr := strconv.Atoi(sleepDurationStr)

		if parseIntErr != nil {
			fmt.Println("invalid sleep duration:", sleepDurationStr)

			closeErr := conn.Close()
			if closeErr != nil {
				return true
			}

			return true
		}

		fmt.Println("received request for", sleepDurationStr, "seconds")

		sch.AddTask(s.NewTask(functions.Sleep(sleepDuration), func() {
			_, writeErr := conn.Write([]byte("HTTP/1.1 200 OK\n\nHandled request for " + sleepDurationStr + " seconds"))

			if writeErr != nil {
				fmt.Println("failed to write response:", writeErr.Error())
				return
			}

			closeErr := conn.Close()

			if closeErr != nil {
				fmt.Println("failed to close connection:", closeErr.Error())
				return
			}

			fmt.Println("handled request for", sleepDurationStr, "seconds")
		}))

		return true
	}
}

func ConnectionListener(sch *s.Scheduler, listener net.Listener) s.Task {
	return s.NewTask(func() bool {
		deadlineErr := listener.(*net.TCPListener).SetDeadline(time.Now().Add(time.Millisecond))

		if deadlineErr != nil {
			fmt.Println("failed to set deadline:", deadlineErr.Error())
			return true
		}

		conn, err := listener.Accept()

		if errors.Is(err, os.ErrDeadlineExceeded) {
			return false
		} else if err != nil {
			fmt.Println("failed to accept connection:", err.Error())
			return true
		}

		sch.AddTask(s.NewTask(HandleConnection(sch, conn)))

		return false
	})
}
