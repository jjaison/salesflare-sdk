# TasksGet200ResponseInner


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** |  | 
**type** | **str** |  | 
**account** | [**TasksGet200ResponseInnerAccount**](TasksGet200ResponseInnerAccount.md) |  | 
**creator** | **int** |  | 
**description** | **str** |  | 
**reminder_date** | **str** |  | 
**email** | [**TasksGet200ResponseInnerEmail**](TasksGet200ResponseInnerEmail.md) |  | 
**company** | **object** |  | 
**meeting** | **object** |  | 
**completed** | **bool** |  | 
**completion_date** | **object** |  | 
**completor** | **object** |  | 
**archived** | **bool** |  | 
**archive_date** | **object** |  | 
**archivor** | **object** |  | 
**last_modified_by** | **int** |  | 
**creation_date** | **str** |  | 
**modification_date** | **str** |  | 
**assignees** | [**List[TasksGet200ResponseInnerAssigneesInner]**](TasksGet200ResponseInnerAssigneesInner.md) |  | 
**can_edit** | **bool** |  | 

## Example

```python
from salesflare.models.tasks_get200_response_inner import TasksGet200ResponseInner

# TODO update the JSON string below
json = "{}"
# create an instance of TasksGet200ResponseInner from a JSON string
tasks_get200_response_inner_instance = TasksGet200ResponseInner.from_json(json)
# print the JSON string representation of the object
print TasksGet200ResponseInner.to_json()

# convert the object into a dict
tasks_get200_response_inner_dict = tasks_get200_response_inner_instance.to_dict()
# create an instance of TasksGet200ResponseInner from a dict
tasks_get200_response_inner_form_dict = tasks_get200_response_inner.from_dict(tasks_get200_response_inner_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


