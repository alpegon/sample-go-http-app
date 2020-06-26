# sample-go-http-app
## Overview
Sample application which can be used to test various scenarios in Kubernetes or OpenShift.

## Running
The image can be retrieved from dockerhub:

https://hub.docker.com/r/serbangilvitu/sample-go-http-app

```
docker run -p 8081:8081 serbangilvitu/sample-go-http-app:latest
```

## Endpoints

### /health
Returns 'OK'

### /ip
Returns the IP of the requester

### /version
Returns the value stored in VERSION, using the colour stored in COLOUR

This can be useful in simulating deploying 2 different versions.

## Configuration
Configuration can be done via the following environment variables, which all have defaults in the Dockerfile.

### COLOUR
The colour used for /version

### LOAD_TIME
Time to wait before service requests

### RESPONSE_TIME
Time to wait before returning a response

### PORT
Port to listen on

### VERSION
The text returned in /version