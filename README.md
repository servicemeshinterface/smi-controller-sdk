# SMI Controller SDK

Projects that would like to build a SMI Spec compliant controller only needs to 
define a plugin that implement the extension points defined by the SDK. Core Kubernetes 
controller methods that handle the lifecylcle are implemented by the SDK and 
the methods in your API are called accordingly.

The intention behind this project is to simplify the process of implementing an SMI controller.
Base controller setup, conversion webhooks and validation are handled by the SDK, 
the end user only needs to implement the business  logic for the base level API in the smi spec `v1aplha1`.

Since the conversion webhooks convert other API versions to the base API which is then passed as 
a value to a function in your implementation. The design intention is to make it a trivial operation to
handle all SMI versions and to give a simple and easy upgrade path.

The other intention behind this SDK is to enable separation between the logic of writing a Kubernetes 
controller and the logic executed upon receiving an event. It should be possible to write a high level
of unit tests without needing to run the controller and Kubernetes.

## EXAMPLE: Implementing the SDK and creating a controller to handle Upsert and Delete for access.smi-spec.io/v1alpha1.TrafficTarget resources

To implement the SDK you need to create structs that implement the callback methods you would
like to receive. The SDK handles the Kubernetes lifecycle of the SMI resources as they are added 
and deleted from the Kubernets cluster saving you the job of writing custom controller code.

You only need to write the custom logic you would like to execute when for example a new `TrafficTarget`
resource is dded to the cluster. In addition by using the SDK you only need to implement methods you
are insterested in, not the full API.

For example, to create a simple implemntation which logs when a SMI resource is added or removed from
the system, you create a struct like the following example.

```go
type logger struct {}

func (l *logger) UpsertTrafficTarget(
	ctx context.Context,
	c client.Client,
	log logr.Logger,
	tt *accessv1alpha1.TrafficTarget,
) (ctrl.Result, error) {

	log.Info("UpsertTrafficTarget", "api", "v1alpha", "target", tt)

	return ctrl.Result{}, nil
}

func (l *logger) DeleteTrafficTarget(
	ctx context.Context,
	c client.Client,
	log logr.Logger,
	tt *accessv1alpha1.TrafficTarget,
) (ctrl.Result, error) {

	log.Info("DeleteTrafficTarget", "api", "v1alpha", "target", tt)

	return ctrl.Result{}, nil
}
```

You then register these callbacks with the controller API, when the controller needs to action a change 
to a SMI Spec resource it calls the apropriate callback. For example, when a new `TrafficTarget` is 
received by the controller the `logger.UpsertTrafficTarget` method is called.

```go
sdk.API().RegisterV1Alpha(&logger{})
```

Once registered you can start the controller

```go
controller.Start()
```

The full example can be seen below

```go
package main

import (
	"github.com/nicholasjackson/smi-controller-sdk/sdk"
	"github.com/nicholasjackson/smi-controller-sdk/sdk/controller"
)

type logger struct {}

func (l *logger) UpsertTrafficTarget(
	ctx context.Context,
	c client.Client,
	log logr.Logger,
	tt *accessv1alpha1.TrafficTarget,
) (ctrl.Result, error) {

	log.Info("UpsertTrafficTarget", "api", "v1alpha", "target", tt)

	return ctrl.Result{}, nil
}

func (l *logger) DeleteTrafficTarget(
	ctx context.Context,
	c client.Client,
	log logr.Logger,
	tt *accessv1alpha.TrafficTarget,
) (ctrl.Result, error) {

	log.Info("DeleteTrafficTarget", "api", "v1alpha", "target", tt)

	return ctrl.Result{}, nil
}

func main() {
	// register our lifecycle callbacks with the controller
	sdk.API().RegisterV1Alpha(&logger{})

	// create and start a the controller
	controller.Start()
}
```

The below output shows the controller in action when a new `TrafficTarget` resource is added to the cluster.

