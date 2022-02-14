package main

import (
	"fmt"
	"syscall"
	"time"
)

const (
	MOUSEEVENTF_ABSOLUTE = 0x8000
	MOUSEEVENTF_MOVE     = 0x0001
)

var (
	moduser32      = syscall.NewLazyDLL("user32.dll")
	procMouseEvent = moduser32.NewProc("mouse_event")
)

func main() {
	purupuru()

	// 1分ごとに発火
	ticker := time.NewTicker(1 * time.Minute)
	for {
		<-ticker.C
		purupuru()
	}
}

func purupuru() {
	t := time.Now() // 現在の時刻
	fmt.Println(t.Format("2006/01/02 15:04:05") + " プルプル")
	MouseMove(false, 0, 0)
}

func MouseMove(abs bool, x int32, y int32) {
	var intype uint32 = MOUSEEVENTF_MOVE
	if abs {
		intype |= MOUSEEVENTF_ABSOLUTE
	}
	procMouseEvent.Call(uintptr(intype), uintptr(x), uintptr(y), 0, 0)
}
