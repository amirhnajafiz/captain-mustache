build:
	go build -o captain-mustache
	chmod +x ./captain-mustache

clear:
	rm -rf build
	rm docker-compose.yaml
	rm .dockerignore
