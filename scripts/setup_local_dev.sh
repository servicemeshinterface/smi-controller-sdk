#!/bin/bash -e

export SERVICE_NAME=smi-webhook
export NAMESPACE=shipyard

# Create the namespace
kubectl apply -f ../k8s_config/namespace.yaml

# Generate the Certificate for the Webhook
./generate_cert.sh --service ${SERVICE_NAME} --namespace ${NAMESPACE} --secret local-webhook-certs

# Install the certs to the expected temp directory
mkdir -p /tmp/k8s-webhook-server/serving-certs/
kubectl get secrets local-webhook-certs -n shipyard -o json | jq -r '.data."cert.pem"' | base64 -d > /tmp/k8s-webhook-server/serving-certs/tls.crt
kubectl get secrets local-webhook-certs -n shipyard -o json | jq -r '.data."key.pem"' | base64 -d > /tmp/k8s-webhook-server/serving-certs/tls.key


# Add the CRDs
# Setup the webhook config
cat ../crds/access.yaml | ./patch_webhook_yaml.sh | kubectl apply -f -
