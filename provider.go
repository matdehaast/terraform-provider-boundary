package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-boundary/internal/provider"
)

func NewProvider() *schema.Provider {
	return provider.New()
}
