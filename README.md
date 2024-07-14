# healy-medicare-app-microservice-clean-arcitecture

install kubectl 
install helm

install ingress-nginx before applying k8s deployments

```
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx

```
create namespace
```
kubectl create ns ingress-nginx
```

if k8s/ingress folder not  exist create one

```
helm template ingress-nginx ingress-nginx \
--repo https://kubernetes.github.io/ingress-nginx  \
--version 4.11.0 \
--namespace ingress-nginx \
> ./k8s/ingress/ingress-nginx-1.11.0.yaml
```

apply the ingress-controller to our cluster

```
kubectl apply -f ./k8s/ingress/ingress-nginx-1.11.0.yaml
```
After a while, they should all be running. The following command will wait for the ingress controller pod to be up, running, and ready:
```
kubectl get svc -n ingress-nginx
```
please add external ip to dns provider
---------------------------------------------------------------
create a k8s/ingress/ingress.yaml
```
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: ingress-service
  annotations:
      kubernetes.io/ingress.class: nginx
spec:
  rules:
  - host: ajay404.online # change to your domain
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: gateway-srv
            port:
              number: 80
```
apply this ingress 
```
kubectl apply -f k8s/ingress/ingress.yaml
```

now u can access your service http

now need to secure our secure

download cert-manager
```
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.15.1/cert-manager.yaml
```

check all cert pods are running 

```
kubectl get pods --namespace cert-manager
```

this should be the similar output

```
NAME                                       READY   STATUS    RESTARTS   AGE
cert-manager-5c6866597-zw7kh               1/1     Running   0          2m
cert-manager-cainjector-577f6d9fd7-tr77l   1/1     Running   0          2m
cert-manager-webhook-787858fcdb-nlzsq      1/1     Running   0          2m
```

create cert issuer,  k8s/certificate/issuer.yaml

```
apiVersion: cert-manager.io/v1
kind: ClusterIssuer
metadata:
  name: letsencrypt-prod
  namespace: cert-manager
spec:
  acme:
    server: https://acme-v02.api.letsencrypt.org/directory
    email: example@gmail.com # change to your email
    privateKeySecretRef:
      name: letsencrypt-prod-key
    solvers:
    - http01:
        ingress:
          class: nginx
```

create a certificate , k8s/certificate/certificate.yaml

```
apiVersion: cert-manager.io/v1
kind: Certificate
metadata:
  name: healy-cert
spec:
  secretName: healy-tls
  dnsNames:
  - ajay404.online #change to your dns
  issuerRef:
    name: letsencrypt-prod # add your issuer name
    kind: ClusterIssuer 

```

apply both

```
kubectl apply -f k8s/certificate/issuer.yaml
kubectl apply -f k8s/certificate/certificate.yaml
```

# GREATE JOB!
