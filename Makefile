NAME   := austin1237/gifbot
TAG    := $$(git rev-parse HEAD)
IMG    := ${NAME}:${TAG}
LATEST := ${NAME}:latest

build:
	@docker build -t ${IMG} ./gifbot
	@docker tag ${IMG} ${LATEST}
 
hub_push:
	@docker push ${NAME}
