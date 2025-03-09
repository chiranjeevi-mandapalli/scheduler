#!/bin/bash
set -e

echo "building docker image for application"
cd meeting-scheduler
docker build -t scheduler .
cd ..

## get the commands from the instance and update accordingly to the commands
echo "login and push to container registry"
aws ecr get-login-password --region ap-south-1 | docker login --username AWS --password-stdin 718634073122.dkr.ecr.ap-south-1.amazonaws.com
docker tag scheduler:latest 718634073122.dkr.ecr.ap-south-1.amazonaws.com/scheduler:latest
docker push 718634073122.dkr.ecr.ap-south-1.amazonaws.com/scheduler:latest

# Run these manually
aws eks update-kubeconfig --region ap-south-1 --name "stack-gen-eks-cluster"

cd scheduler_deployment_helm/
helm install scheduler .

cd ../postgresql
helm install postgres . -f prod.values.yaml

echo "Deployment complete!"