# locust-operator

## Introduction
The purpose of this project is to provide a easy to deploy version of locust.io which also contains additional/useful features that are required.

Locust can be created in two different deployments
 - cluster: creates a single master with multiple workers.
 - standalone: only master instance of locust created.
 
## status

![](https://github.com/amila-ku/locust-operator/workflows/build/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/amila-ku/locust-operator)](https://goreportcard.com/report/github.com/amila-ku/locust-operator)

## operator build

commands

```
operator-sdk build quay.io/amila_ku/locust-operator:v0.0.1

```

## push docker image to repo

```
docker push quay.io/amila_ku/locust-operator:v0.0.1

docker tag quay.io/amila_ku/locust-operator:v0.0.1 amilaku/locust-operator:v0.0.1

docker push amilaku/locust-operator:v0.0.1
```

## Update CRD

```
sed -i 's|REPLACE_IMAGE|quay.io/amila_ku/locust-operator:v0.0.1|g' deploy/operator.yaml
```

## Deploy CRD

```
kubectl apply -f deploy/crds/locustload.cndev.io_locusts_crd.yaml 

kubectl get crds
kubectl create -f deploy/service_account.yaml
kubectl create -f deploy/role.yaml
kubectl create -f deploy/role_binding.yaml
kubectl create -f deploy/operator.yaml

```

### check if operator pod is running 

```
kubectl get pods
NAME                                     READY   STATUS    RESTARTS   AGE
locust-operator-5fb99cfd9b-k5w4b   1/1     Running   0          118s

```

### create CR

```
apiVersion: locustload.cndev.io/v1alpha1
kind: Locust
metadata:
  name: example-locust
spec:
  # Add fields here
  size: 3
  image: amilaku/locust:v0.0.1
  hosturl: https://www.google.com
  users: 2
  hatchrate: 1
```

### Regenerate resources

```
operator-sdk generate k8s
operator-sdk generate crds

```

### Add new controller

```
operator-sdk add controller --api-version=locustload.cndev.io/v1alpha1 --kind=Locust

operator-sdk build quay.io/amila_ku/locust-operator:v0.0.1

sed -i 's|REPLACE_IMAGE|quay.io/amila_ku/locust-operator:v0.0.1|g' deploy/operator.yaml
```
