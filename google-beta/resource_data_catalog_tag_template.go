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
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDataCatalogTagTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceDataCatalogTagTemplateCreate,
		Read:   resourceDataCatalogTagTemplateRead,
		Update: resourceDataCatalogTagTemplateUpdate,
		Delete: resourceDataCatalogTagTemplateDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDataCatalogTagTemplateImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"fields": {
				Type:        schema.TypeSet,
				Required:    true,
				ForceNew:    true,
				Description: `Set of tag template field IDs and the settings for the field. This set is an exhaustive list of the allowed fields. This set must contain at least one field and at most 500 fields.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"type": {
							Type:        schema.TypeList,
							Required:    true,
							Description: `The type of value this tag field can contain.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enum_type": {
										Type:     schema.TypeList,
										Optional: true,
										Description: `Represents an enum type.
 Exactly one of 'primitive_type' or 'enum_type' must be set`,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"allowed_values": {
													Type:     schema.TypeSet,
													Required: true,
													Description: `The set of allowed values for this enum. The display names of the
values must be case-insensitively unique within this set. Currently,
enum values can only be added to the list of allowed values. Deletion
and renaming of enum values are not supported.
Can have up to 500 allowed values.`,
													Elem: datacatalogTagTemplateFieldsFieldsTypeEnumTypeAllowedValuesSchema(),
													// Default schema.HashSchema is used.
												},
											},
										},
									},
									"primitive_type": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringInSlice([]string{"DOUBLE", "STRING", "BOOL", "TIMESTAMP", ""}, false),
										Description: `Represents primitive types - string, bool etc.
 Exactly one of 'primitive_type' or 'enum_type' must be set Possible values: ["DOUBLE", "STRING", "BOOL", "TIMESTAMP"]`,
									},
								},
							},
						},
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `A description for this field.`,
						},
						"display_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `The display name for this field.`,
						},
						"is_required": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: `Whether this is a required field. Defaults to false.`,
						},
						"order": {
							Type:     schema.TypeInt,
							Optional: true,
							Description: `The order of this field with respect to other fields in this tag template.
A higher value indicates a more important field. The value can be negative.
Multiple fields can have the same order, and field orders within a tag do not have to be sequential.`,
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The resource name of the tag template field in URL format. Example: projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}/fields/{field}`,
						},
					},
				},
			},
			"tag_template_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateRegexp(`^[a-z_][a-z0-9_]{0,63}$`),
				Description:  `The id of the tag template to create.`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The display name for this template.`,
			},
			"force_delete": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `This confirms the deletion of any possible tags using this template. Must be set to true in order to delete the tag template.`,
				Default:     false,
			},
			"region": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `Template location region.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The resource name of the tag template in URL format. Example: projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func datacatalogTagTemplateFieldsFieldsTypeEnumTypeAllowedValuesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The display name of the enum value.`,
			},
		},
	}
}

func resourceDataCatalogTagTemplateCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandDataCatalogTagTemplateDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	fieldsProp, err := expandDataCatalogTagTemplateFields(d.Get("fields"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("fields"); !isEmptyValue(reflect.ValueOf(fieldsProp)) && (ok || !reflect.DeepEqual(v, fieldsProp)) {
		obj["fields"] = fieldsProp
	}

	url, err := replaceVars(d, config, "{{DataCatalogBasePath}}projects/{{project}}/locations/{{region}}/tagTemplates?tagTemplateId={{tag_template_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new TagTemplate: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TagTemplate: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating TagTemplate: %s", err)
	}
	if err := d.Set("name", flattenDataCatalogTagTemplateName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating TagTemplate %q: %#v", d.Id(), res)

	return resourceDataCatalogTagTemplateRead(d, meta)
}

func resourceDataCatalogTagTemplateRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{DataCatalogBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TagTemplate: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("DataCatalogTagTemplate %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading TagTemplate: %s", err)
	}

	region, err := getRegion(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("region", region); err != nil {
		return fmt.Errorf("Error reading TagTemplate: %s", err)
	}

	if err := d.Set("name", flattenDataCatalogTagTemplateName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading TagTemplate: %s", err)
	}
	if err := d.Set("display_name", flattenDataCatalogTagTemplateDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading TagTemplate: %s", err)
	}
	if err := d.Set("fields", flattenDataCatalogTagTemplateFields(res["fields"], d, config)); err != nil {
		return fmt.Errorf("Error reading TagTemplate: %s", err)
	}

	return nil
}

func resourceDataCatalogTagTemplateUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TagTemplate: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandDataCatalogTagTemplateDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}

	url, err := replaceVars(d, config, "{{DataCatalogBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating TagTemplate %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating TagTemplate %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating TagTemplate %q: %#v", d.Id(), res)
	}

	return resourceDataCatalogTagTemplateRead(d, meta)
}

func resourceDataCatalogTagTemplateDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TagTemplate: %s", err)
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{DataCatalogBasePath}}{{name}}?force={{force_delete}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting TagTemplate %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "TagTemplate")
	}

	log.Printf("[DEBUG] Finished deleting TagTemplate %q: %#v", d.Id(), res)
	return nil
}

func resourceDataCatalogTagTemplateImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := parseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}

	name := d.Get("name").(string)
	egRegex := regexp.MustCompile("projects/(.+)/locations/(.+)/tagTemplates/(.+)")

	parts := egRegex.FindStringSubmatch(name)
	if len(parts) != 4 {
		return nil, fmt.Errorf("tag template name does not fit the format %s", egRegex)
	}
	if err := d.Set("project", parts[1]); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("region", parts[2]); err != nil {
		return nil, fmt.Errorf("Error setting region: %s", err)
	}
	if err := d.Set("tag_template_id", parts[3]); err != nil {
		return nil, fmt.Errorf("Error setting tag_template_id: %s", err)
	}
	return []*schema.ResourceData{d}, nil
}

