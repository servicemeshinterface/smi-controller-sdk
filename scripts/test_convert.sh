#!/bin/bash

curl -k https://localhost:9533/convert -d \
  '{
      "apiVersion": "admission.k8s.io/v1",
      "kind": "AdmissionReview",
      "request": {
          "uuid":"abc", 
          "kind": {"group":"access.smi-spec.io","version":"v1alpha1","kind":"TrafficTarget"}, 
          "resource": {"group":"access.smi-spec.io","version":"v1alpha1","kind":"traffictargets"}, 
          "operation": "CREATE",
          "name": "my-test",
          "namespace": "default",
          "object": {"a": 2}
        }
    }'
