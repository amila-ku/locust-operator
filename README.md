# locust-operator

## Introduction
Performing loadtests require lot of effort and preperation, inorder to generate traffic similar to production environment a distributed load generation tool would have to be used. [Locust](https://locust.io/) is a very popular load generation tool which supports writing test cases in python. The purpose of this project is to provide ready to use solution of locust.io for performing distributed load testing.

## Guides on load testing 
 - [Distributed Loadtesting with Locust](https://cloud.google.com/solutions/distributed-load-testing-using-gke)

Locust Operator supports two different deployments setups for locust.
 - cluster: creates a single master with multiple workers.
 - standalone: only master instance of locust created.
 
## status

![](https://github.com/amila-ku/locust-operator/workflows/build/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/amila-ku/locust-operator)](https://goreportcard.com/report/github.com/amila-ku/locust-operator)

# How to install Locust Operator

## Create related dependencies for the operator

```
make apply-resources
```

## Deploy CRD to a kubernetes cluster

```
make create-crd
```

### check if operator pod is running 

```
kubectl get pods
NAME                                     READY   STATUS    RESTARTS   AGE
locust-operator-5fb99cfd9b-k5w4b   1/1     Running   0          118s

```

# Create Locust Custom Resources

### load test scripts

Load test scripts has to be created with locust, instrictions on setting up a few test scenarios are shown [here](https://docs.locust.io/en/stable/quickstart.html)

Once the locust script is completed it has to be packaged as a docker container, we use this container when defining the custom resource locust operator. An example is available [here](https://github.com/amila-ku/locust-operator/tree/master/helpers/locust)

### create CR

Deployment of locust master only

```
apiVersion: locustload.cndev.io/v1alpha1
kind: Locust
metadata:
  name: example-locust
spec:
  image: <dockerrepository>/<image-name>:v0.0.1
  hosturl: https://postman-echo.com


```

### create CR with

Distributed locust deployment with workers

```
apiVersion: locustload.cndev.io/v1alpha1
kind: Locust
metadata:
  name: example-locust
spec:
  slaves: 2
  image: amilaku/locust:v0.0.1
  hosturl: https://postman-echo.com

```

## To Do 
- Autoscale locust workers with HPA.
- Load Generation per given rate.
- Prescale application deployments in a Kubernetes cluster before starting load test.
- Generate Report

# Developers

## build operator

### Prerequisites

Before starting you must have golang 1.13 or higher installed, then install operator framework by following [installation instructions](https://sdk.operatorframework.io/docs/install-operator-sdk/).


commands

```
make build-img

```

### push docker image to repo

```
make push-img
```
