resource "aws_db_instance" "rds" {
  identifier            = var.db_instance_name
  allocated_storage     = var.db_allocated_storage
  storage_type          = "gp2"
  engine               = var.db_engine
  engine_version       = var.db_engine_version
  instance_class       = var.db_instance_class
  username            = var.db_username
  password            = var.db_password
  publicly_accessible  = true
  skip_final_snapshot  = true
}