func flattenDataCatalogTagTemplateName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataCatalogTagTemplateDisplayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataCatalogTagTemplateFields(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.(map[string]interface{})
	transformed := make([]interface{}, 0, len(l))
	for k, raw := range l {
		original := raw.(map[string]interface{})
		transformed = append(transformed, map[string]interface{}{
			"field_id":     k,
			"name":         flattenDataCatalogTagTemplateFieldsName(original["name"], d, config),
			"display_name": flattenDataCatalogTagTemplateFieldsDisplayName(original["displayName"], d, config),
			"description":  flattenDataCatalogTagTemplateFieldsDescription(original["description"], d, config),
			"type":         flattenDataCatalogTagTemplateFieldsType(original["type"], d, config),
			"is_required":  flattenDataCatalogTagTemplateFieldsIsRequired(original["isRequired"], d, config),
			"order":        flattenDataCatalogTagTemplateFieldsOrder(original["order"], d, config),
		})
	}
	return transformed
}
func flattenDataCatalogTagTemplateFieldsName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataCatalogTagTemplateFieldsDisplayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataCatalogTagTemplateFieldsDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataCatalogTagTemplateFieldsType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["primitive_type"] =
		flattenDataCatalogTagTemplateFieldsTypePrimitiveType(original["primitiveType"], d, config)
	transformed["enum_type"] =
		flattenDataCatalogTagTemplateFieldsTypeEnumType(original["enumType"], d, config)
	return []interface{}{transformed}
}
func flattenDataCatalogTagTemplateFieldsTypePrimitiveType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataCatalogTagTemplateFieldsTypeEnumType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["allowed_values"] =
		flattenDataCatalogTagTemplateFieldsTypeEnumTypeAllowedValues(original["allowedValues"], d, config)
	return []interface{}{transformed}
}
func flattenDataCatalogTagTemplateFieldsTypeEnumTypeAllowedValues(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(schema.HashResource(datacatalogTagTemplateFieldsFieldsTypeEnumTypeAllowedValuesSchema()), []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"display_name": flattenDataCatalogTagTemplateFieldsTypeEnumTypeAllowedValuesDisplayName(original["displayName"], d, config),
		})
	}
	return transformed
}
func flattenDataCatalogTagTemplateFieldsTypeEnumTypeAllowedValuesDisplayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataCatalogTagTemplateFieldsIsRequired(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenDataCatalogTagTemplateFieldsOrder(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func expandDataCatalogTagTemplateDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogTagTemplateFields(v interface{}, d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	if v == nil {
		return map[string]interface{}{}, nil
	}
	m := make(map[string]interface{})
	for _, raw := range v.(*schema.Set).List() {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandDataCatalogTagTemplateFieldsName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedDisplayName, err := expandDataCatalogTagTemplateFieldsDisplayName(original["display_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDisplayName); val.IsValid() && !isEmptyValue(val) {
			transformed["displayName"] = transformedDisplayName
		}

		transformedDescription, err := expandDataCatalogTagTemplateFieldsDescription(original["description"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !isEmptyValue(val) {
			transformed["description"] = transformedDescription
		}

		transformedType, err := expandDataCatalogTagTemplateFieldsType(original["type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedType); val.IsValid() && !isEmptyValue(val) {
			transformed["type"] = transformedType
		}

		transformedIsRequired, err := expandDataCatalogTagTemplateFieldsIsRequired(original["is_required"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIsRequired); val.IsValid() && !isEmptyValue(val) {
			transformed["isRequired"] = transformedIsRequired
		}

		transformedOrder, err := expandDataCatalogTagTemplateFieldsOrder(original["order"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedOrder); val.IsValid() && !isEmptyValue(val) {
			transformed["order"] = transformedOrder
		}

		transformedFieldId, err := expandString(original["field_id"], d, config)
		if err != nil {
			return nil, err
		}
		m[transformedFieldId] = transformed
	}
	return m, nil
}

func expandDataCatalogTagTemplateFieldsName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogTagTemplateFieldsDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogTagTemplateFieldsDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogTagTemplateFieldsType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPrimitiveType, err := expandDataCatalogTagTemplateFieldsTypePrimitiveType(original["primitive_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPrimitiveType); val.IsValid() && !isEmptyValue(val) {
		transformed["primitiveType"] = transformedPrimitiveType
	}

	transformedEnumType, err := expandDataCatalogTagTemplateFieldsTypeEnumType(original["enum_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnumType); val.IsValid() && !isEmptyValue(val) {
		transformed["enumType"] = transformedEnumType
	}

	return transformed, nil
}

func expandDataCatalogTagTemplateFieldsTypePrimitiveType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogTagTemplateFieldsTypeEnumType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAllowedValues, err := expandDataCatalogTagTemplateFieldsTypeEnumTypeAllowedValues(original["allowed_values"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowedValues); val.IsValid() && !isEmptyValue(val) {
		transformed["allowedValues"] = transformedAllowedValues
	}

	return transformed, nil
}

func expandDataCatalogTagTemplateFieldsTypeEnumTypeAllowedValues(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDisplayName, err := expandDataCatalogTagTemplateFieldsTypeEnumTypeAllowedValuesDisplayName(original["display_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDisplayName); val.IsValid() && !isEmptyValue(val) {
			transformed["displayName"] = transformedDisplayName
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDataCatalogTagTemplateFieldsTypeEnumTypeAllowedValuesDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogTagTemplateFieldsIsRequired(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogTagTemplateFieldsOrder(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
