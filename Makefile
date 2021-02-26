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

fetch_certs:
	mkdir -p /tmp/k8s-webhook-server/serving-certs/
	
	kubectl get secret controller-webhook-certificate -n smi -o json | \
		jq -r '.data."tls.crt"' | \
		base64 -d > /tmp/k8s-webhook-server/serving-certs/tls.crt
	
	kubectl get secret controller-webhook-certificate -n smi -o json | \
		jq -r '.data."tls.key"' | \
		base64 -d > /tmp/k8s-webhook-server/serving-certs/tls.key

run_local: fetch_certs
	go run .

functional_test: fetch_certs
	export KUBECONFIG=$(shipyard output KUBECONFIG) 
	cd test && go run .