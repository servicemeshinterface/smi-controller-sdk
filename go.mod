module github.com/nicholasjackson/smi-controller-sdk

go 1.15

require (
	github.com/go-logr/logr v0.3.0
	github.com/go-logr/zapr v0.3.0 // indirect
	github.com/onsi/ginkgo v1.14.1
	github.com/onsi/gomega v1.10.2
	github.com/servicemeshinterface/smi-sdk-go v0.4.1
	github.com/stretchr/testify v1.6.1
	k8s.io/apimachinery v0.18.8
	k8s.io/client-go v0.18.8
	sigs.k8s.io/controller-runtime v0.6.0
)

//replace github.com/servicemeshinterface/smi-sdk-go v0.4.1 => ../../servicemeshinterface/smi-sdk-go

replace github.com/servicemeshinterface/smi-sdk-go v0.4.1 => github.com/nicholasjackson/smi-sdk-go v0.0.0-20210123215756-d8c5233cc084
