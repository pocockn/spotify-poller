PACKAGES ?= "./..."
DOCKERNAME = "pococknick91"
REPONAME ?= "spotify-poller"
IMG ?= ${DOCKERNAME}/${REPONAME}:${VERSION}
LATEST ?= ${DOCKERNAME}/${REPONAME}:latest
SPOTIFY_CLIENT_SECRET ?= ""
VERSION = $(shell cat ./VERSION)


DEFAULT: test

build:
	@GO111MODULE=on go build -race -o $(REPONAME) ./cmd/$(REPONAME)/main.go

build-image:
	@docker build -t ${IMG} -f builds/Dockerfile .

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BINARY_UNIX) -v ./cmd/$(REPONAME)/main.go

install:
	@echo "=> Install dependencies"
	@GO111MODULE=on go mod download

push-to-registry:
	@docker login -u ${DOCKER_USER} -p ${DOCKER_PASS}
	@docker build -t ${IMG} . -f builds/Dockerfile .
	@docker tag ${IMG} ${LATEST}
	echo "=> Pushing ${IMG} & ${LATEST} to docker"
	@docker push ${DOCKERNAME}/${REPONAME}

run:
	@go build -ldflags "-X main.Version=dev" -o ${REPONAME} ./cmd/$(REPONAME)/main.go
	@ENV=development SPOTIFY_CLIENT_SECRET=${SPOTIFY_CLIENT_SECRET} AWS_REGION=eu-west-1 ./${REPONAME}

test:
	@GO111MODULE=on go test "${PACKAGES}" -cover

vet:
	@@GO111MODULE=on go vet "${PACKAGES}"