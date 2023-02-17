package google_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceComputeBackendBucket_basic(t *testing.T) {
	t.Parallel()

	backendBucketName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	bucketName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:     func() { acctest.TestAccPreCheck(t) },
		Providers:    acctest.TestAccProviders,
		CheckDestroy: testAccCheckComputeBackendBucketDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceComputeBackendBucket_basic(backendBucketName, bucketName),
				Check:  CheckDataSourceStateMatchesResourceState("data.google_compute_backend_bucket.baz", "google_compute_backend_bucket.foobar"),
			},
		},
	})
}

func testAccDataSourceComputeBackendBucket_basic(backendBucketName, bucketName string) string {
	return fmt.Sprintf(`
resource "google_compute_backend_bucket" "foobar" {
  name        = "%s"
  description = "Contains beautiful images"
  bucket_name = google_storage_bucket.image_bucket.name
  enable_cdn  = true
}
resource "google_storage_bucket" "image_bucket" {
  name     = "%s"
  location = "EU"
}
data "google_compute_backend_bucket" "baz" {
  name = google_compute_backend_bucket.foobar.name
}
`, backendBucketName, bucketName)
}
