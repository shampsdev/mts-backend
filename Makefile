ifneq (,$(wildcard ./.env))
	include .env
	export
endif

TAG ?= latest
IMG=shampiniony/mts-backend:$(TAG)

docker-build:
	docker build -t $(IMG) .

docker-push:
	docker push $(IMG)

docker-run:
	docker run -d -p 8000:8000 $(IMG)

docker-stop:
	docker stop $$(docker ps -q --filter ancestor=$(IMG))
