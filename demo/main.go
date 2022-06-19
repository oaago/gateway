package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/oaago/gateway/discovery"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
)

var etcdAddrs = []string{"http://0.0.0.0:2379"}

func main() {
	r := discovery.NewResolver(etcdAddrs, zap.NewNop())
	resolver.Register(r)

	// etcd中注册5个服务
	go newServer(":1001", "1.0.0", 1)
	go newServer(":1002", "1.0.0", 1)
	go newServer(":1003", "1.0.0", 1)
	go newServer(":1004", "1.0.0", 1)
	go newServer(":1006", "1.0.0", 10)

	conn, err := grpc.Dial("etcd:///hello", grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name))
	if err != nil {
		fmt.Println(err.Error())
	}
	defer conn.Close()

	c := NewFooClient(conn)
	time.Sleep(10 * time.Second)
	// 进行十次数据请求
	for i := 0; i < 10; i++ {
		resp, err := c.Greet(context.Background(), &GreetReq{MyName: "abc"})
		if err != nil {
			fmt.Println("say hello failed %v" + err.Error())
		}
		fmt.Println(resp.Msg)
		time.Sleep(300 * time.Millisecond)
	}
	time.Sleep(10 * time.Second)
}

type server struct {
	Port string
}

// Greet implements helloworld.GreeterServer
func (s *server) Greet(ctx context.Context, in *GreetReq) (*GreetResp, error) {
	return &GreetResp{Msg: fmt.Sprintf("Hello From %s", s.Port)}, nil
}

func newServer(port string, version string, weight int64) {
	register := discovery.NewRegister(etcdAddrs, zap.NewNop())
	defer register.Stop()

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	s := grpc.NewServer()
	RegisterFooServer(s, &server{Port: port})

	info := discovery.Server{
		Name:    "hello",
		Addr:    fmt.Sprintf("127.0.0.1%s", port),
		Version: version,
		Weight:  weight,
	}

	register.Register(info, 10)

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}
