# healy-medicare-app-microservice-clean-arcitecture

install kubectl first 

install ingress-nginx before applying k8s deployments or skaffold

```
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.10.1/deploy/static/provider/cloud/deploy.yaml

```

A few pods should start in the ingress-nginx namespace:

```
kubectl get pods --namespace=ingress-nginx

```

After a while, they should all be running. The following command will wait for the ingress controller pod to be up, running, and ready:

```
kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=120s
```

then run skaffold or apply k8s deployments

```
skaffold dev

```