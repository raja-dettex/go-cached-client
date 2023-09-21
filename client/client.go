package client

import (
	"errors"
	"fmt"
	"io"
	"net"
)

type CacheClient struct {
	Port string
}

func New(port string) (*CacheClient, error) {
	if port == "" {
		return nil, errors.New("port can not be empty string")
	}
	return &CacheClient{
		Port: port,
	}, nil
}

func (cc *CacheClient) Open() (net.Conn, error) {
	conn, err := net.Dial("tcp", cc.Port)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("connection refused %v", err))
	}
	return conn, nil
}

func (cc *CacheClient) Set(key, value string, ttlInMinutes int) (string, error) {
	conn, err := cc.Open()
	if err != nil {
		return "", err
	}
	_, err = conn.Write([]byte(fmt.Sprintf("SET %v %v %v", key, value, ttlInMinutes)))
	if err != nil {
		return "", err
	}
	receivedBuff := make([]byte, 2048)
	n, err := conn.Read(receivedBuff)

	if err == io.EOF {
		cc.Close(conn)
	}
	if err != nil {
		return "", err
	}
	return formatAndRespond(receivedBuff[:n]), nil

}

func (cc *CacheClient) Get(key string) (string, error) {
	conn, err := cc.Open()
	if err != nil {
		return "", err
	}
	_, err = conn.Write([]byte(fmt.Sprintf("GET %v", key)))

	if err != nil {
		return "", err
	}
	receivedBuff := make([]byte, 2048)
	n, err := conn.Read(receivedBuff)
	if err != nil {
		return "", err
	}
	if err == io.EOF {
		err := cc.Close(conn)
		if err != nil {
			fmt.Println(err)
		}
	}
	return formatAndRespond(receivedBuff[:n]), nil

}
func (cc *CacheClient) Delete(key string) (string, error) {
	conn, err := cc.Open()
	if err != nil {
		return "", err
	}
	_, err = conn.Write([]byte(fmt.Sprintf("DELETE %v", key)))

	if err != nil {
		return "", err
	}
	receivedBuff := make([]byte, 2048)
	n, err := conn.Read(receivedBuff)
	if err != nil {
		return "", err
	}
	if err == io.EOF {
		err := cc.Close(conn)
		if err != nil {
			fmt.Println(err)
		}
	}
	return formatAndRespond(receivedBuff[:n]), nil

}
func (cc *CacheClient) Has(key string) (string, error) {
	conn, err := cc.Open()
	if err != nil {
		return "", err
	}
	_, err = conn.Write([]byte(fmt.Sprintf("HAS %v", key)))

	if err != nil {
		return "", err
	}
	receivedBuff := make([]byte, 2048)
	n, err := conn.Read(receivedBuff)
	if err != nil {
		return "", err
	}
	if err == io.EOF {
		err := cc.Close(conn)
		if err != nil {
			fmt.Println(err)
		}
	}
	return formatAndRespond(receivedBuff[:n]), nil

}
func (cc *CacheClient) Close(conn net.Conn) error {
	err := conn.Close()
	if err != nil {
		return errors.New(fmt.Sprintf("can not close the connection %v", err))
	}
	return nil
}

func formatAndRespond(raw []byte) string {
	return string(raw)
}
