package api

import (
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func ListInstances(cli bce.Client, args *ListDomainArgs) (*ListDomainResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getDomainUri())
	req.SetMethod(http.GET)

	// Optional arguments settings
	if args != nil {
		if len(args.Domain) != 0 {
			req.SetParam("domain", args.Domain)
		}
		if len(args.DomainPunycode) != 0 {
			req.SetParam("domainPunycode", args.DomainPunycode)
		}
		if len(args.DomainTld) != 0 {
			req.SetParam("domainTld", args.DomainTld)
		}
		if len(args.ResourceId) != 0 {
			req.SetParam("resourceId", args.ResourceId)
		}
		if len(args.TagKey) != 0 {
			req.SetParam("tagKey", args.TagKey)
		}
		if len(args.TagValue) != 0 {
			req.SetParam("tagValue", args.TagValue)
		}
		if len(args.DomainStatus) != 0 {
			req.SetParam("domainStatus", args.DomainStatus)
		}
		if args.PageNo != 0 {
			req.SetParam("pageNo", strconv.Itoa(args.PageNo))
		}
		if args.PageSize != 0 {
			req.SetParam("pageSize", strconv.Itoa(args.PageSize))
		}
	}
	if args == nil || args.PageSize == 0 {
		req.SetParam("pageNo", strconv.Itoa(1))
		req.SetParam("pageSize", strconv.Itoa(100))
	}

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &ListDomainResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}
