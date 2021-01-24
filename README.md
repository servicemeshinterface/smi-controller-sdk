# SMI Controller SDK

This project is a scaffold Kubernetes controller for the SMI specification.

Projects that would like to build a SMI Spec compliant controller only needs to 
define a plugin that implement the extension points defined by the SDK. 

Core Kubernetes controller methods that handle the lifecylcle are implemented by the SDK and 
the methods in your API are called accordingly.

## Implementing the SDK and creating a controller to handle Upsert and Delete for 
## access.smi-spec.io/v1alpha2.TrafficTarget resources

To implement the SDK you need to create structs that implement the callback methods you would
like to receive. The SDK handles the Kubernetes lifecycle of the SMI resources as they are added 
and deleted from the Kubernets cluster saving you the job of writing custom controller code.

You only need to write the custom logic you would like to execute when for example a new TrafficTarget
resource is dded to the cluster. In addition by using the SDK you only need to implement methods you
are insterested in, not the full API.

For example, to create a simple implemntation which logs when a SMI resource is added or removed from
the system, you create a struct like the following example.

```go
type loggerV2 struct {}

func (l *loggerV2) UpsertTrafficTarget(
	ctx context.Context,
	c client.Client,
	log logr.Logger,
	tt *accessv1alpha2.TrafficTarget,
) (ctrl.Result, error) {

	log.Info("UpsertTrafficTarget", "api", "v1alpha2", "target", tt)

	return ctrl.Result{}, nil
}

func (l *loggerV2) DeleteTrafficTarget(
	ctx context.Context,
	c client.Client,
	log logr.Logger,
	tt *accessv1alpha2.TrafficTarget,
) (ctrl.Result, error) {

	log.Info("DeleteTrafficTarget", "api", "v1alpha2", "target", tt)

	return ctrl.Result{}, nil
}
```

You then register these callbacks with the controller API, when the controller needs to action a change 
to a SMI Spec resource it calls the apropriate callback. For example, when a new `TrafficTarget` is 
received by the controller the `loggerV2.UpsertTrafficTarget` method is called.

```go
sdk.API().RegisterV1Alpha2(&loggerV2{})
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

type loggerV2 struct {}

func (l *loggerV2) UpsertTrafficTarget(
	ctx context.Context,
	c client.Client,
	log logr.Logger,
	tt *accessv1alpha2.TrafficTarget,
) (ctrl.Result, error) {

	log.Info("UpsertTrafficTarget", "api", "v1alpha2", "target", tt)

	return ctrl.Result{}, nil
}

func (l *loggerV2) DeleteTrafficTarget(
	ctx context.Context,
	c client.Client,
	log logr.Logger,
	tt *accessv1alpha2.TrafficTarget,
) (ctrl.Result, error) {

	log.Info("DeleteTrafficTarget", "api", "v1alpha2", "target", tt)

	return ctrl.Result{}, nil
}

func main() {
	// register our lifecycle callbacks with the controller
	sdk.API().RegisterV1Alpha2(&loggerV2{})

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

You can then install the example using the helm chart

```shell
$ helm install smi-controller --namespace smi --debug --create-namespace .
```
