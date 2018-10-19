Source: https://github.com/operator-framework/operator-sdk


How to Deploy Bruslan Operator
Deploy the app-operator
$ kubectl create -f deploy/rbac.yaml
$ kubectl create -f deploy/operator.yaml

# By default, creating a custom resource (App) triggers the app-operator to deploy a busybox pod
$ kubectl create -f deploy/cr.yaml



$ operator-sdk build {hier adresse für public registry z.B. dockerhub -> bruslan/demo-operator}
$ docker push {hier adresse für public registry z.B. dockerhub -> bruslan/demo-operator}


Clean Up The Deployment

$ kubectl delete -f deploy/cr.yaml
$ kubectl delete -f deploy/operator.yaml
$ kubectl delete -f deploy/rbac.yaml
