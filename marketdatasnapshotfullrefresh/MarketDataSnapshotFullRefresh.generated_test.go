package marketdatasnapshotfullrefresh

import (
	"bytes"
	"testing"

	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/quickfix"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_MarketDataSnapshotFullRefresh(t *testing.T) {
	Convey("NoMDEntriesRepeatingGroup", t, func() {
		Convey("Group data with side field", func() {
			rowData := bytes.NewBufferString("8=FIX.4.49=20835=W34=53949=SenderCompID52=20230628-08:29:47.30056=TargetCompID55=ETHUSD262=staging-1687940826268=2269=2270=1857.1271=2.019273=08:29:47336=CONTINUOUS54=1269=B271=31767.82357346336=CONTINUOUS10=165")

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
			rowData := bytes.NewBufferString("8=FIX.4.49=20335=W34=53949=SenderCompID52=20230628-08:29:47.30056=TargetCompID55=ETHUSD262=staging-1687940826268=2269=2270=1857.1271=2.019273=08:29:47336=CONTINUOUS269=B271=31767.82357346336=CONTINUOUS10=165")

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
