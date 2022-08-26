# metamirror

A Kubernetes operator to synchronize selected annotation and label values in running Pods.

## Usage

The operator requires `get`, `list`, `patch`, `update` and `watch` permissions on Pods cluster wide. Minimal RBAC resource and operator manifests can be found in [k8s-resources](https://github.com/neilharris123/metamirror/tree/main/k8s-resources).

Set the required environment variables in the metamirror-operator pod:

```Bash
...
  env:
  - name: MM_ANNOTATION
    value: "example/annotation" # an annotation key present in other pod(s). The operator will copy the corresponding annotation value.
  - name: MM_LABEL
    value: "examplelabel"       # the label name to be added to pod(s). The value of the label will be that of the copied annotation value.
```

Deploy a seperate pod with the relevent annotation key and an arbitary value:

```Bash
apiVersion: v1
kind: Pod
metadata:
  annotations:
    example/annotation: "foo"
  labels: {}
  name: test
spec:
  containers:
  - image: alpine:3.16
    imagePullPolicy: Always
    name: test
```

### Results

The operator adds a label to the new pod, combining the MM_LABEL env var and extracted annotation value.
```Bash
kubectl get pod test --show-labels

NAME   READY   STATUS    RESTARTS   AGE   LABELS
test   1/1     Running   0          12s   examplelabel=foo
```
