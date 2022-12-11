package winapi

import (
	"syscall"
	"unsafe"
)

var kernel32 = syscall.NewLazyDLL("kernel32.dll")
var setConsoleCursorInfo = kernel32.NewProc("SetConsoleCursorInfo")
var setConsoleCursorposition = kernel32.NewProc("SetConsoleCursorPosition")
var setConsoleTextAttribute = kernel32.NewProc("SetConsoleTextAttribute")
var sleep = kernel32.NewProc("Sleep")

var user32 = syscall.NewLazyDLL("user32.dll")
var getAsyncKeyState = user32.NewProc("GetAsyncKeyState")

func GetAsyncKeyState(vkey int32) uint32 {
	state, _, _ := getAsyncKeyState.Call(uintptr(vkey))

	return (uint32)(state)
}

func Gotoxy(x int32, y int32) {
	hwnd, _ := syscall.GetStdHandle(syscall.STD_OUTPUT_HANDLE)

	setConsoleCursorposition.Call(uintptr(hwnd), uintptr(x|(y<<16)))
}

func ShowCursor(val int32) {
	type ConsoleCursorInfo struct {
		dwSize   int32
		bVisible int32
	}

	info := ConsoleCursorInfo{
		dwSize:   100,
		bVisible: val,
	}

	hwnd, _ := syscall.GetStdHandle(syscall.STD_OUTPUT_HANDLE)

	setConsoleCursorInfo.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&info)))
}

func SetTextColor(color int32) {
	hwnd, _ := syscall.GetStdHandle(syscall.STD_OUTPUT_HANDLE)

	setConsoleTextAttribute.Call(uintptr(hwnd), uintptr(color))
}

func Sleep(ms int32) {
	sleep.Call(uintptr(ms))
}
