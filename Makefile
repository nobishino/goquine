all:
	go fmt ./...
	find ./*/main.go -type f | xargs -I{} wc -c {}
	find ./unofficial/*/main.go -type f | xargs -I{} wc -c {}
	find ./*/main.go -type f | xargs -I{} sh -c 'echo checking {} && go run {} > out && diff {} out'
	find ./unofficial/*/main.go -type f | xargs -I{} sh -c 'echo checking {} && go run {} > out && diff {} out'
	find ./unofficial/stderr/*/main.go -type f | xargs -I{} wc -c {}
	find ./unofficial/stderr/*/main.go -type f | xargs -I{} sh -c 'echo checking {} && go run {} >& out && diff {} out'

clean:
	go clean
	rm -f out