---
layout: "mongodbatlas"
page_title: "MongoDB Atlas: project_ip_whitelist"
sidebar_current: "docs-mongodbatlas-resource-project-ip-whitelist"
description: |-
    Provides an IP Whitelist resource.
---

# mongodbatlas_project_ip_whitelist

`mongodbatlas_project_ip_whitelist` provides an IP Whitelist entry resource. The whitelist grants access from IPs or CIDRs to clusters within the Project.

-> **NOTE:** Groups and projects are synonymous terms. You may find `groupId` in the official documentation.

## Examples Usage

### Using Cidr Block
```hcl
resource "mongodbatlas_project_ip_whitelist" "test" {
  project_id = "<PROJECT-ID>"
  cidr_block = "1.2.3.4/32"
  comment    = "cidr block for tf acc testing"
}
```

### Using IP Address
```hcl
resource "mongodbatlas_project_ip_whitelist" "test" {
  project_id = "<PROJECT-ID>"
  ip_address = "2.3.4.5"
  comment    = "ip address for tf acc testing"
}
```

### Using AWS Security Group
```hcl
resource "mongodbatlas_network_container" "test" {
  project_id       = "<PROJECT-ID>"
  atlas_cidr_block = "192.168.208.0/21"
  provider_name    = "AWS"
  region_name      = "US_EAST_1"
}

resource "mongodbatlas_network_peering" "test" {
  project_id             = "<PROJECT-ID>"
  container_id           = mongodbatlas_network_container.test.container_id
  accepter_region_name   = "us-east-1"
  provider_name          = "AWS"
  route_table_cidr_block = "172.31.0.0/16"
  vpc_id                 = "vpc-0d93d6f69f1578bd8"
  aws_account_id         = "232589400519"
}

resource "mongodbatlas_project_ip_whitelist" "test" {
  project_id         = "<PROJECT-ID>"
  aws_security_group = "sg-0026348ec11780bd1"
  comment            = "TestAcc for awsSecurityGroup"

  depends_on = ["mongodbatlas_network_peering.test"]
}
```

## Argument Reference

* `project_id` - (Required) The ID of the project in which to add the whitelist entry.
* `aws_security_group` - (Optional) ID of the whitelisted AWS security group. Mutually exclusive with `cidr_block` and `ip_address`.
* `cidr_block` - (Optional) Whitelist entry in Classless Inter-Domain Routing (CIDR) notation. Mutually exclusive with `aws_security_group` and `ip_address`.
* `ip_address` - (Optional) Whitelisted IP address. Mutually exclusive with `aws_security_group` and `cidr_block`.
* `comment` - (Optional) Comment to add to the whitelist entry.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Unique identifier used for terraform for internal manages and can be used to import.

## Import

IP Whitelist entries can be imported using the `project_id`, e.g.

```
$ terraform import mongodbatlas_project_ip_whitelist.test 5d0f1f74cf09a29120e123cd
```

For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/whitelist/)