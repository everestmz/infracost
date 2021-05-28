package aws

import (
	"git.fuzzbuzz.io/fuzz"
	"github.com/infracost/infracost/internal/schema"
)

func FuzzNewECRRepository(f *fuzz.F) {
	rd := &schema.ResourceData{}
	ud := &schema.UsageData{}

	f.Struct("resourceData", &schema.ResourceData{}).Populate(rd)
	f.Struct("usageData", &schema.UsageData{}).Populate(ud)

	NewECRRepository(rd, ud)
}
