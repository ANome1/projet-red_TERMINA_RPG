package RED

import (
	"os"
	"os/exec"
	"runtime"
	"time"
)

func ClearTerminal() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func Pause(seconde int) {
	time.Sleep(time.Duration(seconde) * time.Second)
}
