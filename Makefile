DOCKER_REPO=nicholasjackson/smi-controller-example
DOCKER_VERSION=dev

build_docker:
	docker build -t ${DOCKER_REPO}:${DOCKER_VERSION} .

push_docker:
	docker push ${DOCKER_REPO}:${DOCKER_VERSION}
