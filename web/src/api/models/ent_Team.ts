/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type {ent_TeamEdges} from './ent_TeamEdges';

export type ent_Team = {
    /**
     * Edges holds the relations/edges for other nodes in the graph.
     * The values are being populated by the TeamQuery when eager-loading is set.
     */
    edges?: ent_TeamEdges;
    /**
     * ID of the ent.
     */
    id?: string;
    /**
     * Name holds the value of the "name" field.
     */
    name?: string;
}
