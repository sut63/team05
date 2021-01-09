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
/**
 * 
 * @export
 * @interface ControllersProduct
 */
export interface ControllersProduct {
    /**
     * 
     * @type {number}
     * @memberof ControllersProduct
     */
    genderID?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersProduct
     */
    groupOfAge?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersProduct
     */
    officer?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersProduct
     */
    productName?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersProduct
     */
    productPaymentOfYear?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersProduct
     */
    productPrice?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersProduct
     */
    productTime?: number;
}

export function ControllersProductFromJSON(json: any): ControllersProduct {
    return ControllersProductFromJSONTyped(json, false);
}

export function ControllersProductFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersProduct {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'genderID': !exists(json, 'genderID') ? undefined : json['genderID'],
        'groupOfAge': !exists(json, 'groupOfAge') ? undefined : json['groupOfAge'],
        'officer': !exists(json, 'officer') ? undefined : json['officer'],
        'productName': !exists(json, 'productName') ? undefined : json['productName'],
        'productPaymentOfYear': !exists(json, 'productPaymentOfYear') ? undefined : json['productPaymentOfYear'],
        'productPrice': !exists(json, 'productPrice') ? undefined : json['productPrice'],
        'productTime': !exists(json, 'productTime') ? undefined : json['productTime'],
    };
}

export function ControllersProductToJSON(value?: ControllersProduct | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'genderID': value.genderID,
        'groupOfAge': value.groupOfAge,
        'officer': value.officer,
        'productName': value.productName,
        'productPaymentOfYear': value.productPaymentOfYear,
        'productPrice': value.productPrice,
        'productTime': value.productTime,
    };
}

