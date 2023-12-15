#!/bin/bash

CRD_DIR=./crds
API_DIR=./apis

# Remove previously generated CRDs
rm -f ${CRD_DIR}/*.yaml 

# Generate DeepCopy methods and the CRD
controller-gen paths=${API_DIR}/... object crd output:crd:dir=${CRD_DIR}