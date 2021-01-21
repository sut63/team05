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
} from './';

/**
 * 
 * @export
 * @interface EntCategoryEdges
 */
export interface EntCategoryEdges {
    /**
     * CategoryInquiry holds the value of the category_inquiry edge.
     * @type {Array<EntInquiry>}
     * @memberof EntCategoryEdges
     */
    categoryInquiry?: Array<EntInquiry>;
}

export function EntCategoryEdgesFromJSON(json: any): EntCategoryEdges {
    return EntCategoryEdgesFromJSONTyped(json, false);
}

export function EntCategoryEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntCategoryEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'categoryInquiry': !exists(json, 'categoryInquiry') ? undefined : ((json['categoryInquiry'] as Array<any>).map(EntInquiryFromJSON)),
    };
}

export function EntCategoryEdgesToJSON(value?: EntCategoryEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'categoryInquiry': value.categoryInquiry === undefined ? undefined : ((value.categoryInquiry as Array<any>).map(EntInquiryToJSON)),
    };
}


