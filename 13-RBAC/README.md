## Create jane serviceAccount
```kubectl create serviceaccount jane```

## Get the token of serviceAccount jane
```TOKEN=$(kubectl describe secrets "$(kubectl describe serviceaccount jane | grep -i Tokens | awk '{print $2}')" | grep token: | awk '{print $2}')
```
## Create an user on the kubeconfig file
```kubectl config set-credentials jane --token=$TOKEN
```

## Create a context who matches the cluster with the user
```kubectl config set-context jane --cluster=minikube --user=jane```

## Change the actual context
```kubectl config use-context jane```