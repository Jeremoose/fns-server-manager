package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("/home/steam/steamcmd/sand/Insurgency/Binaries/Linux/InsurgencyServer-Linux-Shipping", "Citadel?Scenario=Scenario_Citadel_Checkpoint_Insurgents?MaxPlayers=8", "-Port=2302", "-GameStatsToken=D2D7FB43DE2649E097E763F8C5DC523D", "-GSLTToken=8AAAA7AF70C9FE4BC4046765C3E2964B", "-log", "-motd=greetings.txt")

	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", stdoutStderr)
}
