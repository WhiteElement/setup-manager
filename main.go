package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {

	// Install programs
	programs := [4]string{"i3", "tmux", "git", "build-essential"}
	for _, prog := range programs {
		basisErr := installBasis(prog)
		if basisErr != nil {
			log.Printf("Error running command(%s)", basisErr.Error())
		}
	}

	// Make ~/Dev
	dirErr := os.Mkdir("/home/manu/Dev", 0777)
	if dirErr != nil {
		log.Printf("Error making dir(%s)", dirErr.Error())
	}

}

func installBasis(programm string) error {

	cmd := exec.Command("sudo", "apt", "install", programm)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	startErr := cmd.Start()
	if startErr != nil {
		return startErr
	}
	waitErr := cmd.Wait()
	if waitErr != nil {
		return waitErr
	}

	return nil
}
