package main

import (
	"encoding/gob"
	"errors"
	"log"

	"github.com/Jiahonzheng/Tiny-RPC"
	"github.com/Jiahonzheng/Tiny-RPC/example/public"
)

func queryUser(uid int) (public.ResponseQueryUser, error) {
	db := make(map[int]public.User)
	db[0] = public.User{Name: "Jiahonzheng", Age: 70}
	db[1] = public.User{Name: "ChiuSinYing", Age: 75}
	if u, ok := db[uid]; ok {
		return public.ResponseQueryUser{User: u, Msg: "success"}, nil
	}
	return public.ResponseQueryUser{User: public.User{}, Msg: "fail"}, errors.New("uid is not in database")
}

func main() {
	gob.Register(public.ResponseQueryUser{})

	addr := "0.0.0.0:2333"
	srv := tiny.NewServer(addr)
	srv.Register("queryUser", queryUser)
	log.Println("service is running")
	go srv.Run()

	for {
	}
}
