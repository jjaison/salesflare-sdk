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
import { TasksGet200ResponseInnerEmailAnyOf } from './tasksGet200ResponseInnerEmailAnyOf';

export class TasksGet200ResponseInnerEmail {
    'id': any | null;
    'emailMessageId': any | null;
    'serviceMessageId': any | null;
    'dataSource': any | null;
    'subject': any | null;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "id",
            "baseName": "id",
            "type": "any"
        },
        {
            "name": "emailMessageId",
            "baseName": "email_message_id",
            "type": "any"
        },
        {
            "name": "serviceMessageId",
            "baseName": "service_message_id",
            "type": "any"
        },
        {
            "name": "dataSource",
            "baseName": "data_source",
            "type": "any"
        },
        {
            "name": "subject",
            "baseName": "subject",
            "type": "any"
        }    ];

    static getAttributeTypeMap() {
        return TasksGet200ResponseInnerEmail.attributeTypeMap;
    }
}

