APP_NAME = heroku-practice
GO_BUILD_ENV := CGO_ENABLED=0 GOOS=darwin GOARCH=amd64

build:
	$(GO_BUILD_ENV) go build -v -o $(APP_NAME) .

clean:
	rm -rf ./$(APP_NAME)