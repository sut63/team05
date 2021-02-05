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
    EntPaybackEdges,
    EntPaybackEdgesFromJSON,
    EntPaybackEdgesFromJSONTyped,
    EntPaybackEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPayback
 */
export interface EntPayback {
    /**
     * 
     * @type {EntPaybackEdges}
     * @memberof EntPayback
     */
    edges?: EntPaybackEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntPayback
     */
    id?: number;
    /**
     * PaybackAccountiden holds the value of the "payback_accountiden" field.
     * @type {string}
     * @memberof EntPayback
     */
    paybackAccountiden?: string;
    /**
     * PaybackAccountname holds the value of the "payback_accountname" field.
     * @type {string}
     * @memberof EntPayback
     */
    paybackAccountname?: string;
    /**
     * PaybackAccountnumber holds the value of the "payback_accountnumber" field.
     * @type {string}
     * @memberof EntPayback
     */
    paybackAccountnumber?: string;
    /**
     * PaybackTransfertime holds the value of the "payback_transfertime" field.
     * @type {string}
     * @memberof EntPayback
     */
    paybackTransfertime?: string;
}

export function EntPaybackFromJSON(json: any): EntPayback {
    return EntPaybackFromJSONTyped(json, false);
}

export function EntPaybackFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPayback {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntPaybackEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'paybackAccountiden': !exists(json, 'payback_accountiden') ? undefined : json['payback_accountiden'],
        'paybackAccountname': !exists(json, 'payback_accountname') ? undefined : json['payback_accountname'],
        'paybackAccountnumber': !exists(json, 'payback_accountnumber') ? undefined : json['payback_accountnumber'],
        'paybackTransfertime': !exists(json, 'payback_transfertime') ? undefined : json['payback_transfertime'],
    };
}

export function EntPaybackToJSON(value?: EntPayback | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntPaybackEdgesToJSON(value.edges),
        'id': value.id,
        'payback_accountiden': value.paybackAccountiden,
        'payback_accountname': value.paybackAccountname,
        'payback_accountnumber': value.paybackAccountnumber,
        'payback_transfertime': value.paybackTransfertime,
    };
}


