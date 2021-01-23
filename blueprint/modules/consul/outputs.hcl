
output "CONSUL_HTTP_ADDR" {
  value = "${docker_ip()}:8500"
}
