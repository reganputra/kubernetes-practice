# Kubernetes Resource Management

## What it is

In Kubernetes, you can specify how much CPU and memory (RAM) each container in a Pod needs. When you specify the resource _request_ for containers in a Pod, the scheduler uses this information to decide which node to place the Pod on. When you specify a resource _limit_ for a container, the kubelet enforces those limits so that the running container is not allowed to use more than the limit you set. This is crucial for managing cluster resources effectively, ensuring application stability, and preventing any single application from consuming all available resources.

## Key Concepts

- **Requests:** The amount of resources that are guaranteed for a container. If a container requests a resource, Kubernetes will only schedule it on a node that can provide that amount.

  - `cpu`: Measured in "cpu units". `100m` represents 0.1 CPU cores. `1` represents 1 full CPU core.
  - `memory`: Measured in bytes. You can use suffixes like `Mi` (Mebibytes) or `Gi` (Gibibytes).

- **Limits:** The maximum amount of resources that a container is allowed to use.

  - `cpu`: If a container tries to use more CPU than its limit, it will be throttled.
  - `memory`: If a container tries to use more memory than its limit, it will be terminated (OOMKilled).

- **Quality of Service (QoS) Classes:** Kubernetes categorizes Pods into QoS classes based on their resource requests and limits. This determines how they are prioritized for scheduling and eviction.

  - **Guaranteed:** Pods where every container has a memory and CPU request that matches its limit. These are the highest priority.
  - **Burstable:** Pods where at least one container has a CPU or memory request that is less than its limit.
  - **BestEffort:** Pods with no CPU or memory requests or limits. These are the lowest priority.

- **ResourceQuota:** A tool for administrators to constrain resource consumption per namespace. It can limit the total amount of CPU and memory that can be requested or limited by all Pods in a namespace.

- **LimitRange:** A policy to constrain resource allocations (to Pods or Containers) in a namespace. It can set default request/limit values for containers that don't specify their own, and enforce minimum/maximum resource values.

## Common `kubectl` Commands

- **Describe a node to see its resource capacity and allocations:**

  ```bash
  kubectl describe node <node-name>
  ```

  _Explanation: Shows how much CPU and memory is available on a node and how much is currently requested by running Pods._

- **View resource usage for Pods:**

  ```bash
  # Requires metrics-server to be installed
  kubectl top pods -n <namespace>
  ```

  _Explanation: Displays the current CPU and memory consumption of Pods._

- **Get a list of ResourceQuotas:**

  ```bash
  kubectl get resourcequotas -n <namespace>
  ```

  _Explanation: Lists all ResourceQuotas in the specified namespace._

- **Describe a ResourceQuota:**
  ```bash
  kubectl describe resourcequota <quota-name> -n <namespace>
  ```
  _Explanation: Shows the details of a ResourceQuota, including its limits and current usage._

## Example Manifests

### Pod with Resource Requests and Limits

```yaml
# pod-with-resources.yaml
apiVersion: v1
kind: Pod
metadata:
  name: frontend
spec:
  containers:
    - name: app
      image: images.my-company.example/app:v4
      resources:
        requests:
          memory: "64Mi"
          cpu: "250m" # 0.25 CPU core
        limits:
          memory: "128Mi"
          cpu: "500m" # 0.5 CPU core
```

### ResourceQuota for a Namespace

```yaml
# quota-example.yaml
apiVersion: v1
kind: ResourceQuota
metadata:
  name: mem-cpu-quota
  namespace: my-namespace
spec:
  hard:
    requests.cpu: "1"
    requests.memory: 1Gi
    limits.cpu: "2"
    limits.memory: 2Gi
```

### LimitRange for a Namespace

```yaml
# limitrange-example.yaml
apiVersion: v1
kind: LimitRange
metadata:
  name: mem-limit-range
  namespace: my-namespace
spec:
  limits:
    - default:
        memory: "512Mi"
        cpu: "500m"
      defaultRequest:
        memory: "256Mi"
        cpu: "250m"
      type: Container
```
