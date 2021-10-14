GOROOT := $(shell go env GOROOT)

run: assets/app.wasm assets/wasm_exec.js
	go run ./

assets/app.wasm: ./app
	mkdir -p ./server/assets
	GOARCH=wasm GOOS=js go build -o assets/app.wasm ./app/

assets/wasm_exec.js: $(GOROOT)/misc/wasm/wasm_exec.js
	cp "$(GOROOT)/misc/wasm/wasm_exec.js" assets/wasm_exec.js
