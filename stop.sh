#!/bin/bash

# Exit script on any error
set -e

echo "Stop deployment"
cd terraform
terraform destroy --auto-approve


