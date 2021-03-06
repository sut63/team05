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
    EntHospital,
    EntHospitalFromJSON,
    EntHospitalFromJSONTyped,
    EntHospitalToJSON,
    EntMember,
    EntMemberFromJSON,
    EntMemberFromJSONTyped,
    EntMemberToJSON,
    EntOfficer,
    EntOfficerFromJSON,
    EntOfficerFromJSONTyped,
    EntOfficerToJSON,
    EntPayment,
    EntPaymentFromJSON,
    EntPaymentFromJSONTyped,
    EntPaymentToJSON,
    EntProduct,
    EntProductFromJSON,
    EntProductFromJSONTyped,
    EntProductToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntInsuranceEdges
 */
export interface EntInsuranceEdges {
    /**
     * 
     * @type {EntHospital}
     * @memberof EntInsuranceEdges
     */
    hospital?: EntHospital;
    /**
     * InsurancePayment holds the value of the insurance_payment edge.
     * @type {Array<EntPayment>}
     * @memberof EntInsuranceEdges
     */
    insurancePayment?: Array<EntPayment>;
    /**
     * 
     * @type {EntMember}
     * @memberof EntInsuranceEdges
     */
    member?: EntMember;
    /**
     * 
     * @type {EntOfficer}
     * @memberof EntInsuranceEdges
     */
    officer?: EntOfficer;
    /**
     * 
     * @type {EntProduct}
     * @memberof EntInsuranceEdges
     */
    product?: EntProduct;
}

export function EntInsuranceEdgesFromJSON(json: any): EntInsuranceEdges {
    return EntInsuranceEdgesFromJSONTyped(json, false);
}

export function EntInsuranceEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntInsuranceEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'hospital': !exists(json, 'Hospital') ? undefined : EntHospitalFromJSON(json['Hospital']),
        'insurancePayment': !exists(json, 'InsurancePayment') ? undefined : ((json['InsurancePayment'] as Array<any>).map(EntPaymentFromJSON)),
        'member': !exists(json, 'Member') ? undefined : EntMemberFromJSON(json['Member']),
        'officer': !exists(json, 'Officer') ? undefined : EntOfficerFromJSON(json['Officer']),
        'product': !exists(json, 'Product') ? undefined : EntProductFromJSON(json['Product']),
    };
}

export function EntInsuranceEdgesToJSON(value?: EntInsuranceEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'hospital': EntHospitalToJSON(value.hospital),
        'insurancePayment': value.insurancePayment === undefined ? undefined : ((value.insurancePayment as Array<any>).map(EntPaymentToJSON)),
        'member': EntMemberToJSON(value.member),
        'officer': EntOfficerToJSON(value.officer),
        'product': EntProductToJSON(value.product),
    };
}


