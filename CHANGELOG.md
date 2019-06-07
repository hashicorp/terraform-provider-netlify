## 0.4.0 (Unreleased)
## 0.3.0 (June 07, 2019)

* provider: This release includes only a Terraform SDK upgrade with compatibility for Terraform v0.12. The provider remains backwards compatible with Terraform v0.11 and this update should have no significant changes in behavior for the provider. Please report any unexpected behavior in new GitHub issues (Terraform core: https://github.com/hashicorp/terraform/issues or Terraform Netlify Provider: https://github.com/terraform-providers/terraform-provider-netlify/issues).

## 0.2.0 (June 07, 2019)

### Features

* `resource/netlify_deploy_key`: support importing ([#3](https://github.com/terraform-providers/terraform-provider-aws/issues/3))
* automatically retry requests ([#2](https://github.com/terraform-providers/terraform-provider-aws/issues/2))

### Improvements

* upgrade to Go 1.11 ([#12](https://github.com/terraform-providers/terraform-provider-aws/issues/12))
* use Go modules with `go mod vendor` ([#16](https://github.com/terraform-providers/terraform-provider-aws/issues/16))
* fix indentation issues in Makefile ([#23](https://github.com/terraform-providers/terraform-provider-aws/issues/23))
* upgrade `netlify/open-api` dependency to v0.11.4 ([#20](https://github.com/terraform-providers/terraform-provider-aws/issues/20))
* upgrade `sirupsen/logrus` dependency to v1.4.2 ([#27](https://github.com/terraform-providers/terraform-provider-aws/issues/27))

### Bug fixes

* docs fixes ([#9](https://github.com/terraform-providers/terraform-provider-aws/issues/9)] [[#10](https://github.com/terraform-providers/terraform-provider-aws/issues/10)] [[#11](https://github.com/terraform-providers/terraform-provider-aws/issues/11)] [[#19](https://github.com/terraform-providers/terraform-provider-aws/issues/19))
* only set repo attribute on read if there is a repo path ([#22](https://github.com/terraform-providers/terraform-provider-aws/issues/22))

## 0.1.0 (September 13, 2018)

Initial Release:

* `resource/netlify_build_hook`
* `resource/netlify_deploy_key`
* `resource/netlify_hook`
* `resource/netlify_site`
