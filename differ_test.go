package protodiffer

import (
	"testing"

	"google.golang.org/protobuf/proto"

	pb "github.com/brotherlogic/protodiffer/proto"
)

var table = []struct {
	old  proto.Message
	new  proto.Message
	diff proto.Message
}{
	{
		old:  &pb.TestProto{Number1: 123},
		new:  &pb.TestProto{Number1: 125},
		diff: &pb.TestProto{Number1: 125},
	},
	{
		old:  &pb.TestProto{Number1: 123},
		new:  &pb.TestProto{Number1: 123},
		diff: &pb.TestProto{},
	},
	{
		old:  &pb.TestProto{Number2: 123},
		new:  &pb.TestProto{Number2: 125},
		diff: &pb.TestProto{Number2: 125},
	},
	{
		old:  &pb.TestProto{Number2: 123},
		new:  &pb.TestProto{Number2: 123},
		diff: &pb.TestProto{},
	},
	{
		old:  &pb.TestProto{Number3: 123},
		new:  &pb.TestProto{Number3: 125},
		diff: &pb.TestProto{Number3: 125},
	},
	{
		old:  &pb.TestProto{Number3: 123},
		new:  &pb.TestProto{Number3: 123},
		diff: &pb.TestProto{},
	},
	{
		old:  &pb.TestProto{Number4: 123},
		new:  &pb.TestProto{Number4: 125},
		diff: &pb.TestProto{Number4: 125},
	},
	{
		old:  &pb.TestProto{Number4: 123},
		new:  &pb.TestProto{Number4: 123},
		diff: &pb.TestProto{},
	},
	{
		old:  &pb.TestProto{Word: "123"},
		new:  &pb.TestProto{Word: "125"},
		diff: &pb.TestProto{Word: "125"},
	},
	{
		old:  &pb.TestProto{Word: "123"},
		new:  &pb.TestProto{Word: "123"},
		diff: &pb.TestProto{},
	},
}

func TestDiff(t *testing.T) {
	for _, test := range table {
		d := &Differ{}
		diff, err := d.Diff(test.old, test.new)
		if err != nil {
			t.Errorf("Unable to diff %v -> %v: %v", test.old, test.new, err)
		}
		if !proto.Equal(diff, test.diff) {
			t.Errorf("%v (Got) != %v (Expected)", diff, test.diff)
		}
	}
}
