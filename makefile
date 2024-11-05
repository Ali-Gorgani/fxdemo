generate_initialization:
	@[ -n "$(name)" ] || (echo "Error: name variable is required. Usage: make generate_initialization name=User"; exit 1)
	go run -mod=mod entgo.io/ent/cmd/ent new "$(name)"

generate_sql_migration:
	@[ -n "$(name)" ] || (echo "Error: name variable is required. Usage: make generate_migration name=my_custom_migration"; exit 1)
	atlas migrate diff "$(name)" \
		--dir "file://ent/migrate/migrations" \
		--to "ent://ent/schema" \
		--dev-url "docker://postgres/15/test?search_path=public"

generate_go_migration:
	go generate ./ent

apply_migration:
	atlas migrate apply \
  		--dir "file://ent/migrate/migrations" \
  		--url "postgres://root:secret@localhost:5432/psql_db?search_path=public&sslmode=disable"

status_migration:
	atlas migrate status \
		--dir "file://ent/migrate/migrations" \
		--url "postgres://root:secret@localhost:5432/psql_db?search_path=public&sslmode=disable"

rollback_migration:
	atlas migrate down \
		--dir "file://ent/migrate/migrations" \
		--url "postgres://root:secret@localhost:5432/psql_db?search_path=public&sslmode=disable" \
		--dev-url "docker://postgres/15/test?search_path=public"

.PHONY: generate_go generate_migration apply_migration status_migration rollback_migration
