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
                400: `Wrong login information provided`,
                429: `Rate limit reached`,
                500: `Something went wrong`,
            },
        });
        return result.body;
    }

}
