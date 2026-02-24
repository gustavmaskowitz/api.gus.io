build:
	cp $$(go env GOROOT)/lib/wasm/wasm_exec.js web/
	GOOS=js GOARCH=wasm go build -o web/main.wasm web/main.go
