module "consul" {
  source = "./modules/consul"
}

module "monitoring" {
  depends_on = ["module.consul"]
  source = "./modules/monitoring"
}
