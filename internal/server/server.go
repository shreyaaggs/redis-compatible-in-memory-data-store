package server

import (
	"bufio"
	"io"
	"log"
	"net"
)

type Server struct {
	addr string
}

func New(addr string) *Server {
	return &Server{addr: addr}
}
func (s *Server) Run() error {
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	defer ln.Close()
	log.Printf("listening on %s", s.addr)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}
		go s.handleConn(conn)
	}
}
func (s *Server) handleConn(conn net.Conn) {
	defer conn.Close()
	log.Printf("client connected: %s", conn.RemoteAddr())
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			log.Printf("client disconnected: %s", conn.RemoteAddr())
			return
		}
		if err != nil {
			log.Printf("read error: %v", err)
			return
		}
		log.Printf("received: %q", line)
		writer.WriteString(line)
		writer.Flush()
	}
}
