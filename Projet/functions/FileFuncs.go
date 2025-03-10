package functions

import (
	"os"
	"os/exec"
	"runtime"
)

// GetArgs renvoie la liste des arguments fournie au programme
func GetArgs() []string {
	return os.Args[1:]
}

// ClearCmd permet de vider le terminal
func ClearCmd() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			return
		}
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			return
		}
	}
}
