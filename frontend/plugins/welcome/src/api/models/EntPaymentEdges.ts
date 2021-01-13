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
    EntBank,
    EntBankFromJSON,
    EntBankFromJSONTyped,
    EntBankToJSON,
    EntInsurance,
    EntInsuranceFromJSON,
    EntInsuranceFromJSONTyped,
    EntInsuranceToJSON,
    EntMember,
    EntMemberFromJSON,
    EntMemberFromJSONTyped,
    EntMemberToJSON,
    EntMoneyTransfer,
    EntMoneyTransferFromJSON,
    EntMoneyTransferFromJSONTyped,
    EntMoneyTransferToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPaymentEdges
 */
export interface EntPaymentEdges {
    /**
     * 
     * @type {EntBank}
     * @memberof EntPaymentEdges
     */
    bank?: EntBank;
    /**
     * 
     * @type {EntInsurance}
     * @memberof EntPaymentEdges
     */
    insurance?: EntInsurance;
    /**
     * 
     * @type {EntMember}
     * @memberof EntPaymentEdges
     */
    member?: EntMember;
    /**
     * 
     * @type {EntMoneyTransfer}
     * @memberof EntPaymentEdges
     */
    moneyTransfer?: EntMoneyTransfer;
}

export function EntPaymentEdgesFromJSON(json: any): EntPaymentEdges {
    return EntPaymentEdgesFromJSONTyped(json, false);
}

export function EntPaymentEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPaymentEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'bank': !exists(json, 'Bank') ? undefined : EntBankFromJSON(json['Bank']),
        'insurance': !exists(json, 'Insurance') ? undefined : EntInsuranceFromJSON(json['Insurance']),
        'member': !exists(json, 'Member') ? undefined : EntMemberFromJSON(json['Member']),
        'moneyTransfer': !exists(json, 'MoneyTransfer') ? undefined : EntMoneyTransferFromJSON(json['MoneyTransfer']),
    };
}

export function EntPaymentEdgesToJSON(value?: EntPaymentEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'bank': EntBankToJSON(value.bank),
        'insurance': EntInsuranceToJSON(value.insurance),
        'member': EntMemberToJSON(value.member),
        'moneyTransfer': EntMoneyTransferToJSON(value.moneyTransfer),
    };
}


