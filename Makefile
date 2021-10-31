


arm:
	env GOOS=linux GOARCH=arm64 go build -o battlesnake_arm64

intel:
	go build
