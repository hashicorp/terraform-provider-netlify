---
layout: "netlify"
page_title: "Netlify: netlify_deploy_key"
sidebar_current: "docs-netlify-resource-deploy-key"
description: |-
  Provides an deploy key resource.
---

# netlify_deploy_key

Creates a new netlify deploy key, typically used by the `netlify_site` resource.

## Example Usage

```hcl
resource "netlify_deploy_key" "key" {}

resource "netlify_site" "main" {
  // ...
  repo {
    // ...
    deploy_key_id = "${netlify_deploy_key.key.id}"
  }
}
```




## Attribute Reference

The following additional attributes are exported:

* `public_key` - Public Key
