package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type Server interface {
	Address() string
	IsAlive() bool
	Serve(rw http.ResponseWriter, req *http.Request)
}

type simpleServer struct {
	address string
	proxy   *httputil.ReverseProxy
}

type LoadBalancer struct {
	port            string
	roundRobinCount int
	servers         []Server
}

func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port:            port,
		roundRobinCount: 0,
		servers:         servers,
	}
}

func newSimpleServer(address string) *simpleServer {
	serverUrl, err := url.Parse(address)
	handlerErr(err)
	return &simpleServer{
		address: address,
		proxy:   httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

func handlerErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (s *simpleServer) Address() string {
	return s.address
}

func (s *simpleServer) IsAlive() bool {
	return true
}

func (s *simpleServer) Serve(rw http.ResponseWriter, req *http.Request) {
	s.proxy.ServeHTTP(rw, req)
}

func (lb *LoadBalancer) getNextAvailableServer() Server {
	server := lb.servers[lb.roundRobinCount%len(lb.servers)]

	for !server.IsAlive() {
		lb.roundRobinCount++
		server = lb.servers[lb.roundRobinCount%len(lb.servers)]
	}
	lb.roundRobinCount++
	return server
}

func (lb *LoadBalancer) serveProxy(rw http.ResponseWriter, req *http.Request) {
	lb.roundRobinCount++
	server := lb.getNextAvailableServer()
	fmt.Println("Forwarding request to", server.Address())
	if server != nil {
		server.Serve(rw, req)
		return
	}
	http.Error(rw, "Service not available", http.StatusServiceUnavailable)
}

func main() {
	servers := []Server{
		newSimpleServer("https://google.com"),
		newSimpleServer("https://reddit.com"),
		newSimpleServer("https://youtube.com"),
		newSimpleServer("https://facebook.com"),
		newSimpleServer("https://duckduckgo.com"),
	}

	lb := NewLoadBalancer(":8080", servers)
	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		lb.serveProxy(rw, req)
	}

	http.HandleFunc("/", handleRedirect)
	fmt.Println("Server started at port " + lb.port)
	http.ListenAndServe(lb.port, nil)
}
