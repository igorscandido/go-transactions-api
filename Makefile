up:
	docker-compose -f build/docker-compose.yaml up -d

down:
	docker-compose -f build/docker-compose.yaml down

tests:
	docker-compose -f build/docker-compose.yaml up -d > /dev/null 2>&1 &&\
	docker exec -it transactions-api go test -v -timeout 30s ./...

logs:
	docker logs -f transactions-api
