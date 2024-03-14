package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"syscall"
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
	server          *http.Server
}

func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port:            port,
		roundRobinCount: 0,
		servers:         servers,
		server:          &http.Server{Addr: port},
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
	resp, err := http.Get(s.address)
	if err != nil {
		fmt.Printf("[Health Check] Server %s - Error: %s\n", s.address, err.Error())
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return true
	}

	fmt.Printf("[Health Check] Server %s - Status Code: %d\n", s.address, resp.StatusCode)
	return false
}

func (s *simpleServer) Serve(rw http.ResponseWriter, req *http.Request) {
	s.proxy.ServeHTTP(rw, req)
}

func (lb *LoadBalancer) getNextAvailableServer() Server {
	server := lb.servers[lb.roundRobinCount%len(lb.servers)]
	fmt.Printf("[Server Check] Checking server %s\n", server.Address())

	for !server.IsAlive() {
		fmt.Printf("[Server Check] Server %s is not alive\n", server.Address())
		lb.roundRobinCount++
		server = lb.servers[lb.roundRobinCount%len(lb.servers)]
	}
	fmt.Printf("[Server Check] Server %s is alive\n", server.Address())
	lb.roundRobinCount++
	return server
}

func (lb *LoadBalancer) serveProxy(rw http.ResponseWriter, req *http.Request) {
	lb.roundRobinCount++
	server := lb.getNextAvailableServer()
	fmt.Printf("[Request Forwarding] Forwarding request to %s\n", server.Address())
	if server != nil {
		server.Serve(rw, req)
		return
	}
	http.Error(rw, "Service not available", http.StatusServiceUnavailable)
}

func (lb *LoadBalancer) start() {
	http.HandleFunc("/", lb.serveProxy)
	fmt.Println("Server started at port " + lb.port)

	go func() {
		if err := lb.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Error starting server:", err)
			os.Exit(1)
		}
	}()
}

func (lb *LoadBalancer) stop() {
	fmt.Println("Shutting down server...")
	if err := lb.server.Shutdown(context.Background()); err != nil {
		fmt.Println("Error shutting down server:", err)
	}
}

func (lb *LoadBalancer) displayServerHealthAndOrder() {
	fmt.Println("[Server Health Status]")
	for _, server := range lb.servers {
		status := "not alive"
		if server.IsAlive() {
			status = "alive"
		}
		fmt.Printf("Server %s is %s\n", server.Address(), status)
	}
}

func main() {
	servers := []Server{
		newSimpleServer("https://example.com"),
		newSimpleServer("https://facebook.com"),
		newSimpleServer("https://instagram.com"),
		newSimpleServer("https://jsonplaceholder.typicode.com"),
		newSimpleServer("https://api.publicapis.org"),
		newSimpleServer("https://dog.ceo/api/breeds/list/all"),
		newSimpleServer("https://nonexistentwebsite123.com"), // Non-existent server
	}

	lb := NewLoadBalancer(":8081", servers)
	lb.start()

	// Listen for interrupt signal to gracefully shutdown
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	<-interrupt
	lb.stop()
}
