package main

import (
	"net/http"
	"net/http/pprof"
	_ "net/http/pprof"

	"github.com/AugustineAurelius/DSS/internal/node"
)

func main() {

	n1 := node.New(":4001")
	err := n1.Serve()
	if err != nil {
		return
	}

	n2 := node.New(":4002")

	n2.Serve()

	err = n2.Dial(":4001")
	if err != nil {
		return
	}

	r := http.NewServeMux()

	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	http.ListenAndServe(":8080", r)
}
