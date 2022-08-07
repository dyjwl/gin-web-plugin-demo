
.PHONY: all
all: lint test build

# ==============================================================================
# Build Options

ROOT_PACKAGE=$(CURDIR)
VERSION_PACKAGE=$(ROOT_PACKAGE)/pkg/app/version

# ==============================================================================
# Includes

include scripts/make-rules/common.mk
include scripts/make-rules/golang.mk
include scripts/make-rules/image.mk

# ==============================================================================
# Targets

## build: Build source code for host platform.
.PHONY: build
build:
	@$(MAKE) go.build

## build.all: Build source code for all platforms.
.PHONY: build.all
build.all:
	@$(MAKE) go.build.all
	
## image: Build docker images and push to registry.
.PHONY: image
image:
	@$(MAKE) image.push

## clean: Remove all files that are created by building.
.PHONY: clean
clean:
	@$(MAKE) go.clean

## lint: Check syntax and styling of go sources.
.PHONY: lint
lint:
	@$(MAKE) go.lint

## test: Run unit test.
.PHONY: test
test:
	@$(MAKE) go.test

## help: Show this help info.
.PHONY: help
help: Makefile
	@echo -e "\nUsage: make <OPTIONS> ... <TARGETS>\n\nTargets:"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'


# ==============================================================================
# Swagger

.PHONY: install-swagger
install-swagger:
	@go install github.com/go-swagger/go-swagger/cmd/swagger@latest

.PHONY: swagger 
swagger:
	@echo "===========> Generating swagger API docs" 
	@swagger generate spec --scan-models -w $(ROOT_PACKAGE)/cmd/genswaggertypedocs -o $(ROOT_PACKAGE)/api/swagger.yaml 

.PHONY: serve-swagger
serve-swagger:
	@swagger serve -F=swagger --no-open --port 36666 $(ROOT_PACKAGE)/api/swagger.yaml

