package main

import (
	"fmt"
	"net/http"
	"net/http/pprof"
	_ "net/http/pprof"
	"os"

	"github.com/AugustineAurelius/DSS/internal/node"
)

func main() {
	// f, _ := os.Create("trace.out")
	// defer f.Close()
	// trace.Start(f)
	// defer trace.Stop()
	agrs := os.Args

	if agrs[1] == "server" {
		fmt.Println("server arg")
		n1 := node.New(":3998")
		err := n1.Serve()
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
	} else {
		for i := 0; i < 100; i++ {
			node.New(fmt.Sprintf(":400%d", i)).Dial(":3998")
		}
		r := http.NewServeMux()

		r.HandleFunc("/debug/pprof/", pprof.Index)
		r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		r.HandleFunc("/debug/pprof/profile", pprof.Profile)
		r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		r.HandleFunc("/debug/pprof/trace", pprof.Trace)

		http.ListenAndServe(":8081", r)

	}

}
