module github.com/nicholasjackson/smi-controller-sdk

go 1.15

require (
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/engine v1.13.1
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/go-logr/logr v0.3.0
	github.com/go-logr/zapr v0.3.0
	github.com/hashicorp/go-hclog v0.15.0
	github.com/onsi/ginkgo v1.14.1
	github.com/onsi/gomega v1.10.2
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/servicemeshinterface/smi-sdk-go v0.4.1
	github.com/stretchr/testify v1.6.1
	k8s.io/apimachinery v0.18.8
	k8s.io/client-go v0.18.8
	sigs.k8s.io/controller-runtime v0.6.0
)

replace github.com/servicemeshinterface/smi-sdk-go v0.4.1 => ../../servicemeshinterface/smi-sdk-go

//replace github.com/servicemeshinterface/smi-sdk-go v0.4.1 => github.com/nicholasjackson/smi-sdk-go v0.0.0-20210115203718-ba2bc66de8a0
