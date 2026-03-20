REPO=https://github.com/IdzAnAG1/mapps-contracts.git#branch=main

buf_gen:
	buf generate $(REPO) --template buf.gen.yaml --path proto/products/v1

local:
	go run cmd/main/main.go
