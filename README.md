# Go Port Scanner

A simple, high-performance port scanner written in Go. This tool demonstrates the power of Go's concurrency model by using Goroutines to scan multiple network ports simultaneously.

## Features

- **Concurrent Scanning:** Scans multiple ports in parallel using Goroutines for maximum speed.
- **Configurable Targets:** Easily change the target host and port range.
- **Fast Execution:** Built-in timeout handling ensures the scan doesn't hang on closed ports.

## Requirements

- [Go](https://go.dev/dl/) installed on your machine.

## Getting Started

### 1. Installation

Ensure Go is installed by checking the version in your terminal:

```bash
go version
```

### 2. Setup

Clone the repository and navigate into the project directory:

```bash
git clone https://github.com/papinashvilisergi/go-port-scanner.git
cd go-port-scanner
```

### 3. Running the Scanner

Open your terminal in the project folder and run:

```bash
go run main.go
```

You should see output like:

```
Port 22 is open!
Port 80 is open!
Port 443 is open!
```

## How It Works

- **Goroutines (`go scanPort`):** Each port scan runs in its own lightweight thread, allowing the program to scan hundreds of ports in the time it would take to scan just one sequentially.
- **Channels (`chan int`):** Results are sent through a channel, allowing the main function to collect them as soon as they are ready — a safe way to communicate between concurrent Goroutines.
- **Networking (`net.DialTimeout`):** Attempts a TCP connection to the target on each port. If a connection is established within 2 seconds, the port is considered open. If it times out or is refused, the port is considered closed and a `0` is sent back through the channel.

## Customization

- **Change Target:** Modify the `address` variable inside `scanPort` to point to a different host. For example, use `127.0.0.1` to scan your own machine:
  ```go
  address := fmt.Sprintf("127.0.0.1:%d", port)
  ```
- **Change Port Range:** Adjust the loop limit in `main` to scan more or fewer ports:
  ```go
  for i := 1; i <= 65535; i++ { // scan all ports
  ```
- **Change Timeout:** Increase or decrease the `2*time.Second` timeout in `net.DialTimeout` depending on your network speed and desired accuracy.

## Legal & Ethical Notice

> **Only scan hosts you own or have explicit permission to scan.**
> Unauthorized port scanning may be illegal in your jurisdiction and against the terms of service of network providers.
> The default target, `scanme.nmap.org`, is provided by the Nmap project specifically for testing purposes.

## License

This project is open source and available under the [MIT License](https://opensource.org/licenses/MIT).