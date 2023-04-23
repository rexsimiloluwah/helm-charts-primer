IMAGE_NAME?=go-docker-sample-app
DOCKER_IMAGE_REPOSITORY_NAME?=similoluwaokunowo/go-docker-sample-app
IMAGE_TAG?=latest

build-docker-image:
	docker build -t ${IMAGE_NAME}:${IMAGE_TAG} .

run-docker-container:
	docker run -p 5045:5045 -d ${IMAGE_NAME}:${IMAGE_TAG}

push-docker-image:
	docker tag ${IMAGE_NAME}:${IMAGE_TAG} ${DOCKER_IMAGE_REPOSITORY_NAME}:${IMAGE_TAG}
	docker push  ${DOCKER_IMAGE_REPOSITORY_NAME}:${IMAGE_TAG}

