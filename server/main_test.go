package main

import (
	"net"
	"testing"
)

type StubServer struct {
	conn net.Conn
}

func TestReceive(t *testing.T) {
	serverConn, clientConn := net.Pipe()	
	go func () {
		defer serverConn.Close()
		
		serverConn.Write([]byte("teste"))

		go HandleConnection(serverConn)

		buf := make([]byte, 1024)
		n, err := serverConn.Read(buf)
		assetNotError(t, err)

		got := buf[:n]
		want := "Received:"


		if string(got) != want {
			t.Errorf("want %q, got %q", want, got)
		}
	}()
	
	defer clientConn.Close()
	clientConn.Write([]byte("test client")) 
}


func assetNotError(t *testing.T, err error){
	t.Helper()

	if err != nil {
		t.Fatalf("was not told to err: %q", err)
	}

}
