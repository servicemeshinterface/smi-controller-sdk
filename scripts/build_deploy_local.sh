#!/bin/bash -e

# Create the namespace
kubectl apply -f ../k8s_config/namespace.yaml

# Generate the Certificate for the Webhook
./generate_cert.sh --service smi-controller --namespace smi --secret smi-webhook-certs

# Add the certificate to the raw config
cat ../k8s_config/webhook.yaml | ./patch_webhook_yaml.sh | kubectl apply -f -

# Deploy the server
kubectl apply -f ../k8s_config/deployment.yaml
