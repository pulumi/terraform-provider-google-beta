---
subcategory: "Cloud IAM"
page_title: "Google: google_iam_workload_identity_pool"
description: |-
  Get a IAM workload identity pool from Google Cloud
---

# google\_iam\_workload_\identity\_pool

Get a IAM workload identity pool from Google Cloud by its id.
~> **Note:** The following resource requires the Beta IAM role `roles/iam.workloadIdentityPoolAdmin` in order to succeed. `OWNER` and `EDITOR` roles do not include the necessary permissions.
## Example Usage

```tf
data "google_iam_workload_identity_pool" "foo" {
  workload_identity_pool_id = "foo-pool"
}
```

## Argument Reference

The following arguments are supported:

* `workload_identity_pool_id` - (Required) The id of the pool which is the
    final component of the resource name.

- - -

* `project` - (Optional) The project in which the resource belongs. If it
    is not provided, the provider project is used.

## Attributes Reference
See google_iam_workload_identity_pool resource for details of all the available attributes.
