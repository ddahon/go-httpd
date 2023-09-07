package main

import (
	"httpd"
)

const REQ_MAX_LENGTH_KB = 10

func main()  {
    routes := map[string]string{"test": "test"}
    serv := httpd.NewServer(routes)
    serv.Start(":8080")
}
