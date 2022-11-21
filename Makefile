all:
	go fmt ./...
	find ./*/main.go -type f | xargs -I{} wc -c {}
	find ./*/main.go -type f | xargs -I{} sh -c 'echo checking {} && go run {} > out && diff {} out'

clean:
	go clean
	rm -f out