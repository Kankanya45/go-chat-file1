package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func sendFile(conn net.Conn, filePath string) {
	defer conn.Close()

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	//send file name to server
	_, fileName := filepath.Split(filePath)

	conn.Write([]byte(fileName))

	// create buffer to read file
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Send file success")
			} else {
				fmt.Println(err)
			}
			return
		}
		conn.Write(buffer[:n])
	}
}

func main() {
	// Connect to server
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	fmt.Println("Connect to server success")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter file pat+name: ")
	filePath, _ := reader.ReadString('\n')
	filePath = strings.TrimSpace(filePath)

	sendFile(conn, filePath)
}
