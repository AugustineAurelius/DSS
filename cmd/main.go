package main

import (
	"fmt"
	"net/http"
	"net/http/pprof"
	_ "net/http/pprof"

	"github.com/AugustineAurelius/DSS/internal/node"
)

func main() {

	n1 := node.New(":3998")
	err := n1.Serve()
	if err != nil {
		return
	}

	n2 := node.New(":3999")
	n2.Serve()

	err = n2.Dial(":3998")
	if err != nil {
		return
	}

	nodes := make([]*node.Node, 100)

	for i := 0; i < 100; i++ {
		nodes[i] = node.New(fmt.Sprintf(":400%d", i))
	}

	for i := 0; i < 100; i++ {
		nodes[i].Dial(":3998")
	}

	r := http.NewServeMux()

	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	http.ListenAndServe(":8080", r)
}
