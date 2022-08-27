# metamirror

![Version: 0.0.1](https://img.shields.io/badge/Version-0.0.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.0.1](https://img.shields.io/badge/AppVersion-0.0.1-informational?style=flat-square)

A Kubernetes operator to synchronize selected annotation and label values in running Pods.

## Usage

1. Add the helm repository:

```Bash
helm repo add metamirror https://neilharris123.github.io/metamirror-helm/
```

2. Set the required `MM_ANNOTATION` and `MM_LABEL` environment variables in the metamirror-operator pod (these can be managed through `controller.mmAnnotation` and `controller.mmLabel` in the `values.yaml`):


```Bash
...
controller:
# controller.mmAnnotation -- represents an annotation key present in other pod(s). The operator will copy the corresponding annotation value.
  mmAnnotation: "example/annotation"
# controller.mmLabel -- the label name to be added to pod(s) with the above annotation. The value of the label will be that of the copied annotation value.
  mmLabel: "examplelabel"
```

Alternatively, set these when installing the chart:

```Bash
helm install my-metamirror metamirror/metamirror -set controller.mmAnnotation=example/annotation --set controller.mmLabel=examplelabel
```

3. Test the operator by deploying a seperate pod with the relevent annotation key and an arbitary value:

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

The operator adds a label to the new pod with the relevent values:
```Bash
kubectl get pod test --show-labels

NAME   READY   STATUS    RESTARTS   AGE   LABELS
test   1/1     Running   0          12s   examplelabel=foo
```

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` |  |
  autoscaling.enabled | bool | `false` |
| container.extraEnvs | object | `{}` | Additional environment variables |
| controller.leaderElection | bool | `false` | Enable leader election for running multiple controller pods |
| controller.mmAnnotation | string | `""` | The annotation key present in other pod(s). The operator will copy the corresponding annotation value. |
| controller.mmLabel | string | `""` |  The label name to be added to pod(s) that contain mmAnnotation. The operator will add the copied annotation value as the labels value. |
| fullnameOverride | string | `""` |  |
| image.pullPolicy | string | `"IfNotPresent"` |  |
| image.repository | string | `"neilharris123/metamirror"` |  |
| image.tag | string | `v0.0.1` |  |
| imagePullSecrets | list | `[]` |  |
| nameOverride | string | `""` |  |
| nodeSelector | object | `{}` |  |
| podAnnotations | object | `{}` |  |
| podSecurityContext | object | `{}` |  |
| replicaCount | int | `1` |  |
| resources | object | `{}` |  |
| securityContext | object | `{}` |  |
| tolerations | list | `[]` |  |
