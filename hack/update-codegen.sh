#!/usr/bin/env bash

vendor/k8s.io/code-generator/generate-groups.sh all \
  github.com/sakiib/crd/pkg/client github.com/sakiib/crd/pkg/apis \
  book.com:v1alpha1
