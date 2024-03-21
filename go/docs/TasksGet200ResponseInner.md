# TasksGet200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Type** | **string** |  | 
**Account** | [**TasksGet200ResponseInnerAccount**](TasksGet200ResponseInnerAccount.md) |  | 
**Creator** | **NullableInt32** |  | 
**Description** | **string** |  | 
**ReminderDate** | **string** |  | 
**Email** | [**TasksGet200ResponseInnerEmail**](TasksGet200ResponseInnerEmail.md) |  | 
**Company** | **interface{}** |  | 
**Meeting** | **interface{}** |  | 
**Completed** | **bool** |  | 
**CompletionDate** | **interface{}** |  | 
**Completor** | **interface{}** |  | 
**Archived** | **bool** |  | 
**ArchiveDate** | **interface{}** |  | 
**Archivor** | **interface{}** |  | 
**LastModifiedBy** | **NullableInt32** |  | 
**CreationDate** | **string** |  | 
**ModificationDate** | **string** |  | 
**Assignees** | [**[]TasksGet200ResponseInnerAssigneesInner**](TasksGet200ResponseInnerAssigneesInner.md) |  | 
**CanEdit** | **bool** |  | 

## Methods

### NewTasksGet200ResponseInner

`func NewTasksGet200ResponseInner(id int32, type_ string, account TasksGet200ResponseInnerAccount, creator NullableInt32, description string, reminderDate string, email TasksGet200ResponseInnerEmail, company interface{}, meeting interface{}, completed bool, completionDate interface{}, completor interface{}, archived bool, archiveDate interface{}, archivor interface{}, lastModifiedBy NullableInt32, creationDate string, modificationDate string, assignees []TasksGet200ResponseInnerAssigneesInner, canEdit bool, ) *TasksGet200ResponseInner`

NewTasksGet200ResponseInner instantiates a new TasksGet200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTasksGet200ResponseInnerWithDefaults

`func NewTasksGet200ResponseInnerWithDefaults() *TasksGet200ResponseInner`

NewTasksGet200ResponseInnerWithDefaults instantiates a new TasksGet200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TasksGet200ResponseInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TasksGet200ResponseInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TasksGet200ResponseInner) SetId(v int32)`

SetId sets Id field to given value.


### GetType

`func (o *TasksGet200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TasksGet200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TasksGet200ResponseInner) SetType(v string)`

SetType sets Type field to given value.


### GetAccount

`func (o *TasksGet200ResponseInner) GetAccount() TasksGet200ResponseInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *TasksGet200ResponseInner) GetAccountOk() (*TasksGet200ResponseInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *TasksGet200ResponseInner) SetAccount(v TasksGet200ResponseInnerAccount)`

SetAccount sets Account field to given value.


### GetCreator

`func (o *TasksGet200ResponseInner) GetCreator() int32`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *TasksGet200ResponseInner) GetCreatorOk() (*int32, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *TasksGet200ResponseInner) SetCreator(v int32)`

SetCreator sets Creator field to given value.


### SetCreatorNil

`func (o *TasksGet200ResponseInner) SetCreatorNil(b bool)`

 SetCreatorNil sets the value for Creator to be an explicit nil

### UnsetCreator
`func (o *TasksGet200ResponseInner) UnsetCreator()`

UnsetCreator ensures that no value is present for Creator, not even an explicit nil
### GetDescription

`func (o *TasksGet200ResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TasksGet200ResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TasksGet200ResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetReminderDate

`func (o *TasksGet200ResponseInner) GetReminderDate() string`

GetReminderDate returns the ReminderDate field if non-nil, zero value otherwise.

### GetReminderDateOk

`func (o *TasksGet200ResponseInner) GetReminderDateOk() (*string, bool)`

GetReminderDateOk returns a tuple with the ReminderDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderDate

`func (o *TasksGet200ResponseInner) SetReminderDate(v string)`

SetReminderDate sets ReminderDate field to given value.


### GetEmail

`func (o *TasksGet200ResponseInner) GetEmail() TasksGet200ResponseInnerEmail`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TasksGet200ResponseInner) GetEmailOk() (*TasksGet200ResponseInnerEmail, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TasksGet200ResponseInner) SetEmail(v TasksGet200ResponseInnerEmail)`

SetEmail sets Email field to given value.


### GetCompany

`func (o *TasksGet200ResponseInner) GetCompany() interface{}`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *TasksGet200ResponseInner) GetCompanyOk() (*interface{}, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *TasksGet200ResponseInner) SetCompany(v interface{})`

SetCompany sets Company field to given value.


### SetCompanyNil

`func (o *TasksGet200ResponseInner) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *TasksGet200ResponseInner) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetMeeting

`func (o *TasksGet200ResponseInner) GetMeeting() interface{}`

GetMeeting returns the Meeting field if non-nil, zero value otherwise.

### GetMeetingOk

`func (o *TasksGet200ResponseInner) GetMeetingOk() (*interface{}, bool)`

GetMeetingOk returns a tuple with the Meeting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeeting

`func (o *TasksGet200ResponseInner) SetMeeting(v interface{})`

SetMeeting sets Meeting field to given value.


### SetMeetingNil

`func (o *TasksGet200ResponseInner) SetMeetingNil(b bool)`

 SetMeetingNil sets the value for Meeting to be an explicit nil

### UnsetMeeting
`func (o *TasksGet200ResponseInner) UnsetMeeting()`

UnsetMeeting ensures that no value is present for Meeting, not even an explicit nil
### GetCompleted

