package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
	"unsafe"
	"github.com/abdabTheCreator/minecraft-control/mouse-control"
)

const (
	VK_LEFT  = 0x25
	VK_UP    = 0x26
	VK_RIGHT = 0x27
	VK_DOWN  = 0x28
)

var (
	user32                   = syscall.NewLazyDLL("user32.dll")
	procGetForegroundWindow  = user32.NewProc("GetForegroundWindow")
	procGetWindowTextW       = user32.NewProc("GetWindowTextW")
)

func getActiveWindowTitle() (string, error) {
	ret, _, _ := procGetForegroundWindow.Call()
	if ret == 0 {
		return "", fmt.Errorf("failed to get active window")
	}

	var buf [256]uint16
	procGetWindowTextW.Call(ret, uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))

	title := syscall.UTF16ToString(buf[:])
	return title, nil
}

func isMinecraftActive() bool {
	title, err := getActiveWindowTitle()
	if err != nil {
		return false
	}

	return strings.HasPrefix(title, "Minecraft ")
}

func simulateMouseMovement(x, y int) {
	// Implement the mouse movement logic here
	fmt.Printf("Simulating mouse movement: x=%d, y=%d\n", x, y)
}

func handleKeyPress(keyCode int, mouseMoveCallback func(int, int)) {
	switch keyCode {
	case VK_LEFT:
		mouseMoveCallback(-10, 0) // Move left
	case VK_UP:
		mouseMoveCallback(0, -10) // Move up
	case VK_RIGHT:
		mouseMoveCallback(10, 0) // Move right
	case VK_DOWN:
		mouseMoveCallback(0, 10) // Move down
	default:
		// Ignore other keypresses
	}
}

func main() {
	// Detect the operating system to set appropriate command for clearing console
	clearCmd := ""
	switch runtime.GOOS {
	case "windows":
		clearCmd = "cls"
	case "linux", "darwin":
		clearCmd = "clear"
	default:
		fmt.Println("Unsupported operating system")
		os.Exit(1)
	}

	// Start an infinite loop to handle keypresses
	for {
		// Clear the console
		cmd := exec.Command(clearCmd)
		cmd.Stdout = os.Stdout
		cmd.Run()

		// Check if Minecraft is active
		if !isMinecraftActive() {
			fmt.Println("Minecraft is not active. Exiting...")
			break
		}

		// Prompt for keypress
		fmt.Println("Press arrow keys to move the mouse. Press any other key to exit.")
		var keyCode int
		fmt.Scanf("%d\n", &keyCode)

		handleKeyPress(keyCode, simulateMouseMovement)
	}
}

