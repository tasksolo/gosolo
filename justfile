go := env_var_or_default('GOCMD', 'go')

default: tidy test

tidy:
	{{go}} mod tidy
	goimports -l -w .
	gofumpt -l -w .
	{{go}} fmt ./...

test:
	{{go}} vet ./...
	golangci-lint run ./...
	{{go}} test -race -coverprofile=cover.out -timeout=45s -parallel=10 ./...
	{{go}} tool cover -html=cover.out -o=cover.html

todo:
	git grep -e TODO --and --not -e ignoretodo

update-client: && default
	curl --silent --output client.go 'https://api.solø.com/v1/_goclient?packageName=gosolo'
