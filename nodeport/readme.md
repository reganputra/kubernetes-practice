# Kubernetes Service: NodePort

## What it is

A `NodePort` is a type of Kubernetes Service that exposes an application on a static port on each Node's IP address. When you create a `NodePort` service, Kubernetes allocates a port from a configurable range (typically 30000-32767), and any traffic sent to that port on any Node is forwarded to the service. This makes the service accessible from outside the cluster by using `<NodeIP>:<NodePort>`.

## Key Characteristics

- **External Exposure:** Provides a simple way to expose a service to the outside world for development, testing, or demonstration purposes.
- **Static Port on Each Node:** The same port is opened on every node in the cluster.
- **Builds on ClusterIP:** A `NodePort` service is an extension of `ClusterIP`. When you create a `NodePort` service, Kubernetes automatically creates a `ClusterIP` service as well, so the `NodePort` service is also accessible internally via its `ClusterIP`.
- **Limited Port Range:** The port must be within the configured `nodePort` range, which limits its use for well-known ports (like 80 or 443).
- **Not for Production:** While useful, `NodePort` is not typically recommended for production use cases. It couples the service to specific nodes and can be inefficient. For production, `LoadBalancer` or `Ingress` are the preferred methods for exposing services.

## Common Use Cases

- **Development and Testing:** Quickly exposing an application during the development lifecycle without setting up a full load balancer.
- **Demos:** Providing a simple, stable endpoint for demonstrating an application.
- **When an External Load Balancer is Not Available:** In on-premises or bare-metal environments where a cloud provider's load balancer is not an option.

## Common `kubectl` Commands

- **Create a Service from a manifest file:**

  ```bash
  kubectl apply -f service-manifest.yaml
  ```

  _Explanation: Creates or updates a Service using a YAML manifest file._

- **Get a list of Services:**

  ```bash
  kubectl get services -n <namespace>
  ```

  _Explanation: Lists all Services, showing their type, ClusterIP, and the NodePort mapping._

- **Describe a Service to see its details:**

  ```bash
  kubectl describe service <service-name> -n <namespace>
  ```

  _Explanation: Shows detailed information about a Service, including its NodePort, endpoints, and selector._

- **Get the IP address of your nodes:**

  ```bash
  kubectl get nodes -o wide
  ```

  _Explanation: Shows the internal and external IP addresses of the cluster nodes, which you can use to access the NodePort service._

- **Delete a Service:**
  ```bash
  kubectl delete service <service-name> -n <namespace>
  ```
  _Explanation: Deletes the specified Service._

## Example Manifest

Here is an example of a `NodePort` Service that exposes a web server running in a set of Pods with the label `app: my-web-app`.

```yaml
# web-app-service.yaml
apiVersion: v1
kind: Service
metadata:
  name: my-web-app-service
spec:
  # Set the service type to NodePort
  type: NodePort
  selector:
    # This selector must match the labels of the Pods you want to expose.
    app: my-web-app
  ports:
    - protocol: TCP
      # The port on the Pods that the traffic will be forwarded to.
      targetPort: 80
      # The port that the Service will be exposed on internally (ClusterIP).
      port: 80
      # Optional: Specify a static port for the NodePort.
      # If not specified, Kubernetes will allocate a free port from the range.
      nodePort: 30007
```
