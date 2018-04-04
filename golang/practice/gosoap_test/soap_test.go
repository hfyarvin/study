package gosoap_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetAllData(t *testing.T) {
	Convey("TestGetAllData", t, func() {
		err := GetAllData()
		So(err, ShouldBeNil, nil)
	})
}
