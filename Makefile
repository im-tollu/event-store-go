define enforce-db-url
	$(if $(value ESTORE_DB_URL),,$(error "ESTORE_DB_URL is not set!"))
endef

.PHONY: test
test: out/gotestsum out/test-reports
	@echo "\n>>> Running tests..."
	$(call enforce-db-url)
	set -euo pipefail && \
	cd src && \
	go test ./... | ../out/gotestsum --junitfile ../out/test-reports/unit-tests.xml

.PHONY: dbclean
dbclean: migrate
	@echo "\n>>> Cleaning database and applying migrations from scratch..."
	$(call enforce-db-url)
	out/migrate -clean -db=$(ESTORE_DB_URL) -migrations=out/migrations

.PHONY: clean
clean:
	rm -r out

.PHONY: migrate
migrate: out/migrate out/migrations/*.sql
	@echo "\n>>> Preparing database migration..."

out/migrate: out
	@echo "\n>>> Building migration tool..."
	cd src && \
	go build -o ../out/migrate cmd/migrate/main.go

out/migrations/*.sql: out/migrations
	@echo "\n>>> Copying migration scripts..."
	cp src/migrations/*.sql out/migrations

out/migrations: out
	@echo "\n>>> Preparing directory for migration scripts..."
	mkdir out/migrations

out:
	@echo "\n>>> Preparing output directory..."
	mkdir out

out/gotestsum: out
	@echo "\n>>> Building gotestsum..."
	cd src && \
	go build -o ../out/gotestsum gotest.tools/gotestsum

out/test-reports: out
	@echo "\n>>> Preparing test-reports directory..."
	mkdir out/test-reports
