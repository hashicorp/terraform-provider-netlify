---
layout: "netlify"
page_title: "Netlify: netlify_hook"
sidebar_current: "docs-netlify-resource-hook"
description: |-
  Provides an hook resource.
---

# netlify_hook

An [outgoing webhook](https://www.netlify.com/docs/webhooks/#outgoing-webhooks-and-notifications), typically used to notify a third party service about deploys.

## Example Usage

```hcl
resource "netlify_hook" "email_on_deploy" {
  site_id = "12345"
  type    = "email"
  event   = "deploy_created"

  data {
    email = "test@test.com"
  }
}
```

## Argument Reference

The following arguments are supported:

* `site_id` - (Required) - id of the site on netlify
* `type` - (Required) - type of outgoing webhook, for example slack, email, github commit status, etc
* `event` - (Required) - when to send the data, for example on deploy create, succeed, fail, etc
* `data` - (Required) object/hash of data to be sent along with the webhook. this varies depending on the `type`
