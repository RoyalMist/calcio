/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type {ent_User} from '../models/ent_User';
import {request as __request} from '../core/request';

export class PlayersService {

    /**
     * Fetch all Calcio's users.
     * Retrieves all Calcio's users as a json list.
     * @param authorization The authentication token
     * @returns ent_User The list of users
     * @throws ApiError
     */
    public static async getPlayersService(
        authorization: string,
    ): Promise<Array<ent_User>> {
        const result = await __request({
            method: 'GET',
            path: `/api/users`,
            headers: {
                'Authorization': authorization,
            },
            errors: {
                400: `When the token is absent`,
                401: `When the token is invalid`,
                500: `When something went wrong`,
            },
        });
        return result.body;
    }

}
