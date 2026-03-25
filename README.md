# Proxy Server

## Description

A high-performance, concurrent HTTP proxy server written in Go, featuring a possible extensible caching layer

## Overview

It acts as an intermediary for requests from clients seeking resources from other servers.

    Full HTTP Proxying: Transparently forwards GET, POST, and other HTTP methods.

    Caching Engine (Planned): In-memory storage to reduce latency for frequent requests.

    Concurrency: Leverages Go routines for handling multiple simultaneous connections.


## Project Structure

This project follows the standard Go project layout to ensure clean separation between the entry point and the business logic.

```bash
proxy-server
├── main.go              # Application entry point 
├── proxy/
│   ├── server/          # Core proxy logic
│   │   ├── proxy.go
│   │   └── proxy_test.go # Unit tests for forwarding logic
│   └── cache/           # Caching implementation
│       ├── cache.go
│       └── cache_test.go # Unit tests for TTL and eviction
├── go.mod               # Module definition
└── README.md
```

## Installation & Usage

### Prerequisites
Go 1.21+

### Setup

1. Clone the repository:

```bash
git clone https://github.com/jdmwihia/proxy-server.git
cd proxy-server
```

2. Initialize dependencies:

```bash
go mod download
```

3. Run the Server

```bash
go run .
```
