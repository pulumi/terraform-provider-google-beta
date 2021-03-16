// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudAssetOrganizationFeed() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudAssetOrganizationFeedCreate,
		Read:   resourceCloudAssetOrganizationFeedRead,
		Update: resourceCloudAssetOrganizationFeedUpdate,
		Delete: resourceCloudAssetOrganizationFeedDelete,

		Importer: &schema.ResourceImporter{
			State: resourceCloudAssetOrganizationFeedImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"billing_project": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The project whose identity will be used when sending messages to the
destination pubsub topic. It also specifies the project for API 
enablement check, quota, and billing.`,
			},
			"feed_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.`,
			},
			"feed_output_config": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `Output configuration for asset feed destination.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pubsub_destination": {
							Type:        schema.TypeList,
							Required:    true,
							Description: `Destination on Cloud Pubsub.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"topic": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `Destination on Cloud Pubsub topic.`,
									},
								},
							},
						},
					},
				},
			},
			"org_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The organization this feed should be created in.`,
			},
			"asset_names": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `A list of the full names of the assets to receive updates. You must specify either or both of 
assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"asset_types": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `A list of types of the assets to receive updates. You must specify either or both of assetNames
and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
the feed. For example: "compute.googleapis.com/Disk"
See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
supported asset types.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"condition": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `A condition which determines whether an asset update should be published. If specified, an asset
will be returned only when the expression evaluates to true. When set, expression field
must be a valid CEL expression on a TemporalAsset with name temporal_asset. Example: a Feed with
expression "temporal_asset.deleted == true" will only publish Asset deletions. Other fields of
condition are optional.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"expression": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `Textual representation of an expression in Common Expression Language syntax.`,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `Description of the expression. This is a longer text which describes the expression,
e.g. when hovered over it in a UI.`,
						},
						"location": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `String indicating the location of the expression for error reporting, e.g. a file 
name and a position in the file.`,
						},
						"title": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `Title for the expression, i.e. a short string describing its purpose.
