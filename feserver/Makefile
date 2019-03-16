
.PHONY: build
build:
	GOARCH=amd64 GOOS=linux go build -o feserver

.PHONY: deploy
deploy: build
	scp -i $(KEYPATH) feserver $(DEST)

