build:
	go build -o ./_output/heycni -trimpath -ldflags '-s -w' cmd/heycni/main.go
	go build -o ./_output/heyipam -trimpath -ldflags '-s -w' cmd/heyipam/main.go

package:
	docker build -t ghcr.io/hey-io/heycni:v0.1.0 -f ./manifests/Dockerfile .

