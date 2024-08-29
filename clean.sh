#!/bin/bash



images=(
    "auth-service-builder:latest"
    "auth-service-status:latest"
)


for i in "${images[@]}"; do
   docker rmi $i --force
done