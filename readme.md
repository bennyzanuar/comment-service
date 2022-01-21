# Comment Service

Repository for demo comment service

## requirement
- node v 16
- docker
- kubernetes


## how to
- deploy
envsubst < config/deployment.yml | kubectl apply -f
kubectl apply -f config/service.yml

- check deployment
kubectl get deployments

- check pod
kubectl get pods

- check service
kubectl get service

- check apps
kubectl logs comments-api-xxxxxx

- port forward from docker to kube
kubectl port-forward service/comments-api 8009:8008