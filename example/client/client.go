package main

import (
	"encoding/gob"
	"log"
	"net"

	"github.com/Jiahonzheng/Tiny-RPC"
	"github.com/Jiahonzheng/Tiny-RPC/example/public"
)

func main() {
	gob.Register(public.ResponseQueryUser{})

	addr := "0.0.0.0:2333"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Printf("dial error: %v\n", err)
	}
	cli := tiny.NewClient(conn)

	var correctQuery func(int) (public.ResponseQueryUser, error)
	var wrongQuery func(int) (public.ResponseQueryUser, error)

	cli.Call("queryUser", &correctQuery)
	u, err := correctQuery(1)
	if err != nil {
		log.Printf("query error: %v\n", err)
	} else {
		log.Println(u.Name, u.Age, u.Msg)
	}
	u, err = correctQuery(2)
	if err != nil {
		log.Printf("query error: %v\n", err)
	} else {
		log.Printf("query result: %v %v %v\n", u.Name, u.Age, u.Msg)
	}

	cli.Call("queryUse", &wrongQuery)
	u, err = wrongQuery(1)
	if err != nil {
		log.Printf("query error: %v\n", err)
	} else {
		log.Println(u)
	}

	conn.Close()
}
