package morpheus

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The URL of the Morpheus Data Appliance where requests will be directed.",
				DefaultFunc: schema.EnvDefaultFunc("MORPHEUS_API_URL", nil),
			},

			"access_token": {
				Type:          schema.TypeString,
				Optional:      true,
				Sensitive:     true,
				Description:   "Access Token of Morpheus user. This can be used instead of authenticating with Username and Password.",
				DefaultFunc:   schema.EnvDefaultFunc("MORPHEUS_API_TOKEN", nil),
				ConflictsWith: []string{"username", "password"},
			},

			"username": {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   "Username of Morpheus user for authentication",
				DefaultFunc:   schema.EnvDefaultFunc("MORPHEUS_API_USERNAME", nil),
				ConflictsWith: []string{"access_token"},
			},

			"password": {
				Type:          schema.TypeString,
				Optional:      true,
				Sensitive:     true,
				Description:   "Password of Morpheus user for authentication",
				DefaultFunc:   schema.EnvDefaultFunc("MORPHEUS_API_PASSWORD", nil),
				ConflictsWith: []string{"access_token"},
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"morpheus_ansible_playbook_task":         resourceAnsiblePlaybookTask(),
			"morpheus_arm_app_blueprint":             resourceArmAppBlueprint(),
			"morpheus_arm_spec_template":             resourceArmSpecTemplate(),
			"morpheus_backup_creation_policy":        resourceBackupCreationPolicy(),
			"morpheus_budget_policy":                 resourceBudgetPolicy(),
			"morpheus_checkbox_option_type":          resourceCheckboxOptionType(),
			"morpheus_cloud_formation_app_blueprint": resourceCloudFormationAppBlueprint(),
			"morpheus_cloud_formation_spec_template": resourceCloudFormationSpecTemplate(),
			"morpheus_cluster_resource_name_policy":  resourceClusterResourceNamePolicy(),
			"morpheus_contact":                       resourceContact(),
			"morpheus_environment":                   resourceEnvironment(),
			"morpheus_execute_schedule":              resourceExecuteSchedule(),
			"morpheus_groovy_script_task":            resourceGroovyScriptTask(),
			"morpheus_group":                         resourceMorpheusGroup(),
			"morpheus_helm_app_blueprint":            resourceHelmAppBlueprint(),
			"morpheus_helm_spec_template":            resourceHelmSpecTemplate(),
			"morpheus_hidden_option_type":            resourceHiddenOptionType(),
			"morpheus_hostname_policy":               resourceHostNamePolicy(),
			"morpheus_instance_name_policy":          resourceInstanceNamePolicy(),
			"morpheus_javascript_task":               resourceJavaScriptTask(),
			"morpheus_kubernetes_app_blueprint":      resourceKubernetesAppBlueprint(),
			"morpheus_kubernetes_spec_template":      resourceKubernetesSpecTemplate(),
			"morpheus_manual_option_list":            resourceManualOptionList(),
			"morpheus_max_containers_policy":         resourceMaxContainersPolicy(),
			"morpheus_max_cores_policy":              resourceMaxCoresPolicy(),
			"morpheus_max_hosts_policy":              resourceMaxHostsPolicy(),
			"morpheus_max_memory_policy":             resourceMaxMemoryPolicy(),
			"morpheus_max_storage_policy":            resourceMaxStoragePolicy(),
			"morpheus_max_vms_policy":                resourceMaxVmsPolicy(),
			"morpheus_network_domain":                resourceNetworkDomain(),
			"morpheus_network_quota_policy":          resourceNetworkQuotaPolicy(),
			"morpheus_number_option_type":            resourceNumberOptionType(),
			"morpheus_operational_workflow":          resourceOperationalWorkflow(),
			"morpheus_password_option_type":          resourcePasswordOptionType(),
			"morpheus_powershell_script_task":        resourcePowerShellScriptTask(),
			"morpheus_price":                         resourcePrice(),
			"morpheus_price_set":                     resourcePriceSet(),
			"morpheus_provisioning_workflow":         resourceProvisioningWorkflow(),
			"morpheus_python_script_task":            resourcePythonScriptTask(),
			"morpheus_rest_option_list":              resourceRestOptionList(),
			"morpheus_restart_task":                  resourceRestartTask(),
			"morpheus_router_quota_policy":           resourceRouterQuotaPolicy(),
			"morpheus_ruby_script_task":              resourceRubyScriptTask(),
			"morpheus_select_list_option_type":       resourceSelectListOptionType(),
			"morpheus_service_plan":                  resourceServicePlan(),
			"morpheus_shell_script_task":             resourceShellScriptTask(),
			"morpheus_task_job":                      resourceTaskJob(),
			"morpheus_tenant":                        resourceTenant(),
			"morpheus_terraform_app_blueprint":       resourceTerraformAppBlueprint(),
			"morpheus_terraform_spec_template":       resourceTerraformSpecTemplate(),
			"morpheus_text_option_type":              resourceTextOptionType(),
			"morpheus_typeahead_option_type":         resourceTypeAheadOptionType(),
			"morpheus_user_creation_policy":          resourceUserCreationPolicy(),
			"morpheus_vsphere_cloud":                 resourceVsphereCloud(),
			"morpheus_vsphere_instance":              resourceVsphereInstance(),
			"morpheus_wiki_page":                     resourceWikiPage(),
			"morpheus_workflow_catalog_item":         resourceWorkflowCatalogItem(),
			"morpheus_workflow_policy":               resourceWorkflowPolicy(),
			"morpheus_write_attributes_task":         resourceWriteAttributesTask(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"morpheus_cloud":            dataSourceMorpheusCloud(),
			"morpheus_contact":          dataSourceMorpheusContact(),
			"morpheus_environment":      dataSourceMorpheusEnvironment(),
			"morpheus_execute_schedule": dataSourceMorpheusExecuteSchedule(),
			"morpheus_group":            dataSourceMorpheusGroup(),
			"morpheus_instance_type":    dataSourceMorpheusInstanceType(),
			"morpheus_integration":      dataSourceMorpheusIntegration(),
			"morpheus_instance_layout":  dataSourceMorpheusInstanceLayout(),
			"morpheus_network":          dataSourceMorpheusNetwork(),
			"morpheus_option_type":      dataSourceMorpheusOptionType(),
			"morpheus_plan":             dataSourceMorpheusPlan(),
			"morpheus_price":            dataSourceMorpheusPrice(),
			"morpheus_price_set":        dataSourceMorpheusPriceSet(),
			"morpheus_resource_pool":    dataSourceMorpheusResourcePool(),
			"morpheus_spec_template":    dataSourceMorpheusSpecTemplate(),
			"morpheus_task":             dataSourceMorpheusTask(),
			"morpheus_tenant":           dataSourceMorpheusTenant(),
			"morpheus_tenant_role":      dataSourceMorpheusTenantRole(),
			"morpheus_workflow":         dataSourceMorpheusWorkflow(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	config := Config{
		Url:         d.Get("url").(string),
		AccessToken: d.Get("access_token").(string),
		Username:    d.Get("username").(string),
		Password:    d.Get("password").(string),
		//Insecure:                d.Get("insecure").(bool), //.(bool),
	}
	return config.Client()
}