This can be used e.g. in UIs which allow to enter the expression.`,
						},
					},
				},
			},
			"content_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"CONTENT_TYPE_UNSPECIFIED", "RESOURCE", "IAM_POLICY", "ORG_POLICY", "ACCESS_POLICY", ""}, false),
				Description:  `Asset content type. If not specified, no content but the asset name and type will be returned. Possible values: ["CONTENT_TYPE_UNSPECIFIED", "RESOURCE", "IAM_POLICY", "ORG_POLICY", "ACCESS_POLICY"]`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The format will be organizations/{organization_number}/feeds/{client-assigned_feed_identifier}.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceCloudAssetOrganizationFeedCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	assetNamesProp, err := expandCloudAssetOrganizationFeedAssetNames(d.Get("asset_names"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("asset_names"); !isEmptyValue(reflect.ValueOf(assetNamesProp)) && (ok || !reflect.DeepEqual(v, assetNamesProp)) {
		obj["assetNames"] = assetNamesProp
	}
	assetTypesProp, err := expandCloudAssetOrganizationFeedAssetTypes(d.Get("asset_types"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("asset_types"); !isEmptyValue(reflect.ValueOf(assetTypesProp)) && (ok || !reflect.DeepEqual(v, assetTypesProp)) {
		obj["assetTypes"] = assetTypesProp
	}
	contentTypeProp, err := expandCloudAssetOrganizationFeedContentType(d.Get("content_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("content_type"); !isEmptyValue(reflect.ValueOf(contentTypeProp)) && (ok || !reflect.DeepEqual(v, contentTypeProp)) {
		obj["contentType"] = contentTypeProp
	}
	feedOutputConfigProp, err := expandCloudAssetOrganizationFeedFeedOutputConfig(d.Get("feed_output_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("feed_output_config"); !isEmptyValue(reflect.ValueOf(feedOutputConfigProp)) && (ok || !reflect.DeepEqual(v, feedOutputConfigProp)) {
		obj["feedOutputConfig"] = feedOutputConfigProp
	}
	conditionProp, err := expandCloudAssetOrganizationFeedCondition(d.Get("condition"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("condition"); !isEmptyValue(reflect.ValueOf(conditionProp)) && (ok || !reflect.DeepEqual(v, conditionProp)) {
		obj["condition"] = conditionProp
	}

	obj, err = resourceCloudAssetOrganizationFeedEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{CloudAssetBasePath}}organizations/{{org_id}}/feeds?feedId={{feed_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new OrganizationFeed: %#v", obj)
	billingProject := ""

	if parts := regexp.MustCompile(`projects\/([^\/]+)\/`).FindStringSubmatch(url); parts != nil {
		billingProject = parts[1]
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// Send the project ID in the X-Goog-User-Project header.
	origUserProjectOverride := config.UserProjectOverride
	config.UserProjectOverride = true
	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating OrganizationFeed: %s", err)
	}
	if err := d.Set("name", flattenCloudAssetOrganizationFeedName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Restore the original value of user_project_override.
	config.UserProjectOverride = origUserProjectOverride

	log.Printf("[DEBUG] Finished creating OrganizationFeed %q: %#v", d.Id(), res)

	return resourceCloudAssetOrganizationFeedRead(d, meta)
}

func resourceCloudAssetOrganizationFeedRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{CloudAssetBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	if parts := regexp.MustCompile(`projects\/([^\/]+)\/`).FindStringSubmatch(url); parts != nil {
		billingProject = parts[1]
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("CloudAssetOrganizationFeed %q", d.Id()))
	}

	if err := d.Set("name", flattenCloudAssetOrganizationFeedName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationFeed: %s", err)
	}
	if err := d.Set("asset_names", flattenCloudAssetOrganizationFeedAssetNames(res["assetNames"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationFeed: %s", err)
	}
	if err := d.Set("asset_types", flattenCloudAssetOrganizationFeedAssetTypes(res["assetTypes"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationFeed: %s", err)
	}
	if err := d.Set("content_type", flattenCloudAssetOrganizationFeedContentType(res["contentType"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationFeed: %s", err)
	}
	if err := d.Set("feed_output_config", flattenCloudAssetOrganizationFeedFeedOutputConfig(res["feedOutputConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationFeed: %s", err)
	}
	if err := d.Set("condition", flattenCloudAssetOrganizationFeedCondition(res["condition"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationFeed: %s", err)
	}

	return nil
}

func resourceCloudAssetOrganizationFeedUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	assetNamesProp, err := expandCloudAssetOrganizationFeedAssetNames(d.Get("asset_names"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("asset_names"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, assetNamesProp)) {
		obj["assetNames"] = assetNamesProp
	}
	assetTypesProp, err := expandCloudAssetOrganizationFeedAssetTypes(d.Get("asset_types"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("asset_types"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, assetTypesProp)) {
		obj["assetTypes"] = assetTypesProp
	}
	contentTypeProp, err := expandCloudAssetOrganizationFeedContentType(d.Get("content_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("content_type"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, contentTypeProp)) {
		obj["contentType"] = contentTypeProp
	}
	feedOutputConfigProp, err := expandCloudAssetOrganizationFeedFeedOutputConfig(d.Get("feed_output_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("feed_output_config"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, feedOutputConfigProp)) {
		obj["feedOutputConfig"] = feedOutputConfigProp
	}
	conditionProp, err := expandCloudAssetOrganizationFeedCondition(d.Get("condition"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("condition"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, conditionProp)) {
		obj["condition"] = conditionProp
	}

	obj, err = resourceCloudAssetOrganizationFeedEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{CloudAssetBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating OrganizationFeed %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("asset_names") {
		updateMask = append(updateMask, "assetNames")
	}

	if d.HasChange("asset_types") {
		updateMask = append(updateMask, "assetTypes")
	}

	if d.HasChange("content_type") {
		updateMask = append(updateMask, "contentType")
	}

	if d.HasChange("feed_output_config") {
		updateMask = append(updateMask, "feedOutputConfig")
	}

	if d.HasChange("condition") {
		updateMask = append(updateMask, "condition")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	if parts := regexp.MustCompile(`projects\/([^\/]+)\/`).FindStringSubmatch(url); parts != nil {
		billingProject = parts[1]
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating OrganizationFeed %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating OrganizationFeed %q: %#v", d.Id(), res)
	}

	return resourceCloudAssetOrganizationFeedRead(d, meta)
}

func resourceCloudAssetOrganizationFeedDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := replaceVars(d, config, "{{CloudAssetBasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	if parts := regexp.MustCompile(`projects\/([^\/]+)\/`).FindStringSubmatch(url); parts != nil {
		billingProject = parts[1]
	}

	log.Printf("[DEBUG] Deleting OrganizationFeed %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "OrganizationFeed")
	}

	log.Printf("[DEBUG] Finished deleting OrganizationFeed %q: %#v", d.Id(), res)
	return nil
}

func resourceCloudAssetOrganizationFeedImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	if err := d.Set("name", d.Id()); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func flattenCloudAssetOrganizationFeedName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudAssetOrganizationFeedAssetNames(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudAssetOrganizationFeedAssetTypes(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudAssetOrganizationFeedContentType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudAssetOrganizationFeedFeedOutputConfig(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["pubsub_destination"] =
		flattenCloudAssetOrganizationFeedFeedOutputConfigPubsubDestination(original["pubsubDestination"], d, config)
	return []interface{}{transformed}
}
func flattenCloudAssetOrganizationFeedFeedOutputConfigPubsubDestination(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["topic"] =
		flattenCloudAssetOrganizationFeedFeedOutputConfigPubsubDestinationTopic(original["topic"], d, config)
	return []interface{}{transformed}
}
func flattenCloudAssetOrganizationFeedFeedOutputConfigPubsubDestinationTopic(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudAssetOrganizationFeedCondition(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["expression"] =
		flattenCloudAssetOrganizationFeedConditionExpression(original["expression"], d, config)
	transformed["title"] =
		flattenCloudAssetOrganizationFeedConditionTitle(original["title"], d, config)
	transformed["description"] =
		flattenCloudAssetOrganizationFeedConditionDescription(original["description"], d, config)
	transformed["location"] =
		flattenCloudAssetOrganizationFeedConditionLocation(original["location"], d, config)
	return []interface{}{transformed}
}
func flattenCloudAssetOrganizationFeedConditionExpression(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudAssetOrganizationFeedConditionTitle(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudAssetOrganizationFeedConditionDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudAssetOrganizationFeedConditionLocation(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandCloudAssetOrganizationFeedAssetNames(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetOrganizationFeedAssetTypes(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetOrganizationFeedContentType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetOrganizationFeedFeedOutputConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPubsubDestination, err := expandCloudAssetOrganizationFeedFeedOutputConfigPubsubDestination(original["pubsub_destination"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPubsubDestination); val.IsValid() && !isEmptyValue(val) {
		transformed["pubsubDestination"] = transformedPubsubDestination
	}

	return transformed, nil
}

func expandCloudAssetOrganizationFeedFeedOutputConfigPubsubDestination(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTopic, err := expandCloudAssetOrganizationFeedFeedOutputConfigPubsubDestinationTopic(original["topic"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTopic); val.IsValid() && !isEmptyValue(val) {
		transformed["topic"] = transformedTopic
	}

	return transformed, nil
}

func expandCloudAssetOrganizationFeedFeedOutputConfigPubsubDestinationTopic(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetOrganizationFeedCondition(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedExpression, err := expandCloudAssetOrganizationFeedConditionExpression(original["expression"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExpression); val.IsValid() && !isEmptyValue(val) {
		transformed["expression"] = transformedExpression
	}

	transformedTitle, err := expandCloudAssetOrganizationFeedConditionTitle(original["title"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTitle); val.IsValid() && !isEmptyValue(val) {
		transformed["title"] = transformedTitle
	}

	transformedDescription, err := expandCloudAssetOrganizationFeedConditionDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !isEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	transformedLocation, err := expandCloudAssetOrganizationFeedConditionLocation(original["location"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocation); val.IsValid() && !isEmptyValue(val) {
		transformed["location"] = transformedLocation
	}

	return transformed, nil
}

func expandCloudAssetOrganizationFeedConditionExpression(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetOrganizationFeedConditionTitle(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetOrganizationFeedConditionDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetOrganizationFeedConditionLocation(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceCloudAssetOrganizationFeedEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// Remove the "folders/" prefix from the folder ID
	if folder, ok := d.GetOkExists("folder"); ok {
		if err := d.Set("folder_id", strings.TrimPrefix(folder.(string), "folders/")); err != nil {
			return nil, fmt.Errorf("Error setting folder_id: %s", err)
		}
	}
	// The feed object must be under the "feed" attribute on the request.
	newObj := make(map[string]interface{})
	newObj["feed"] = obj
	return newObj, nil
}
