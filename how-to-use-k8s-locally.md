## Running Kubernetes Locally with Minikube, kubectl, kubens, and kubectx

### 1. Install the Tools

- **Minikube:** Local Kubernetes cluster for testing/development  
  https://minikube.sigs.k8s.io/docs/start/
- **kubectl:** Kubernetes CLI tool  
  https://kubernetes.io/docs/tasks/tools/
- **kubectx & kubens:** Easy context and namespace switching  
  https://github.com/ahmetb/kubectx

### 2. Start Minikube

```bash
minikube start
```

This will spin up a local Kubernetes cluster.

### 3. Use kubectl to Interact with Your Cluster

- Check cluster nodes:
  ```bash
  kubectl get nodes
  ```
- Deploy your app:
  ```bash
  kubectl apply -f your-app.yaml
  ```
- List pods and services:
  ```bash
  kubectl get pods
  kubectl get services
  ```

### 4. Switch Contexts and Namespaces

- List all contexts:
  ```bash
  kubectx
  ```
- Switch to a context:
  ```bash
  kubectx <context-name>
  ```
- List namespaces:
  ```bash
  kubens
  ```
- Switch namespace:
  ```bash
  kubens <namespace>
  ```

### 5. Stop & Delete the Cluster

```bash
minikube stop
minikube delete
```
