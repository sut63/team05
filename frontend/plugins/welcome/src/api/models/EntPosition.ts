/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Playlist Vidoe
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntPositionEdges,
    EntPositionEdgesFromJSON,
    EntPositionEdgesFromJSONTyped,
    EntPositionEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPosition
 */
export interface EntPosition {
    /**
     * 
     * @type {EntPositionEdges}
     * @memberof EntPosition
     */
    edges?: EntPositionEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntPosition
     */
    id?: number;
    /**
     * PositionName holds the value of the "position_name" field.
     * @type {string}
     * @memberof EntPosition
     */
    positionName?: string;
}

export function EntPositionFromJSON(json: any): EntPosition {
    return EntPositionFromJSONTyped(json, false);
}

export function EntPositionFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPosition {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntPositionEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'positionName': !exists(json, 'position_name') ? undefined : json['position_name'],
    };
}

export function EntPositionToJSON(value?: EntPosition | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntPositionEdgesToJSON(value.edges),
        'id': value.id,
        'position_name': value.positionName,
    };
}

