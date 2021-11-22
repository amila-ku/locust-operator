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

### create CR

Deployment of locust master only

```
apiVersion: locustload.cndev.io/v1alpha1
kind: Locust
metadata:
  name: example-locust
spec:
  image: amilaku/locust:v0.0.1
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

commands

```
make build

```

### build docker image

```
make docker-build
```