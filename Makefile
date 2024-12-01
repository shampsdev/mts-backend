ifneq (,$(wildcard ./.env))
	include .env
	export
endif

TAG ?= 0.1.9
IMG=shampiniony/mts-backend:$(TAG)

docker-build:
	docker build -t $(IMG) .

docker-push:
	docker push $(IMG)

docker-run:
	docker run -d -p 8000:8000 $(IMG)

docker-stop:
	docker stop $$(docker ps -q --filter ancestor=$(IMG))

### GO

swag: swag-install ## Generate swag documentation (github.com/swaggo/swag)
	$(SWAG) init -g cmd/main.go -o docs

### Dependencies

SWAG = $(shell pwd)/bin/swag
swag-install:
	$(call go-get-tool,$(SWAG),github.com/swaggo/swag/cmd/swag@v1.16.3)

PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
define go-get-tool
@[ -f $(1) ] || { \
set -e ;\
TMP_DIR=$$(mktemp -d) ;\
cd $$TMP_DIR ;\
go mod init tmp ;\
echo "Downloading $(2)" ;\
GOBIN=$(PROJECT_DIR)/bin go install $(2) ;\
rm -rf $$TMP_DIR ;\
}
endef
