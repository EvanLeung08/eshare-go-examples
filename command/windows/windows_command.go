package windows

import (
	"log"
	"os/exec"
	"runtime"
)

func ExecuteWinCmd(cmd string) {
	systemType := runtime.GOOS

	if systemType == "windows" {

		out, err := exec.Command("cmd", "/c", "echo", cmd).Output()
		if err != nil {
			log.Fatal(err)
		}

		log.Printf(string(out))
	}

}
