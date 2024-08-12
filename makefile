build-zip-view:
	GOOS=linux GOARCH=arm64 go build -o bootstrap ./cmd/view/main.go
	zip view-lambda-handler.zip bootstrap

build-zip-upload:
	GOOS=linux GOARCH=arm64 go build -o bootstrap ./cmd/upload/main.go
	zip upload-lambda-handler.zip bootstrap

build-zip: build-zip-view build-zip-upload