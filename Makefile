# start/stop docker container by docker-compose
up:
	docker-compose down
	docker-compose up
down:
	docker-compose down


# run shell in container
.PHONY: backend
backend:
	docker exec -it backend-api sh

.PHONY: test
test:
	docker exec -it backend-api go test -run ''

.PHONY: lint
lint:
	docker exec --tty backend-api golangci-lint run ./...

.PHONY: dog
dog:
	docker exec --tty backend-api golint $(go list ./...) | reviewdog -name=golangci-lint -efm -diff='git diff master'




# golangci-lint run ./... | reviewdog -efm="%E%f:%l: %m" -diff="git diff master"

# golangci-lint run --out-format=checkstyle ./... | reviewdog -f=golangci -diff="git diff master"
# golangci-lint run --out-format=checkstyle ./... | reviewdog -f=checkstyle -diff="git diff"


# '%E%f:%l: %m'