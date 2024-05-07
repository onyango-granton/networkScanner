# Network Scanner

## Introduction
Network Scanner is a Go program designed to scan a network by sending ping requests to other members to check which IP addresses are active.

## Features
- Obtains your IP address and separates it from the subnet mask.
- Scans the network by sending ping requests to other members.
- Identifies active IP addresses on the network.

## Requirements
- Linux machine
- Go installed 

## Usage
1. Clone the repository:
   ```bash
   git clone https://github.com/onyango-granton/networkScanner.git

2. Navigate to project directory"
    ```bash
    cd networkScanner

3. Build the executable:
    ```bash
    go build

4. Run the program
    ```bash
    ./networkScanner

## Output --> new.txt
    --- 192.168.89.1 ping statistics ---
    1 packets transmitted, 1 received, 0% packet loss, time 0ms
    rtt min/avg/max/mdev = 257.064/257.064/257.064/0.000 ms
    PING 192.168.89.2 (192.168.89.2) 56(84) bytes of data.

    --- 192.168.89.2 ping statistics ---
    1 packets transmitted, 0 received, 100% packet loss, time 0ms

    PING 192.168.89.3 (192.168.89.3) 56(84) bytes of data.
    From 192.168.89.173 icmp_seq=1 Destination Host Unreachable

    --- 192.168.89.3 ping statistics ---
    1 packets transmitted, 0 received, +1 errors, 100% packet loss, time 0ms

    PING 192.168.89.4 (192.168.89.4) 56(84) bytes of data.
    From 192.168.89.173 icmp_seq=1 Destination Host Unreachable