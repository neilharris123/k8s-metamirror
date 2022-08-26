# metamirror

![Version: 0.0.1](https://img.shields.io/badge/Version-0.0.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.0.1](https://img.shields.io/badge/AppVersion-0.0.1-informational?style=flat-square)

A Kubernetes operator to synchronize selected annotation and label values in running Pods.

## Source Code

* <https://github.com/neilharris123/metamirror>

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
