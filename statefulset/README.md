# Kubernetes StatefulSet

## What it is

A StatefulSet is a Kubernetes workload API object used to manage stateful applications. It manages the deployment and scaling of a set of Pods, and provides guarantees about the ordering and uniqueness of these Pods. Unlike a Deployment, a StatefulSet maintains a sticky identity for each of its Pods. These pods are created from the same spec, but are not interchangeable: each has a persistent identifier that it maintains across any rescheduling.

## Key Characteristics

- **Stable, Unique Network Identifiers:** Each Pod in a StatefulSet has a stable, unique hostname based on its ordinal index (e.g., `web-0`, `web-1`).
- **Stable, Persistent Storage:** Each Pod gets its own persistent storage volume that is tied to its identity. If the Pod is rescheduled, the new Pod will be connected to the same storage.
- **Ordered, Graceful Deployment and Scaling:** Pods are created, updated, and deleted in a strict, ordered sequence. For example, when scaling up, `web-0` is deployed before `web-1`. When scaling down, `web-n` is terminated before `web-n-1`.
- **Ordered, Automated Rolling Updates:** Updates to StatefulSets are rolled out in order, from the highest ordinal to the lowest.
- **Headless Service Requirement:** StatefulSets require a Headless Service to control the domain of its Pods and provide the stable network identities.

## Common `kubectl` Commands

- **Create a StatefulSet from a manifest file:**

  ```bash
  kubectl apply -f manifest.yaml
  ```

  _Explanation: Creates or updates a StatefulSet and its associated Headless Service._

- **Get a list of StatefulSets:**

  ```bash
  kubectl get statefulsets -n <namespace>
  ```

  _Explanation: Lists all StatefulSets in the specified namespace._

- **Describe a StatefulSet to see its status:**

  ```bash
  kubectl describe statefulset <statefulset-name> -n <namespace>
  ```

  _Explanation: Shows detailed information about a StatefulSet, including its current and desired replica counts and events._

- **Scale a StatefulSet:**

  ```bash
  kubectl scale statefulset <statefulset-name> --replicas=<number> -n <namespace>
  ```

  _Explanation: Changes the number of Pods in the StatefulSet, following the ordered scaling guarantees._

- **Check the rollout status of a StatefulSet:**

  ```bash
  kubectl rollout status statefulset/<statefulset-name> -n <namespace>
  ```

  _Explanation: Watches the status of a StatefulSet's rolling update until it's complete._

- **Delete a StatefulSet:**
  ```bash
  kubectl delete statefulset <statefulset-name> -n <namespace>
  ```
  _Explanation: Deletes the StatefulSet. The associated PersistentVolumeClaims (PVCs) are not deleted by default, to prevent data loss._

## Example Manifest

Here is a basic example of a StatefulSet that runs a simple web application, along with its required Headless Service.

```yaml
# statefulset-example.yaml

# First, define the Headless Service for network identity
apiVersion: v1
kind: Service
metadata:
  name: nginx-headless-service
  labels:
    app: nginx
spec:
  ports:
    - port: 80
      name: web
  # clusterIP: None makes this a "headless" service
  clusterIP: None
  selector:
    app: nginx
---
# Then, define the StatefulSet
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: web-statefulset
spec:
  selector:
    matchLabels:
      app: nginx # has to match .spec.template.metadata.labels
  serviceName: "nginx-headless-service"
  replicas: 3 # by default is 1
  template:
    metadata:
      labels:
        app: nginx # has to match .spec.selector.matchLabels
    spec:
      terminationGracePeriodSeconds: 10
      containers:
        - name: nginx
          image: k8s.gcr.io/nginx-slim:0.8
          ports:
            - containerPort: 80
              name: web
          volumeMounts:
            - name: www
              mountPath: /usr/share/nginx/html
  # Define a PersistentVolumeClaim template.
  # This will create a new PVC for each Pod.
  volumeClaimTemplates:
    - metadata:
        name: www
      spec:
        accessModes: ["ReadWriteOnce"]
        storageClassName: "my-storage-class" # Make sure you have a StorageClass
        resources:
          requests:
            storage: 1Gi
```
