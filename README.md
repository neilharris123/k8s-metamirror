# metamirror

A Kubernetes operator to synchronize selected annotation and label values in running Pods.

## Quickstart

This quickstart uses individual yaml files. If you would prefer to install and configure metamirror via helm, see [here](https://artifacthub.io/packages/helm/metamirror/metamirror).

Navigate to the [k8s-resources](https://github.com/neilharris123/metamirror/tree/main/k8s-resources) directory where minimal RBAC and pod manifests are located.

Set the required environment variables in the metamirror-operator [Pod](https://github.com/neilharris123/metamirror/blob/main/k8s-resources/metamirror-operator.yaml) manifest:

```Bash
...
  env:
  - name: MM_ANNOTATION
    value: "example/annotation" # an annotation key present in other pod(s). The operator will copy the corresponding annotation value.
  - name: MM_LABEL
    value: "examplelabel"       # the label name to be added to pod(s) with the annotation. The value of the label will be the same as the copied annotation value.
```
Deploy all metamirror quickstart resources:

```Bash
kubectl create -f clusterrole.yaml -f clusterrolebinding.yaml -f metamirror-operator.yaml -f serviceaccount.yaml
```

## Testing

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

## Results

The operator adds a label to the new Pod with the relevent values:
```Bash
kubectl get pod test --show-labels

NAME   READY   STATUS    RESTARTS   AGE   LABELS
test   1/1     Running   0          12s   examplelabel=foo
```

## Limitations

The current version of the operator can only mirror a single annotation/label value. A future release will allow synchronization of multiple values.
