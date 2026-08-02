package server

import (
	"net",
	"fmt",
	"bufio",
	"log",
	"strings",
)

type Server struct {
	address string
}

func New(address string) *Server {
	return &Server{
		address: address,
	}
}

func (s *Server) Run() error {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		return fmt.Errorf("Listen on %s: %w", s.adderss, err)
	}

	defer listener.Close()

	fmt.Printf("Listening on %s\n", s.address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			return fmt.Errorf("accept connection: %w", err)
		}

		go handleConnection(conn)
}

func handleConnection(conn net.Conn) (reader bufio.Reader) {
	defer connection.Close()

	scanner := bufio.NewScanner(connection)

	for scanner.Scan() {
		message := scanner.Text()
		fmt.Printf("Recieved: %s\n", message)

		_, err := fmt.Fprintf(connection, "Echo: %s\n", message)
		if err != nil {
			return
		}
	}
}
