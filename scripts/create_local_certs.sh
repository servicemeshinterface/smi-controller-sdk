#!/bin/bash -e

# Generate the Certificate for the Webhook
./generate_cert.sh --service smi-controller --namespace smi --secret local-webhook-certs
