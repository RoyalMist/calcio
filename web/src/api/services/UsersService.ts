/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type {ent_User} from '../models/ent_User';
import {request as __request} from '../core/request';

export class UsersService {

    /**
     * Fetch all Calcio's users.
     * Retrieves all Calcio's users as a json list.
     * @param authorization The authentication token
     * @returns ent_User The list of users
     * @throws ApiError
     */
    public static async getUsersService(
        authorization: string,
    ): Promise<Array<ent_User>> {
        const result = await __request({
            method: 'GET',
            path: `/api/users`,
            headers: {
                'Authorization': authorization,
            },
            errors: {
                400: `Authentication token is absent`,
                401: `Invalid authentication token`,
                500: `Something went wrong`,
            },
        });
        return result.body;
    }

}
