# locust-operator

## Introduction
The purpose of this project is to provide a easy to deploy version of locust.io.

Locust can be created in two different deployments
 - cluster: creates a single master with multiple workers.
 - standalone: only master instance of locust created.
 
## status

![](https://github.com/amila-ku/locust-operator/workflows/build/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/amila-ku/locust-operator)](https://goreportcard.com/report/github.com/amila-ku/locust-operator)

## build operator

commands

```
make build-img

```

### push docker image to repo

```
make push-img
```

## Create related dependencies

```
make apply-resources
```

## Deploy CRDs

```
make create-crd
```

### check if operator pod is running 

```
kubectl get pods
NAME                                     READY   STATUS    RESTARTS   AGE
locust-operator-5fb99cfd9b-k5w4b   1/1     Running   0          118s

```

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