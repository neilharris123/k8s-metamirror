# metamirror

![Version: 0.0.2](https://img.shields.io/badge/Version-0.0.2-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.0.2](https://img.shields.io/badge/AppVersion-0.0.2-informational?style=flat-square)

A Kubernetes operator to synchronize selected annotation and label values in running Pods.

## Usage

1. Add the helm repository:

```Bash
helm repo add metamirror https://neilharris123.github.io/metamirror-helm/
```

2. Set the required `MM_ANNOTATIONS` and `MM_LABELS` environment variables in the metamirror-operator pod (these can be managed through `controller.mmAnnotations` and `controller.mmLabels` in the `values.yaml`):


```Bash
...
controller:
# controller.mmAnnotations -- represents annotation keys present in other pod(s). The operator will copy the corresponding annotation values. Multiple values should be seperated by a single comma.
  mmAnnotations: "environment/annotation,project/annotation"
# controller.mmLabels -- the label names to be added to pod(s) with the above annotations. The value of the labels will be that of the copied annotation values. Multiple values should be seperated by a single comma.
  mmLabels: "environment,project"
```

Alternatively, set these when installing the chart:

```Bash
helm install my-metamirror metamirror/metamirror --set controller.mmAnnotations=environment/annotation --set controller.mmLabels=environment
```

3. Test the operator by deploying a seperate pod with the relevent annotation key and any values of interest:

```Bash
apiVersion: v1
kind: Pod
metadata:
  annotations:
    environment/annotation: "dev"
    project/annotation: "myproject"
  labels: {}
  name: test
spec:
  containers:
  - image: alpine:3.16
    imagePullPolicy: Always
    name: test
```

### Results

The operator adds label to the new pod with the relevent values:
```Bash
kubectl get pod test --show-labels

NAME   READY   STATUS    RESTARTS   AGE   LABELS
test   1/1     Running   0          12s   environment=dev,project=myproject
```

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` |  |
  autoscaling.enabled | bool | `false` |
| container.extraEnvs | object | `{}` | Additional environment variables |
| controller.leaderElection | bool | `false` | Enable leader election for running multiple controller pods |
| controller.mmAnnotations | string | `""` | The annotation key(s) present in other pods you would like to modify. You can either add a single annotation key, or multiple annotation keys seperated by commas. The operator will copy the corresponding annotation values. |
| controller.mmLabels | string | `""` |  The label names to be added to pods that contain any of the mmAnnotations. Multiple values should be seperated by a comma. The operator will add the copied annotation value as the labels value. |
| fullnameOverride | string | `""` |  |
| image.pullPolicy | string | `"IfNotPresent"` |  |
| image.repository | string | `"neilharris123/metamirror"` |  |
| image.tag | string | `v0.0.2` |  |
| imagePullSecrets | list | `[]` |  |
| nameOverride | string | `""` |  |
| nodeSelector | object | `{}` |  |
| podAnnotations | object | `{}` |  |
| podSecurityContext | object | `{}` |  |
| replicaCount | int | `1` |  |
| resources | object | `{}` |  |
| securityContext | object | `{}` |  |
| tolerations | list | `[]` |  |
