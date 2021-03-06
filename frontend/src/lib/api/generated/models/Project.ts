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
import {
    Task,
    TaskFromJSON,
    TaskFromJSONTyped,
    TaskToJSON,
} from './';

/**
 * 
 * @export
 * @interface Project
 */
export interface Project {
    /**
     * 
     * @type {string}
     * @memberof Project
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof Project
     */
    ownerId?: string;
    /**
     * 
     * @type {string}
     * @memberof Project
     */
    name: string;
    /**
     * 
     * @type {string}
     * @memberof Project
     */
    elevatorPitch: string;
    /**
     * 
     * @type {Array<Task>}
     * @memberof Project
     */
    tasks?: Array<Task>;
    /**
     * 
     * @type {Date}
     * @memberof Project
     */
    createdAt?: Date;
    /**
     * 
     * @type {Date}
     * @memberof Project
     */
    updatedAt?: Date;
}

export function ProjectFromJSON(json: any): Project {
    return ProjectFromJSONTyped(json, false);
}

export function ProjectFromJSONTyped(json: any, ignoreDiscriminator: boolean): Project {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': !exists(json, 'id') ? undefined : json['id'],
        'ownerId': !exists(json, 'owner_id') ? undefined : json['owner_id'],
        'name': json['name'],
        'elevatorPitch': json['elevator_pitch'],
        'tasks': !exists(json, 'tasks') ? undefined : ((json['tasks'] as Array<any>).map(TaskFromJSON)),
        'createdAt': !exists(json, 'created_at') ? undefined : (new Date(json['created_at'])),
        'updatedAt': !exists(json, 'updated_at') ? undefined : (new Date(json['updated_at'])),
    };
}

export function ProjectToJSON(value?: Project | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'owner_id': value.ownerId,
        'name': value.name,
        'elevator_pitch': value.elevatorPitch,
        'tasks': value.tasks === undefined ? undefined : ((value.tasks as Array<any>).map(TaskToJSON)),
        'created_at': value.createdAt === undefined ? undefined : (value.createdAt.toISOString()),
        'updated_at': value.updatedAt === undefined ? undefined : (value.updatedAt.toISOString()),
    };
}


