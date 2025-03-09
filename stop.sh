#!/bin/bash

# Exit script on any error
set -e

docker-compose down -v 

echo "Stop deployment"
cd terraform
terraform destroy


