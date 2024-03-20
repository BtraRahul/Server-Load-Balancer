from flask import Flask, jsonify, redirect, request
import requests
import time
import random
import threading
from flask_cors import CORS  # Import CORS module

app = Flask(__name__)
CORS(app)  # Enable CORS for all routes

# List of backend servers
backend_servers = [
    "https://example.com",
    "https://facebook.com",
    "https://instagram.com",
    "https://jsonplaceholder.typicode.com",
    "https://api.publicapis.org",
    "https://dog.ceo/api/breeds/list/all",
]

# Dictionary to store server response times
server_response_times = {server: 0 for server in backend_servers}

def get_server_response(server):
    start_time = time.time()
    try:
        response = requests.get(server)
        time_to_connect = round((time.time() - start_time) * 1000, 2)  # Convert to milliseconds
        server_response_times[server] = time_to_connect
    except requests.RequestException as e:
        server_response_times[server] = -1  # Indicates connection failure

def update_server_stats():
    while True:
        for server in backend_servers:
            get_server_response(server)
        time.sleep(3)  # Update server stats every 3 seconds

@app.route("/")
def proxy_server():
    server = random.choice(backend_servers)
    return redirect(server)

@app.route("/stats")
def server_stats():
    return jsonify(server_response_times)

if __name__ == "__main__":
    update_thread = threading.Thread(target=update_server_stats)
    update_thread.daemon = True
    update_thread.start()

    app.run(port=8081)
