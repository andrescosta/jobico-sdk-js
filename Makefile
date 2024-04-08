.PHONY: install

# https://github.com/mozilla-spidermonkey/sm-wasi-demo/blob/main/data.json

install:
	docker exec jobico-control-plane mkdir -p /data/volumes/pv1/js/prg
	docker exec jobico-control-plane curl -o /data/volumes/pv1/js/js.wasm -LJ https://firefoxci.taskcluster-artifacts.net/OWBTW-rCTqGtqtvO4rpJkA/0/public/build/js.wasm
	docker cp sdk/ jobico-control-plane:/data/volumes/pv1/js/prg

install-file:
	docker cp test.js jobico-control-plane:/data/volumes/pv1/js/prg

install-job:
	kubectl apply -f k8s-greet.yml