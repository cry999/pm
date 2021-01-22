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


import * as runtime from '../runtime';
import {
    Task,
    TaskFromJSON,
    TaskToJSON,
    TaskForm,
    TaskFormFromJSON,
    TaskFormToJSON,
    TaskList,
    TaskListFromJSON,
    TaskListToJSON,
} from '../models';

export interface CancelTaskRequest {
    taskId: string;
}

export interface CreateTaskRequest {
    taskForm: TaskForm;
}

export interface DetailTaskRequest {
    taskId: string;
}

export interface DoneTaskRequest {
    taskId: string;
}

export interface ProgressTaskRequest {
    taskId: string;
}

export interface PutOnHoldTaskRequest {
    taskId: string;
}

export interface RedoTaskRequest {
    taskId: string;
}

/**
 * 
 */
export class TaskApi extends runtime.BaseAPI {

    /**
     * The Task is not done for the time being
     */
    async cancelTaskRaw(requestParameters: CancelTaskRequest): Promise<runtime.ApiResponse<Task>> {
        if (requestParameters.taskId === null || requestParameters.taskId === undefined) {
            throw new runtime.RequiredError('taskId','Required parameter requestParameters.taskId was null or undefined when calling cancelTask.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = typeof token === 'function' ? token("bearer", []) : token;

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/tasks/{task_id}/cancel`.replace(`{${"task_id"}}`, encodeURIComponent(String(requestParameters.taskId))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TaskFromJSON(jsonValue));
    }

    /**
     * The Task is not done for the time being
     */
    async cancelTask(requestParameters: CancelTaskRequest): Promise<Task> {
        const response = await this.cancelTaskRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create a new task
     */
    async createTaskRaw(requestParameters: CreateTaskRequest): Promise<runtime.ApiResponse<Task>> {
        if (requestParameters.taskForm === null || requestParameters.taskForm === undefined) {
            throw new runtime.RequiredError('taskForm','Required parameter requestParameters.taskForm was null or undefined when calling createTask.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = typeof token === 'function' ? token("bearer", []) : token;

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/tasks`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: TaskFormToJSON(requestParameters.taskForm),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TaskFromJSON(jsonValue));
    }

    /**
     * Create a new task
     */
    async createTask(requestParameters: CreateTaskRequest): Promise<Task> {
        const response = await this.createTaskRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async detailTaskRaw(requestParameters: DetailTaskRequest): Promise<runtime.ApiResponse<Task>> {
        if (requestParameters.taskId === null || requestParameters.taskId === undefined) {
            throw new runtime.RequiredError('taskId','Required parameter requestParameters.taskId was null or undefined when calling detailTask.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = typeof token === 'function' ? token("bearer", []) : token;

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/tasks/{task_id}`.replace(`{${"task_id"}}`, encodeURIComponent(String(requestParameters.taskId))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TaskFromJSON(jsonValue));
    }

    /**
     */
    async detailTask(requestParameters: DetailTaskRequest): Promise<Task> {
        const response = await this.detailTaskRaw(requestParameters);
        return await response.value();
    }

    /**
     * Work in progress Task is done
     */
    async doneTaskRaw(requestParameters: DoneTaskRequest): Promise<runtime.ApiResponse<Task>> {
        if (requestParameters.taskId === null || requestParameters.taskId === undefined) {
            throw new runtime.RequiredError('taskId','Required parameter requestParameters.taskId was null or undefined when calling doneTask.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = typeof token === 'function' ? token("bearer", []) : token;

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/tasks/{task_id}/done`.replace(`{${"task_id"}}`, encodeURIComponent(String(requestParameters.taskId))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TaskFromJSON(jsonValue));
    }

    /**
     * Work in progress Task is done
     */
    async doneTask(requestParameters: DoneTaskRequest): Promise<Task> {
        const response = await this.doneTaskRaw(requestParameters);
        return await response.value();
    }

    /**
     * List all tasks associated with authenticated user.
     */
    async listAllAssociatedWithUserRaw(): Promise<runtime.ApiResponse<TaskList>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = typeof token === 'function' ? token("bearer", []) : token;

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/tasks`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TaskListFromJSON(jsonValue));
    }

    /**
     * List all tasks associated with authenticated user.
     */
    async listAllAssociatedWithUser(): Promise<TaskList> {
        const response = await this.listAllAssociatedWithUserRaw();
        return await response.value();
    }

    /**
     * Make TODO task work in progress
     */
    async progressTaskRaw(requestParameters: ProgressTaskRequest): Promise<runtime.ApiResponse<Task>> {
        if (requestParameters.taskId === null || requestParameters.taskId === undefined) {
            throw new runtime.RequiredError('taskId','Required parameter requestParameters.taskId was null or undefined when calling progressTask.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = typeof token === 'function' ? token("bearer", []) : token;

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/tasks/{task_id}/progress`.replace(`{${"task_id"}}`, encodeURIComponent(String(requestParameters.taskId))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TaskFromJSON(jsonValue));
    }

    /**
     * Make TODO task work in progress
     */
    async progressTask(requestParameters: ProgressTaskRequest): Promise<Task> {
        const response = await this.progressTaskRaw(requestParameters);
        return await response.value();
    }

    /**
     * The Task will be put on hold for a while.
     */
    async putOnHoldTaskRaw(requestParameters: PutOnHoldTaskRequest): Promise<runtime.ApiResponse<Task>> {
        if (requestParameters.taskId === null || requestParameters.taskId === undefined) {
            throw new runtime.RequiredError('taskId','Required parameter requestParameters.taskId was null or undefined when calling putOnHoldTask.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = typeof token === 'function' ? token("bearer", []) : token;

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/tasks/{task_id}/hold`.replace(`{${"task_id"}}`, encodeURIComponent(String(requestParameters.taskId))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TaskFromJSON(jsonValue));
    }

    /**
     * The Task will be put on hold for a while.
     */
    async putOnHoldTask(requestParameters: PutOnHoldTaskRequest): Promise<Task> {
        const response = await this.putOnHoldTaskRaw(requestParameters);
        return await response.value();
    }

    /**
     * Redo the Task
     */
    async redoTaskRaw(requestParameters: RedoTaskRequest): Promise<runtime.ApiResponse<Task>> {
        if (requestParameters.taskId === null || requestParameters.taskId === undefined) {
            throw new runtime.RequiredError('taskId','Required parameter requestParameters.taskId was null or undefined when calling redoTask.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = typeof token === 'function' ? token("bearer", []) : token;

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/tasks/{task_id}/redo`.replace(`{${"task_id"}}`, encodeURIComponent(String(requestParameters.taskId))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TaskFromJSON(jsonValue));
    }

    /**
     * Redo the Task
     */
    async redoTask(requestParameters: RedoTaskRequest): Promise<Task> {
        const response = await this.redoTaskRaw(requestParameters);
        return await response.value();
    }

}
