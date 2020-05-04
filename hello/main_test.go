
package main

import (
	"testing"
)

func TestGCE(t *testing.T) {
	i := newInstance()
	//if !metadata.OnGCE() && i.Error != "Not running on GCE" {
		t.log("Unit tests completed, but error does not indicate that fact.")
	//}
}
