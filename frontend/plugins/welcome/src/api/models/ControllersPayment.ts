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
 * @interface ControllersPayment
 */
export interface ControllersPayment {
    /**
     * 
     * @type {string}
     * @memberof ControllersPayment
     */
    accountName?: string;
    /**
     * 
     * @type {string}
     * @memberof ControllersPayment
     */
    accountNumber?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersPayment
     */
    bankID?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersPayment
     */
    insuranceID?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersPayment
     */
    memberID?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersPayment
     */
    moneytransferID?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersPayment
     */
    phoneNumber?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersPayment
     */
    price?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersPayment
     */
    transferTime?: string;
}

export function ControllersPaymentFromJSON(json: any): ControllersPayment {
    return ControllersPaymentFromJSONTyped(json, false);
}

export function ControllersPaymentFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersPayment {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'accountName': !exists(json, 'accountName') ? undefined : json['accountName'],
        'accountNumber': !exists(json, 'accountNumber') ? undefined : json['accountNumber'],
        'bankID': !exists(json, 'bankID') ? undefined : json['bankID'],
        'insuranceID': !exists(json, 'insuranceID') ? undefined : json['insuranceID'],
        'memberID': !exists(json, 'memberID') ? undefined : json['memberID'],
        'moneytransferID': !exists(json, 'moneytransferID') ? undefined : json['moneytransferID'],
        'phoneNumber': !exists(json, 'phoneNumber') ? undefined : json['phoneNumber'],
        'price': !exists(json, 'price') ? undefined : json['price'],
        'transferTime': !exists(json, 'transferTime') ? undefined : json['transferTime'],
    };
}

export function ControllersPaymentToJSON(value?: ControllersPayment | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'accountName': value.accountName,
        'accountNumber': value.accountNumber,
        'bankID': value.bankID,
        'insuranceID': value.insuranceID,
        'memberID': value.memberID,
        'moneytransferID': value.moneytransferID,
        'phoneNumber': value.phoneNumber,
        'price': value.price,
        'transferTime': value.transferTime,
    };
}


