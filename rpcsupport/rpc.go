package rpcsupport

import (
	"fmt"
	"go-crawler-distributed/mylog"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServeRpc(host string,
	service interface{}) error {
	err := rpc.Register(service)
	if err != nil {
		return err
	}
	mylog.LogInfo("rpcsupport.rpc.ServeRpc", fmt.Sprintf("listing on %s", host))

	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("error %v", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
	return nil
}

func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		return nil, err
	}

	client := jsonrpc.NewClient(conn)
	return client, nil
}
