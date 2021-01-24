ingress "smi-webhook" {
  source {
    driver = "k8s"
    
    config {
      cluster = "k8s_cluster.dc1"
      port = 9443
    }
  }
  
  destination {
    driver = "local"
    
    config { 
      address = "localhost"
      port = 9443
    }
  
  }
}

ingress "local-to" {
  source {
    driver = "local"
    
    config {
      port = 9533
    }
  }
  
  destination {
    driver = "k8s"
    
    config {
      cluster = "k8s_cluster.${var.consul_k8s_cluster}"
      address = "smi-webhook.shipyard.svc"
      port = 9443
    }
  }
}
