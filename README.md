# KRR

## steps

### install minikube:
```curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64```
```sudo install minikube-linux-amd64 /usr/local/bin/minikube && rm minikube-linux-amd64```

### start minikube:
```minikube start```

### minikube cmds:
#### loadbalance service get IP?
```minikube service service/prometheus-kube-prometheus-prometheus```

### install metrics-server
```kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml```

### skip tls check in metrics-server
``kubectl edit deployment metrics-server -n kube-system``
add
```yaml
spec:
  containers:
  - name: metrics-server
    image: <metrics-server-image>
    args:
    - --kubelet-insecure-tls
    - --kubelet-preferred-address-types=InternalIP,Hostname,ExternalIP
```

### install helm
```
curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
chmod 700 get_helm.sh
./get_helm.sh
helm version
```

### Add the default Helm chart repository:
```
helm repo add stable https://charts.helm.sh/stable
helm repo update
```
### Add the Prometheus Community Helm repository:
```
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
```

### install prometheus:
Search for the Prometheus Helm chart:
```
helm search repo prometheus
```
Install Prometheus using the kube-prometheus-stack chart:
```
helm install prometheus prometheus-community/kube-prometheus-stack
```
prometheus: This is the release name.
prometheus-community/kube-prometheus-stack: This is the chart name.


### fix krr error:
The error message you're seeing is related to RBAC (Role-Based Access Control) permissions in Kubernetes. Specifically, the service account system:serviceaccount:default:default does not have sufficient permissions to list the HorizontalPodAutoscalers (HPA) resource, which is part of the autoscaling API group.

To fix this, you need to grant the necessary permissions to the default service account (or whichever service account you're using for the krr pod) to access the HPA resources.

Steps to Grant Permissions
Create a RoleBinding for the Service Account

You will need to create a RoleBinding (for a specific namespace) or a ClusterRoleBinding (for the entire cluster) to grant the necessary permissions to the service account.

Since your service account is default:default, you can bind the ClusterRole view (or create a custom role if necessary) to this service account.

Here's an example of how you can create a ClusterRoleBinding for the service account default:
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
name: krr-clusterrolebinding
subjects:
- kind: ServiceAccount
  name: default
  namespace: default
  roleRef:
  kind: ClusterRole
  name: view  # 'view' gives read-only access to resources
  apiGroup: rbac.authorization.k8s.io
```
This ClusterRoleBinding binds the view role (which grants read-only access to all resources) to the default service account in the default namespace.
Apply this configuration using kubectl:
``kubectl apply -f krr-rolebinding.yaml``

### run krr one time
```bash
kubectl run -i --tty mypod --image=us-central1-docker.pkg.dev/genuine-flight-317411/devel/krr:v1.8.3 --rm --restart=Never
```

### krr to file json format
```bash
python krr.py simple -f json --fileoutput krr-report.json
```







## steps v2

### 

sum(irate(container_cpu_usage_seconds_total{pod="krr-app-85cddf485f-wpcdv"}[10m]))

python krr.py simple --prometheus-url=http://prometheus-kube-prometheus-prometheus.default.svc.cluster.local:9090






TODO: захардкоженный путь до конфига в докерфайле приложения
TODO: захардкоженный порт в deployment файле приложения.
TODO: сервисный аккаунт для пода приложения https://chatgpt.com/c/67514dae-d274-800b-92ec-af26e56d97f2 (права на запуск пода)

TODO: сервисный аккаунт для пода крр https://chatgpt.com/c/67516411-9bdc-800b-afe5-82a063e217c2 (права на получение данных прома)
krr-service-account.yaml не работает
krr-rolebinding.yaml работает