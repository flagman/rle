package codec


import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"bytes"
)


func TestEnc(t *testing.T) {
	Convey("Decode of data length should agree with 'Batman'", t, func() {
		name := []byte("Bruce")
		data := []byte("Batman")
		run, _ := NewRun(String, name, data)
		So(run.DataLength(), ShouldEqual, len(data))
	})

	Convey("Decode of name length should agree with 'Bruce'", t, func() {
		name := []byte("Bruce")
		data := []byte("Batman")
		run, _ := NewRun(String, name, data)
		So(run.NameLength(), ShouldEqual, len(name))
	})

	Convey("Decode of name should produce 'Bruce'", t, func() {
		name := []byte("Bruce")
		data := []byte("Batman")
		run, _ := NewRun(String, name, data)
		So(bytes.Equal(run.Name(), name), ShouldBeTrue)
	})

	Convey("New run length and size should agree", t, func() {
		name := []byte("Bruce")
		data := []byte("Batman")
		size, err := RunSize(name, data)
		So(err, ShouldBeNil)

		run, err := NewRun(String, name, data)
		So(err, ShouldBeNil)
		So(size, ShouldEqual, run.Len())
	})

	Convey("Size of run based on parts", t, func() {
		size, err := RunSize([]byte("Bruce"), []byte("Batman"))
		So(err, ShouldBeNil)
		So(size, ShouldEqual, 1+1+5+4+6)
	})

	Convey("NewRun should error when data is nil", t, func() {
		run, err := NewRun(String, []byte("Bruce"), nil)
		So(err, ShouldNotBeNil)
		So(len(run), ShouldEqual, 0)
	})

	Convey("NewRun should error when name is nil", t, func() {
		run, err := NewRun(String, nil, []byte("Batman"))
		So(err, ShouldNotBeNil)
		So(len(run), ShouldEqual, 0)
	})
}


