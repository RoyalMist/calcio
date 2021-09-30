.PHONY: documentation ent dev dev-front dev-back build run

documentation:
	go install github.com/swaggo/swag/cmd/swag@latest
	go run github.com/swaggo/swag/cmd/swag@latest init
	npm run gen:api --prefix web

ent:
	go install entgo.io/ent/cmd/ent@latest
	@echo "New ent model:"; read MOD; go run entgo.io/ent/cmd/ent@latest init $$MOD

dev-front:
	npm run dev --prefix web

dev-back:
	go run .

dev: documentation
	make -j 2 dev-front dev-back

build: documentation
	npm run build --prefix web
	go build .

quality:
	go vet ./...
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	go run github.com/securego/gosec/v2/cmd/gosec@latest -exclude-dir ent ./...
	go test -v ./...
