NAME   := austin1237/gifbot
TAG    := $$(git rev-parse HEAD)
IMG    := ${NAME}:${TAG}
LATEST := ${NAME}:latest
DEV    := ${NAME}:dev


build_prod:
	@docker build -t ${IMG} ./gifbot
	@docker tag ${IMG} ${LATEST}

build_dev:
	@docker build -t ${IMG}_dev ./gifbot
	@docker tag ${IMG}_dev ${DEV}

hub_push:
	@docker push ${NAME}

deploy_dev:
	cd terraform/dev && terraform init
	cd terraform/dev && terraform apply

deploy_prod:
	cd terraform/prod && terraform init
	cd terraform/prod && terraform apply