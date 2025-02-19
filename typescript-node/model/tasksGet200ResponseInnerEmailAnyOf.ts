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

export class TasksGet200ResponseInnerEmailAnyOf {
    'id': number;
    'emailMessageId': string;
    'serviceMessageId': string;
    'dataSource': number;
    'subject': string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "id",
            "baseName": "id",
            "type": "number"
        },
        {
            "name": "emailMessageId",
            "baseName": "email_message_id",
            "type": "string"
        },
        {
            "name": "serviceMessageId",
            "baseName": "service_message_id",
            "type": "string"
        },
        {
            "name": "dataSource",
            "baseName": "data_source",
            "type": "number"
        },
        {
            "name": "subject",
            "baseName": "subject",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return TasksGet200ResponseInnerEmailAnyOf.attributeTypeMap;
    }
}