`func (o *TasksGet200ResponseInner) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *TasksGet200ResponseInner) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *TasksGet200ResponseInner) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.


### GetCompletionDate

`func (o *TasksGet200ResponseInner) GetCompletionDate() interface{}`

GetCompletionDate returns the CompletionDate field if non-nil, zero value otherwise.

### GetCompletionDateOk

`func (o *TasksGet200ResponseInner) GetCompletionDateOk() (*interface{}, bool)`

GetCompletionDateOk returns a tuple with the CompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDate

`func (o *TasksGet200ResponseInner) SetCompletionDate(v interface{})`

SetCompletionDate sets CompletionDate field to given value.


### SetCompletionDateNil

`func (o *TasksGet200ResponseInner) SetCompletionDateNil(b bool)`

 SetCompletionDateNil sets the value for CompletionDate to be an explicit nil

### UnsetCompletionDate
`func (o *TasksGet200ResponseInner) UnsetCompletionDate()`

UnsetCompletionDate ensures that no value is present for CompletionDate, not even an explicit nil
### GetCompletor

`func (o *TasksGet200ResponseInner) GetCompletor() interface{}`

GetCompletor returns the Completor field if non-nil, zero value otherwise.

### GetCompletorOk

`func (o *TasksGet200ResponseInner) GetCompletorOk() (*interface{}, bool)`

GetCompletorOk returns a tuple with the Completor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletor

`func (o *TasksGet200ResponseInner) SetCompletor(v interface{})`

SetCompletor sets Completor field to given value.


### SetCompletorNil

`func (o *TasksGet200ResponseInner) SetCompletorNil(b bool)`

 SetCompletorNil sets the value for Completor to be an explicit nil

### UnsetCompletor
`func (o *TasksGet200ResponseInner) UnsetCompletor()`

UnsetCompletor ensures that no value is present for Completor, not even an explicit nil
### GetArchived

`func (o *TasksGet200ResponseInner) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *TasksGet200ResponseInner) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *TasksGet200ResponseInner) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetArchiveDate

`func (o *TasksGet200ResponseInner) GetArchiveDate() interface{}`

GetArchiveDate returns the ArchiveDate field if non-nil, zero value otherwise.

### GetArchiveDateOk

`func (o *TasksGet200ResponseInner) GetArchiveDateOk() (*interface{}, bool)`

GetArchiveDateOk returns a tuple with the ArchiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveDate

`func (o *TasksGet200ResponseInner) SetArchiveDate(v interface{})`

SetArchiveDate sets ArchiveDate field to given value.


### SetArchiveDateNil

`func (o *TasksGet200ResponseInner) SetArchiveDateNil(b bool)`

 SetArchiveDateNil sets the value for ArchiveDate to be an explicit nil

### UnsetArchiveDate
`func (o *TasksGet200ResponseInner) UnsetArchiveDate()`

UnsetArchiveDate ensures that no value is present for ArchiveDate, not even an explicit nil
### GetArchivor

`func (o *TasksGet200ResponseInner) GetArchivor() interface{}`

GetArchivor returns the Archivor field if non-nil, zero value otherwise.

### GetArchivorOk

`func (o *TasksGet200ResponseInner) GetArchivorOk() (*interface{}, bool)`

GetArchivorOk returns a tuple with the Archivor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivor

`func (o *TasksGet200ResponseInner) SetArchivor(v interface{})`

SetArchivor sets Archivor field to given value.


### SetArchivorNil

`func (o *TasksGet200ResponseInner) SetArchivorNil(b bool)`

 SetArchivorNil sets the value for Archivor to be an explicit nil

### UnsetArchivor
`func (o *TasksGet200ResponseInner) UnsetArchivor()`

UnsetArchivor ensures that no value is present for Archivor, not even an explicit nil
### GetLastModifiedBy

`func (o *TasksGet200ResponseInner) GetLastModifiedBy() int32`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *TasksGet200ResponseInner) GetLastModifiedByOk() (*int32, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *TasksGet200ResponseInner) SetLastModifiedBy(v int32)`

SetLastModifiedBy sets LastModifiedBy field to given value.


### SetLastModifiedByNil

`func (o *TasksGet200ResponseInner) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *TasksGet200ResponseInner) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
### GetCreationDate

`func (o *TasksGet200ResponseInner) GetCreationDate() string`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *TasksGet200ResponseInner) GetCreationDateOk() (*string, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *TasksGet200ResponseInner) SetCreationDate(v string)`

SetCreationDate sets CreationDate field to given value.


### GetModificationDate

`func (o *TasksGet200ResponseInner) GetModificationDate() string`

GetModificationDate returns the ModificationDate field if non-nil, zero value otherwise.

### GetModificationDateOk

`func (o *TasksGet200ResponseInner) GetModificationDateOk() (*string, bool)`

GetModificationDateOk returns a tuple with the ModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationDate

`func (o *TasksGet200ResponseInner) SetModificationDate(v string)`

SetModificationDate sets ModificationDate field to given value.


### GetAssignees

`func (o *TasksGet200ResponseInner) GetAssignees() []TasksGet200ResponseInnerAssigneesInner`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *TasksGet200ResponseInner) GetAssigneesOk() (*[]TasksGet200ResponseInnerAssigneesInner, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *TasksGet200ResponseInner) SetAssignees(v []TasksGet200ResponseInnerAssigneesInner)`

SetAssignees sets Assignees field to given value.


### GetCanEdit

`func (o *TasksGet200ResponseInner) GetCanEdit() bool`

GetCanEdit returns the CanEdit field if non-nil, zero value otherwise.

### GetCanEditOk

`func (o *TasksGet200ResponseInner) GetCanEditOk() (*bool, bool)`

GetCanEditOk returns a tuple with the CanEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEdit

`func (o *TasksGet200ResponseInner) SetCanEdit(v bool)`

SetCanEdit sets CanEdit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


