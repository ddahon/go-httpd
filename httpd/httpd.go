package httpd

import (
	"bytes"
	"fmt"
	"net"
)

const REQ_MAX_LENGTH_KB = 10

type Httpd struct {
    routes map[string]string
    conn net.Conn
    ln net.Listener
}

func NewServer(routes map[string]string) Httpd {
    return Httpd{routes, nil, nil}
}

func getPath([]byte) string {
    //TODO get path from request
    return "tes"
}

func (httpd *Httpd) handleRequest(req []byte) {
    fmt.Println(string(req))
    res, ok := httpd.routes[getPath(req)]
    if !ok {
        res = "Route not found"        
    }
    fmt.Println(res)
    httpd.conn.Write([]byte(res))
}


func (httpd *Httpd) Start(address string) {
    ln, err := net.Listen("tcp", address)
    if err != nil {
        panic(err)
    }
    httpd.ln = ln
    b := bytes.NewBuffer(make([]byte, REQ_MAX_LENGTH_KB * 1000))
    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println(err.Error())
        }
        httpd.conn = conn
        n, err := conn.Read(b.Bytes())
        httpd.handleRequest(b.Next(n))
    }
}
