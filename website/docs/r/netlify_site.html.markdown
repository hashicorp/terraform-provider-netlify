---
layout: "netlify"
page_title: "Netlify: netlify_site"
sidebar_current: "docs-netlify-resource-site"
description: |-
  Provides an site resource.
---

# netlify_site

[DESCRIPTION]

## Example Usage

```hcl
resource "netlify_site" "bar" {


}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required)
* `custom_domain` - (Required)
* `deploy_url` - (Required)
* `repo` - (Required) See [Repository](#repo)

### Repository

`repo` supports the following argument

* `command` - (Optional)
* `deploy_key_id` - (Optional)
* `dir` - (Optional)
* `provider` - (Required)
* `repo_path` - (Required)
* `repo_branch` - (Required)
