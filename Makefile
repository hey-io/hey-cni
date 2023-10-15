build:
	go build -o ./_output/hey-cni -trimpath -ldflags '-s -w' cmd/hey-cni/main.go