# metamirror

A Kubernetes operator to synchronize selected annotation and label values in running Pods.

## Quickstart

This quickstart uses individual yaml files. If you would prefer to install and configure metamirror via helm, see [here](https://artifacthub.io/packages/helm/metamirror/metamirror).

Navigate to the [k8s-resources](https://github.com/neilharris123/metamirror/tree/main/k8s-resources) directory where minimal RBAC and pod manifests are located.

Set the MM_ANNOTATIONS and MM_LABELS environment variables for the metamirror-operator [Pod](https://github.com/neilharris123/metamirror/blob/main/k8s-resources/metamirror-operator.yaml). Multiple values should be seperated by a single comma, with no whitespace. The order matters, as the values are converted into slices, and the operator modifies values based on their index position. For example, the value of annotation in position [0] is added as the value of the label in position [0]. For this reason, The number of annotation and label values added to the env vars must be equal, otherwise the config is invalid, and the controller will panic and exit:

```Bash
...
  env:
  - name: MM_ANNOTATIONS
    value: "environment/annotation,project/annotation" # annotation keys present in other pod(s) (multiple values should be seperated by a comma). The operator will copy the corresponding annotation values.
  - name: MM_LABELS
    value: "environment,project" # the label names to be added to pod(s) deployed with any of the MM_ANNOTATIONS. The value of the labels will be the same as the copied annotation values.
```
Deploy all metamirror quickstart resources:

```Bash
kubectl create -f clusterrole.yaml -f clusterrolebinding.yaml -f metamirror-operator.yaml -f serviceaccount.yaml
```

## Testing

Deploy a seperate pod with the relevent annotation(s):

```Bash
apiVersion: v1
kind: Pod
metadata:
  annotations:
    environemnt/annotation: "dev"
    project/annotation: "myproject"
  labels: {}
  name: test
spec:
  containers:
  - image: alpine:3.16
    imagePullPolicy: Always
    command: ["/bin/sh", "-ec", "sleep 1000"]
    name: test
```

## Results

The operator adds labels to the new Pod with the relevent values:
```Bash
kubectl get pod test --show-labels

NAME   READY   STATUS    RESTARTS   AGE   LABELS
test   1/1     Running   0          12s   environment=dev,project=myproject
```
