---

# Load Balancer Project Documentation

## Overview
This project implements a simple load balancer using Python and Flask. A load balancer is a critical component in distributed systems that helps distribute incoming network traffic across multiple backend servers to ensure optimal resource utilization, maximize throughput, and minimize response time. In this project, we have created a basic load balancer that randomly selects backend servers from a predefined list and redirects incoming requests to one of these servers.

## Load Balancer Functionality
The project consists of two main components:

1. **`server.py`**: This file defines a Flask application that acts as a proxy server and load balancer. It maintains a list of backend servers and periodically updates the response times of these servers. The `/stats` endpoint provides real-time information about the response times of each backend server. The load balancer randomly selects a backend server from the list and redirects incoming requests to that server.

2. **`visual.py`**: This script fetches the server statistics from the `/stats` endpoint of the Flask application and visualizes the response times using matplotlib. It continuously fetches the server stats and updates the visualization every 3 seconds.

## Usage
To use the load balancer project:
1. Run the `server.py` script to start the Flask application.
2. Run the `visual.py` script to fetch and visualize the server statistics.

## CORS Policy
To address the Cross-Origin Resource Sharing (CORS) policy issue, we have enabled CORS for all routes in the Flask application. This allows requests from different origins, such as frontend applications running on `http://127.0.0.1:5500`, to access the `/stats` endpoint without encountering CORS errors.

## Future Improvements
Some potential improvements for this project include:
- Implementing more advanced load balancing algorithms (e.g., round-robin, least connections).
- Adding health checks to monitor the status of backend servers.
- Enhancing the visualization of server statistics for better insights.

## Conclusion
This project provides a basic implementation of a load balancer using Python and Flask. By understanding the principles of load balancing and exploring this project, developers can gain valuable insights into building scalable and efficient distributed systems.

Feel free to contribute, provide feedback, or explore further enhancements to this project!

---
