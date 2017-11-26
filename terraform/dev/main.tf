terraform {
  required_version = "> 0.10.0"
  backend "s3" {
    bucket     = "austin1237-gifbot-state-dev"
    key        = "global/s3/terraform.tfstate"
    region     = "us-east-1"
    encrypt    = "true"
    lock_table = "gifbot-state-lock-dev"
  }
}

provider "aws" {
  region = "${var.region}"
}

# ---------------------------------------------------------------------------------------------------------------------
# CREATE THE ECS CLUSTER
# ---------------------------------------------------------------------------------------------------------------------

module "ecs_cluster" {
  source = "./ecs-cluster"

  name = "gifbot-ecs-${var.env}"
  size = 1
  instance_type = "t2.nano"
  key_pair_name = "${var.key_pair_name}"

  vpc_id = "${data.aws_vpc.default.id}"
  subnet_ids = ["${data.aws_subnet.default.*.id}"]

}

# ---------------------------------------------------------------------------------------------------------------------
# CREATE THE GO Discord Bot service
# ---------------------------------------------------------------------------------------------------------------------

module "gifbot" {
  source = "./ecs-service"

  name = "gifbot-${var.env}"
  ecs_cluster_id = "${module.ecs_cluster.ecs_cluster_id}"
  
  image = "${var.gifbot_image}"
  version = "${var.gifbot_version}"
  cpu = 1024
  memory = 400
  desired_count = 1
  
  container_port = "${var.gifbot_port}"
  host_port = "${var.gifbot_port}"

  num_env_vars = 1
  env_vars = "${map("BOT_TOKEN", "${var.BOT_TOKEN_DEV}")}"
}

# ---------------------------------------------------------------------------------------------------------------------
# DEPLOY THIS EXAMPLE IN THE DEFAULT SUBNETS OF THE DEFAULT VPC
# To keep this example as easy to use as possible, we deploy into the default subnets of your default VPC. That means
# everything is accessible publicy, which is fine for learning/experimenting, but NOT a good practice for production.
# In real world use cases, you should run your code in the private subnets of a custom VPC.
#
# Note that if you do not have a default VPC (i.e. you have an older AWS account or you deleted the VPC), you will
# need to manually fill in the VPC and subnet IDs above.
# ---------------------------------------------------------------------------------------------------------------------

data "aws_vpc" "default" {
  default = true
}

data "aws_availability_zones" "available" {}

# Look up the default subnets in the AZs available to this account (up to a max of 3)
data "aws_subnet" "default" {
  count = "${min(length(data.aws_availability_zones.available.names), 3)}"
  default_for_az = true
  vpc_id = "${data.aws_vpc.default.id}"
  availability_zone = "${element(data.aws_availability_zones.available.names, count.index)}"
}