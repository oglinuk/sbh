output:
	go build

docker: build
	docker run -it --name go-sbh-docker go-sbh:0.1.0

docker-web: build
	docker run --name go-sbh-docker go-sbh:0.1.0 /bin/bash -c "./go-sbh -w"

build:
	docker build . -t go-sbh:0.1.0

clean:
	rm go-sbh

docker-clean:
	docker stop go-sbh-docker
	docker rm go-sbh-docker
