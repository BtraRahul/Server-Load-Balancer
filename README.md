# Simple Load Balancer in Go

This project implements a basic load balancer in Go that uses a round-robin algorithm to distribute incoming HTTP requests among a predefined list of servers. The primary goal is to demonstrate how load balancing works and how it can be implemented in Go. The load balancer ensures that the requests are evenly distributed, thus preventing any single server from being overwhelmed with too many requests.

## Features
### Round-Robin Scheduling: 
The core of the load balancing algorithm, which distributes incoming requests evenly across all available servers in a circular order. This ensures fair workload distribution and maximizes resource utilization.
### Server Health Check: 
Although simplified, the framework includes a mechanism to check if a server is alive before forwarding requests. This ensures reliability and availability, as requests are only sent to operational servers.
### Reverse Proxying: 
Utilizes the httputil.ReverseProxy from Go's standard library to forward requests to the target server. This acts as a reverse proxy, making the load balancer transparent to the end user.
### Simple Server Implementation: 
Demonstrates how to set up basic servers that the load balancer will manage. These servers are represented by various popular websites for demonstration purposes.
### Error Handling: 
Includes basic error handling to deal with scenarios where the target servers are unreachable.

## Usage
The load balancer is designed to be straightforward to use. It initializes a list of servers (in this case, popular websites) and starts a web server that listens on port 8080. All incoming requests to this port are automatically distributed among the available servers based on the round-robin algorithm.

## Implementation Details
### Server Interface: 
Defines the essential functions that a server must implement to be managed by the load balancer.
### LoadBalancer Struct: 
Holds the load balancer's state, including the list of servers and the index for the next server to use.
### SimpleServer Struct: 
A basic implementation of the Server interface, representing an individual server managed by the load balancer.

## Code Structure
- main.go: Contains the main function and the implementation of the load balancer, server interface, and simple server.
Running the Project

- To run the project, simply execute: `go run main.go`

This will start the load balancer on port 8080, and you can begin making requests to it.