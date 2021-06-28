helm "smi-controller" {
  depends_on = ["helm.cert-manager"]

  cluster = "k8s_cluster.dc1"
  namespace = "shipyard"

  chart = "../../../helm/smi-controller"
  values = "./smi-controller-helm-values.yaml" 
}