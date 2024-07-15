cek:
	go version
fe: 
	cd fe-bapsi
	yarn serve

be:
	cd be-bapsi
	go run .
migrate:
	cd be-bapsi
	go run migration/migrate.go
	go run migration/seed.go 

bingung:
	cd fe-bapsi
	git fetch
	git pull --all

	cd ../be-bapsi
	git fetch
	git pull --all