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
    Project,
    ProjectFromJSON,
    ProjectFromJSONTyped,
    ProjectToJSON,
} from './';

/**
 * 
 * @export
 * @interface ProjectList
 */
export interface ProjectList {
    /**
     * 
     * @type {Array<Project>}
     * @memberof ProjectList
     */
    results?: Array<Project>;
}

export function ProjectListFromJSON(json: any): ProjectList {
    return ProjectListFromJSONTyped(json, false);
}

export function ProjectListFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProjectList {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'results': !exists(json, 'results') ? undefined : ((json['results'] as Array<any>).map(ProjectFromJSON)),
    };
}

export function ProjectListToJSON(value?: ProjectList | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'results': value.results === undefined ? undefined : ((value.results as Array<any>).map(ProjectToJSON)),
    };
}


