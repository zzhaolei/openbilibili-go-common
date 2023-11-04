package archive

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	xsql "go-common/library/database/sql"

	"bou.ke/monkey"
	"github.com/smartystreets/goconvey/convey"
)

func TestStaff(t *testing.T) {
	convey.Convey("RawStaff", t, func(ctx convey.C) {
		c := context.Background()
		ctx.Convey("When everything gose positive", func(ctx convey.C) {
			guard := monkey.PatchInstanceMethod(reflect.TypeOf(d.db), "Query", func(_ *xsql.DB, _ context.Context, _ string, _ ...interface{}) (*xsql.Rows, error) {
				return nil, fmt.Errorf("db.Query error")
			})
			defer guard.Unpatch()
			_, err := d.RawStaffData(c, 10110195)
			ctx.Convey("Then err should be nil.bizs should not be nil.", func(ctx convey.C) {
				ctx.So(err, convey.ShouldNotBeNil)
			})
		})
	})
}
