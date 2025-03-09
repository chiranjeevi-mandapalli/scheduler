#!/bin/bash
set -e

echo "Initializing Terraform..."
cd terraform
terraform init

echo "Validate Terraform ..."
terraform validate

echo "Running Terraform Plan..."
terraform plan

echo "Applying Terraform..."
terraform apply -auto-approve
