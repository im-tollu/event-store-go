define enforce-db-url
	$(if $(value ESTORE_DB_URL),,$(error "ESTORE_DB_URL is not set!"))
endef

.PHONY: test
test: out/gotestsum
	@echo "\n>>> Running tests..."
	mkdir -p out/test-reports
	$(call enforce-db-url)
	set -eu pipefail && \
	cd src && \
	go test ./... | ../out/gotestsum --junitfile ../out/test-reports/unit-tests.xml

.PHONY: dbclean
dbclean: migrate
	@echo "\n>>> Cleaning database and applying migrations from scratch..."
	$(call enforce-db-url)
	out/migrate -clean -db=$(ESTORE_DB_URL) -migrations=out/migrations

.PHONY: clean
clean:
	@echo "\n>>> Cleaning build artifacts..."
	rm -r out

.PHONY: migrate
migrate: out/migrate out/migrations/*.sql
	@echo "\n>>> Preparing database migration..."

out/.f:
	@echo "\n>>> Preparing output directory..."
	mkdir out
	touch out/.f

out/migrate: out/.f
	@echo "\n>>> Building migration tool..."
	cd src && \
	go build -o ../out/migrate cmd/migrate/main.go

out/migrations/*.sql:
	@echo "\n>>> Copying migration scripts..."
	mkdir -p out/migrations
	cp src/migrations/*.sql out/migrations/

out/gotestsum: out/.f
	@echo "\n>>> Building gotestsum..."
	cd src && \
	go build -o ../out/gotestsum gotest.tools/gotestsum

