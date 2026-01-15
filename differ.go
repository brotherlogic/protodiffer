package protodiffer

import (
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Differ struct{}

func (d *Differ) Diff(old, new proto.Message) (proto.Message, error) {
	if old.ProtoReflect().Type().Descriptor().Name() != new.ProtoReflect().Type().Descriptor().Name() {
		return nil, status.Errorf(codes.FailedPrecondition, "%v (old) should match %v (new)", old.ProtoReflect().Type().Descriptor().Name(), new.ProtoReflect().Type().Descriptor().Name())
	}

	// Clone m and then remove equal fields (cmp to old)
	m := proto.Clone(new)

	for field := 0; field < old.ProtoReflect().Type().Descriptor().Fields().Len(); field++ {
		descriptor := old.ProtoReflect().Type().Descriptor().Fields().Get(field)
		switch v := old.ProtoReflect().Type().Descriptor().Fields().Get(field).Kind(); v {
		case protoreflect.MessageKind:
			// We currently ignore nested fields
			continue
		case protoreflect.Int32Kind:
			if old.ProtoReflect().Get(descriptor).Int() == new.ProtoReflect().Get(descriptor).Int() {
				m.ProtoReflect().Set(descriptor, protoreflect.ValueOfInt32(0))
			}
		case protoreflect.Int64Kind:
			if old.ProtoReflect().Get(descriptor).Int() == new.ProtoReflect().Get(descriptor).Int() {
				m.ProtoReflect().Set(descriptor, protoreflect.ValueOfInt64(0))
			}
		case protoreflect.FloatKind:
			if old.ProtoReflect().Get(descriptor).Float() == new.ProtoReflect().Get(descriptor).Float() {
				m.ProtoReflect().Set(descriptor, protoreflect.ValueOfFloat32(0))
			}
		case protoreflect.DoubleKind:
			if old.ProtoReflect().Get(descriptor).Float() == new.ProtoReflect().Get(descriptor).Float() {
				m.ProtoReflect().Set(descriptor, protoreflect.ValueOfFloat64(0))
			}
		case protoreflect.StringKind:
			if old.ProtoReflect().Get(descriptor).String() == new.ProtoReflect().Get(descriptor).String() {
				m.ProtoReflect().Set(descriptor, protoreflect.ValueOfString(""))
			}
		default:
			log.Printf("Unable to process: %v", v)
		}
	}
	return m, nil
}
