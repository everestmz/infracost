package aws_test

import (
	"testing"

	fuzztest "git.fuzzbuzz.io/fuzz/testing"
	"github.com/infracost/infracost/internal/providers/terraform/aws"
	"github.com/infracost/infracost/internal/providers/terraform/tftest"
	"github.com/infracost/infracost/internal/schema"
)

func TestEcrRepositoryGoldenFile(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	tftest.GoldenFileResourceTests(t, "ecr_repository_test")
}

func TestFuzzNewECRRepository(t *testing.T) {
	f := fuzztest.NewChecker(t)

	fuzztest.Check(f, aws.FuzzNewECRRepository, fuzztest.Test{
		f.Struct("resourceData", &schema.ResourceData{}): &schema.ResourceData{
			Type:         "abc",
			ProviderName: "def",
		},
		f.Struct("usageData", &schema.UsageData{}): &schema.UsageData{
			Address: "dfghjk",
		},
	})

	fuzztest.Randomize(f, aws.FuzzNewECRRepository, 1000)
}
