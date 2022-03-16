package files

import "fmt"

func Main(projectName *string) string {
	return fmt.Sprintf(`
package main

import (
	"%s/src/router"
	"fmt"
	"net"
)

func main() {
	fmt.Println("GYS")

	outboundIP, err := getOutboundIP()
	if err != nil {
		panic(err)
	}

	fmt.Println("Outbound IP - ", outboundIP)

	if err := router.Router().Run(); err != nil {
		panic(err)
	}
}

func getOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP, nil
}
`, *projectName)
}
