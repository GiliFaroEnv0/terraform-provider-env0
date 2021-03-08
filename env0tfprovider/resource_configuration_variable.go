package env0tfprovider

import (
	"context"
	"errors"

	"github.com/env0/terraform-provider-env0/env0apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceConfigurationVariable() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceConfigurationVariableCreate,
		ReadContext:   resourceConfigurationVariableRead,
		UpdateContext: resourceConfigurationVariableUpdate,
		DeleteContext: resourceConfigurationVariableDelete,

		Importer: &schema.ResourceImporter{StateContext: resourceConfigurationVariableImport},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "name to give the configuration variable",
				Required:    true,
			},
			"value": {
				Type:        schema.TypeString,
				Description: "value for the configuration variable",
				Required:    true,
			},
			"is_sensitive": {
				Type:        schema.TypeBool,
				Description: "is the variable sensitive, defaults to false",
				Optional:    true,
				Default:     false,
			},
			"project_id": {
				Type:          schema.TypeString,
				Description:   "create the variable under this project, not globally",
				Optional:      true,
				ConflictsWith: []string{"blueprint_id", "environment_id"},
			},
			"blueprint_id": {
				Type:          schema.TypeString,
				Description:   "create the variable under this blueprint, not globally",
				Optional:      true,
				ConflictsWith: []string{"project_id", "environment_id"},
			},
			"environment_id": {
				Type:          schema.TypeString,
				Description:   "create the variable under this environment, not globally",
				Optional:      true,
				ConflictsWith: []string{"blueprint_id", "project_id"},
			},
			"type": {
				Type:        schema.TypeString,
				Description: "default 'environment'. set to 'terraform' to create a terraform variable",
				Optional:    true,
				Default:     "environment",
			},
			"enum": {
				Type:        schema.TypeList,
				Description: "limit possible values to values from this list",
				Optional:    true,
				Elem: &schema.Schema{
					Type:        schema.TypeString,
					Description: "name to give the configuration variable",
				},
			},
		},
	}
}

func whichScope(d *schema.ResourceData) (env0apiclient.Scope, string) {
	scope := env0apiclient.ScopeGlobal
	scopeId := ""
	if projectId, ok := d.GetOk("project_id"); ok {
		scope = env0apiclient.ScopeProject
		scopeId = projectId.(string)
	}
	if blueprintId, ok := d.GetOk("blueprint_id"); ok {
		scope = env0apiclient.ScopeBlueprint
		scopeId = blueprintId.(string)
	}
	if environmentId, ok := d.GetOk("environment_id"); ok {
		scope = env0apiclient.ScopeEnvironment
		scopeId = environmentId.(string)
	}

	return scope, scopeId
}

func resourceConfigurationVariableCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apiClient := meta.(*env0apiclient.ApiClient)

	scope, scopeId := whichScope(d)
	name := d.Get("name").(string)
	value := d.Get("value").(string)
	isSensitive := d.Get("is_sensitive").(bool)
	typeAsString := d.Get("type").(string)
	var type_ env0apiclient.ConfigurationVariableType
	switch typeAsString {
	case "environment":
		type_ = env0apiclient.ConfigurationVariableTypeEnvironment
	case "terraform":
		type_ = env0apiclient.ConfigurationVariableTypeTerraform
	default:
		return diag.Errorf("'type' can only receive either 'environment' or 'terraform': %s", typeAsString)
	}
	var enumValues []string = nil
	if specified, ok := d.GetOk("enum_values"); ok {
		enumValues = specified.([]string)
	}
	configurationVariable, err := apiClient.ConfigurationVariableCreate(name, value, isSensitive, scope, scopeId, type_, enumValues)
	if err != nil {
		return diag.Errorf("could not create configurationVariable: %v", err)
	}

	d.SetId(configurationVariable.Id)

	return nil
}

func resourceConfigurationVariableRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apiClient := meta.(*env0apiclient.ApiClient)

	id := d.Id()
	scope, scopeId := whichScope(d)
	variables, err := apiClient.ConfigurationVariables(scope, scopeId)
	if err != nil {
		return diag.Errorf("could not get configurationVariable: %v", err)
	}
	for _, variable := range variables {
		if variable.Id == id {
			d.Set("name", variable.Name)
			d.Set("value", variable.Value)
			d.Set("is_sensitive", variable.IsSensitive)
			if variable.Type == int64(env0apiclient.ConfigurationVariableTypeTerraform) {
				d.Set("type", "terraform")
			} else {
				d.Set("type", "environment")
			}
			return nil
		}
	}
	return diag.Errorf("variable %s not found (under this scope): %v", id, err)
}

func resourceConfigurationVariableUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Errorf("not implemented")
}

func resourceConfigurationVariableDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apiClient := meta.(*env0apiclient.ApiClient)

	id := d.Id()
	err := apiClient.ConfigurationVariableDelete(id)
	if err != nil {
		return diag.Errorf("could not delete configurationVariable: %v", err)
	}
	return nil
}

func resourceConfigurationVariableImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return nil, errors.New("Not implemented")
	// apiClient := meta.(*env0apiclient.ApiClient)

	// id := d.Id()
	// configurationVariable, err := apiClient.ConfigurationVariable(id)
	// if err != nil {
	// 	return nil, err
	// }

	// d.Set("name", configurationVariable.Name)

	// return []*schema.ResourceData{d}, nil
}
