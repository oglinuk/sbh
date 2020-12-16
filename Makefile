output:
	go build

docker: build
	docker run -it --name sbh-docker sbh:0.1.0

docker-web: build
	docker run --name sbh-docker sbh:0.1.0 /bin/bash -c "./sbh -w"

build:
	docker build . -t sbh:0.1.0

clean:
	rm sbh

docker-clean:
	docker stop sbh-docker
	docker rm sbh-docker
