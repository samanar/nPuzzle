package utils

import (
	"fmt"
	"strings"
	
	"github.com/gorilla/websocket"
)

type Reply struct {
	Title    string   `json:"title"`
	Body     string   `json:"body"`
	Solution []string `json:"solution"`
}

func SendBfsLogs(logOutput chan Reply, title, body string, solutions []string) {
	logOutput <- Reply{Title: title, Body: body, Solution: solutions}
}

func SendLogs(logOutput chan Reply, conn *websocket.Conn) {
	for {
		select {
		case val, ok := <-logOutput:
			if ok {
				err := conn.WriteJSON(val)
				if err != nil {
					fmt.Println("something wen wrong trying to send data to socket ", err)
				}
			} else {
				return
			}
		}
	}
}

func ConvertArrayToString(array [][]byte) []string {
	var result []string
	for _, n := range array {
		result = append(result, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(n)), "-"), "[]"))
	}
	return result
}

func UpdateFilter(count int) int {
	if count >= 10000000000000 {
		return 10000000000000
	}
	if count >= 1000000000000 {
		return 1000000000000
	}
	if count >= 100000000000 {
		return 100000000000
	}
	if count >= 10000000000 {
		return 10000000000
	}
	if count >= 1000000000 {
		return 1000000000
	}
	if count >= 100000000 {
		return 100000000
	}
	if count >= 10000000 {
		return 10000000
	}
	if count >= 1000000 {
		return 1000000
	}
	if count >= 100000 {
		return 100000
	}
	if count >= 10000 {
		return 10000
	}
	if count >= 1000 {
		return 1000
	} else {
		return 100
	}
}
