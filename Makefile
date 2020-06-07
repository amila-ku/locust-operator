SHELL = bash
VERSION ?= v0.0.3

clean:
	kubectl delete deployment locust-operator
	kubectl delete locust example-locust

generate:
	operator-sdk generate k8s
	operator-sdk generate crds

delete-crd:
	kubectl delete crd locustload.cndev.io

create-cr:
	kubectl create -f deploy/operator.yaml
	kubectl create -f deploy/crds/locustload.cndev.io_v1alpha1_locust_cr.yaml 

apply-resources:
	kubectl create -f deploy/service_account.yaml
	kubectl create -f deploy/role.yaml
	kubectl create -f deploy/role_binding.yaml
	kubectl create -f deploy/operator.yaml

create-crd:
	kubectl create -f deploy/crds/locustload.cndev.io_locusts_crd.yaml 
	kubectl create -f deploy/operator.yaml

push-img: 
	docker push amilaku/locust-operator:${VERSION}

build-img:
	operator-sdk build amilaku/locust-operator:${VERSION}

rel:
	git tag ${VERSION}
	git push origin --tags


