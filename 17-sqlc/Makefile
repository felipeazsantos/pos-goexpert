createmigration:
	migrate create -ext sql -dir sql/migrations -seq init_schema

migrateup:
	migrate -path sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" up

migratedown:
	migrate -path sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" down

.PHONY: createmigration migrateup migratedown
# .PHONY is used to indicate that these targets do not represent files, but rather commands to be executed.