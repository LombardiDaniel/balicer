package main

import (
	"flag"
	"fmt"
	"log/slog"

	"github.com/patos-ufscar/balicer/cli"
	"github.com/patos-ufscar/balicer/common"
	"github.com/patos-ufscar/balicer/handlers"
	"github.com/patos-ufscar/balicer/servers"
	"github.com/patos-ufscar/balicer/utils"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "configPath", "./defaultConf.yml", "the config.yml file path")
	flag.Parse()

}

func main() {
	utils.InitSlogger()
	serverConfs, err := cli.ParseConfig(configPath)
	if err != nil {
		slog.Error("error on parsing")
		return
	}

	// TODO: fazer a serverConfs retornar só a conf
	// ai usar diversas factories pra utilizar a conf e iniciar os módulos
	// TODO: fazer as paths só permitirem continuar se finalizar com "/", mas pensar em param tb

	for _, v := range serverConfs {
		lis, err := common.Bind(v.Port)
		if err != nil {
			slog.Error(fmt.Sprintf("Could not bind to port: %d", v.Port))
			panic("could not bind to port")
		}

		slog.Info(fmt.Sprintf("Listening on port %d", v.Port))

		hs := []handlers.Handler{}
		for _, locConf := range v.Locations {
			h, err := handlers.HandlerFactory(locConf)
			if err != nil {
				slog.Error(err.Error())
				return
			}
			hs = append(hs, h)
		}

		server := servers.NewServer(
			v.Port,
			v.HostsRegs,
			hs,
		)

		go server.Serve(*lis)
	}

	hangChannel := make(chan int)

	<-hangChannel
}

// package main

// import (
// 	"log"
// 	"net/http"
// 	"net/http/httputil"
// 	"net/url"
// )

// func main() {
//         remote, err := url.Parse("http://google.com")
//         if err != nil {
// 			panic(err)
//         }

//         handler := func(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
// 			return func(w http.ResponseWriter, r *http.Request) {
// 				log.Println(r.URL)
// 				r.Host = remote.Host
// 				w.Header().Set("X-Ben", "Rad")
// 				p.ServeHTTP(w, r)
// 			}
//         }

//         proxy := httputil.NewSingleHostReverseProxy(remote)
//         http.HandleFunc("/", handler(proxy))
//         err = http.ListenAndServe(":8080", nil)
//         if err != nil {
// 			panic(err)
//         }
// }
