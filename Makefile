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

test: clean critic security lint
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

build: test
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

run: swag build
	$(BUILD_DIR)/$(APP_NAME)

docker.run: docker.network docker.mysql swag docker.fiber docker.redis

docker.network:
	docker network inspect dev-network >/dev/null 2>&1 || \
	docker network create -d bridge dev-network

docker.fiber.build:
	docker build -t fiber .

docker.fiber: docker.fiber.build
	docker run --rm -d \
		--name AIM_Coding_Test-fiber \
		--network dev-network \
		-p 3000:3000 \
		fiber

docker.mysql:
	docker run --rm -d \
		--name AIM_Coding_Test-mysql \
		--network dev-network \
		-e MYSQL_USER_USER=dev \
		-e MYSQL_PASSWORD= \
		-e MYSQL_DB=AIM_Coding_Test_DB \
		-p 3306:3306 \
		mysql

docker.redis:
	docker run --rm -d \
		--name AIM_Coding_Test-redis \
		--network dev-network \
		-p 6379:6379 \
		redis

docker.stop: docker.stop.fiber docker.stop.mysql docker.stop.redis

docker.stop.fiber:
	docker stop AIM_Coding_Test-fiber

docker.stop.mysql:
	docker stop AIM_Coding_Test-mysql

docker.stop.redis:
	docker stop AIM_Coding_Test-redis

swag:
	swag init
