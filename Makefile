SRC := $(CURDIR)/src
OUT := $(CURDIR)/out

define enforce-db-url
	$(if $(value DB_URL),,$(error "DB_URL is not set!"))
endef

.PHONY: test
test:
	@echo "\n>>> Running tests..."
	$(call enforce-db-url)
	cd $(SRC) && \
	go test ./... -args -db=$(DB_URL)

.PHONY: dbclean
dbclean: migrate
	@echo "\n>>> Cleaning database and applying migrations from scratch..."
	$(call enforce-db-url)
	cd $(OUT) && ./migrate -clean -db=$(DB_URL)

.PHONY: clean
clean:
	rm $(OUT)

.PHONY: migrate
migrate: $(OUT)/migrate $(OUT)/migrations/*.sql
	@echo "\n>>> Preparing database migration..."

$(OUT)/migrate: $(OUT)
	@echo "\n>>> Building migration tool..."
	cd $(SRC) && \
	go build -o $(OUT)/migrate cmd/migrate/main.go

$(OUT)/migrations/*.sql: $(OUT)/migrations
	@echo "\n>>> Copying migration scripts..."
	cp $(SRC)/migrations/*.sql \
	$(OUT)/migrations

$(OUT)/migrations: $(OUT)
	@echo "\n>>> Preparing directory for migration scripts..."
	mkdir $(OUT)/migrations

$(OUT):
	@echo "\n>>> Preparing output directory..."
	mkdir $(OUT)