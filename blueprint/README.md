---
title: Consul Service Mesh on Kubernetes with Monitoring
author: Nic Jackson
slug: k8s_consul_stack
---

# Consul Service Mesh on Kubernetes with Monitoring

shipyard_version: ">= 0.1.18"

This blueprint creates a Kubernetes cluster and installs the following elements:

* Consul Service Mesh With CRDs
* Prometheus
* Loki
* Grafana

To access Grafana the following details can be used:

* user: admin
* pass: admin

ACLs are disabled for Consul
