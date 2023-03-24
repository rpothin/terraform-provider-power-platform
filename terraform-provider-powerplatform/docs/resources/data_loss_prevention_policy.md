---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "powerplatform_data_loss_prevention_policy Resource - terraform-provider-powerplatform"
subcategory: ""
description: |-
  Resource to manage Data Loss Prevention policies.
---

# powerplatform_data_loss_prevention_policy (Resource)

Resource to manage Data Loss Prevention policies.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `default_connectors_classification` (String)
- `display_name` (String)
- `environment_type` (String)

### Optional

- `connector_group` (Block List) (see [below for nested schema](#nestedblock--connector_group))
- `environment` (Block List) (see [below for nested schema](#nestedblock--environment))

### Read-Only

- `created_by` (String)
- `created_time` (String)
- `e_tag` (String)
- `id` (String) The ID of this resource.
- `last_modified_by` (String)
- `last_modified_time` (String)
- `name` (String)

<a id="nestedblock--connector_group"></a>
### Nested Schema for `connector_group`

Optional:

- `classification` (String)
- `connector` (Block List) (see [below for nested schema](#nestedblock--connector_group--connector))

<a id="nestedblock--connector_group--connector"></a>
### Nested Schema for `connector_group.connector`

Optional:

- `name` (String)

Read-Only:

- `id` (String) The ID of this resource.
- `type` (String)



<a id="nestedblock--environment"></a>
### Nested Schema for `environment`

Optional:

- `name` (String)

Read-Only:

- `id` (String) The ID of this resource.
- `type` (String)

