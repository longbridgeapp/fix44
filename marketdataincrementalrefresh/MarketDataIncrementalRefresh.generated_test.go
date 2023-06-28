package marketdataincrementalrefresh

import (
	"bytes"
	"testing"

	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/quickfix"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_MarketDataIncrementalRefresh(t *testing.T) {
	Convey("NoMDEntriesRepeatingGroup", t, func() {
		Convey("Group data with side field", func() {
			rowData := bytes.NewBufferString("8=FIX.4.49=21835=X34=1749349=SenderCompID52=20230627-08:40:16.03956=TargetCompID262=staging-1687854966268=2279=0269=255=ETHUSD270=1874.9271=2273=08:40:16336=CONTINUOUS54=1279=1269=B271=31435.24557346336=CONTINUOUS10=086")

			msg := quickfix.NewMessage()
			err := quickfix.ParseMessage(msg, rowData)
			So(err, ShouldBeNil)

			group := NewNoMDEntriesRepeatingGroup()
			err = msg.ToMessage().Body.GetGroup(group)
			So(err, ShouldBeNil)
			So(group.Len(), ShouldEqual, 2)

			entry := group.Get(0)
			ok := entry.HasSide()
			side, err := entry.GetSide()
			So(ok, ShouldBeTrue)
			So(err, ShouldBeNil)
			So(side, ShouldEqual, enum.Side_BUY)
		})

		Convey("Group data without side field", func() {
			rowData := bytes.NewBufferString("8=FIX.4.49=21335=X34=1749349=SenderCompID52=20230627-08:40:16.03956=TargetCompID262=staging-1687854966268=2279=0269=255=ETHUSD270=1874.9271=2273=08:40:16336=CONTINUOUS279=1269=B271=31435.24557346336=CONTINUOUS10=086")

			msg := quickfix.NewMessage()
			err := quickfix.ParseMessage(msg, rowData)
			So(err, ShouldBeNil)

			group := NewNoMDEntriesRepeatingGroup()
			err = msg.ToMessage().Body.GetGroup(group)
			So(err, ShouldBeNil)
			So(group.Len(), ShouldEqual, 2)

			newEntry := group.Add()
			ok := newEntry.HasSide()
			So(ok, ShouldBeFalse)

			newEntry.SetSide(enum.Side_SELL)
			ok = newEntry.HasSide()
			side, err := newEntry.GetSide()
			So(ok, ShouldBeTrue)
			So(side, ShouldEqual, enum.Side_SELL)
		})
	})
}
