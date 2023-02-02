run: assets/app.wasm assets/wasm_exec.js
	go run ./

assets/app.wasm: ./app
	mkdir -p ./server/assets
	tinygo build -o assets/app.wasm -target wasm ./app/

assets/wasm_exec.js:
	cp "$$(tinygo env TINYGOROOT)/targets/wasm_exec.js" assets/wasm_exec.js
