---
subcategory: "Compute Engine"
description: |-
  Provide access to a Resource Policy's attributes
---

# google\_compute\_resource\_policy

Provide access to a Resource Policy's attributes. For more information see [the official documentation](https://cloud.google.com/compute/docs/disks/scheduled-snapshots) or the [API](https://cloud.google.com/compute/docs/reference/rest/beta/resourcePolicies).

```hcl
provider "google-beta" {
  region = "us-central1"
  zone   = "us-central1-a"
}

data "google_compute_resource_policy" "daily" {
  provider = google-beta
  name     = "daily"
  region   = "us-central1"
}
```

## Argument Reference

The following arguments are supported:

* `name` (Required) - The name of the Resource Policy.
* `project` (Optional) - Project from which to list the Resource Policy. Defaults to project declared in the provider.
* `region` (Required) - Region where the Resource Policy resides.

## Attributes Reference

The following attributes are exported:

* `description` - Description of this Resource Policy.
* `self_link` - The URI of the resource.
