.DEFAULT_GOAL := all
m?="commit message not given"

all:
	@GOTMPDIR=/var/tmp go run main.go

upd:
	@go1.26.1 get -u github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs

git:
	@git diff-index --quiet HEAD || git add . && git commit -m $(m)
	@git tag -a "$(v)"
	@git push

win:
	@GOTMPDIR=/var/tmp GOOS=windows GOARCH=amd64 go1.26.1 build -buildvcs=true -o build/tv-listen-voice.exe
	@tar czf tv-listen-voice.tgz -C build tv-listen-voice.exe 
