# Developer Guide

## Getting Started

We're just getting started as a project and this guide is still a work in progress. Please review and follow the instructions in the [README.md](./README.md) to get started with the project. Ask questions in GitHub issues.

## Building the APIs

The SMI APIs live in the [apis/](./apis/) directory of this repository. This project uses [Kubebuilder v3](https://book.kubebuilder.io/) to scaffold APIs. If you're building a new API, install Kubebuilder v3 and use the `kubebuilder create api` command. For example:

```console
$ kubebuilder create api --group access --version v1alpha1 --kind TrafficTarget
Create Resource [y/n]
y
Create Controller [y/n]
y
```

Along with defining the Kubernetes resource, we'll also want to define a controller which Kubebuilder allows us to do with their CLI. The controller for the resource contains reconciliation logic (actions that take place when the resource is created, deleted, updated, etc.).

The `PROJECT` file at the root of the project directory houses Kubebuilder metadata for scaffoling new components.

After a new API is scaffolded, add the appropriate structs and attributes necessary to the `<api-kind>_types.go` file. Then, run `make generate`. This will automatically generate some necessary files for the API. Remember to run `make generate` any time you modify the structs related to the API.
