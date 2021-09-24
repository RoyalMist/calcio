/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type {api_login} from '../models/api_login';
import {request as __request} from '../core/request';

export class AuthenticationService {

    /**
     * Permits a user to log in to Calcio if credentials are valid.
     * Log in and retrieve the PASETO token signed. This method is rate limited.
     * @param login Login json object
     * @returns string Paseto Token
     * @throws ApiError
     */
    public static async postAuthenticationService(
        login: api_login,
    ): Promise<string> {
        const result = await __request({
            method: 'POST',
            path: `/api/auth/login`,
            body: login,
            errors: {
                400: `When the token is absent or malformed`,
                401: `When the token is invalid`,
                429: `When the rate limit is reached`,
                500: `When something went wrong`,
            },
        });
        return result.body;
    }

}
