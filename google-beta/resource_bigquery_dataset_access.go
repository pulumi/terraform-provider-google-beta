// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
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
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/googleapi"
)

var bigqueryAccessRoleToPrimitiveMap = map[string]string{
	"roles/bigquery.dataOwner":  "OWNER",
	"roles/bigquery.dataEditor": "WRITER",
	"roles/bigquery.dataViewer": "READER",
}

func resourceBigQueryDatasetAccessRoleDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	if primitiveRole, ok := bigqueryAccessRoleToPrimitiveMap[new]; ok {
		return primitiveRole == old
	}
	return false
}

// we want to diff suppress any iam_members that are configured as `iam_member`, but stored in state as a different member type
func resourceBigQueryDatasetAccessIamMemberDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	if primitiveRole, ok := bigqueryAccessRoleToPrimitiveMap[new]; ok {
		return primitiveRole == old
	}

	if d.Get("api_updated_member") == true {
		expectedIamMember := d.Get("iam_member").(string)
		parts := strings.SplitAfter(expectedIamMember, ":")

		strippedIamMember := parts[0]
		if len(parts) > 1 {
			strippedIamMember = parts[1]
		}

		if memberInState := d.Get("user_by_email").(string); memberInState != "" {
			return memberInState == strippedIamMember
		}

		if memberInState := d.Get("group_by_email").(string); memberInState != "" {
			return memberInState == strippedIamMember
		}

		if memberInState := d.Get("domain").(string); memberInState != "" {
			return memberInState == strippedIamMember
		}

		if memberInState := d.Get("special_group").(string); memberInState != "" {
			return memberInState == strippedIamMember
		}
	}

	return false
}

// this function will go through a response's access list and see if the iam_member has been reassigned to a different member_type
// if it has, it will return the member type, and the member
func resourceBigQueryDatasetAccessReassignIamMemberInNestedObjectList(d *schema.ResourceData, meta interface{}, items []interface{}) (member_type string, member interface{}, err error) {
	expectedRole, err := expandNestedBigQueryDatasetAccessRole(d.Get("role"), d, meta.(*Config))
	if err != nil {
		return "", nil, err
	}
	expectedFlattenedRole := flattenNestedBigQueryDatasetAccessRole(expectedRole, d, meta.(*Config))

	expectedIamMember, err := expandNestedBigQueryDatasetAccessIamMember(d.Get("iam_member"), d, meta.(*Config))
	if err != nil {
		return "", nil, err
	}
	expectedFlattenedIamMember := flattenNestedBigQueryDatasetAccessIamMember(expectedIamMember, d, meta.(*Config))

	parts := strings.SplitAfter(expectedFlattenedIamMember.(string), ":")

	expectedStrippedIamMember := parts[0]
	if len(parts) > 1 {
		expectedStrippedIamMember = parts[1]
	}

	// Search list for this resource.
	for _, itemRaw := range items {
		if itemRaw == nil {
			continue
		}
		item := itemRaw.(map[string]interface{})

		itemRole := flattenNestedBigQueryDatasetAccessRole(item["role"], d, meta.(*Config))
		// isEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(isEmptyValue(reflect.ValueOf(itemRole)) && isEmptyValue(reflect.ValueOf(expectedFlattenedRole))) && !reflect.DeepEqual(itemRole, expectedFlattenedRole) {
			log.Printf("[DEBUG] Skipping item with role= %#v, looking for %#v)", itemRole, expectedFlattenedRole)
			continue
		}

		itemUserByEmail := flattenNestedBigQueryDatasetAccessUserByEmail(item["userByEmail"], d, meta.(*Config))
		if reflect.DeepEqual(itemUserByEmail, expectedStrippedIamMember) {
			log.Printf("[DEBUG] Iam Member changed to userByEmail= %#v)", itemUserByEmail)
			return "user_by_email", itemUserByEmail, nil
		}
		itemGroupByEmail := flattenNestedBigQueryDatasetAccessGroupByEmail(item["groupByEmail"], d, meta.(*Config))
		if reflect.DeepEqual(itemGroupByEmail, expectedStrippedIamMember) {
			log.Printf("[DEBUG] Iam Member changed to groupByEmail= %#v)", itemGroupByEmail)
			return "group_by_email", itemGroupByEmail, nil
		}
		itemDomain := flattenNestedBigQueryDatasetAccessDomain(item["domain"], d, meta.(*Config))
		if reflect.DeepEqual(itemDomain, expectedStrippedIamMember) {
			log.Printf("[DEBUG] Iam Member changed to domain= %#v)", itemDomain)
			return "domain", itemDomain, nil
		}
		itemSpecialGroup := flattenNestedBigQueryDatasetAccessSpecialGroup(item["specialGroup"], d, meta.(*Config))
		if reflect.DeepEqual(itemSpecialGroup, expectedStrippedIamMember) {
			log.Printf("[DEBUG] Iam Member changed to specialGroup= %#v)", itemSpecialGroup)
			return "special_group", itemSpecialGroup, nil
		}
		itemIamMember := flattenNestedBigQueryDatasetAccessIamMember(item["iamMember"], d, meta.(*Config))
		if reflect.DeepEqual(itemIamMember, expectedFlattenedIamMember) {
			log.Printf("[DEBUG] Iam Member stayed as iamMember= %#v)", itemIamMember)
			return "", nil, nil
		}
		continue
	}
	log.Printf("[DEBUG] Did not find item for resource %q)", d.Id())
	return "", nil, nil
}

