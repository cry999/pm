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
    Credential,
    CredentialFromJSON,
    CredentialToJSON,
    Token,
    TokenFromJSON,
    TokenToJSON,
} from '../models';

export interface SigninRequest {
    credential: Credential;
}

export interface SignupRequest {
    credential: Credential;
}

/**
 * 
 */
export class OAuthApi extends runtime.BaseAPI {

    /**
     * Registered user sign in
     */
    async signinRaw(requestParameters: SigninRequest): Promise<runtime.ApiResponse<Token>> {
        if (requestParameters.credential === null || requestParameters.credential === undefined) {
            throw new runtime.RequiredError('credential','Required parameter requestParameters.credential was null or undefined when calling signin.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/oauth/signin`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: CredentialToJSON(requestParameters.credential),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TokenFromJSON(jsonValue));
    }

    /**
     * Registered user sign in
     */
    async signin(requestParameters: SigninRequest): Promise<Token> {
        const response = await this.signinRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async signupRaw(requestParameters: SignupRequest): Promise<runtime.ApiResponse<Token>> {
        if (requestParameters.credential === null || requestParameters.credential === undefined) {
            throw new runtime.RequiredError('credential','Required parameter requestParameters.credential was null or undefined when calling signup.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/oauth/signup`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: CredentialToJSON(requestParameters.credential),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TokenFromJSON(jsonValue));
    }

    /**
     */
    async signup(requestParameters: SignupRequest): Promise<Token> {
        const response = await this.signupRaw(requestParameters);
        return await response.value();
    }

}