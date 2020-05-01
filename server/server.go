package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

type SocksServer struct {
	Version int
}

type Request struct {
	Version byte
	Command byte
	AddrType byte
	Ip4Addr []byte
	Ip6Addr []byte
	Host []byte
	Port int
}

type Response struct {
	Version byte
	ResponseCode byte
	AddrType byte
	AddrDest []byte
	Port []byte
}

func (s *SocksServer) ListenAndServer(addr string) error {

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer listen.Close()

	for {
		conn,err := listen.Accept()
		if err != nil {
			return err
		}

		go s.handleConnection(conn)
	}
	return nil
}

func (s *SocksServer) handleConnection(conn net.Conn) error {
	reader := bufio.NewReader(conn)
	version, err := reader.ReadByte()
	if err != nil {
		return err
	}

	if int(version) != s.Version {
		return fmt.Errorf("version not supported")
	}

	authTypeCnt,err := reader.ReadByte()
	authTypes := make([]byte,int(authTypeCnt))
	_, err = io.ReadFull(conn,authTypes)
	if err != nil {
		return err
	}

	_, err = conn.Write([]byte{byte(s.Version), 0})
	if err != nil {
		return err
	}

	return s.request(conn,reader)
}

func (s *SocksServer) request(conn net.Conn,bConn *bufio.Reader) error {

	req := &Request{}

	header := make([]byte, 4)
	_, err := io.ReadAtLeast(conn, header, 4)
	if err != nil {
		return err
	}

	addrLen := 0

	switch req.AddrType {
	case 0x1:
		addrLen = 4
	case 0x4:
		addrLen = 16
	case 0x3:
		length,err := bConn.ReadByte()
		if err != nil {
			return err
		}
		addrLen = int(length)
	}

	portBytes := 2
	b := make([]byte,addrLen + portBytes)

	_, err = io.ReadFull(conn, b)
	if err != nil {
		return err
	}

	switch req.AddrType {
	case 0x1:
		req.Ip4Addr = b[:addrLen]
	case 0x4:
		req.Ip6Addr = b[:addrLen]
	case 0x3:
		req.Host = b[:addrLen]
	}

	req.Port = int(b[addrLen]) << 8 | int(b[addrLen + 1])

	switch req.Command {
	case 0x1:
		return nil
	default:
		return s.error(conn,Comm)
	}
	
	return nil
}

func (s *SocksServer) error(conn net.Conn,respCode byte) error {
	resp := &Response{
		Version:s.Version,

	}


}