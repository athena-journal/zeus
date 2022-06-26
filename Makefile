.PHONY: install build serve clean pack deploy ship

TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)

export TAG

install:
	go get .

build: install
	go build -ldflags "-X main.version=$(TAG)" -o zeus .

serve: build
	./zeus

clean:
	rm ./zeus

pack:
	GOOS=linux make build
	docker build -t athena/zeus:$(TAG) .

deploy:
	envsubst < k8s/deployment.yaml | kubectl apply -f -

ship: pack deploy clean