func resourceBigQueryDatasetAccess() *schema.Resource {
	return &schema.Resource{
		Create: resourceBigQueryDatasetAccessCreate,
		Read:   resourceBigQueryDatasetAccessRead,
		Delete: resourceBigQueryDatasetAccessDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"dataset_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `A unique ID for this dataset, without the project name. The ID
must contain only letters (a-z, A-Z), numbers (0-9), or
underscores (_). The maximum length is 1,024 characters.`,
			},
			"domain": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: resourceBigQueryDatasetAccessIamMemberDiffSuppress,
				Description: `A domain to grant access to. Any users signed in with the
domain specified will be granted the specified access`,
				ExactlyOneOf: []string{"user_by_email", "group_by_email", "domain", "special_group", "iam_member", "view"},
			},
			"group_by_email": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: resourceBigQueryDatasetAccessIamMemberDiffSuppress,
				Description:      `An email address of a Google Group to grant access to.`,
				ExactlyOneOf:     []string{"user_by_email", "group_by_email", "domain", "special_group", "iam_member", "view"},
			},
			"iam_member": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: resourceBigQueryDatasetAccessIamMemberDiffSuppress,
				Description: `Some other type of member that appears in the IAM Policy but isn't a user,
group, domain, or special group. For example: 'allUsers'`,
				ExactlyOneOf: []string{"user_by_email", "group_by_email", "domain", "special_group", "iam_member", "view"},
			},
			"role": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: resourceBigQueryDatasetAccessRoleDiffSuppress,
				Description: `Describes the rights granted to the user specified by the other
member of the access object. Basic, predefined, and custom roles are
supported. Predefined roles that have equivalent basic roles are
swapped by the API to their basic counterparts, and will show a diff
post-create. See
[official docs](https://cloud.google.com/bigquery/docs/access-control).`,
			},
			"special_group": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: resourceBigQueryDatasetAccessIamMemberDiffSuppress,
				Description: `A special group to grant access to. Possible values include:


* 'projectOwners': Owners of the enclosing project.


* 'projectReaders': Readers of the enclosing project.


* 'projectWriters': Writers of the enclosing project.


* 'allAuthenticatedUsers': All authenticated BigQuery users.`,
				ExactlyOneOf: []string{"user_by_email", "group_by_email", "domain", "special_group", "iam_member", "view"},
			},
			"user_by_email": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: resourceBigQueryDatasetAccessIamMemberDiffSuppress,
				Description: `An email address of a user to grant access to. For example:
fred@example.com`,
				ExactlyOneOf: []string{"user_by_email", "group_by_email", "domain", "special_group", "iam_member", "view"},
			},
			"view": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `A view from a different dataset to grant access to. Queries
executed against that view will have read access to tables in
this dataset. The role field is not required when this field is
set. If that view is updated by any user, access to the view
needs to be granted again via an update operation.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dataset_id": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: `The ID of the dataset containing this table.`,
						},
						"project_id": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: `The ID of the project containing this table.`,
						},
						"table_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
							Description: `The ID of the table. The ID must contain only letters (a-z,
A-Z), numbers (0-9), or underscores (_). The maximum length
is 1,024 characters.`,
						},
					},
				},
				ExactlyOneOf: []string{"user_by_email", "group_by_email", "domain", "special_group", "iam_member", "view"},
			},
			"api_updated_member": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "If true, represents that that the iam_member in the config was translated to a different member type by the API, and is stored in state as a different member type",
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

func resourceBigQueryDatasetAccessCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	datasetIdProp, err := expandNestedBigQueryDatasetAccessDatasetId(d.Get("dataset_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("dataset_id"); !isEmptyValue(reflect.ValueOf(datasetIdProp)) && (ok || !reflect.DeepEqual(v, datasetIdProp)) {
		obj["datasetId"] = datasetIdProp
	}
	roleProp, err := expandNestedBigQueryDatasetAccessRole(d.Get("role"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("role"); !isEmptyValue(reflect.ValueOf(roleProp)) && (ok || !reflect.DeepEqual(v, roleProp)) {
		obj["role"] = roleProp
	}
	userByEmailProp, err := expandNestedBigQueryDatasetAccessUserByEmail(d.Get("user_by_email"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("user_by_email"); !isEmptyValue(reflect.ValueOf(userByEmailProp)) && (ok || !reflect.DeepEqual(v, userByEmailProp)) {
		obj["userByEmail"] = userByEmailProp
	}
	groupByEmailProp, err := expandNestedBigQueryDatasetAccessGroupByEmail(d.Get("group_by_email"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("group_by_email"); !isEmptyValue(reflect.ValueOf(groupByEmailProp)) && (ok || !reflect.DeepEqual(v, groupByEmailProp)) {
		obj["groupByEmail"] = groupByEmailProp
	}
	domainProp, err := expandNestedBigQueryDatasetAccessDomain(d.Get("domain"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("domain"); !isEmptyValue(reflect.ValueOf(domainProp)) && (ok || !reflect.DeepEqual(v, domainProp)) {
		obj["domain"] = domainProp
	}
	specialGroupProp, err := expandNestedBigQueryDatasetAccessSpecialGroup(d.Get("special_group"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("special_group"); !isEmptyValue(reflect.ValueOf(specialGroupProp)) && (ok || !reflect.DeepEqual(v, specialGroupProp)) {
		obj["specialGroup"] = specialGroupProp
	}
	iamMemberProp, err := expandNestedBigQueryDatasetAccessIamMember(d.Get("iam_member"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("iam_member"); !isEmptyValue(reflect.ValueOf(iamMemberProp)) && (ok || !reflect.DeepEqual(v, iamMemberProp)) {
		obj["iamMember"] = iamMemberProp
	}
	viewProp, err := expandNestedBigQueryDatasetAccessView(d.Get("view"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("view"); !isEmptyValue(reflect.ValueOf(viewProp)) && (ok || !reflect.DeepEqual(v, viewProp)) {
		obj["view"] = viewProp
	}

	lockName, err := replaceVars(d, config, "{{dataset_id}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{BigQueryBasePath}}projects/{{project}}/datasets/{{dataset_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new DatasetAccess: %#v", obj)

	obj, err = resourceBigQueryDatasetAccessPatchCreateEncoder(d, meta, obj)
	if err != nil {
		return err
	}
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DatasetAccess: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate), isBigqueryIAMQuotaError)
	if err != nil {
		return fmt.Errorf("Error creating DatasetAccess: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/datasets/{{dataset_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// by default, we are not updating the member
	if err := d.Set("api_updated_member", false); err != nil {
		return fmt.Errorf("Error setting api_updated_member: %s", err)
	}

	// iam_member is a generalized attribute, if the API can map it to a different member type on the backend, it will return
	// the correct member_type in the response. If it cannot be mapped to a different member type, it will stay in iam_member.
	if iamMemberProp != "" {
		member_type, member, err := resourceBigQueryDatasetAccessReassignIamMemberInNestedObjectList(d, meta, res["access"].([]interface{}))
		if err != nil {
			fmt.Println(err)
		}

		// if the member type changed, we set that member_type in state (it's already in the response) and we clear iam_member
		// and we set "api_updated_member" to true to acknowledge that we are making this change
		if member_type != "" {
			if err := d.Set(member_type, member.(string)); err != nil {
				return fmt.Errorf("Error setting member_type: %s", err)
			}
			if err := d.Set("iam_member", ""); err != nil {
				return fmt.Errorf("Error setting iam_member: %s", err)
			}
			if err := d.Set("api_updated_member", true); err != nil {
				return fmt.Errorf("Error setting api_updated_member: %s", err)
			}
		}
	}

	log.Printf("[DEBUG] Finished creating DatasetAccess %q: %#v", d.Id(), res)

	return resourceBigQueryDatasetAccessRead(d, meta)
}

func resourceBigQueryDatasetAccessRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{BigQueryBasePath}}projects/{{project}}/datasets/{{dataset_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DatasetAccess: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil, isBigqueryIAMQuotaError)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("BigQueryDatasetAccess %q", d.Id()))
	}

	res, err = flattenNestedBigQueryDatasetAccess(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Object isn't there any more - remove it from the state.
		log.Printf("[DEBUG] Removing BigQueryDatasetAccess because it couldn't be matched.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading DatasetAccess: %s", err)
	}

	if err := d.Set("role", flattenNestedBigQueryDatasetAccessRole(res["role"], d, config)); err != nil {
		return fmt.Errorf("Error reading DatasetAccess: %s", err)
	}
	if err := d.Set("user_by_email", flattenNestedBigQueryDatasetAccessUserByEmail(res["userByEmail"], d, config)); err != nil {
		return fmt.Errorf("Error reading DatasetAccess: %s", err)
	}
	if err := d.Set("group_by_email", flattenNestedBigQueryDatasetAccessGroupByEmail(res["groupByEmail"], d, config)); err != nil {
		return fmt.Errorf("Error reading DatasetAccess: %s", err)
	}
	if err := d.Set("domain", flattenNestedBigQueryDatasetAccessDomain(res["domain"], d, config)); err != nil {
		return fmt.Errorf("Error reading DatasetAccess: %s", err)
	}
	if err := d.Set("special_group", flattenNestedBigQueryDatasetAccessSpecialGroup(res["specialGroup"], d, config)); err != nil {
		return fmt.Errorf("Error reading DatasetAccess: %s", err)
	}
	if err := d.Set("iam_member", flattenNestedBigQueryDatasetAccessIamMember(res["iamMember"], d, config)); err != nil {
		return fmt.Errorf("Error reading DatasetAccess: %s", err)
	}
	if err := d.Set("view", flattenNestedBigQueryDatasetAccessView(res["view"], d, config)); err != nil {
		return fmt.Errorf("Error reading DatasetAccess: %s", err)
	}

	return nil
}

func resourceBigQueryDatasetAccessDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DatasetAccess: %s", err)
	}
	billingProject = project

	lockName, err := replaceVars(d, config, "{{dataset_id}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{BigQueryBasePath}}projects/{{project}}/datasets/{{dataset_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	obj, err = resourceBigQueryDatasetAccessPatchDeleteEncoder(d, meta, obj)
	if err != nil {
		return handleNotFoundError(err, d, "DatasetAccess")
	}
	log.Printf("[DEBUG] Deleting DatasetAccess %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete), isBigqueryIAMQuotaError)
	if err != nil {
		return handleNotFoundError(err, d, "DatasetAccess")
	}

	log.Printf("[DEBUG] Finished deleting DatasetAccess %q: %#v", d.Id(), res)
	return nil
}

func flattenNestedBigQueryDatasetAccessRole(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNestedBigQueryDatasetAccessUserByEmail(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNestedBigQueryDatasetAccessGroupByEmail(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNestedBigQueryDatasetAccessDomain(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNestedBigQueryDatasetAccessSpecialGroup(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNestedBigQueryDatasetAccessIamMember(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNestedBigQueryDatasetAccessView(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["dataset_id"] =
		flattenNestedBigQueryDatasetAccessViewDatasetId(original["datasetId"], d, config)
	transformed["project_id"] =
		flattenNestedBigQueryDatasetAccessViewProjectId(original["projectId"], d, config)
	transformed["table_id"] =
		flattenNestedBigQueryDatasetAccessViewTableId(original["tableId"], d, config)
	return []interface{}{transformed}
}
func flattenNestedBigQueryDatasetAccessViewDatasetId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNestedBigQueryDatasetAccessViewProjectId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNestedBigQueryDatasetAccessViewTableId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandNestedBigQueryDatasetAccessDatasetId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNestedBigQueryDatasetAccessRole(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	if v == nil {
		return nil, nil
	}

	if primitiveRole, ok := bigqueryAccessRoleToPrimitiveMap[v.(string)]; ok {
		return primitiveRole, nil
	}
	return v, nil
}

func expandNestedBigQueryDatasetAccessUserByEmail(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNestedBigQueryDatasetAccessGroupByEmail(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNestedBigQueryDatasetAccessDomain(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNestedBigQueryDatasetAccessSpecialGroup(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNestedBigQueryDatasetAccessIamMember(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNestedBigQueryDatasetAccessView(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDatasetId, err := expandNestedBigQueryDatasetAccessViewDatasetId(original["dataset_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDatasetId); val.IsValid() && !isEmptyValue(val) {
		transformed["datasetId"] = transformedDatasetId
	}

	transformedProjectId, err := expandNestedBigQueryDatasetAccessViewProjectId(original["project_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !isEmptyValue(val) {
		transformed["projectId"] = transformedProjectId
	}

	transformedTableId, err := expandNestedBigQueryDatasetAccessViewTableId(original["table_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTableId); val.IsValid() && !isEmptyValue(val) {
		transformed["tableId"] = transformedTableId
	}

	return transformed, nil
}

func expandNestedBigQueryDatasetAccessViewDatasetId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNestedBigQueryDatasetAccessViewProjectId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNestedBigQueryDatasetAccessViewTableId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func flattenNestedBigQueryDatasetAccess(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	var v interface{}
	var ok bool

	v, ok = res["access"]
	if !ok || v == nil {
		return nil, nil
	}

	switch v.(type) {
	case []interface{}:
		break
	case map[string]interface{}:
		// Construct list out of single nested resource
		v = []interface{}{v}
	default:
		return nil, fmt.Errorf("expected list or map for value access. Actual value: %v", v)
	}

	_, item, err := resourceBigQueryDatasetAccessFindNestedObjectInList(d, meta, v.([]interface{}))
	if err != nil {
		return nil, err
	}
	return item, nil
}

func resourceBigQueryDatasetAccessFindNestedObjectInList(d *schema.ResourceData, meta interface{}, items []interface{}) (index int, item map[string]interface{}, err error) {
	expectedRole, err := expandNestedBigQueryDatasetAccessRole(d.Get("role"), d, meta.(*Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedRole := flattenNestedBigQueryDatasetAccessRole(expectedRole, d, meta.(*Config))
	expectedUserByEmail, err := expandNestedBigQueryDatasetAccessUserByEmail(d.Get("user_by_email"), d, meta.(*Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedUserByEmail := flattenNestedBigQueryDatasetAccessUserByEmail(expectedUserByEmail, d, meta.(*Config))
	expectedGroupByEmail, err := expandNestedBigQueryDatasetAccessGroupByEmail(d.Get("group_by_email"), d, meta.(*Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedGroupByEmail := flattenNestedBigQueryDatasetAccessGroupByEmail(expectedGroupByEmail, d, meta.(*Config))
	expectedDomain, err := expandNestedBigQueryDatasetAccessDomain(d.Get("domain"), d, meta.(*Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedDomain := flattenNestedBigQueryDatasetAccessDomain(expectedDomain, d, meta.(*Config))
	expectedSpecialGroup, err := expandNestedBigQueryDatasetAccessSpecialGroup(d.Get("special_group"), d, meta.(*Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedSpecialGroup := flattenNestedBigQueryDatasetAccessSpecialGroup(expectedSpecialGroup, d, meta.(*Config))
	expectedIamMember, err := expandNestedBigQueryDatasetAccessIamMember(d.Get("iam_member"), d, meta.(*Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedIamMember := flattenNestedBigQueryDatasetAccessIamMember(expectedIamMember, d, meta.(*Config))
	expectedView, err := expandNestedBigQueryDatasetAccessView(d.Get("view"), d, meta.(*Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedView := flattenNestedBigQueryDatasetAccessView(expectedView, d, meta.(*Config))

	// Search list for this resource.
	for idx, itemRaw := range items {
		if itemRaw == nil {
			continue
		}
		item := itemRaw.(map[string]interface{})

		itemRole := flattenNestedBigQueryDatasetAccessRole(item["role"], d, meta.(*Config))
		// isEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(isEmptyValue(reflect.ValueOf(itemRole)) && isEmptyValue(reflect.ValueOf(expectedFlattenedRole))) && !reflect.DeepEqual(itemRole, expectedFlattenedRole) {
			log.Printf("[DEBUG] Skipping item with role= %#v, looking for %#v)", itemRole, expectedFlattenedRole)
			continue
		}
		itemUserByEmail := flattenNestedBigQueryDatasetAccessUserByEmail(item["userByEmail"], d, meta.(*Config))
		// isEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(isEmptyValue(reflect.ValueOf(itemUserByEmail)) && isEmptyValue(reflect.ValueOf(expectedFlattenedUserByEmail))) && !reflect.DeepEqual(itemUserByEmail, expectedFlattenedUserByEmail) {
			log.Printf("[DEBUG] Skipping item with userByEmail= %#v, looking for %#v)", itemUserByEmail, expectedFlattenedUserByEmail)
			continue
		}
		itemGroupByEmail := flattenNestedBigQueryDatasetAccessGroupByEmail(item["groupByEmail"], d, meta.(*Config))
		// isEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(isEmptyValue(reflect.ValueOf(itemGroupByEmail)) && isEmptyValue(reflect.ValueOf(expectedFlattenedGroupByEmail))) && !reflect.DeepEqual(itemGroupByEmail, expectedFlattenedGroupByEmail) {
			log.Printf("[DEBUG] Skipping item with groupByEmail= %#v, looking for %#v)", itemGroupByEmail, expectedFlattenedGroupByEmail)
			continue
		}
		itemDomain := flattenNestedBigQueryDatasetAccessDomain(item["domain"], d, meta.(*Config))
		// isEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(isEmptyValue(reflect.ValueOf(itemDomain)) && isEmptyValue(reflect.ValueOf(expectedFlattenedDomain))) && !reflect.DeepEqual(itemDomain, expectedFlattenedDomain) {
			log.Printf("[DEBUG] Skipping item with domain= %#v, looking for %#v)", itemDomain, expectedFlattenedDomain)
			continue
		}
		itemSpecialGroup := flattenNestedBigQueryDatasetAccessSpecialGroup(item["specialGroup"], d, meta.(*Config))
		// isEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(isEmptyValue(reflect.ValueOf(itemSpecialGroup)) && isEmptyValue(reflect.ValueOf(expectedFlattenedSpecialGroup))) && !reflect.DeepEqual(itemSpecialGroup, expectedFlattenedSpecialGroup) {
			log.Printf("[DEBUG] Skipping item with specialGroup= %#v, looking for %#v)", itemSpecialGroup, expectedFlattenedSpecialGroup)
			continue
		}
		itemIamMember := flattenNestedBigQueryDatasetAccessIamMember(item["iamMember"], d, meta.(*Config))
		// isEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(isEmptyValue(reflect.ValueOf(itemIamMember)) && isEmptyValue(reflect.ValueOf(expectedFlattenedIamMember))) && !reflect.DeepEqual(itemIamMember, expectedFlattenedIamMember) {
			log.Printf("[DEBUG] Skipping item with iamMember= %#v, looking for %#v)", itemIamMember, expectedFlattenedIamMember)
			continue
		}
		itemView := flattenNestedBigQueryDatasetAccessView(item["view"], d, meta.(*Config))
		// isEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(isEmptyValue(reflect.ValueOf(itemView)) && isEmptyValue(reflect.ValueOf(expectedFlattenedView))) && !reflect.DeepEqual(itemView, expectedFlattenedView) {
			log.Printf("[DEBUG] Skipping item with view= %#v, looking for %#v)", itemView, expectedFlattenedView)
			continue
		}
		log.Printf("[DEBUG] Found item for resource %q: %#v)", d.Id(), item)
		return idx, item, nil
	}
	return -1, nil, nil
}

// PatchCreateEncoder handles creating request data to PATCH parent resource
// with list including new object.
func resourceBigQueryDatasetAccessPatchCreateEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	currItems, err := resourceBigQueryDatasetAccessListForPatch(d, meta)
	if err != nil {
		return nil, err
	}

	_, found, err := resourceBigQueryDatasetAccessFindNestedObjectInList(d, meta, currItems)
	if err != nil {
		return nil, err
	}

	// Return error if item already created.
	if found != nil {
		return nil, fmt.Errorf("Unable to create DatasetAccess, existing object already found: %+v", found)
	}

	// Return list with the resource to create appended
	res := map[string]interface{}{
		"access": append(currItems, obj),
	}

	return res, nil
}

// PatchDeleteEncoder handles creating request data to PATCH parent resource
// with list excluding object to delete.
func resourceBigQueryDatasetAccessPatchDeleteEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	currItems, err := resourceBigQueryDatasetAccessListForPatch(d, meta)
	if err != nil {
		return nil, err
	}

	idx, item, err := resourceBigQueryDatasetAccessFindNestedObjectInList(d, meta, currItems)
	if err != nil {
		return nil, err
	}
	if item == nil {
		// Spoof 404 error for proper handling by Delete (i.e. no-op)
		return nil, &googleapi.Error{
			Code:    404,
			Message: "DatasetAccess not found in list",
		}
	}

	updatedItems := append(currItems[:idx], currItems[idx+1:]...)
	res := map[string]interface{}{
		"access": updatedItems,
	}

	return res, nil
}

// ListForPatch handles making API request to get parent resource and
// extracting list of objects.
func resourceBigQueryDatasetAccessListForPatch(d *schema.ResourceData, meta interface{}) ([]interface{}, error) {
	config := meta.(*Config)
	url, err := replaceVars(d, config, "{{BigQueryBasePath}}projects/{{project}}/datasets/{{dataset_id}}")
	if err != nil {
		return nil, err
	}
	project, err := getProject(d, config)
	if err != nil {
		return nil, err
	}

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return nil, err
	}

	res, err := sendRequest(config, "GET", project, url, userAgent, nil, isBigqueryIAMQuotaError)
	if err != nil {
		return nil, err
	}

	var v interface{}
	var ok bool

	v, ok = res["access"]
	if ok && v != nil {
		ls, lsOk := v.([]interface{})
		if !lsOk {
			return nil, fmt.Errorf(`expected list for nested field "access"`)
		}
		return ls, nil
	}
	return nil, nil
}
