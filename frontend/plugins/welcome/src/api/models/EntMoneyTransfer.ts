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
    EntMoneytransferEdges,
    EntMoneytransferEdgesFromJSON,
    EntMoneytransferEdgesFromJSONTyped,
    EntMoneytransferEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntMoneytransfer
 */
export interface EntMoneytransfer {
    /**
     * 
     * @type {EntMoneytransferEdges}
     * @memberof EntMoneytransfer
     */
    edges?: EntMoneytransferEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntMoneytransfer
     */
    id?: number;
    /**
     * MoneytransferType holds the value of the "moneytransfer_type" field.
     * @type {string}
     * @memberof EntMoneytransfer
     */
    moneytransferType?: string;
}

export function EntMoneytransferFromJSON(json: any): EntMoneytransfer {
    return EntMoneytransferFromJSONTyped(json, false);
}

export function EntMoneytransferFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntMoneytransfer {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntMoneytransferEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'moneytransferType': !exists(json, 'moneytransfer_type') ? undefined : json['moneytransfer_type'],
    };
}

export function EntMoneytransferToJSON(value?: EntMoneytransfer | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntMoneytransferEdgesToJSON(value.edges),
        'id': value.id,
        'moneytransfer_type': value.moneytransferType,
    };
}


