package linux

import (
	"log"
	"os/exec"
	"runtime"
)

func ExecuteLinuxCmd(cmd string) {
	systemType := runtime.GOOS

	if systemType == "linux" {

		out, err := exec.Command("echo", cmd).Output()
		if err != nil {
			log.Fatal(err)
		}

		log.Printf(string(out))
	}

}
