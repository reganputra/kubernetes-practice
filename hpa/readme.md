# Horizontal Pod Autoscaler (HPA)

## What it is

The Horizontal Pod Autoscaler (HPA) is a Kubernetes resource that automatically scales the number of pod replicas in a Deployment, ReplicaSet, or StatefulSet based on observed CPU utilization or other select metrics. The HPA is implemented as a control loop that periodically queries the resource utilization against the metrics specified in each HPA definition. Its goal is to increase the number of pods when the load is high and decrease it when the load is low, ensuring performance and resource efficiency.

## Key Characteristics

- **Metric-Based Scaling:** Can scale based on CPU and memory usage, as well as custom metrics exposed via the custom metrics API.
- **Target-Based:** You define a target value for a metric (e.g., 80% CPU utilization), and the HPA adjusts the number of replicas to maintain that target.
- **Configurable Limits:** You specify the minimum and maximum number of replicas the autoscaler can set.
- **Cooldown Period:** Includes a configurable stabilization window to prevent rapid scaling up and down (thrashing).
- **Works with Controllers:** It targets scalable controllers like Deployments and StatefulSets, not individual Pods.

## Common `kubectl` Commands

- **Create an HPA imperatively:**

  ```bash
  kubectl autoscale deployment <deployment-name> --cpu-percent=80 --min=1 --max=10
  ```

  _Explanation: Creates an HPA that targets the specified deployment, scaling it between 1 and 10 replicas to maintain an average CPU utilization of 80%._

- **Apply an HPA from a manifest file:**

  ```bash
  kubectl apply -f hpa-manifest.yaml
  ```

  _Explanation: Creates or updates an HPA using a YAML manifest file._

- **Get a list of HPAs:**

  ```bash
  kubectl get hpa -n <namespace>
  ```

  _Explanation: Lists all HPAs in the specified namespace._

- **Describe an HPA to see its status and events:**

  ```bash
  kubectl describe hpa <hpa-name> -n <namespace>
  ```

  _Explanation: Shows detailed information about an HPA, including current metrics, target values, and scaling events._

- **Delete an HPA:**
  ```bash
  kubectl delete hpa <hpa-name> -n <namespace>
  ```
  _Explanation: Deletes the specified HPA. The workload it was managing will stop autoscaling and remain at its current replica count._

## Example Manifest

Here is a basic example of an HPA that targets a Deployment named `php-apache`.

```yaml
# hpa-example.yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: php-apache-hpa
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: php-apache # The name of the deployment to scale
  minReplicas: 1
  maxReplicas: 10
  metrics:
    - type: Resource
      resource:
        name: cpu
        target:
          type: Utilization
          # Target average CPU utilization at 50%
          averageUtilization: 50
```
