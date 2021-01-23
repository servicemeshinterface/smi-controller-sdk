//
// Install Consul using the helm chart.
//
helm "consul" {
  cluster = "k8s_cluster.${var.consul_k8s_cluster}"

  // chart = "github.com/hashicorp/consul-helm?ref=crd-controller-base"
  chart = "github.com/hashicorp/consul-helm?ref=v0.28.0"
  values = "./helm/consul-values.yaml"

  health_check {
    timeout = "60s"
    pods = ["app=consul"]
  }
}

ingress "consul" {
  source {
    driver = "local"
    
    config {
      port = 8500
    }
  }
  
  destination {
    driver = "k8s"
    
    config {
      cluster = "k8s_cluster.${var.consul_k8s_cluster}"
      address = "consul-ui.default.svc"
      port = 80
    }
  }
}
