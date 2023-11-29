package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	defer conn.Close() // Close connection before exit

	// Create Buffer to Store date
	buffer := make([]byte, 1024) //1024 bytes

	// Receive filename from Client

	fileNameBuffer := make([]byte, 64)

	n, err := conn.Read(fileNameBuffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	fileName := string(fileNameBuffer[:n])
	fmt.Println("Receive File Name:", fileName)

	// create file to store data
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer file.Close() //close file brfore exit

	// Receive and erite data to file
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Transfer Complete")
			} else {
				fmt.Println(err)
			}
			return
		}
		file.Write(buffer[:n])
	}

}

func main() {
	listener, err := net.Listen("tcp", "10.6.5.139:5000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening on port 5000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("client Connected:", conn.RemoteAddr())

		go handleConnection(conn)
	}
}
