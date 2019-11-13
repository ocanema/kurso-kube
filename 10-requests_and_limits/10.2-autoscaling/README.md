# On kubernetes  
## Clone from metrics server repo

```git clone https://github.com/kubernetes-incubator/metrics-server.git```
```kubectl create -f metrics-server/deploy/1.8+/```

# On minikube
```minikube addons enable metrics-server```

# Create the autoscaling components

```kubectl create -f .```

