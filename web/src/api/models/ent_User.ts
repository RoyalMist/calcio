/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type {ent_UserEdges} from './ent_UserEdges';

export type ent_User = {
    /**
     * Admin holds the value of the "admin" field.
     */
    admin?: boolean;
    /**
     * Edges holds the relations/edges for other nodes in the graph.
     * The values are being populated by the UserQuery when eager-loading is set.
     */
    edges?: ent_UserEdges;
    /**
     * ID of the ent.
     */
    id?: string;
    /**
     * Name holds the value of the "name" field.
     */
    name?: string;
}
