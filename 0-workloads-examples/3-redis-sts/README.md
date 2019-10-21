## create redis cluster: 
```
kubectl exec -it redis-cluster-0 -- redis-cli --cluster create --cluster-replicas 1 $(kubectl get pods -l app=redis-cluster -o jsonpath='{range.items[*]}{.status.podIP}:6379 ')
```
## view cluster info:
```
kubectl exec -it redis-cluster-0 -- redis-cli cluster info
```

## view cluster members info:
```
for x in $(seq 0 5); do echo "redis-cluster-$x"; kubectl exec redis-cluster-$x -- redis-cli role; echo; done
```