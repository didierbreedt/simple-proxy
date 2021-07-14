This extremely simple Golang app allows one to easily expose web data via a proxy. 
Useful for i.e. a Kubernetes environment that uses annotations for Prometheus 
service discovery on resources not in the Kubernetes environment 
(i.e. AWS managed services with JMX metric endpoints)

``
docker run -e PROXY_TARGET=http://example.org/ -e HOST_PORT=8088 -p 8088:8088 didierbreedthyve/simple-proxy
``

You can then request this url via `http://localhost:8088/`.