/* tslint:disable */
/* eslint-disable */
/**
 * PM API Specifications
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 0.0.1
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface Task
 */
export interface Task {
    /**
     * 
     * @type {string}
     * @memberof Task
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof Task
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof Task
     */
    description?: string;
    /**
     * 
     * @type {string}
     * @memberof Task
     */
    status?: TaskStatusEnum;
    /**
     * 
     * @type {string}
     * @memberof Task
     */
    ownerId?: string;
    /**
     * 
     * @type {string}
     * @memberof Task
     */
    assigneeId?: string | null;
    /**
     * 
     * @type {Date}
     * @memberof Task
     */
    deadline?: Date | null;
    /**
     * 
     * @type {Date}
     * @memberof Task
     */
    createdAt?: Date;
    /**
     * 
     * @type {Date}
     * @memberof Task
     */
    updatedAt?: Date;
}

/**
* @export
* @enum {string}
*/
export enum TaskStatusEnum {
    Todo = 'TODO',
    Wip = 'WIP',
    Done = 'DONE',
    Cancel = 'CANCEL',
    Pending = 'PENDING'
}

export function TaskFromJSON(json: any): Task {
    return TaskFromJSONTyped(json, false);
}

export function TaskFromJSONTyped(json: any, ignoreDiscriminator: boolean): Task {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': !exists(json, 'id') ? undefined : json['id'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'status': !exists(json, 'status') ? undefined : json['status'],
        'ownerId': !exists(json, 'owner_id') ? undefined : json['owner_id'],
        'assigneeId': !exists(json, 'assignee_id') ? undefined : json['assignee_id'],
        'deadline': !exists(json, 'deadline') ? undefined : (json['deadline'] === null ? null : new Date(json['deadline'])),
        'createdAt': !exists(json, 'created_at') ? undefined : (new Date(json['created_at'])),
        'updatedAt': !exists(json, 'updated_at') ? undefined : (new Date(json['updated_at'])),
    };
}

export function TaskToJSON(value?: Task | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'name': value.name,
        'description': value.description,
        'status': value.status,
        'owner_id': value.ownerId,
        'assignee_id': value.assigneeId,
        'deadline': value.deadline === undefined ? undefined : (value.deadline === null ? null : value.deadline.toISOString()),
        'created_at': value.createdAt === undefined ? undefined : (value.createdAt.toISOString()),
        'updated_at': value.updatedAt === undefined ? undefined : (value.updatedAt.toISOString()),
    };
}


