ATLAS_DATABASE_URL=postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@db:5432/$(POSTGRES_DB)?sslmode=disable

migrate-diff:
	@echo "Generating migration files..."
	atlas schema inspect --url "$(ATLAS_DATABASE_URL)" | \
	atlas migrate diff --dir "file://infrastructure/ent/migrations" --dev-url "$(ATLAS_DATABASE_URL)" --to "file://infrastructure/ent/schema/"

migrate-apply:
	@echo "Applying migrations to database..."
	atlas migrate apply --dir "file://infrastructure/ent/migrations" --url "$(ATLAS_DATABASE_URL)"

diff:
	@echo "Checking schema diff for $(ATLAS_DATABASE_URL)"
	atlas schema diff --url "$(ATLAS_DATABASE_URL)" --to "file://infrastructure/ent/schema/"