package cmd

import (
	"fmt"
	"lb/lb"
	"lb/srv"
	"log"
	"net/http"
	"net/url"
)

func Execute() error {

	srvs := []string{
		"http://127.0.0.1:8082",
		"http://127.0.0.1:8081",
	}

	servers := []srv.Server{}
	for _, val := range srvs {
		serverUrl, err := url.Parse(val)
		if err != nil {
			return err
		}
		servers = append(servers, srv.Server{
			URL: serverUrl,
		})
	}

	loadBalancer := lb.NewLoadBalancer(servers)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		srv := loadBalancer.NextServerLeastConnection()
		fmt.Println(srv.URL.String())
		srv.AddConnectionCount()

		srv.Proxy().ServeHTTP(w, r)
		srv.SubstractConnectionCount()
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ok"))
	})

	log.Println("Starting server on port")
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}

	return nil
}
