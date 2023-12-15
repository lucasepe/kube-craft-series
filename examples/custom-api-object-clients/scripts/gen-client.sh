#!/bin/bash

# Go installs in $GOBIN if defined, and $GOPATH/bin otherwise
gobin="${GOBIN:-$(go env GOPATH)/bin}"

# Project root folder
project_root=custom-api-object-clients

# Number for the log level verbosity
v="${VERBOSE:-1}"

# File containing boilerplate header text
boilerplate="./hack/boilerplate.go.txt"

# Destination folder for generated code
clientset_subdir=generated

# Name of the generated clientset
clientset_name="metals"

# Go package root for the generated code
in_pkg_root=github.com/lucasepe/kube-craft-series/examples

# Go package where to create the ${clientset_subdir}
out_pkg_root=${in_pkg_root}/${project_root}

# Go package root segment (used for clean up at the end)
pkg_root_segment=${in_pkg_root%%/*}

# Remove previosly generated folder
rm -rf "./${clientset_subdir}"

# Execute client-gen tool
"${gobin}/client-gen" \
    -v "${v}" \
    --fake-clientset "false" \
    --go-header-file "${boilerplate}" \
    --clientset-name "${clientset_name}" \
    --input-base "./" \
    --output-base "./" \
    --output-package "${out_pkg_root}/${clientset_subdir}" \
    --input "${in_pkg_root}/custom-api-object/apis/metals/v1alpha1"

# HACK to do some clean up!
# Move the folder with generated code
mv "./${out_pkg_root}/${clientset_subdir}" "./"

# Remove the unused "package" folder
rm -r "./${pkg_root_segment}"
