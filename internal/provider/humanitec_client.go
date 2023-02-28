package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/humanitec/humanitec-go-autogen"
)

const (
	app = "terraform-provider-humanitec"
)

func NewHumanitecClient(host, token, version string) (*humanitec.Client, diag.Diagnostics) {
	client, err := humanitec.NewClient(&humanitec.Config{
		Token:       token,
		URL:         host,
		InternalApp: fmt.Sprintf("%s/%s", app, version),
		RequestLogger: func(req *humanitec.RequestDetails) {
			tflog.Debug(req.Context, "api req", map[string]interface{}{"method": req.Method, "uri": req.URL.String(), "body": req.Body})
		},
		ResponseLogger: func(res *humanitec.ResponseDetails) {
			tflog.Debug(res.Context, "api res", map[string]interface{}{"status": res.StatusCode, "body": res.Body})
		},
	})

	var diags diag.Diagnostics
	if err != nil {
		diags.AddError("Unable to create Humanitec client", err.Error())
		return nil, diags
	}

	return client, diags
}