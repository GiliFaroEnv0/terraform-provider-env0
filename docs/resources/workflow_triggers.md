---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "env0_workflow_triggers Resource - terraform-provider-env0"
subcategory: ""
description: |-
  
---

# env0_workflow_triggers (Resource)



## Example Usage

```terraform
data "env0_project" "default" {
  name = "Default Organization Project"
}

resource "env0_template" "template" {
  name              = "Template for environment resource"
  type              = "terraform"
  repository        = "https://github.com/env0/templates"
  path              = "misc/null-resource"
  terraform_version = "0.15.1"
}

resource "env0_environment" "the_trigger" {
  force_destroy              = true
  name                       = "the_trigger"
  project_id                 = data.env0_project.default.id
  template_id                = env0_template.template.id
  approve_plan_automatically = true
}

resource "env0_environment" "downstream_environment" {
  force_destroy              = true
  name                       = "downstream_environment"
  project_id                 = data.env0_project.default.id
  template_id                = env0_template.template.id
  approve_plan_automatically = true
}

resource "env0_workflow_triggers" "trigger_link" {
  environment_id = env0_environment.the_trigger.id
  downstream_environment_ids = [
    env0_environment.downstream_environment.id
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **downstream_environment_ids** (List of String) environments to trigger
- **environment_id** (String) id of the source environment

### Optional

- **id** (String) The ID of this resource.


