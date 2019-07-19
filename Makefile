IMAGE_VERSION = 1.0.0
TAG           = kodam/papertrail-exporter

.PHONY: test
test:
	go test

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build papertrail_exporter.go
	docker build --no-cache -t $(TAG) .

.PHONY: publish
publish:
	docker tag ${TAG}:${IMAGE_VERSION} ${TAG}:latest 
	docker push ${TAG}:${IMAGE_VERSION}
	docker push ${TAG}:latest 