```shell
➜ go run .
2021-01-14T19:02:59.569Z        INFO    controller-runtime.metrics      metrics server is starting to listen{"addr": ":8080"}
2021-01-14T19:02:59.575Z        INFO    setup   starting manager
2021-01-14T19:02:59.576Z        INFO    controller-runtime.manager      starting metrics server {"path": "/metrics"}
2021-01-14T19:02:59.576Z        INFO    controller-runtime.controller   Starting EventSource    {"controller": "traffictarget", "source": "kind source: /, Kind="}
2021-01-14T19:02:59.677Z        INFO    controller-runtime.controller   Starting Controller     {"controller": "traffictarget"}
2021-01-14T19:02:59.677Z        INFO    controller-runtime.controller   Starting workers        {"controller": "traffictarget", "worker count": 1}
2021-01-14T19:03:14.797Z        INFO    controllers.TrafficTarget       UpsertTrafficTarget     {"api": "v1alpha2", "target": {"kind":"TrafficTarget","apiVersion":"access.smi-spec.io/v1alpha2","metadata":{"name":"path-specific","namespace":"default","selfLink":"/apis/access.smi-spec.io/v1alpha2/namespaces/default/traffictargets/path-specific","uid":"61ac9d1a-0be5-43d6-8eb5-00e1c55d01a9","resourceVersion":"30406","generation":1,"creationTimestamp":"2021-01-14T19:03:16Z","annotations":{"kubectl.kubernetes.io/last-applied-configuration":"{\"apiVersion\":\"access.smi-spec.io/v1alpha2\",\"kind\":\"TrafficTarget\",\"metadata\":{\"annotations\":{},\"name\":\"path-specific\",\"namespace\":\"default\"},\"spec\":{\"destination\":{\"kind\":\"ServiceAccount\",\"name\":\"service-a\",\"namespace\":\"default\",\"port\":8080},\"rules\":[{\"kind\":\"HTTPRouteGroup\",\"matches\":[\"metrics\"],\"name\":\"the-routes\"}],\"sources\":[{\"kind\":\"ServiceAccount\",\"name\":\"prometheus\",\"namespace\":\"default\"}]}}\n"},"finalizers":["traffictarget.finalizers.smi-controller"]},"spec":{"destination":{"kind":"ServiceAccount","name":"service-a","namespace":"default","port":8080},"sources":[{"kind":"ServiceAccount","name":"prometheus","namespace":"default"}],"rules":[{"kind":"HTTPRouteGroup","name":"the-routes","matches":["metrics"]}]}}}
2021-01-14T19:03:14.797Z        DEBUG   controller-runtime.controller   Successfully Reconciled {"controller": "traffictarget", "request": "default/path-specific"}
```

## Installing the example logging controller

To install the example logging SMI controller you can use the provided Helm chart. Before installing the chart ensure you have Cert Manager running on your system or execut the following commands to install:

```shell
$ kubectl apply --validate=false -f https://github.com/jetstack/cert-manager/releases/download/v1.1.0/cert-manager.crds.yaml

$ helm repo add jetstack https://charts.jetstack.io

$ helm install --name my-release --namespace cert-manager jetstack/cert-manager
```

You can then install the example using the helm chart:

```shell
helm repo add smi-controler https://nicholasjackson.io/smi-controller-sdk/

helm install smi-controller smi-controller/smi-controller
```

## Helm Chart - Building

The SMI controller SDK provides a Helm chart for installation. Like the main project aim, the aim of
the chart is to provide a generic resource that any implementation of the Controller SDK can use.
It should be possible for any project which leveraging the SDK to use the generic Helm chart to install
their controller. This like the SDK itself saves maintenance and allows implementors to concentrate
on writing mesh specific logic.

To build and update the chart the following command can be used:

```shell
make update_helm
```

```shell
helm package ./helm/smi-controller
Successfully packaged chart and saved it to: /home/nicj/code/src/github.com/nicholasjackson/smi-controller-sdk/smi-controller-0.1.0.tgz
mv smi-controller-0.1.0.tgz ./docs/
cd ./docs && helm repo index .
```

This repository uses GitHub pages to serve a Helm chart repo from the docs folder, commiting 
the changes created with the previous command to the `main` branch will automatically update the 
repository and start serving the new chart version.

## Developing Locally

To create a tests Kubernetes cluster on Docker that you can use when developing the Controller SDK
you can use the included Shipyard blueprint that will create the cluster and install any 
pre-requistes.


```shell
➜ shipyard run ./shipyard 
Running configuration from:  ./shipyard

2021-02-26T17:18:44.822Z [INFO]  Creating Output: ref=KUBECONFIG
2021-02-26T17:18:44.822Z [INFO]  Creating Network: ref=dc1
2021-02-26T17:18:44.845Z [INFO]  Creating Cluster: ref=dc1
2021-02-26T17:19:38.654Z [INFO]  Create Ingress: ref=smi-webhook
2021-02-26T17:19:38.654Z [INFO]  Applying Kubernetes configuration: ref=cert-manager config=[/home/nicj/go/src/github.com/nicholasjackson/smi-controller-sdk/shipyard/modules/smi-controller/cert-manager.crds.yaml]
2021-02-26T17:19:38.858Z [INFO]  Creating Helm chart: ref=cert-manager
2021-02-26T17:19:53.291Z [INFO]  Creating Helm chart: ref=smi-controler

########################################################

Title Consul Service Mesh on Kubernetes with Monitoring
Author Nic Jackson

shipyard_version: ">= 0.2.1"

1 Service Mesh Interface Controller SDK

....
```

