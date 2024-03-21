# OpenapiClient::TasksGet200ResponseInner

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  |  |
| **type** | **String** |  |  |
| **account** | [**TasksGet200ResponseInnerAccount**](TasksGet200ResponseInnerAccount.md) |  |  |
| **creator** | **Integer** |  |  |
| **description** | **String** |  |  |
| **reminder_date** | **String** |  |  |
| **email** | [**TasksGet200ResponseInnerEmail**](TasksGet200ResponseInnerEmail.md) |  |  |
| **company** | **Object** |  |  |
| **meeting** | **Object** |  |  |
| **completed** | **Boolean** |  |  |
| **completion_date** | **Object** |  |  |
| **completor** | **Object** |  |  |
| **archived** | **Boolean** |  |  |
| **archive_date** | **Object** |  |  |
| **archivor** | **Object** |  |  |
| **last_modified_by** | **Integer** |  |  |
| **creation_date** | **String** |  |  |
| **modification_date** | **String** |  |  |
| **assignees** | [**Array&lt;TasksGet200ResponseInnerAssigneesInner&gt;**](TasksGet200ResponseInnerAssigneesInner.md) |  |  |
| **can_edit** | **Boolean** |  |  |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::TasksGet200ResponseInner.new(
  id: null,
  type: null,
  account: null,
  creator: null,
  description: null,
  reminder_date: null,
  email: null,
  company: null,
  meeting: null,
  completed: null,
  completion_date: null,
  completor: null,
  archived: null,
  archive_date: null,
  archivor: null,
  last_modified_by: null,
  creation_date: null,
  modification_date: null,
  assignees: null,
  can_edit: null
)
```

