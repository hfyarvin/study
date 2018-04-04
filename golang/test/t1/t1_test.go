//go test -v
package t1

import(
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAdd(t *testing.T) {
	Convey("将两数相加", t, func() {
		So(Add(1,2), ShouldEqual, 3)
	})
}

func TestSubtract(t *testing.T) {
    Convey("将两数相减", t, func() {
        So(Subtract(1, 2), ShouldEqual, -1)
    })
}

func TestMultiply(t *testing.T) {//每个单元测试需要以Test开头
    Convey("将两数相乘", t, func() {
        So(Multiply(3, 2), ShouldEqual, 6)
    })
}

func TestDivision(t *testing.T) {
    Convey("将两数相除", t, func() {//只有最外层convey需要传入t

        Convey("除以非 0 数", func() {
            num, err := Division(10, 2)
            So(err, ShouldBeNil)
            So(num, ShouldEqual, 5)
        })

        Convey("除以 0", func() {
            _, err := Division(10, 0)
            So(err, ShouldNotBeNil)
        })
    })
}