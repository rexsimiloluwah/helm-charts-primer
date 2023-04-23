## Learning how to use HELM charts

1. To  Create `dev` and `prod` namespaces
```bash
$kubectl create namespace dev
$kubectl create namespace prod 
```

2. To install the Helm chart for the multiple environments
```bash
# for default environment 
$helm install go-sample-app-release ./go-docker-sample-app \
    --values ./go-docker-sample-app/values.yaml

# for dev environment 
$helm install go-sample-app-release ./go-docker-sample-app \
    --values ./go-docker-sample-app/values.yaml \
    -f ./go-docker-sample-app/values-dev.yaml -n dev

# for prod environment
$helm install go-sample-app-release ./go-docker-sample-app \
    --values ./go-docker-sample-app/values.yaml \
    -f ./go-docker-sample-app/values-prod.yaml -n prod
```

3. To list the installed Helm charts for all namespaces
```bash
$helm list --all-namespaces
```

4. To view the created deployment, pods, and services for a namespace i.e. dev,prod
```bash
$kubectl get pods --namespace <namespace> 
$kubectl get deployment --namespace <namespace> 
$kubectl get service --namespace <namespace>
```

5. To access the server via your local network 
```bash
$minikube tunnel
# i.e. for the service in the 'dev' namespace
$kubectl port-forward service/go-sample-app :5045 --namespace dev
```
