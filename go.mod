module github.com/nicholasjackson/smi-controller

go 1.15

require (
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/engine v1.13.1
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/go-logr/logr v0.1.0
	github.com/hashicorp/go-hclog v0.15.0
	github.com/onsi/ginkgo v1.11.0
	github.com/onsi/gomega v1.8.1
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/servicemeshinterface/smi-sdk-go v0.4.1
	github.com/stretchr/testify v1.4.0
	k8s.io/apimachinery v0.18.0
	k8s.io/client-go v0.18.0
	sigs.k8s.io/controller-runtime v0.5.0
)

replace k8s.io/apimachinery v0.18.0 => k8s.io/apimachinery v0.17.2

replace k8s.io/client-go v0.18.0 => k8s.io/client-go v0.17.2
