/*
 * Copyright 2017 Baidu, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
 * except in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the
 * License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions
 * and limitations under the License.
 */

// model.go - definitions of the request arguments and results data structure model

package api

import "time"

type ListDomainResult struct {
	TotalCount int `json:"totalCount"`
	Result     []struct {
		UserID             string    `json:"userId"`
		ResourceID         string    `json:"resourceId"`
		ContactID          string    `json:"contactId"`
		DomainName         string    `json:"domainName"`
		Domain             string    `json:"domain"`
		Provider           string    `json:"provider"`
		DNS                []string  `json:"dns"`
		RegisterTime       time.Time `json:"registerTime"`
		ExpireTime         time.Time `json:"expireTime"`
		AuditStartTime     time.Time `json:"auditStartTime"`
		AuditStatus        string    `json:"auditStatus"`
		RealAuditStatus    string    `json:"realAuditStatus"`
		ExpireStatus       string    `json:"expireStatus"`
		EmailVerifyStatus  string    `json:"emailVerifyStatus"`
		NamingVerifyStatus string    `json:"namingVerifyStatus"`
		PrivacyProtected   bool      `json:"privacyProtected"`
		BeianStatus        string    `json:"beianStatus"`
		DomainStatus       string    `json:"domainStatus"`
		PrivacySupported   bool      `json:"privacySupported"`
		PanelOpened        bool      `json:"panelOpened"`
		AuditEnabled       bool      `json:"auditEnabled"`
		RenewEnabled       bool      `json:"renewEnabled"`
		UserType           string    `json:"userType"`
		Email              string    `json:"email"`
	} `json:"result"`
}

type ListDomainArgs struct {
	Domain         string
	DomainPunycode string
	DomainTld      string
	ResourceId     string
	TagKey         string
	TagValue       string
	DomainStatus   string
	PageNo         int
	PageSize       int
}
