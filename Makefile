.PHONY: clean critic security lint test build run

APP_NAME = apiserver
BUILD_DIR = $(PWD)/build
MIGRATIONS_FOLDER = $(PWD)/platform/migrations
DATABASE_URL = mysql://AIM_Coding_Test-mysql/AIM_Coding_Test_DB?serverTimezone=UTC

clean:
	rm -rf ./build

critic:
	gocritic check -enableAll ./...

security:
	gosec ./...

lint:
	golangci-lint run ./...


docker.run: swag



docker.stop: docker.stop.fiber docker.stop.mysql docker.stop.redis

docker.stop.fiber:
	docker stop AIM_Coding_Test-fiber

docker.stop.mysql:
	docker stop AIM_Coding_Test-mysql

docker.stop.redis:
	docker stop AIM_Coding_Test-redis

swag:
	swag init
