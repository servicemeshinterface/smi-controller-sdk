---
title: Consul Service Mesh on Kubernetes with Monitoring
author: Nic Jackson
slug: k8s_consul_stack
---

shipyard_version: ">= 0.2.1"
# Service Mesh Interface Controller SDK

This blueprint creates a Kubernetes cluster and installs the following elements:

* Cert Manager
* SMI Controller CRDs and webhook config
* Local ingress exposing port 9443 for webhook to local machine
