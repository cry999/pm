FROM golang:1.15.6-alpine

ARG MIGRATE_VER=4.12.2
WORKDIR /app
RUN apk add --no-cache \
    # required by build-phase
    git \
    gcc \
    g++ \
    # required by dev-phase
    curl \
    gomplate \
    jq \
    openssl \
    && set -x \
    && export GO111MODULE=off \
    # embedding static files tool
    && go get -v github.com/markbates/pkger/cmd/pkger \
    # migrate tool
    && curl -L https://github.com/golang-migrate/migrate/releases/download/v${MIGRATE_VER}/migrate.linux-amd64.tar.gz | tar xvz \
    && mv migrate.linux-amd64 /usr/local/bin/migrate \
    # orm generator
    && go get -u -t github.com/volatiletech/sqlboiler \
    && go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql \
    # auto-reload server
    && go get -u -v github.com/cosmtrek/air \
    # development tools
    && go get -u -v golang.org/x/tools/gopls \
    && go get -u -v github.com/mdempsky/gocode \
    && go get -u -v github.com/uudashr/gopkgs/v2/cmd/gopkgs \
    && go get -u -v github.com/ramya-rao-a/go-outline \
    && go get -u -v github.com/stamblerre/gocode \
    && go get -u -v github.com/rogpeppe/godef \
    && go get -u -v golang.org/x/tools/cmd/goimports \
    && go get -u -v golang.org/x/lint/golint \
    && go get -u -v github.com/cweill/gotests \
    && go get -u -v github.com/fatih/gomodifytags \
    && go get -u -v github.com/godoctor/godoctor \
    # mock generator
    && GO111MODULE=on go get github.com/golang/mock/mockgen@v1.4.4

ADD . /app
RUN go build -o $GOPATH/bin/watcher ./cmd/tools/watcher

ENTRYPOINT [ "watcher", "-config", "configs/watcher.config.json" ]
