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
