package Server

import (
	"github.com/soy-oslab/soy-log-collector"
	"github.com/smallnest/rpcx/server"
)

var addr = flag.String("addr", "localhost:8972", "server address")

func main() {
	s := server.NewServer()
	s.Register(new(HotPort), "")
	s.Serve("tcp", *addr)
}
