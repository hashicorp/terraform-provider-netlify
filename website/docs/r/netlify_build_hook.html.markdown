---
layout: "netlify"
page_title: "Netlify: netlify_build_hook"
sidebar_current: "docs-netlify-resource-build-hook"
description: |-
  Provides an build hook resource.
---

# netlify_build_hook

[DESCRIPTION]

## Example Usage

```hcl
resource "netlify_build_hook" "bar" {
  site_id =
  branch =
  title =
}
```

## Argument Reference

The following arguments are supported:

* `site_id` - (Required) [add description]
* `branch` - (Required) [add description]
* `title` - (Required) [add description]


## Attribute Reference

The following additional attributes are exported:

* `url` - URL of the project
