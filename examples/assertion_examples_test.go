package examples

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestAssertions(t *testing.T) {
	Convey("All assertions should be accessible", t, func() {
		Convey("Equality assertions should be accessible", func() {
			thing1a := thing{a: "asdf"}
			thing1b := thing{a: "asdf"}
			thing2 := thing{a: "qwer"}

			So(1, ShouldEqual, 1)
			So(1, ShouldNotEqual, 2)
			So(1, ShouldAlmostEqual, 1.000000000000001)
			So(1, ShouldNotAlmostEqual, 2, 0.5)
			So(thing1a, ShouldResemble, thing1b)
			So(thing1a, ShouldNotResemble, thing2)
			So(&thing1a, ShouldPointTo, &thing1a)
			So(&thing1a, ShouldNotPointTo, &thing1b)
			So(nil, ShouldBeNil)
			So(1, ShouldNotBeNil)
			So(true, ShouldBeTrue)
			So(false, ShouldBeFalse)
			So(0, ShouldBeZeroValue)
		})

		Convey("Numeric comparison assertions should be accessible", func() {
			So(1, ShouldBeGreaterThan, 0)
			So(1, ShouldBeGreaterThanOrEqualTo, 1)
			So(1, ShouldBeLessThan, 2)
			So(1, ShouldBeLessThanOrEqualTo, 1)
			So(1, ShouldBeBetween, 0, 2)
			So(1, ShouldNotBeBetween, 2, 4)
		})

		Convey("Container assertions should be accessible", func() {
			So([]int{1, 2, 3}, ShouldContain, 2)
			So([]int{1, 2, 3}, ShouldNotContain, 4)
			So(1, ShouldBeIn, []int{1, 2, 3})
			So(4, ShouldNotBeIn, []int{1, 2, 3})
		})

		Convey("String assertions should be accessible", func() {
			So("asdf", ShouldStartWith, "a")
			So("asdf", ShouldNotStartWith, "z")
			So("asdf", ShouldEndWith, "df")
			So("asdf", ShouldNotEndWith, "as")
			So("", ShouldBeBlank)
			So("asdf", ShouldNotBeBlank)
			So("asdf", ShouldContainSubstring, "sd")
			So("asdf", ShouldNotContainSubstring, "af")
		})

		Convey("Panic recovery assertions should be accessible", func() {
			So(panics, ShouldPanic)
			So(func() {}, ShouldNotPanic)
			So(panics, ShouldPanicWith, "Goofy Gophers!")
			So(panics, ShouldNotPanicWith, "Guileless Gophers!")
		})

		Convey("Type-checking assertions should be accessible", func() {
			So(1, ShouldHaveSameTypeAs, 0)
			So(1, ShouldNotHaveSameTypeAs, "1")
		})

		Convey("Time assertions should be accessible", func() {
			january1, _ := time.Parse(timeLayout, "2013-01-01 00:00")
			january2, _ := time.Parse(timeLayout, "2013-01-02 00:00")
			january3, _ := time.Parse(timeLayout, "2013-01-03 00:00")
			january4, _ := time.Parse(timeLayout, "2013-01-04 00:00")
			january5, _ := time.Parse(timeLayout, "2013-01-05 00:00")
			oneDay, _ := time.ParseDuration("24h0m0s")

			So(january1, ShouldHappenBefore, january4)
			So(january1, ShouldHappenOnOrBefore, january1)
			So(january2, ShouldHappenAfter, january1)
			So(january2, ShouldHappenOnOrAfter, january2)
			So(january3, ShouldHappenBetween, january2, january5)
			So(january3, ShouldHappenOnOrBetween, january3, january5)
			So(january1, ShouldNotHappenOnOrBetween, january2, january5)
			So(january2, ShouldHappenWithin, oneDay, january3)
			So(january5, ShouldNotHappenWithin, oneDay, january1)
		})
	})
}

type thing struct {
	a string
}

func panics() {
	panic("Goofy Gophers!")
}

const timeLayout = "2006-01-02 15:04"
