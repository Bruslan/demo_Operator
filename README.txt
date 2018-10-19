Source: https://github.com/operator-framework/operator-sdk


How to Deploy Bruslan Operator

First, checkout and install the operator-sdk CLI (Source: https://github.com/operator-framework/operator-sdk):

    $ mkdir -p $GOPATH/src/github.com/operator-framework
    $ cd $GOPATH/src/github.com/operator-framework
    $ git clone https://github.com/operator-framework/operator-sdk
    $ git checkout tags/v0.0.5
    $ dep ensure
    $ go install github.com/operator-framework/operator-sdk/commands/operator-sdk

Create Image and Push it into Public Registry:

    $ operator-sdk build {hier adresse für public registry z.B. dockerhub -> bruslan/demo-operator}
    $ docker push {hier adresse für public registry z.B. dockerhub -> bruslan/demo-operator}


Deploy the app-operator:

    $ kubectl create -f deploy/rbac.yaml
    $ kubectl create -f deploy/operator.yaml
    $ kubectl create -f deploy/cr.yaml


Clean Up The Deployment

    $ kubectl delete -f deploy/cr.yaml
    $ kubectl delete -f deploy/operator.yaml
    $ kubectl delete -f deploy/rbac.yaml
