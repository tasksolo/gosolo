go := env_var_or_default('GOCMD', 'go')

default: tidy test

tidy:
	{{go}} mod tidy
	goimports -l -w .
	gofumpt -gofumpt -l -w .
	{{go}} fmt ./...

test:
	{{go}} vet ./...
	golangci-lint run ./...

todo:
	-git grep -e TODO --and --not -e ignoretodo

update-client: && default
	curl --silent --output client.go 'https://a45ad4e2.api.solotask.io/v1/_client.go?packageName=gosolo'
