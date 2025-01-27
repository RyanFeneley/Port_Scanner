# Port Scanner
## Overview
This project is a simple Python-based port scanner that scans a range of ports on a given IP address and reports which ports are open. It can be used to check for open ports on a server and to analyze network security.

## Features
- Scans a specified IP address for open ports.
- Allows scanning of a customizable range of ports.
- Provides a report of open ports.
- Uses multi-threading for faster scanning, allowing multiple ports to be checked concurrently.

## Requirements
- Python 3.x
- socket module (part of Python's standard library)
- 	hreading module (part of Python's standard library)

## Usage
1. Clone the repository or download the code.
2. Run the script with the target IP address and port range:
   \\\ash
   python port_scanner.py <target-ip> <start-port> <end-port>
   \\\
   Example:
   \\\ash
   python port_scanner.py 192.168.1.1 1 1024
   \\\

## How it Works
- The script uses Python's socket module to attempt a connection to each port on the target IP address within the specified range.
- If the connection is successful, it reports the port as open.
- Multi-threading is used to speed up the process by scanning multiple ports simultaneously.
