# ---------- help ----------

.PHONY: help
help: ## display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ {printf "  \033[36m%-24s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)


# ---------- compose ----------

.PHONY: compose-up
compose-up: ## up docker compose
	docker compose up -d --build
	make app-logs

.PHONY: compose-down
compose-down: ## down docker compose
	docker compose down

.PHONY: app-logs
app-logs: ## show app container logs
	docker logs -f zlatoivan_ru-app-1

# ---------- local ----------

.PHONY: compose-up-local
compose-up-local: ## up docker compose local
	docker compose up -d

.PHONY: compose-down-local
compose-down-local: ## down docker compose local
	docker compose down

.PHONY: run
run: ## run local
	CONFIG_PATH=config/values_local.yaml go run cmd/server/main.go

.PHONY: run-test
run-test: ## run local for tests
	CONFIG_PATH=config/config_test.local.yaml go run cmd/main.go

..PHONY: build
build: ## build local
	go build -o main cmd/main.go


# ---------- migrations ----------

POSTGRES_SETUP := user=postgres password=postgres dbname=postgres host=localhost port=5433 sslmode=disable
POSTGRES_SETUP_TEST := user=postgres password=postgres dbname=test host=localhost port=5431 sslmode=disable
MIGRATION_FOLDER=$(CURDIR)/migrations
MIGRATION_NAME=pvz

.PHONY: migration-create
migration-create: ## create new migration
	goose -dir "$(MIGRATION_FOLDER)" create "$(MIGRATION_NAME)" sql

.PHONY: migration-up
migration-up: ## migration up
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP)" up

.PHONY: migration-down
migration-down: ## migration down
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP)" down

.PHONY: migration-up-test
migration-up-test: ## migration up test
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" up

.PHONY: migration-down-test
migration-down-test: ## migration down test
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" down


# ---------- tests ----------

GO_COVER_EXCLUDE:="mock|.*.pb(..*)?.go|config|main.go|postgres.go"

.PHONY: t
t: ## run tests
	go test ./... -v -count=2 -p=3 -race -coverprofile=cover.out.tmp $(RUN_ARGS)

.PHONY: t-integration
t-integration: ## run integration tests
	CONFIG_PATH=$(CURDIR)/config/config_test.yaml go test ./... -tags=integration -v

.PHONY: t-colored
t-colored: ## run tests colored
	grc go test -v -count=2 -p=3 ./...

.PHONY: t-integration-colored
t-integration-colored: ## run integration tests colored
	CONFIG_PATH=$(CURDIR)/config/config_test.yaml grc go test ./... -tags=integration -v

.PHONY: t-cov ## get tests coverage
t-cov:
	make t
	grep -vE ${GO_COVER_EXCLUDE} cover.out.tmp > cover.out || cp cover.out.tmp cover.out
	go tool cover -func=cover.out | sort -k 3 -n
	go tool cover -html=cover.out
	rm cover.out.tmp
	rm cover.out


# ---------- rare ----------

.PHONY: generate-proto
generate-proto: ## generate proto
	rm -rf internal/pkg swagger && \
	mkdir -p internal/pkg swagger
	protoc  --proto_path=internal/api_proto \
			--go_out=internal/pkg \
			--go-grpc_out=internal/pkg \
			--grpc-gateway_out=internal/pkg --grpc-gateway_opt generate_unbound_methods=true \
			--openapiv2_out=swagger \
 			api.proto

.PHONY: lint
lint: ## check by golangci linter
	golangci-lint run

.PHONY: docker-rm-volume
docker-rm-volume: ## remove docker volume
	docker volume rm pg-data

.PHONY: gen-ssl-cert
gen-ssl-cert: ## generate fresh ssl certificate
	openssl genrsa -out server.key 2048  # Сгенерировать приватный ключ (.key)
	openssl req -new -x509 -sha256 -key server.key -out server.crt -days 365 -nodes  # Сгенерировать публичный ключ (.crt), но основе приватного
	mv -f server.key server.crt internal/server/certs/  # Поместить оба файла в папку /certificate


# ---------- ssh ----------

.PHONY: ssh
ssh: ## connect to ssh
	ssh root@147.45.106.47

.PHONY: new-ssh
new-ssh: ## connect to ssh with password
	ssh-keygen -R 147.45.106.47 && \
	ssh-copy-id root@147.45.106.47 && \
	ssh root@147.45.106.47