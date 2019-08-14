# NGINX Reverse Proxy by Request Body

## Description

The forwarding is determined by the `server` field found in the request body.

The proxy server is implemented using [OpenResty NGINX Docker image](https://hub.docker.com/r/openresty/openresty).  This version is used because it contains the NGINX lua module needed for reading the request body.

The web servers under the proxy are Docker containers with a simple server written in GO which responds to any request with the contents of the "WHO" environment variable and the request body.
The binary found in the repository is precompiled for Linux(needed for the Docker image). 

The Docker Compose file starts the proxy server on port 8080, with the `nginx.conf` mounted. It also starts two web servers and sets the "WHO" environment variable to "server_1" and "server_2" respectively. Making use of the Docker internal network, we do not need to expose/map any ports to the host, the names of the containers act as host names.

### Prerequisites

```
docker
docker-compose
curl
```

### Run

#### Spin up the containers

```
docker-compose up
```

#### Send requests:

Targeting Server 1:
```
curl http://localhost:8080 -d '{"server":"1"}'
```
Response: 
```
Hello from server_1
received body: {"server":"1"}
```

Targeting Server 2:
```
curl http://localhost:8080 -d '{"server":"2"}'
```
Response: 
```
Hello from server_2
received body: {"server":"2"}
```


