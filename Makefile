.PHONY : documentation entity build-front run-front build-back run-back

documentation:
	go run github.com/swaggo/swag/cmd/swag@latest init
	npm run gen:api --prefix web

entity:
	@echo "Model:"; read MOD; go run entgo.io/ent/cmd/ent init $$MOD

build-front: documentation
	npm run build --prefix web

run-front:
	npm run dev --prefix web

build-back:
	go build .

run-back:
	go run .
