# Kubernetes ConfigMap

## What it is

A ConfigMap is a Kubernetes API object used to store non-confidential data in key-value pairs. Pods can consume ConfigMaps as environment variables, command-line arguments, or as configuration files in a volume. This allows you to decouple configuration artifacts from image content to keep containerized applications portable.

## Key Characteristics

- **Decoupling:** Separates configuration from application code.
- **Flexibility:** Can be consumed by pods in various ways (environment variables, volume mounts).
- **Centralized Management:** Allows for easier management and updates of configuration data across multiple pods.
- **Non-sensitive Data:** Intended for non-sensitive data. For sensitive data like credentials or keys, use Secrets.
- **Size Limit:** The data size in a ConfigMap is limited to 1 MiB.

## Common `kubectl` Commands

- **Create a ConfigMap from a file or literal values:**

  ```bash
  # From a file
  kubectl create configmap <configmap-name> --from-file=<path-to-file>

  # From literal values
  kubectl create configmap <configmap-name> --from-literal=<key1>=<value1> --from-literal=<key2>=<value2>
  ```

  _Explanation: Creates a new ConfigMap with the specified name and data._

- **Apply a ConfigMap from a manifest file:**

  ```bash
  kubectl apply -f manifest.yaml
  ```

  _Explanation: Creates or updates a ConfigMap using a YAML or JSON manifest file._

- **Get a list of ConfigMaps:**

  ```bash
  kubectl get configmaps -n <namespace>
  ```

  _Explanation: Lists all ConfigMaps in the specified namespace._

- **Describe a ConfigMap to see its details:**

  ```bash
  kubectl describe configmap <configmap-name> -n <namespace>
  ```

  _Explanation: Shows detailed information about a specific ConfigMap, including its data._

- **Edit a ConfigMap:**

  ```bash
  kubectl edit configmap <configmap-name> -n <namespace>
  ```

  _Explanation: Opens the ConfigMap's manifest in the default editor for live edits._

- **Delete a ConfigMap:**
  ```bash
  kubectl delete configmap <configmap-name> -n <namespace>
  ```
  _Explanation: Deletes the specified ConfigMap._

## Example Manifest

Here is a basic example of a ConfigMap manifest:

```yaml
# configmap-example.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: app-config
  namespace: default
data:
  # Key-value pairs
  APP_COLOR: "blue"
  APP_MODE: "production"
  # File-like key
  app.properties: |
    database.host=mysql.example.com
    database.port=3306
    feature.enabled=true
```
