## Create mysql password secret:
```
k create secret generic mysql-pass --from-literal=password=P@55W0Rd --dry-run -o yaml
```