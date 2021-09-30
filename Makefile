documentation:
	go run github.com/swaggo/swag/cmd/swag@latest init
	npm run gen:api --prefix web

new-ent-model:
	@echo "Model:"; read MOD; go run entgo.io/ent/cmd/ent init $$MOD
