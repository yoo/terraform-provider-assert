// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"net"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = IPv4Function{}
)

func NewIPv4Function() function.Function {
	return IPv4Function{}
}

type IPv4Function struct{}

func (r IPv4Function) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "ipv4"
}

func (r IPv4Function) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Checks whether a string is a valid IPv4 address",
		Parameters: []function.Parameter{
			function.StringParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The string to check",
				Name:               "ip_address",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r IPv4Function) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var ip string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &ip))
	if resp.Error != nil {
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, isIPv4(ip)))
}

func isIPv4(ip string) bool {
	return net.ParseIP(ip).To4() != nil
}
