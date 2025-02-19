/**
 * app.salesflare.com
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';
import { TasksGet200ResponseInnerAccount } from './tasksGet200ResponseInnerAccount';
import { TasksGet200ResponseInnerAssigneesInner } from './tasksGet200ResponseInnerAssigneesInner';
import { TasksGet200ResponseInnerEmail } from './tasksGet200ResponseInnerEmail';

export class TasksGet200ResponseInner {
    'id': number;
    'type': string;
    'account': TasksGet200ResponseInnerAccount;
    'creator': number | null;
    'description': string;
    'reminderDate': string;
    'email': TasksGet200ResponseInnerEmail;
    'company': any | null;
    'meeting': any | null;
    'completed': boolean;
    'completionDate': any | null;
    'completor': any | null;
    'archived': boolean;
    'archiveDate': any | null;
    'archivor': any | null;
    'lastModifiedBy': number | null;
    'creationDate': string;
    'modificationDate': string;
    'assignees': Array<TasksGet200ResponseInnerAssigneesInner>;
    'canEdit': boolean;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "id",
            "baseName": "id",
            "type": "number"
        },
        {
            "name": "type",
            "baseName": "type",
            "type": "string"
        },
        {
            "name": "account",
            "baseName": "account",
            "type": "TasksGet200ResponseInnerAccount"
        },
        {
            "name": "creator",
            "baseName": "creator",
            "type": "number"
        },
        {
            "name": "description",
            "baseName": "description",
            "type": "string"
        },
        {
            "name": "reminderDate",
            "baseName": "reminder_date",
            "type": "string"
        },
        {
            "name": "email",
            "baseName": "email",
            "type": "TasksGet200ResponseInnerEmail"
        },
        {
            "name": "company",
            "baseName": "company",
            "type": "any"
        },
        {
            "name": "meeting",
            "baseName": "meeting",
            "type": "any"
        },
        {
            "name": "completed",
            "baseName": "completed",
            "type": "boolean"
        },
        {
            "name": "completionDate",
            "baseName": "completion_date",
            "type": "any"
        },
        {
            "name": "completor",
            "baseName": "completor",
            "type": "any"
        },
        {
            "name": "archived",
            "baseName": "archived",
            "type": "boolean"
        },
        {
            "name": "archiveDate",
            "baseName": "archive_date",
            "type": "any"
        },
        {
            "name": "archivor",
            "baseName": "archivor",
            "type": "any"
        },
        {
            "name": "lastModifiedBy",
            "baseName": "last_modified_by",
            "type": "number"
        },
        {
            "name": "creationDate",
            "baseName": "creation_date",
            "type": "string"
        },
        {
            "name": "modificationDate",
            "baseName": "modification_date",
            "type": "string"
        },
        {
            "name": "assignees",
            "baseName": "assignees",
            "type": "Array<TasksGet200ResponseInnerAssigneesInner>"
        },
        {
            "name": "canEdit",
            "baseName": "can_edit",
            "type": "boolean"
        }    ];

    static getAttributeTypeMap() {
        return TasksGet200ResponseInner.attributeTypeMap;
    }
}

