---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "powerplatform Provider"
subcategory: ""
description: |-
  
---

# powerplatform Provider



## Example Usage

```terraform
provider "powerplatform" {
  username = "${var.username}"
  password = "${var.password}"
  host = "${var.host}"
}


variable "username" {
    default = "user@domain.onmicrosoft.com"
    type = string
}

variable "password" {
    type = string
}

variable "host" {
    default = "http://localhost:8080"
    type = string
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `host` (String)
- `password` (String, Sensitive)
- `username` (String)