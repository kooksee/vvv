.PHONY: test

s:
	open  http://localhost:8080/github.com/kooksee/vvv/example/todomvc
#	http://localhost:8080/github.com/kooksee/vvv/example/todomvc/todomvc.js
	gopherjs serve .

init:
	go get -u github.com/go-bindata/go-bindata

test:
	@echo "test"
	go test -v .

w:
	gopherjs build --watch github.com/kooksee/go-hackernews --output build/script.min.js --minify

build:
	gopherjs build --output build/script.min.js --minify


gopherjs:
	gopherjs serve github.com/nobonobo/vecty-chatapp/app

devrun:
	go run main.go -dev

assets:
	mkdir -p ./dist/app/assets
	gopherjs build -m -o ./dist/app/app.js github.com/nobonobo/vecty-chatapp/app
	cp -f ./app/index.html ./dist/app/
	cp -Rf ./app/assets/* ./dist/app/assets/

build: assets
	go build -o ./dist/server main.go
