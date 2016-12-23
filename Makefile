GO_BUILD_ENV := GOOS=linux GOARCH=amd64

default: buildmac

deps:
	go get github.com/bugsnag/bugsnag-go

bin/fbmessenger-boilerplate-go: *.go
	$(GO_BUILD_ENV) go build -v -o $@ $^

bin/fbmessenger-boilerplate-go-mac: *.go
	go build -race -v -o $@ $^

build: bin/fbmessenger-boilerplate-go

buildmac: bin/fbmessenger-boilerplate-go-mac

runnotify: buildmac
	-killall fbmessenger-boilerplate-go-mac
	-terminal-notifier -title "fbmessenger-boilerplate-go" -message "Built and running!" -remove
	bin/fbmessenger-boilerplate-go-mac

watch:
	supervisor --no-restart-on exit -e go,html -i bin --exec make -- runnotify

clean:
	rm -f bin/*

test:
	go test -v .

run: build init
	bin/fbmessenger-boilerplate-go

heroku: bin/fbmessenger-boilerplate-go
	heroku container:push web
