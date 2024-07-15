setup-air:
	# go get github.com/air-verse/air
	export PATH=$PATH:$(go env GOPATH)/bin
	air

seed:
	go run migration/migrate.go
	go run migration/seed.go

migrate:
	go run migration/migrate.go

run:
	go run .

remove:
	rm -r database.db