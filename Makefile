NAME := apisrv

GOFLAGS=-mod=vendor

PKG := `go list ${GOFLAGS} -f {{.Dir}} ./...`

ifeq ($(RACE),1)
	GOFLAGS+=-race
endif

LINT_VERSION := v1.54.2

MAIN := cmd/${NAME}/main.go

tools:
	@go install github.com/vmkteam/mfd-generator@latest
	@go install github.com/vmkteam/zenrpc/v2/zenrpc@latest
	@curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin ${LINT_VERSION}

fmt:
	@goimports -local ${NAME} -l -w $(PKG)

lint:
	@golangci-lint run -c .golangci.yml

build:
	@CGO_ENABLED=0 go build $(GOFLAGS) -o ${NAME} $(MAIN)

run:
	@echo "Compiling"
	@go run $(GOFLAGS) $(MAIN) -config=cfg/local.toml -verbose -verbose-sql

generate:
	#@go generate ./pkg/rpc
	@go generate ./pkg/vt

test:
	@echo "Running tests"
	@go test -count=1 $(GOFLAGS) -coverprofile=coverage.txt -covermode count $(PKG)

test-short:
	@go test $(GOFLAGS) -v -test.short -test.run="Test[^D][^B]" -coverprofile=coverage.txt -covermode count $(PKG)

mod:
	@go mod tidy
	@go mod vendor
	@git add vendor

NS := "NONE"

MAPPING := "common:users;vfs:vfsFiles,vfsFolders"

mfd-xml:
	@mfd-generator xml -c "postgres://postgres:postgres@localhost:5432/apisrv?sslmode=disable" -m ./docs/model/newsportal.mfd -n $(MAPPING)
mfd-model:
	@mfd-generator model -m ./docs/model/newsportal.mfd -p db -o ./pkg/db
mfd-repo: --check-ns
	@mfd-generator repo -m ./docs/model/newsportal.mfd -p db -o ./pkg/db -n $(NS)
mfd-vt-xml:
	@mfd-generator xml-vt -m ./docs/model/newsportal.mfd
mfd-vt-rpc: --check-ns
	@mfd-generator vt -m docs/model/newsportal.mfd -o pkg/vt -p vt -x apisrv/pkg/db -n $(NS)
mfd-xml-lang:
	#TODO: add namespaces support for xml-lang command
	@mfd-generator xml-lang  -m ./docs/model/newsportal.mfd
mfd-vt-template: --check-ns
	@mfd-generator template -m docs/model/newsportal.mfd  -o ../gold-vt/ -n $(NS)

type-script-client: generate
	@go run $(GOFLAGS) $(MAIN) -config=cfg/local.toml -ts_client > ../gold-vt/src/services/api/factory.ts


--check-ns:
ifeq ($(NS),"NONE")
	$(error "You need to set NS variable before run this command. For example: NS=common make $(MAKECMDGOALS) or: make $(MAKECMDGOALS) NS=common")
endif
