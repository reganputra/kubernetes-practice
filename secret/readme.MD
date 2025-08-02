# Kubernetes Secret

## What it is

A Secret is a Kubernetes object that lets you store and manage sensitive information, such as passwords, OAuth tokens, and ssh keys. Storing confidential data in a Secret is safer and more flexible than putting it verbatim in a Pod definition or in a container image. A Pod can reference the Secret and use its data, for example, as environment variables or as files in a volume.

## Key Characteristics

- **Sensitive Data:** Designed specifically for storing confidential data.
- **Encoding:** Data is stored in a base64 encoded format. This is for encoding, not encryption, and can be easily decoded. For stronger protection, other mechanisms like encryption at rest should be configured.
- **Consumption by Pods:** Pods can consume Secrets as environment variables or as files mounted into the container's filesystem.
- **Decoupling:** Separates sensitive data from the application code and Pod specifications, improving portability and security.
- **Types:** Secrets can have different types (`Opaque`, `kubernetes.io/service-account-token`, `kubernetes.io/dockercfg`, etc.) which help Kubernetes understand how to handle the data. `Opaque` is the default type.

## Common `kubectl` Commands

- **Create a Secret from literal values:**

  ```bash
  kubectl create secret generic <secret-name> --from-literal=<key1>=<value1> --from-literal=<key2>=<value2>
  ```

  _Explanation: Creates a new Secret of type `generic` with the specified key-value pairs._

- **Create a Secret from a file:**

  ```bash
  kubectl create secret generic <secret-name> --from-file=<path-to-file>
  ```

  _Explanation: Creates a new Secret where the key is the filename and the value is the file's content._

- **Apply a Secret from a manifest file:**

  ```bash
  kubectl apply -f manifest.yaml
  ```

  _Explanation: Creates or updates a Secret using a YAML manifest. The data must be base64 encoded._

- **Get a list of Secrets:**

  ```bash
  kubectl get secrets -n <namespace>
  ```

  _Explanation: Lists all Secrets in the specified namespace._

- **Describe a Secret:**

  ```bash
  kubectl describe secret <secret-name> -n <namespace>
  ```

  _Explanation: Shows metadata about a Secret, but does not display the decoded data._

- **Get the decoded data from a Secret:**

  ```bash
  kubectl get secret <secret-name> -n <namespace> -o jsonpath='{.data.<key>}' | base64 --decode
  ```

  _Explanation: Retrieves and decodes a specific value from a Secret._

- **Delete a Secret:**
  ```bash
  kubectl delete secret <secret-name> -n <namespace>
  ```
  _Explanation: Deletes the specified Secret._

## Example Manifest

Here is a basic example of a Secret manifest. The values for `username` and `password` must be base64 encoded.

```yaml
# secret-example.yaml
apiVersion: v1
kind: Secret
metadata:
  name: my-app-secret
type: Opaque # Default type
data:
  # The values are base64 encoded strings.
  # echo -n 'admin' | base64  => YWRtaW4=
  # echo -n 'p@ssw0rd1' | base64 => cEBzc3cwcmQx
  username: YWRtaW4=
  password: cEBzc3cwcmQx
```
