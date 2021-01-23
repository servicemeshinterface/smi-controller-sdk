#!/bin/bash -e

# Create the namespace
kubectl apply -f ../k8s_config/namespace.yaml

# Add the CRDs
kubectl apply -f ../crds/access.yaml

# Generate the Certificate for the Webhook
./generate_cert.sh --service smi-webhook --namespace shipyard --secret local-webhook-certs

# Install the certs to the expected temp directory
mkdir -p /tmp/k8s-webhook-server/serving-certs/
kubectl get secrets local-webhook-certs -n shipyard -o json | jq -r '.data."cert.pem"' | base64 -d > /tmp/k8s-webhook-server/serving-certs/tls.crt
kubectl get secrets local-webhook-certs -n shipyard -o json | jq -r '.data."key.pem"' | base64 -d > /tmp/k8s-webhook-server/serving-certs/tls.key

export SERVICE_NAME=smi-webhook
export NAMESPACE=shipyard

# Setup the webhook config
cat ../k8s_config/webhook.yaml | ./patch_webhook_yaml.sh | kubectl apply -f -
