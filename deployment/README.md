# Kubernetes Deployment

## What it is

A Deployment is a Kubernetes resource that provides declarative updates for Pods and ReplicaSets. You describe a desired state in a Deployment, and the Deployment Controller changes the actual state to the desired state at a controlled rate. You can define Deployments to create new ReplicaSets, or to remove existing Deployments and adopt all their resources with new Deployments.

## Key Characteristics

- **Declarative Updates:** You define the desired state, and Kubernetes handles the rest.
- **Rolling Updates:** Allows for zero-downtime updates by incrementally replacing old Pods with new ones.
- **Rollbacks:** You can easily revert to a previous version of a Deployment if something goes wrong.
- **Scaling:** You can scale the number of replicas up or down as needed.
- **Self-healing:** If a Pod fails, the Deployment will automatically create a new one to maintain the desired number of replicas.

## Common `kubectl` Commands

- **Create a Deployment from a manifest file:**

  ```bash
  kubectl apply -f manifest.yaml
  ```

  _Explanation: Creates or updates a Deployment using a YAML or JSON manifest file._

- **Get a list of Deployments:**

  ```bash
  kubectl get deployments -n <namespace>
  ```

  _Explanation: Lists all Deployments in the specified namespace._

- **Describe a Deployment to see its status:**

  ```bash
  kubectl describe deployment <deployment-name> -n <namespace>
  ```

  _Explanation: Shows detailed information about a Deployment, including its strategy, replicas, and events._

- **Scale a Deployment:**

  ```bash
  kubectl scale deployment <deployment-name> --replicas=<number> -n <namespace>
  ```

  _Explanation: Changes the number of Pods managed by the Deployment._

- **Check the rollout status of a Deployment:**

  ```bash
  kubectl rollout status deployment/<deployment-name> -n <namespace>
  ```

  _Explanation: Watches the status of a Deployment's rolling update until it's complete._

- **View rollout history:**

  ```bash
  kubectl rollout history deployment/<deployment-name> -n <namespace>
  ```

  _Explanation: Shows the history of revisions for a Deployment._

- **Rollback to a previous revision:**

  ```bash
  # Rollback to the previous version
  kubectl rollout undo deployment/<deployment-name> -n <namespace>

  # Rollback to a specific revision
  kubectl rollout undo deployment/<deployment-name> --to-revision=<revision-number> -n <namespace>
  ```

  _Explanation: Reverts the Deployment to a previous state._

- **Delete a Deployment:**
  ```bash
  kubectl delete deployment <deployment-name> -n <namespace>
  ```
  _Explanation: Deletes the specified Deployment, which also deletes its associated ReplicaSets and Pods._

## Example Manifest

Here is a basic example of a Deployment that runs three replicas of an NGINX web server:

```yaml
# deployment-example.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
        - name: nginx
          image: nginx:1.14.2
          ports:
            - containerPort: 80
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 25%
      maxSurge: 25%
```
