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

module "rds" {
  source              = "./modules/rds"
  db_instance_name    = "stack-gen-postgres-rds"
  db_username         = "chiru"
  db_password         = "securepassword123"
  db_instance_class   = "db.t3.micro"
  db_allocated_storage = 20
  db_engine           = "postgres"
  db_engine_version   = "15"
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

