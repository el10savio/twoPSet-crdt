
provision:
	@echo "Provisioning TwoPSet Cluster"	
	bash scripts/provision.sh

twopset-build:
	@echo "Building TwoPSet Docker Image"	
	docker build -t twopset -f Dockerfile .

twopset-run:
	@echo "Running Single TwoPSet Docker Container"
	docker run -p 8080:8080 -d twopset

info:
	echo "TwoPSet Cluster Nodes"
	docker ps | grep 'twopset'
	docker network ls | grep twopset_network

clean:
	@echo "Cleaning TwoPSet Cluster"
	docker ps -a | awk '$$2 ~ /twopset/ {print $$1}' | xargs -I {} docker rm -f {}
	docker network rm twopset_network

build:
	@echo "Building TwoPSet Server"	
	go build -o bin/twopset main.go

fmt:
	@echo "go fmt TwoPSet Server"	
	go fmt ./...

test:
	@echo "Testing TwoPSet"	
	go test -v --cover ./...