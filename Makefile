NAME   := austin1237/gifbot
TAG    := $$(git rev-parse HEAD)
IMG    := ${NAME}:${TAG}
LATEST := ${NAME}:latest


build_prod:
	@docker build -t ${IMG} ./gifbot
	@docker tag ${IMG} ${LATEST}

build_dev:
	@docker build ./gifbot
	@docker tag dev

hub_push:
	@docker push ${NAME}
