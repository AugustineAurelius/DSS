package main

import (
	"fmt"
	"net/http"
	"net/http/pprof"
	_ "net/http/pprof"
	"os"
	"runtime/trace"

	"github.com/AugustineAurelius/DSS/internal/node"
)

func main() {
	f, _ := os.Create("trace.out")
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()

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
		nodes[i].Serve()
	}

	for i := 0; i < 100; i++ {
		nodes[i].Dial(":3998")
	}

	for i := 0; i < 15; i++ {
		for j := 0; j < 15; j++ {
			if j == i {
				continue
			}
			nodes[i].Dial(fmt.Sprintf(":400%d", j))

		}

	}

	r := http.NewServeMux()

	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	http.ListenAndServe(":8080", r)
}
