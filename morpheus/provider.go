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
				ConflictsWith: []string{"username", "password", "tenant_subdomain"},
			},

			"tenant_subdomain": {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   "The tenant subdomain used for authentication",
				DefaultFunc:   schema.EnvDefaultFunc("MORPHEUS_API_TENANT", nil),
				ConflictsWith: []string{"access_token"},
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
			"morpheus_active_directory_identity_source":      resourceActiveDirectoryIdentitySource(),
			"morpheus_ansible_integration":                   resourceAnsibleIntegration(),
			"morpheus_ansible_playbook_task":                 resourceAnsiblePlaybookTask(),
			"morpheus_ansible_tower_integration":             resourceAnsibleTowerIntegration(),
			"morpheus_ansible_tower_task":                    resourceAnsibleTowerTask(),
			"morpheus_api_option_list":                       resourceApiOptionList(),
			"morpheus_app_blueprint_catalog_item":            resourceAppBlueprintCatalogItem(),
			"morpheus_arm_app_blueprint":                     resourceArmAppBlueprint(),
			"morpheus_arm_spec_template":                     resourceArmSpecTemplate(),
			"morpheus_aws_cloud":                             resourceAWSCloud(),
			"morpheus_aws_instance":                          resourceAwsInstance(),
			"morpheus_azure_cloud":                           resourceAzureCloud(),
			"morpheus_backup_creation_policy":                resourceBackupCreationPolicy(),
			"morpheus_backup_setting":                        resourceBackupSetting(),
			"morpheus_boot_script":                           resourceBootScript(),
			"morpheus_budget_policy":                         resourceBudgetPolicy(),
			"morpheus_checkbox_option_type":                  resourceCheckboxOptionType(),
			"morpheus_cloud_formation_app_blueprint":         resourceCloudFormationAppBlueprint(),
			"morpheus_cloud_formation_spec_template":         resourceCloudFormationSpecTemplate(),
			"morpheus_cluster_layout":                        resourceClusterLayout(),
			"morpheus_cluster_package":                       resourceClusterPackage(),
			"morpheus_cluster_resource_name_policy":          resourceClusterResourceNamePolicy(),
			"morpheus_contact":                               resourceContact(),
			"morpheus_credential":                            resourceCredential(),
			"morpheus_cypher_access_policy":                  resourceCypherAccessPolicy(),
			"morpheus_cypher_secret":                         resourceCypherSecret(),
			"morpheus_cypher_tfvars":                         resourceCypherTFVars(),
			"morpheus_delayed_delete_policy":                 resourceDelayedDeletePolicy(),
			"morpheus_delete_approval_policy":                resourceDeleteApprovalPolicy(),
			"morpheus_docker_registry_integration":           resourceDockerRegistryIntegration(),
			"morpheus_email_task":                            resourceEmailTask(),
			"morpheus_environment":                           resourceEnvironment(),
			"morpheus_execute_schedule":                      resourceExecuteSchedule(),
			"morpheus_file_template":                         resourceFileTemplate(),
			"morpheus_git_integration":                       resourceGitIntegration(),
			"morpheus_groovy_script_task":                    resourceGroovyScriptTask(),
			"morpheus_group":                                 resourceMorpheusGroup(),
			"morpheus_guidance_setting":                      resourceGuidanceSetting(),
			"morpheus_helm_app_blueprint":                    resourceHelmAppBlueprint(),
			"morpheus_helm_spec_template":                    resourceHelmSpecTemplate(),
			"morpheus_hidden_option_type":                    resourceHiddenOptionType(),
			"morpheus_hostname_policy":                       resourceHostNamePolicy(),
			"morpheus_instance_catalog_item":                 resourceInstanceCatalogItem(),
			"morpheus_instance_layout":                       resourceInstanceLayout(),
			"morpheus_instance_name_policy":                  resourceInstanceNamePolicy(),
			"morpheus_instance_type":                         resourceInstanceType(),
			"morpheus_ipv4_ip_pool":                          resourceIPv4IPPool(),
			"morpheus_javascript_task":                       resourceJavaScriptTask(),
			"morpheus_library_script_task":                   resourceLibraryScriptTask(),
			"morpheus_library_template_task":                 resourceLibraryTemplateTask(),
			"morpheus_license":                               resourceLicense(),
			"morpheus_key_pair":                              resourceKeyPair(),
			"morpheus_kubernetes_app_blueprint":              resourceKubernetesAppBlueprint(),
			"morpheus_kubernetes_spec_template":              resourceKubernetesSpecTemplate(),
			"morpheus_manual_option_list":                    resourceManualOptionList(),
			"morpheus_max_containers_policy":                 resourceMaxContainersPolicy(),
			"morpheus_max_cores_policy":                      resourceMaxCoresPolicy(),
			"morpheus_max_hosts_policy":                      resourceMaxHostsPolicy(),
			"morpheus_max_memory_policy":                     resourceMaxMemoryPolicy(),
			"morpheus_max_storage_policy":                    resourceMaxStoragePolicy(),
			"morpheus_max_vms_policy":                        resourceMaxVmsPolicy(),
			"morpheus_monitoring_setting":                    resourceMonitoringSetting(),
			"morpheus_motd_policy":                           resourceMotdPolicy(),
			"morpheus_nested_workflow_task":                  resourceNestedWorkflowTask(),
			"morpheus_network_domain":                        resourceNetworkDomain(),
			"morpheus_network_quota_policy":                  resourceNetworkQuotaPolicy(),
			"morpheus_node_type":                             resourceNodeType(),
			"morpheus_number_option_type":                    resourceNumberOptionType(),
			"morpheus_operational_workflow":                  resourceOperationalWorkflow(),
			"morpheus_password_option_type":                  resourcePasswordOptionType(),
			"morpheus_power_schedule_policy":                 resourcePowerSchedulePolicy(),
			"morpheus_powershell_script_task":                resourcePowerShellScriptTask(),
			"morpheus_preseed_script":                        resourcePreseedScript(),
			"morpheus_price_set":                             resourcePriceSet(),
			"morpheus_price":                                 resourcePrice(),
			"morpheus_provision_approval_policy":             resourceProvisionApprovalPolicy(),
			"morpheus_provisioning_setting":                  resourceProvisioningSetting(),
			"morpheus_provisioning_workflow":                 resourceProvisioningWorkflow(),
			"morpheus_puppet_integration":                    resourcePuppetIntegration(),
			"morpheus_python_script_task":                    resourcePythonScriptTask(),
			"morpheus_radio_list_option_type":                resourceRadioListOptionType(),
			"morpheus_resource_pool_group":                   resourceResourcePoolGroup(),
			"morpheus_rest_option_list":                      resourceRestOptionList(),
			"morpheus_restart_task":                          resourceRestartTask(),
			"morpheus_router_quota_policy":                   resourceRouterQuotaPolicy(),
			"morpheus_ruby_script_task":                      resourceRubyScriptTask(),
			"morpheus_saml_identity_source":                  resourceSAMLIdentitySource(),
			"morpheus_scale_threshold":                       resourceScaleThreshold(),
			"morpheus_script_template":                       resourceScriptTemplate(),
			"morpheus_security_package":                      resourceSecurityPackage(),
			"morpheus_select_list_option_type":               resourceSelectListOptionType(),
			"morpheus_service_plan":                          resourceServicePlan(),
			"morpheus_servicenow_integration":                resourceServiceNowIntegration(),
			"morpheus_shell_script_task":                     resourceShellScriptTask(),
			"morpheus_standard_cloud":                        resourceStandardCloud(),
			"morpheus_tag_policy":                            resourceTagPolicy(),
			"morpheus_task_job":                              resourceTaskJob(),
			"morpheus_tenant_role":                           resourceTenantRole(),
			"morpheus_tenant":                                resourceTenant(),
			"morpheus_terraform_app_blueprint":               resourceTerraformAppBlueprint(),
			"morpheus_terraform_spec_template":               resourceTerraformSpecTemplate(),
			"morpheus_text_option_type":                      resourceTextOptionType(),
			"morpheus_textarea_option_type":                  resourceTextAreaOptionType(),
			"morpheus_typeahead_option_type":                 resourceTypeAheadOptionType(),
			"morpheus_user_creation_policy":                  resourceUserCreationPolicy(),
			"morpheus_user_group_creation_policy":            resourceUserGroupCreationPolicy(),
			"morpheus_user":                                  resourceMorpheusUser(),
			"morpheus_user_group":                            resourceUserGroup(),
			"morpheus_user_role":                             resourceUserRole(),
			"morpheus_vro_integration":                       resourceVrealizeOrchestratorIntegration(),
			"morpheus_vro_task":                              resourceVrealizeOrchestratorTask(),
			"morpheus_vsphere_cloud_datastore_configuration": resourceVSphereCloudDatastoreConfiguration(),
			"morpheus_vsphere_cloud":                         resourceVsphereCloud(),
			"morpheus_vsphere_instance":                      resourceVsphereInstance(),
			"morpheus_wiki_page":                             resourceWikiPage(),
			"morpheus_workflow_catalog_item":                 resourceWorkflowCatalogItem(),
			"morpheus_workflow_job":                          resourceWorkflowJob(),
			"morpheus_workflow_policy":                       resourceWorkflowPolicy(),
			"morpheus_write_attributes_task":                 resourceWriteAttributesTask(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"morpheus_ansible_tower_job_template": dataSourceMorpheusAnsibleTowerJobTemplate(),
			"morpheus_ansible_tower_inventory":    dataSourceMorpheusAnsibleTowerInventory(),
			"morpheus_blueprint":                  dataSourceMorpheusBlueprint(),
			"morpheus_budget":                     dataSourceMorpheusBudget(),
			"morpheus_catalog_item_type":          dataSourceMorpheusCatalogItemType(),
			"morpheus_cloud":                      dataSourceMorpheusCloud(),
			"morpheus_cluster_type":               dataSourceMorpheusClusterType(),
			"morpheus_contact":                    dataSourceMorpheusContact(),
			"morpheus_credential":                 dataSourceMorpheusCredential(),
			"morpheus_cypher_secret":              dataSourceMorpheusCypherSecret(),
			"morpheus_domain":                     dataSourceMorpheusDomain(),
			"morpheus_environment":                dataSourceMorpheusEnvironment(),
			"morpheus_execute_schedule":           dataSourceMorpheusExecuteSchedule(),
			"morpheus_file_template":              dataSourceMorpheusFileTemplate(),
			"morpheus_git_integration":            dataSourceMorpheusGitIntegration(),
			"morpheus_group":                      dataSourceMorpheusGroup(),
			"morpheus_instance_layout":            dataSourceMorpheusInstanceLayout(),
			"morpheus_instance_type":              dataSourceMorpheusInstanceType(),
			"morpheus_integration":                dataSourceMorpheusIntegration(),
			"morpheus_job":                        dataSourceMorpheusJob(),
			"morpheus_key_pair":                   dataSourceMorpheusKeyPair(),
			"morpheus_network":                    dataSourceMorpheusNetwork(),
			"morpheus_network_group":              dataSourceMorpheusNetworkGroup(),
			"morpheus_network_subnet":             dataSourceMorpheusNetworkSubnet(),
			"morpheus_node_type":                  dataSourceMorpheusNodeType(),
			"morpheus_option_list":                dataSourceMorpheusOptionList(),
			"morpheus_option_type":                dataSourceMorpheusOptionType(),
			"morpheus_permission_set":             dataSourceMorpheusPermissionSet(),
			"morpheus_plan":                       dataSourceMorpheusPlan(),
			"morpheus_policy":                     dataSourceMorpheusPolicy(),
			"morpheus_power_schedule":             dataSourceMorpheusPowerSchedule(),
			"morpheus_price_set":                  dataSourceMorpheusPriceSet(),
			"morpheus_price":                      dataSourceMorpheusPrice(),
			"morpheus_provision_type":             dataSourceMorpheusProvisionType(),
			"morpheus_resource_pool":              dataSourceMorpheusResourcePool(),
			"morpheus_script_template":            dataSourceMorpheusScriptTemplate(),
			"morpheus_security_package":           dataSourceMorpheusSecurityPackage(),
			"morpheus_servicenow_workflow":        dataSourceMorpheusServiceNowWorkflow(),
			"morpheus_spec_template":              dataSourceMorpheusSpecTemplate(),
			"morpheus_storage_bucket":             dataSourceMorpheusStorageBucket(),
			"morpheus_storage_volume_type":        dataSourceMorpheusStorageVolumeType(),
			"morpheus_task":                       dataSourceMorpheusTask(),
			"morpheus_tenant_role":                dataSourceMorpheusTenantRole(),
			"morpheus_tenant":                     dataSourceMorpheusTenant(),
			"morpheus_user_group":                 dataSourceMorpheusUserGroup(),
			"morpheus_user_role":                  dataSourceMorpheusUserRole(),
			"morpheus_vdi_pool":                   dataSourceMorpheusVDIPool(),
			"morpheus_virtual_image":              dataSourceMorpheusVirtualImage(),
			"morpheus_vro_workflow":               dataSourceMorpheusVrealizeOrchestratorWorkflow(),
			"morpheus_workflow":                   dataSourceMorpheusWorkflow(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	config := Config{
		Url:             d.Get("url").(string),
		AccessToken:     d.Get("access_token").(string),
		TenantSubdomain: d.Get("tenant_subdomain").(string),
		Username:        d.Get("username").(string),
		Password:        d.Get("password").(string),
		//Insecure:                d.Get("insecure").(bool), //.(bool),
	}
	return config.Client()
}
