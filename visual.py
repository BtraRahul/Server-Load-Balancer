# import requests
# import matplotlib.pyplot as plt
# import time

# # List of websites to check
# websites = [
#     "https://example.com",
#     "https://facebook.com",
#     "https://instagram.com",
#     "https://jsonplaceholder.typicode.com",
#     "https://api.publicapis.org",
#     "https://dog.ceo/api/breeds/list/all",
# ]

# # Function to fetch data from websites and measure time to connect
# def measure_time_to_connect(website):
#     try:
#         response = requests.get(website)
#         return response.elapsed.total_seconds() * 1000  # Convert to milliseconds
#     except requests.RequestException:
#         return None

# # Function to update and display the plot
# def update_plot():
#     # Measure time to connect for each website
#     times_to_connect = {website: measure_time_to_connect(website) for website in websites}

#     # Filter out websites with None response times
#     valid_times_to_connect = {website: time for website, time in times_to_connect.items() if time is not None}

#     # Plot the time to connect for each website
#     plt.clf()  # Clear the previous plot
#     plt.bar(valid_times_to_connect.keys(), valid_times_to_connect.values())
#     plt.xlabel('Website')
#     plt.ylabel('Time to Connect (ms)')
#     plt.title('Time to Connect for Each Website')
#     plt.xticks(rotation=45, ha='right')
#     plt.tight_layout()
#     plt.pause(5)  # Pause for 5 seconds before updating

# # Continuous loop to update the plot
# while True:
#     update_plot()
# import requests
# import matplotlib.pyplot as plt
# import time

# def fetch_server_stats():
#     try:
#         response = requests.get("http://127.0.0.1:8081/stats")
#         if response.status_code == 200:
#             return response.json()
#         else:
#             print("Failed to fetch server stats. Status Code:", response.status_code)
#             return {}
#     except requests.RequestException as e:
#         print("Error fetching server stats:", e)
#         return {}

# def visualize_server_stats(server_stats):
#     server_addresses = list(server_stats.keys())
#     response_times = list(server_stats.values())

#     plt.bar(server_addresses, response_times)
#     plt.xlabel('Server Address')
#     plt.ylabel('Time to Connect (ms)')
#     plt.title('Server Response Times')
#     plt.xticks(rotation=45)
#     plt.tight_layout()
#     plt.show()

# if __name__ == "__main__":
#     while True:
#         server_stats = fetch_server_stats()
#         visualize_server_stats(server_stats)
#         time.sleep(3)  # Update every 3 seconds

import requests
import matplotlib.pyplot as plt
import time

def fetch_server_stats():
    try:
        response = requests.get("http://127.0.0.1:8081/stats")
        if response.status_code == 200:
            return response.json()
        else:
            print("Failed to fetch server stats. Status Code:", response.status_code)
            return {}
    except requests.RequestException as e:
        print("Error fetching server stats:", e)
        return {}

def visualize_server_stats(server_stats):
    server_addresses = list(server_stats.keys())
    response_times = list(server_stats.values())

    plt.bar(server_addresses, response_times)
    plt.xlabel('Server')
    plt.ylabel('Response Time (ms)')
    plt.title('Backend Server Response Times')
    plt.xticks(rotation=45)
    plt.show()

while True:
    server_stats = fetch_server_stats()
    visualize_server_stats(server_stats)
    time.sleep(3)  # Update and visualize server stats every 3 seconds
