package metric

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSubscribe(t *testing.T) {
	m := &Metric{
		sub: new(subscriptions),
	}
	m.Subscribe()
	Convey("adds a subscription", t, func() {
		So(m.Subscriptions(), ShouldEqual, 1)
	})
}

func TestUnsubscribe(t *testing.T) {
	m := &Metric{
		sub: new(subscriptions),
	}
	Convey("when there are subscriptions to unsubscribe", t, func() {
		m.Subscribe()
		Convey("It removes a subscription", func() {
			m.Unsubscribe()
			So(m.Subscriptions(), ShouldEqual, 0)
		})
	})
	Convey("when there are not subscriptions to unsubscribe", t, func() {
		Convey("It returns an error", func() {
			err := m.Unsubscribe()
			So(err, ShouldNotBeNil)
		})
	})
}

func TestSubscriptions(t *testing.T) {
	m := &Metric{
		sub: new(subscriptions),
	}
	Convey("should accurately represent the subscription count", t, func() {
		for i := 0; i < 6; i++ {
			m.Subscribe()
		}
		for i := 0; i < 3; i++ {
			m.Unsubscribe()
		}
		So(m.Subscriptions(), ShouldEqual, 3)
	})
}
