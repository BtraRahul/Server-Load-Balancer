import requests
import time

def fetch_server_data():
    try:
        response = requests.get("http://localhost:8081")  # Assuming the Go program is running locally
        if response.status_code == 200:
            return response.json()
        else:
            print("Failed to fetch server data. Status Code:", response.status_code)
            return None
    except requests.RequestException as e:
        print("Error fetching server data:", e)
        return None

def display_server_data(server_data):
    if server_data:
        print("[Server Health Status]")
        for server in server_data:
            status = "alive" if server['alive'] else "not alive"
            print(f"Server {server['address']} is {status}. Connection Time: {server['connection_time']}")

if __name__ == "__main__":
    while True:
        server_data = fetch_server_data()
        display_server_data(server_data)
        time.sleep(5)  # Fetch server data every 5 seconds