Shipyard is available for all platforms and can be installed by following the instructions on the 
Shipyard website: [https://shipyard.run/docs/install](https://shipyard.run/docs/install)

Shipyard exposes the controller and webhooks running on your local machine to the local 
Kubernetes cluster. 

Shipyard places the Kubernetes config file needed for interacting with the server int `$HOME/.shipyard/config/dc1/kubeconfig.yaml`
you can use the command `export KUBECONFIG=$(shipyard output KUBECONFIG)` to set this as an environment variable.

Running the command `make run_local` will automatically fetch the TLS certificates needed for the webhook server
and will start the local code.

```shell
➜ make run_local
mkdir -p /tmp/k8s-webhook-server/serving-certs/
kubectl get secret controller-webhook-certificate -n smi -o json | \
        jq -r '.data."tls.crt"' | \
        base64 -d > /tmp/k8s-webhook-server/serving-certs/tls.crt
kubectl get secret controller-webhook-certificate -n smi -o json | \
        jq -r '.data."tls.key"' | \
        base64 -d > /tmp/k8s-webhook-server/serving-certs/tls.key
go run .
I0226 17:24:40.411612   23455 request.go:621] Throttling request took 1.0387667s, request: GET:https://127.0.0.1:64207/apis/storage.k8s.io/v1beta1?timeout=32s
2021-02-26T17:24:40.413Z        INFO    controller-runtime.metrics      metrics server is starting to listen    {"addr": ":9102"}
```

You can test the setup by applying one of the local configuration files

```shell
➜ k apply -f ./examples/traffictarget_v2.yaml 
traffictarget.access.smi-spec.io/path-specific-v2 created
```

Looking at the logs you will see that the locally running code has handled the webhook conversion and the 
controller code.

```shell
21-02-26T17:27:53.139Z        INFO    traffictarget-resource  ConvertTo v1alpha1
2021-02-26T17:27:53.142Z        INFO    traffictarget-resource  ConvertFrom v1alpha1
2021-02-26T17:27:53.143Z        INFO    traffictarget-resource  ConvertFrom v1alpha1
2021-02-26T17:27:53.144Z        INFO    traffictarget-resource  ConvertFrom v1alpha1
2021-02-26T17:27:53.146Z        INFO    traffictarget-resource  ConvertFrom v1alpha1
2021-02-26T17:27:53.149Z        INFO    controllers.TrafficTarget       UpsertTrafficTarget     {"api": "v1alpha1", "target": {"kind":"TrafficTarget","apiVersion":"access.smi-spec.io/v1alpha1","metadata":{"name":"path-specific-v2","namespace":"default","selfLink":"/apis/access.smi-spec.io/v1alpha1/namespaces/default/traffictargets/path-specific-v2","uid":"ca023c97-633a-44c4-9d89-cdd54b076ada","resourceVersion":"1305","generation":1,"creationTimestamp":"2021-02-26T17:27:53Z","annotations":{"kubectl.kubernetes.io/last-applied-configuration":"{\"apiVersion\":\"access.smi-spec.io/v1alpha2\",\"kind\":\"TrafficTarget\",\"metadata\":{\"annotations\":{},\"name\":\"path-specific-v2\",\"namespace\":\"default\"},\"spec\":{\"destination\":{\"kind\":\"ServiceAccount\",\"name\":\"service-a\",\"namespace\":\"default\",\"port\":8080},\"rules\":[{\"kind\":\"HTTPRouteGroup\",\"matches\":[\"metrics\"],\"name\":\"the-routes\"}],\"sources\":[{\"kind\":\"ServiceAccount\",\"name\":\"prometheus\",\"namespace\":\"default\"}]}}\n"},"finalizers":["traffictarget.finalizers.smi-controller"],"managedFields":[{"manager":"kubectl","operation":"Update","apiVersion":"access.smi-spec.io/v1alpha2","time":"2021-02-26T17:27:53Z","fieldsType":"FieldsV1","fieldsV1":{"f:metadata":{"f:annotations":{".":{},"f:kubectl.kubernetes.io/last-applied-configuration":{}}},"f:spec":{".":{},"f:destination":{".":{},"f:kind":{},"f:name":{},"f:namespace":{},"f:port":{}},"f:rules":{},"f:sources":{}}}},{"manager":"smi-controller-sdk","operation":"Update","apiVersion":"access.smi-spec.io/v1alpha1","time":"2021-02-26T17:27:53Z","fieldsType":"FieldsV1","fieldsV1":{"f:metadata":{"f:finalizers":{".":{},"v:\"traffictarget.finalizers.smi-controller\"":{}}}}}]},"destination":{"kind":"","name":""},"sources":null,"specs":null}}
2021-02-26T17:27:53.149Z        DEBUG   controller-runtime.controller   Successfully Reconciled {"controller": "traffictarget", "request": "default/path-specific-v2"}
2021-02-26T17:27:53.149Z        INFO    controllers.TrafficTarget       UpsertTrafficTarget     {"api": "v1alpha1", "target": {"kind":"TrafficTarget","apiVersion":"access.smi-spec.io/v1alpha1","metadata":{"name":"path-specific-v2","namespace":"default","selfLink":"/apis/access.smi-spec.io/v1alpha1/namespaces/default/traffictargets/path-specific-v2","uid":"ca023c97-633a-44c4-9d89-cdd54b076ada","resourceVersion":"1305","generation":1,"creationTimestamp":"2021-02-26T17:27:53Z","annotations":{"kubectl.kubernetes.io/last-applied-configuration":"{\"apiVersion\":\"access.smi-spec.io/v1alpha2\",\"kind\":\"TrafficTarget\",\"metadata\":{\"annotations\":{},\"name\":\"path-specific-v2\",\"namespace\":\"default\"},\"spec\":{\"destination\":{\"kind\":\"ServiceAccount\",\"name\":\"service-a\",\"namespace\":\"default\",\"port\":8080},\"rules\":[{\"kind\":\"HTTPRouteGroup\",\"matches\":[\"metrics\"],\"name\":\"the-routes\"}],\"sources\":[{\"kind\":\"ServiceAccount\",\"name\":\"prometheus\",\"namespace\":\"default\"}]}}\n"},"finalizers":["traffictarget.finalizers.smi-controller"],"managedFields":[{"manager":"kubectl","operation":"Update","apiVersion":"access.smi-spec.io/v1alpha2","time":"2021-02-26T17:27:53Z","fieldsType":"FieldsV1","fieldsV1":{"f:metadata":{"f:annotations":{".":{},"f:kubectl.kubernetes.io/last-applied-configuration":{}}},"f:spec":{".":{},"f:destination":{".":{},"f:kind":{},"f:name":{},"f:namespace":{},"f:port":{}},"f:rules":{},"f:sources":{}}}},{"manager":"smi-controller-sdk","operation":"Update","apiVersion":"access.smi-spec.io/v1alpha1","time":"2021-02-26T17:27:53Z","fieldsType":"FieldsV1","fieldsV1":{"f:metadata":{"f:finalizers":{".":{},"v:\"traffictarget.finalizers.smi-controller\"":{}}}}}]},"destination":{"kind":"","name":""},"sources":null,"specs":null}}
2021-02-26T17:27:53.149Z        DEBUG   controller-runtime.controller   Successfully Reconciled {"controller": "traffictarget", "request": "default/path-specific-v2"}
2021-02-26T17:27:53.149Z        INFO    traffictarget-resource  ConvertFrom v1alpha1
```

When Shipyard setup the local environment it created a Kubernetes service `smi-webhook.shipyard.svc` that proxies traffic to your local machine. It also configures the CRDs for the SMI resources to use this service for the conversion and validation webhook.
Using this feature you can develop 100% locally without needing to deploy the controller to the Kubernetes cluster for testing.

To remove any resources created by Shipyard you can use the command `shipyard destroy`.

```shell
➜ shipyard destroy 
2021-02-26T17:34:30.516Z [INFO]  Destroy Ingress: ref=smi-webhook id=d199780b-2066-41da-80d4-4294cc13adcb
2021-02-26T17:34:30.516Z [INFO]  Destroy Helm chart: ref=smi-controler
2021-02-26T17:34:30.525Z [INFO]  Destroy Helm chart: ref=cert-manager
2021-02-26T17:34:30.526Z [INFO]  Destroy Kubernetes configuration: ref=cert-manager config=[/home/nicj/go/src/github.com/nicholasjackson/smi-controller-sdk/shipyard/modules/smi-controller/cert-manager.crds.yaml]
2021-02-26T17:34:30.687Z [INFO]  Destroy Cluster: ref=dc1
2021-02-26T17:34:31.215Z [INFO]  Destroy Network: ref=dc1
```