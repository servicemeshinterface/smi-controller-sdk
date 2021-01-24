# Module to install Grafana, Loki, and Prometheus to Kubernetes

This module installs Grafana, Loki, and Prometheus to the specified
kubernetes cluster


## Created resources
* Consul Grafana Helm Chart
* Loki Helm Chart
* Promtail Helm Chart
* Prometheus Helm Chart
* Grafana Ingress running on port 8080

## Variables

To use this module the following variables need to be set:

* monitoring_k8s_cluster - name of the kubernetes cluster
* monitoring_network - name of the network the main application is running on

## Usage

This module can be consumed by using the module stanza

```
module "monitoring" {
  source = "./module_path_or_github"
}
```
