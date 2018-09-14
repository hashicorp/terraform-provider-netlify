---
layout: "netlify"
page_title: "Netlify: netlify_build_hook"
sidebar_current: "docs-netlify-resource-build-hook"
description: |-
  Provides an build hook resource.
---

# netlify_build_hook

Manages build hooks, also known as [incoming webhooks]
(https://www.netlify.com/docs/webhooks/#outgoing-webhooks). These can,
at the time of writing, only be used to trigger new builds of the site.
To create one, provide your site id along with the name of the hook, and
the branch to be built when the hook is triggered.

## Example Usage

```hcl
resource "netlify_build_hook" "trigger" {
  site_id = "12345"
  branch  = "master"
  title   = "Manual Build Trigger"
}
```

## Argument Reference

The following arguments are supported:

* `site_id` - (Required) Your netlify site's unique id
* `branch` - (Required) branch to be built when the hook is triggered
* `title` - (Required) name of the webhook - this is purely for organization and
can be any name you want


## Attribute Reference

The following additional attributes are exported:

* `url` - URL of the project
