DOCKER_REPO=nicholasjackson/smi-controller-example
DOCKER_VERSION=dev

build_docker:
	docker build -t ${DOCKER_REPO}:${DOCKER_VERSION} .

push_docker:
	docker push ${DOCKER_REPO}:${DOCKER_VERSION}

update_helm:
	helm package ./helm/smi-controller
	mv smi-controller-0.1.0.tgz ./docs/
	cd ./docs && helm repo index .
