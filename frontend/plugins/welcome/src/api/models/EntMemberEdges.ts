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
    EntInquiry,
    EntInquiryFromJSON,
    EntInquiryFromJSONTyped,
    EntInquiryToJSON,
    EntInsurance,
    EntInsuranceFromJSON,
    EntInsuranceFromJSONTyped,
    EntInsuranceToJSON,
    EntPayback,
    EntPaybackFromJSON,
    EntPaybackFromJSONTyped,
    EntPaybackToJSON,
    EntPayment,
    EntPaymentFromJSON,
    EntPaymentFromJSONTyped,
    EntPaymentToJSON,
    EntRecordinsurance,
    EntRecordinsuranceFromJSON,
    EntRecordinsuranceFromJSONTyped,
    EntRecordinsuranceToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntMemberEdges
 */
export interface EntMemberEdges {
    /**
     * MemberInquiry holds the value of the member_inquiry edge.
     * @type {Array<EntInquiry>}
     * @memberof EntMemberEdges
     */
    memberInquiry?: Array<EntInquiry>;
    /**
     * MemberInsurance holds the value of the member_insurance edge.
     * @type {Array<EntInsurance>}
     * @memberof EntMemberEdges
     */
    memberInsurance?: Array<EntInsurance>;
    /**
     * MemberPayback holds the value of the member_payback edge.
     * @type {Array<EntPayback>}
     * @memberof EntMemberEdges
     */
    memberPayback?: Array<EntPayback>;
    /**
     * MemberPayment holds the value of the member_payment edge.
     * @type {Array<EntPayment>}
     * @memberof EntMemberEdges
     */
    memberPayment?: Array<EntPayment>;
    /**
     * MemberRecordinsurance holds the value of the member_recordinsurance edge.
     * @type {Array<EntRecordinsurance>}
     * @memberof EntMemberEdges
     */
    memberRecordinsurance?: Array<EntRecordinsurance>;
}

export function EntMemberEdgesFromJSON(json: any): EntMemberEdges {
    return EntMemberEdgesFromJSONTyped(json, false);
}

export function EntMemberEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntMemberEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'memberInquiry': !exists(json, 'memberInquiry') ? undefined : ((json['memberInquiry'] as Array<any>).map(EntInquiryFromJSON)),
        'memberInsurance': !exists(json, 'memberInsurance') ? undefined : ((json['memberInsurance'] as Array<any>).map(EntInsuranceFromJSON)),
        'memberPayback': !exists(json, 'memberPayback') ? undefined : ((json['memberPayback'] as Array<any>).map(EntPaybackFromJSON)),
        'memberPayment': !exists(json, 'memberPayment') ? undefined : ((json['memberPayment'] as Array<any>).map(EntPaymentFromJSON)),
        'memberRecordinsurance': !exists(json, 'memberRecordinsurance') ? undefined : ((json['memberRecordinsurance'] as Array<any>).map(EntRecordinsuranceFromJSON)),
    };
}

export function EntMemberEdgesToJSON(value?: EntMemberEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'memberInquiry': value.memberInquiry === undefined ? undefined : ((value.memberInquiry as Array<any>).map(EntInquiryToJSON)),
        'memberInsurance': value.memberInsurance === undefined ? undefined : ((value.memberInsurance as Array<any>).map(EntInsuranceToJSON)),
        'memberPayback': value.memberPayback === undefined ? undefined : ((value.memberPayback as Array<any>).map(EntPaybackToJSON)),
        'memberPayment': value.memberPayment === undefined ? undefined : ((value.memberPayment as Array<any>).map(EntPaymentToJSON)),
        'memberRecordinsurance': value.memberRecordinsurance === undefined ? undefined : ((value.memberRecordinsurance as Array<any>).map(EntRecordinsuranceToJSON)),
    };
}


