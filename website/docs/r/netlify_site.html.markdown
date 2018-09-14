---
layout: "netlify"
page_title: "Netlify: netlify_site"
sidebar_current: "docs-netlify-resource-site"
description: |-
  Provides an site resource.
---

# netlify_site

Primary settings for the Netlify site - should contain the bulk of your configuration. Allows configuration of most aspects of your Netlify site.

## Example Usage

```hcl
resource "netlify_site" "main" {
  name = "my-site"

  repo {
    command       = "middleman build"
    deploy_key_id = "${netlify_deploy_key.key.id}"
    dir           = "/build"
    provider      = "github"
    repo_path     = "username/repo"
    repo_branch   = "master"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) - name of your site on netlify
* `repo` - (Required) See [Repository](#repo)
* `custom_domain` - (Optional) - a custom domain name, must be configured using a cname in accordance with [netlify's docs](https://www.netlify.com/docs/custom-domains)
* `deploy_url` - (Optional)

### Repository

`repo` supports the following argument

* `command` - (Optional) - shell command run before deployment, typically used to build the site
* `deploy_key_id` - (Optional) - a deploy key id from the `deploy_key` resource
* `dir` - (Optional) - the directory to deploy
* `provider` - (Required) - name of your VCS provider
* `repo_path` - (Required) - path to your repo, typically `username/reponame`
* `repo_branch` - (Required) - branch to be deployed
