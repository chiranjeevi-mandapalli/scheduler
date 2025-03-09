provider "aws" {
  region = "ap-south-1"
}

# Fetch the default VPC or use a specific one
data "aws_vpc" "selected" {
  default = true  # Set to false if using a specific VPC, then use `id = "<your-vpc-id>"`
}

# Get all subnets within the selected VPC
data "aws_subnets" "selected" {
  filter {
    name   = "vpc-id"
    values = [data.aws_vpc.selected.id]
  }
}

module "eks" {
    source = "./modules/eks"
    cluster_name = "stack-gen-eks-cluster"
    cluster_version = "1.29"
    subnet_ids = data.aws_subnets.selected.ids
}


module "ecr" {
  source        = "./modules/ecr"
  ecr_repo_name = "scheduler"
}

