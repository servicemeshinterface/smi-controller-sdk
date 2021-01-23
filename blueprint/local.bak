// ingress exposing a local application
// would enable traffic on the K8s cluster dc1 sent to:
//      my-local-service.shipyard.svc.cluster.local:9090
// to be directed to:
//      localhost:30002
// on the shipyard host
ingress "k8s-to-local" {
  source {
    driver = "k8s"
    
    config {
      cluster = "k8s_cluster.dc1"
      port = 9091
    }
  }
  
  destination {
    driver = "local"
    
    config { 
      address = "localhost"
      port = 30002
    }
  
  }
}

// ingress exposing an application on one K8s cluster to the shipyard host
// would enable traffic on the shipyard host sent to:
//      localhost:9090
// to be directed to:
//      dc1-service.mynamespace.svc.cluster.local:30002
// on the dc1 cluster
ingress "local-to-k8s" {
  source {
    driver = "local"
    
    config {
      port = 9092
    }
  }
  
  destination {
    driver = "k8s"
    
    config {
      cluster = "k8s_cluster.dc1"
      address = "k8s-to-local.shipyard.svc"
      port = 9091
    }
  }
}

//// ingress exposing an application on one K8s cluster to another 
//// would enable traffic on the K8s cluster dc1 sent to:
////      dc1-service.shipyard.svc.cluster.local:9090
//// to be directed to:
////      dc2-service.mynamespace.svc.cluster.local:30002
//// on the dc2 cluster
//ingress "connector-k8-to-k8" {
//  source {
//    driver = "k8s"
//    
//    config = {
//      cluster = "k8s_cluster.dc1"
//      service = "dc1-service"
//      port = 9090
//    }
//
//  }
//  
//  destination {
//    driver = "k8s"
//    
//    config = {
//      cluster = "k8s_cluster.dc2"
//      service = "dc2-service"
//      namespace = "mynamespace"
//      port = 30002
//    }
//
//  }
//}
//
//// ingress exposing an application in a Docker container to the shipyard host
//// would enable traffic on the shipyard host sent to:
////      localhost:9090
//// to be directed to:
////      consul.container.shipyard.run:30002
//// on the Docker host
//ingress "connector-local-to-docker" {
//  destination {
//    driver = "docker"
//    
//    config = {
//      service = "consul.container.shipyard.run"
//      docker_host = docker_host()
//      port = 8500
//    }
//  }
//  
//  source {
//    driver = "local"
//    
//    config = {
//      service = "localhost"
//      port = 9090
//    }
//  }
//}
