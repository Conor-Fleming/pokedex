package main

import (
	"fmt"
)

const (
	BANNER = `
====================================
__________       __          ________                 
\______   \____ |  | __ ____ \______ \   ____ ___  ___
 |     ___/  _ \|  |/ // __ \ |    |  \_/ __ \\  \/  /
 |    |  (  <_> )    <\  ___/ |        \  ___/ >    < 
 |____|   \____/|__|_ \\___  >_______  /\___  >__/\_ \
                     \/    \/        \/     \/      \/
                           
====================================
`
)

func main() {
	fmt.Print(BANNER)
	cfg := &config{}

	startPoke(cfg)
}

func invalidCommand() {
	fmt.Println("Command not found")
}
