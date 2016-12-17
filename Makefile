GO_BUILD_ENV := GOOS=linux GOARCH=amd64

default: buildmac

deps:
	go get github.com/bugsnag/bugsnag-go

bin/spotibotfb: *.go
	$(GO_BUILD_ENV) go build -v -o $@ $^

bin/spotibotfb-mac: *.go
	go build -race -v -o $@ $^

build: bin/spotibotfb

buildmac: bin/spotibotfb-mac

runnotify: buildmac
	-killall spotibotfb-mac
	-terminal-notifier -title "spotibotfb" -message "Built and running!" -remove
	bin/spotibotfb-mac

watch:
	supervisor --no-restart-on exit -e go,html -i bin --exec make -- runnotify

clean:
	rm -f bin/*

test:
	go test -v .

run: build init
	bin/spotibotfb

heroku: bin/spotibotfb
	heroku container:push web
