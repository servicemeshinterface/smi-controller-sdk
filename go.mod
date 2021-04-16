module github.com/servicemeshinterface/smi-controller-sdk

go 1.15

require (
	github.com/cucumber/godog v0.11.0
	github.com/cucumber/messages-go/v10 v10.0.3
	github.com/go-logr/logr v0.3.0
	github.com/go-logr/zapr v0.3.0 // indirect
	github.com/onsi/ginkgo v1.14.1
	github.com/onsi/gomega v1.10.2
	github.com/servicemeshinterface/smi-sdk-go v0.4.1
	github.com/stretchr/testify v1.6.1
	k8s.io/api v0.19.2
	k8s.io/apimachinery v0.19.2
	k8s.io/client-go v0.19.2
	sigs.k8s.io/controller-runtime v0.7.2
)

//replace github.com/servicemeshinterface/smi-sdk-go v0.4.1 => ../servicemeshinterface/smi-sdk-go

replace github.com/servicemeshinterface/smi-sdk-go v0.4.1 => github.com/nicholasjackson/smi-sdk-go v0.0.0-20210331180119-d1be3fc5b8c9
