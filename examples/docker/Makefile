cli: build
	docker run -it --name sbh-docker sbh:0.1.0 /bin/bash -c "./sbh"

web: build
	docker run --name sbh-docker sbh:0.1.0 /bin/bash -c "./sbh -w"

build:
	docker build . -t sbh:0.1.0

clean:
	docker stop sbh-docker
	docker rm sbh-docker
