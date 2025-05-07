#!/bin/bash
# generate_pod_yaml.sh
POD_NAME="fasthttp-pod"

cat <<EOF > pod.yaml
apiVersion: v1
kind: Pod
metadata:
  name: $POD_NAME
spec:
  containers:
    - name: fasthttp
      image: fasthttp-server
      ports:
        - containerPort: 8080
EOF

echo "pod.yaml generated for $POD_NAME"
