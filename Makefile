.PHONY: dev prod clean build

dev:
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml up --build

prod:
	docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d --build

clean:
	docker-compose down -v
	docker system prune -f

build:
	docker-compose build
