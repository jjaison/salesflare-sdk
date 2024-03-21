# TasksGet200ResponseInnerEmail


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **object** |  | 
**email_message_id** | **object** |  | 
**service_message_id** | **object** |  | 
**data_source** | **object** |  | 
**subject** | **object** |  | 

## Example

```python
from salesflare.models.tasks_get200_response_inner_email import TasksGet200ResponseInnerEmail

# TODO update the JSON string below
json = "{}"
# create an instance of TasksGet200ResponseInnerEmail from a JSON string
tasks_get200_response_inner_email_instance = TasksGet200ResponseInnerEmail.from_json(json)
# print the JSON string representation of the object
print TasksGet200ResponseInnerEmail.to_json()

# convert the object into a dict
tasks_get200_response_inner_email_dict = tasks_get200_response_inner_email_instance.to_dict()
# create an instance of TasksGet200ResponseInnerEmail from a dict
tasks_get200_response_inner_email_form_dict = tasks_get200_response_inner_email.from_dict(tasks_get200_response_inner_email_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


