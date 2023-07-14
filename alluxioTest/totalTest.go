package alluxioTest

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"jasondrogba/alluxio-cacheTest/sshTest"
	"log"
	"os"
)

func AlluxioTest(hostname string) {
	config := sshTest.SetupSSH()
	log.Println("start alluxio")
	StopAlluxio(config, hostname)
	FormatAlluxio(config, hostname)
	RunAlluxio(config, hostname)
	log.Println("start test")
	RunTest(config, hostname)
}

func FormatAlluxio(config *ssh.ClientConfig, hostname string) {
	conn, err := ssh.Dial("tcp", hostname+":"+"22", config)
	if err != nil {
		fmt.Println("Failed to establish SSH connection:", err)
		os.Exit(1)
	}
	defer conn.Close()
	session, err := conn.NewSession()
	if err != nil {
		fmt.Println("Failed to create session:", err)
		os.Exit(1)
	}
	defer session.Close()

	cmd := fmt.Sprintf("sudo su alluxio -c \"cd /opt/alluxio && ./bin/alluxio format\"")
	output, err := session.Output(cmd)
	if err != nil {
		fmt.Println("Failed to run command:", err)
		os.Exit(1)
	}
	fmt.Print(string(output))
}
func StopAlluxio(config *ssh.ClientConfig, hostname string) {
	conn, err := ssh.Dial("tcp", hostname+":"+"22", config)
	if err != nil {
		fmt.Println("Failed to establish SSH connection:", err)
		os.Exit(1)
	}
	defer conn.Close()
	session, err := conn.NewSession()
	if err != nil {
		fmt.Println("Failed to create session:", err)
		os.Exit(1)
	}
	defer session.Close()

	cmd := fmt.Sprintf("sudo su alluxio -c \"cd /opt/alluxio && ./bin/alluxio-stop.sh all SudoMount\"")
	output, err := session.Output(cmd)
	if err != nil {
		fmt.Println("Failed to run command:", err)
		os.Exit(1)
	}
	fmt.Print(string(output))
}
func RunAlluxio(config *ssh.ClientConfig, hostname string) {
	conn, err := ssh.Dial("tcp", hostname+":"+"22", config)
	if err != nil {
		fmt.Println("Failed to establish SSH connection:", err)
		os.Exit(1)
	}
	defer conn.Close()
	session, err := conn.NewSession()
	if err != nil {
		fmt.Println("Failed to create session:", err)
		os.Exit(1)
	}
	defer session.Close()

	cmd := fmt.Sprintf("sudo su alluxio -c \"cd /opt/alluxio && ./bin/alluxio-start.sh all\"")
	output, err := session.Output(cmd)
	if err != nil {
		fmt.Println("Failed to run command:", err)
		os.Exit(1)
	}
	fmt.Print(string(output))
}

func RunTest(config *ssh.ClientConfig, hostname string) {
	conn, err := ssh.Dial("tcp", hostname+":"+"22", config)
	if err != nil {
		fmt.Println("Failed to establish SSH connection:", err)
		os.Exit(1)
	}
	defer conn.Close()
	session, err := conn.NewSession()
	if err != nil {
		fmt.Println("Failed to create session:", err)
		os.Exit(1)
	}
	defer session.Close()

	cmd := fmt.Sprintf("./testspeed-linux")
	output, err := session.Output(cmd)
	if err != nil {
		fmt.Println("Failed to run command:", err)
		os.Exit(1)
	}
	fmt.Print(string(output))
}
