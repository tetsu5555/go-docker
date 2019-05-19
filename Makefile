# start/stop docker container by docker-compose
up:
	docker-compose up
down:
	docker-compose down


# run shell in container
.PHONY: backend
backend:
	docker exec -it backend-api sh