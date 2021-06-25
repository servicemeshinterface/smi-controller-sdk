//helm "smi-controler" {
//  depends_on = ["helm.cert-manager"]
//
//  cluster = "k8s_cluster.dc1"
//  namespace = "smi"
//
//  chart = "../../../helm/smi-controller"
//
//  values_string = {
//    "controller.enabled" = "false"
//    "webhook.service" = "smi-webhook"
//    "webhook.namespace" = "shipyard"
//    "webhook.additionalDNSNames.0" = "smi-webhook.shipyard.svc"
//  }
//}