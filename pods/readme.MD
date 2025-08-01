# Kubernetes Pod

## What it is

A Pod is the smallest and simplest unit in the Kubernetes object model that you create or deploy. A Pod represents a single instance of a running process in your cluster and can contain one or more containers, such as Docker containers. When a Pod runs multiple containers, the containers are managed as a single entity and share the Pod's resources, such as storage and networking.

## Key Characteristics

- **Smallest Deployable Unit:** The fundamental building block in Kubernetes.
- **Multi-Container Support:** Can host multiple co-located and co-managed containers that share resources.
- **Shared Context:** Containers within a Pod share the same network namespace (IP address and port space) and can communicate via `localhost`. They can also share storage volumes.
- **Ephemeral Nature:** Pods are considered ephemeral, disposable entities. They are not designed to be long-lived. If a Pod fails, it is not automatically restarted or rescheduled. Instead, higher-level controllers like Deployments or StatefulSets are used to manage Pod lifecycle.
- **Unique IP Address:** Every Pod in a cluster is assigned a unique IP address.

## Common `kubectl` Commands

- **Create a Pod from a manifest file:**

  ```bash
  kubectl apply -f manifest.yaml
  ```

  _Explanation: Creates or updates a Pod using a YAML or JSON manifest file. Note: It's more common to manage Pods via a Deployment, ReplicaSet, or StatefulSet._

- **Get a list of Pods:**

  ```bash
  kubectl get pods -n <namespace> -o wide
  ```

  _Explanation: Lists all Pods in the specified namespace with additional information like the node they are running on._

- **Describe a Pod to see its details:**

  ```bash
  kubectl describe pod <pod-name> -n <namespace>
  ```

  _Explanation: Shows detailed information about a specific Pod, including its status, events, IP address, and containers._

- **View logs from a Pod:**

  ```bash
  # For a single-container pod
  kubectl logs <pod-name> -n <namespace>

  # For a multi-container pod, specify the container name
  kubectl logs <pod-name> -c <container-name> -n <namespace>
  ```

  _Explanation: Fetches the logs from a container within the Pod._

- **Execute a command inside a Pod:**

  ```bash
  kubectl exec -it <pod-name> -n <namespace> -- /bin/sh
  ```

  _Explanation: Opens an interactive shell session inside a running container in the Pod._

- **Delete a Pod:**
  ```bash
  kubectl delete pod <pod-name> -n <namespace>
  ```
  _Explanation: Deletes the specified Pod. If the Pod is managed by a controller, a new one will likely be created to replace it._

## Example Manifest

Here is a basic example of a Pod that runs a single NGINX container:

```yaml
# pod-example.yaml
apiVersion: v1
kind: Pod
metadata:
  name: nginx-pod
  labels:
    app: web
spec:
  containers:
    - name: nginx-container
      image: nginx:1.21.6
      ports:
        - containerPort: 80
          protocol: TCP
  # Pods are not automatically restarted by default if they fail.
  # This behavior is managed by controllers like Deployments.
  restartPolicy: Always
```
