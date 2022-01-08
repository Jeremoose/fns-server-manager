package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	sandstormServer := fnsservermanager
	cmd := exec.Command("/home/steam/steamcmd/sand/Insurgency/Binaries/Linux/InsurgencyServer-Linux-Shipping", "Mountain?Scenario=Scenario_Summit_Checkpoint_Security?MaxPlayers=8?Lighting=\"Night\"", "-Port=2302", "-GameStatsToken=D2D7FB43DE2649E097E763F8C5DC523D", "-GSLTToken=8AAAA7AF70C9FE4BC4046765C3E2964B", "-log")
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
}
