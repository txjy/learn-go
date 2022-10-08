package test

import (
	"cloud_disk/core/etc/helper"
	"testing"
)

func TestRandCode(t *testing.T) {
	println(helper.RandCode())
}
