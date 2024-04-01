package cmd

import (
	"fmt"
	"lb/lb"
	"lb/srv"
	"log"
	"net/http"
	"net/url"
	"os"
)

func Execute() error {

	srvs := []string{
		"http://127.0.0.1:8080",
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
		srv := loadBalancer.NextServerLeastActive()

		srv.AddActiveConnection()
		fmt.Println(srv.URL.Host)
		req, err := http.NewRequest(r.Method, srv.URL.Path, nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}

		req.Header = r.Header
		req.Host = r.Host
		req.RemoteAddr = r.RemoteAddr

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		defer resp.Body.Close()

		fmt.Printf("Response from server: %s %s\n\n", resp.Proto, resp.Status)

		// srv.Proxy().ServeHTTP(w, r)
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
