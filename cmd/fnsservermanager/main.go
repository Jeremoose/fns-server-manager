package main

import (
	"fmt"
	"log"

	fnsservermanager "github.com/jeremoose/fns-server-manager/internal/app"
)

func main() {
	sandstormServer, err := fnsservermanager.NewServer("sandstorm server", "/home/steam/steamcmd/sand/Insurgency/Binaries/Linux/InsurgencyServer-Linux-Shipping", "Mountain?Scenario=Scenario_Summit_Checkpoint_Security?MaxPlayers=8?Lighting=\"Night\"", "-Port=2302", "-GameStatsToken=D2D7FB43DE2649E097E763F8C5DC523D", "-GSLTToken=8AAAA7AF70C9FE4BC4046765C3E2964B", "-log")
	if err != nil {
		log.Fatal(err)
	}

	if err := sandstormServer.Run(); err != nil {
		log.Fatal(err)
	}

	stdoutBuf, err := sandstormServer.GetStdOut()
	if err != nil {
		log.Fatal(err)
	}
	stderrBuf, err := sandstormServer.GetStdErr()
	if err != nil {
		log.Fatal(err)
	}

	outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
}
