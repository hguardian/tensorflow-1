// Code generated by protoc-gen-go.
// source: tensorflow/core/framework/graph.proto
// DO NOT EDIT!

package config

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Represents the graph of operations
type GraphDef struct {
	Node []*NodeDef `protobuf:"bytes,1,rep,name=node" json:"node,omitempty"`
	// Compatibility versions of the graph.  See core/public/version.h for version
	// history.  The GraphDef version is distinct from the TensorFlow version, and
	// each release of TensorFlow will support a range of GraphDef versions.
	Versions *VersionDef `protobuf:"bytes,4,opt,name=versions" json:"versions,omitempty"`
	// Deprecated single version field; use versions above instead.  Since all
	// GraphDef changes before "versions" was introduced were forward
	// compatible, this field is entirely ignored.
	Version int32 `protobuf:"varint,3,opt,name=version" json:"version,omitempty"`
	// EXPERIMENTAL. DO NOT USE OR DEPEND ON THIS YET.
	//
	// "library" provides user-defined functions.
	//
	// Naming:
	//   * library.function.name are in a flat namespace.
	//     NOTE: We may need to change it to be hierarchical to support
	//     different orgs. E.g.,
	//     { "/google/nn", { ... }},
	//     { "/google/vision", { ... }}
	//     { "/org_foo/module_bar", {...}}
	//     map<string, FunctionDefLib> named_lib;
	//   * If node[i].op is the name of one function in "library",
	//     node[i] is deemed as a function call. Otherwise, node[i].op
	//     must be a primitive operation supported by the runtime.
	//
	//
	// Function call semantics:
	//
	//   * The callee may start execution as soon as some of its inputs
	//     are ready. The caller may want to use Tuple() mechanism to
	//     ensure all inputs are ready in the same time.
	//
	//   * The consumer of return values may start executing as soon as
	//     the return values the consumer depends on are ready.  The
	//     consumer may want to use Tuple() mechanism to ensure the
	//     consumer does not start until all return values of the callee
	//     function are ready.
	Library *FunctionDefLibrary `protobuf:"bytes,2,opt,name=library" json:"library,omitempty"`
}

func (m *GraphDef) Reset()                    { *m = GraphDef{} }
func (m *GraphDef) String() string            { return proto.CompactTextString(m) }
func (*GraphDef) ProtoMessage()               {}
func (*GraphDef) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *GraphDef) GetNode() []*NodeDef {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *GraphDef) GetVersions() *VersionDef {
	if m != nil {
		return m.Versions
	}
	return nil
}

func (m *GraphDef) GetLibrary() *FunctionDefLibrary {
	if m != nil {
		return m.Library
	}
	return nil
}

type NodeDef struct {
	// The name given to this operator. Used for naming inputs,
	// logging, visualization, etc.  Unique within a single GraphDef.
	// Must match the regexp "[A-Za-z0-9.][A-Za-z0-9_./]*".
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The operation name.  There may be custom parameters in attrs.
	// Op names starting with an underscore are reserved for internal use.
	Op string `protobuf:"bytes,2,opt,name=op" json:"op,omitempty"`
	// Each input is "node:src_output" with "node" being a string name and
	// "src_output" indicating which output tensor to use from "node". If
	// "src_output" is 0 the ":0" suffix can be omitted.  Regular inputs
	// may optionally be followed by control inputs that have the format
	// "^node".
	Input []string `protobuf:"bytes,3,rep,name=input" json:"input,omitempty"`
	// A (possibly partial) specification for the device on which this
	// node should be placed.
	// The expected syntax for this string is as follows:
	//
	// DEVICE_SPEC ::= COLOCATED_NODE | PARTIAL_SPEC
	//
	// COLOCATED_NODE ::= "@" NODE_NAME  // See NodeDef.name above.
	// PARTIAL_SPEC ::= ("/" CONSTRAINT) *
	// CONSTRAINT ::= ("job:" JOB_NAME)
	//              | ("replica:" [1-9][0-9]*)
	//              | ("task:" [1-9][0-9]*)
	//              | ( ("gpu" | "cpu") ":" ([1-9][0-9]* | "*") )
	//
	// Valid values for this string include:
	// * "@other/node"                         (colocate with "other/node")
	// * "/job:worker/replica:0/task:1/gpu:3"  (full specification)
	// * "/job:worker/gpu:3"                   (partial specification)
	// * ""                                    (no specification)
	//
	// If the constraints do not resolve to a single device (or if this
	// field is empty or not present), the runtime will attempt to
	// choose a device automatically.
	Device string `protobuf:"bytes,4,opt,name=device" json:"device,omitempty"`
	// Operation-specific graph-construction-time configuration.
	// Note that this should include all attrs defined in the
	// corresponding OpDef, including those with a value matching
	// the default -- this allows the default to change and makes
	// NodeDefs easier to interpret on their own.  However, if
	// an attr with a default is not specified in this list, the
	// default will be used.
	// The "names" (keys) must match the regexp "[a-z][a-z0-9_]+" (and
	// one of the names from the corresponding OpDef's attr field).
	// The values must have a type matching the corresponding OpDef
	// attr's type field.
	// TODO(josh11b): Add some examples here showing best practices.
	Attr map[string]*AttrValue `protobuf:"bytes,5,rep,name=attr" json:"attr,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *NodeDef) Reset()                    { *m = NodeDef{} }
func (m *NodeDef) String() string            { return proto.CompactTextString(m) }
func (*NodeDef) ProtoMessage()               {}
func (*NodeDef) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *NodeDef) GetAttr() map[string]*AttrValue {
	if m != nil {
		return m.Attr
	}
	return nil
}

func init() {
	proto.RegisterType((*GraphDef)(nil), "tensorflow.GraphDef")
	proto.RegisterType((*NodeDef)(nil), "tensorflow.NodeDef")
}

func init() { proto.RegisterFile("tensorflow/core/framework/graph.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x91, 0x5f, 0x4b, 0x32, 0x41,
	0x14, 0xc6, 0x99, 0xfd, 0xa3, 0xee, 0x11, 0x5e, 0x5e, 0xa6, 0x92, 0x41, 0x2a, 0x16, 0x21, 0x5a,
	0x2a, 0x56, 0xb2, 0x1b, 0xe9, 0x2e, 0xb1, 0xba, 0x09, 0x91, 0xb9, 0xf0, 0x36, 0x56, 0x9d, 0xb5,
	0x45, 0x9d, 0x59, 0x66, 0x47, 0xc5, 0x4f, 0xd7, 0x37, 0xe9, 0x73, 0x74, 0x19, 0x33, 0xbb, 0xea,
	0x5e, 0x24, 0xdd, 0x9d, 0x99, 0xf3, 0x3b, 0x0f, 0xe7, 0x39, 0x0f, 0x5c, 0x29, 0xc6, 0x33, 0x21,
	0xe3, 0x85, 0xd8, 0xb4, 0x27, 0x42, 0xb2, 0x76, 0x2c, 0xa3, 0x25, 0xdb, 0x08, 0x39, 0x6f, 0xcf,
	0x64, 0x94, 0x7e, 0x84, 0xa9, 0x14, 0x4a, 0x60, 0x38, 0x60, 0xcd, 0x9b, 0xe3, 0x23, 0x91, 0x52,
	0xf2, 0x7d, 0x1d, 0x2d, 0x56, 0x2c, 0x9f, 0x6b, 0x06, 0xc7, 0xd9, 0x78, 0xc5, 0x27, 0x2a, 0x11,
	0xfc, 0x6f, 0x72, 0xcd, 0x64, 0x96, 0x08, 0x9e, 0xe5, 0x64, 0xeb, 0x13, 0x41, 0xed, 0x55, 0xef,
	0xd6, 0x67, 0x31, 0xbe, 0x06, 0x87, 0x8b, 0x29, 0x23, 0xc8, 0xb7, 0x83, 0x7a, 0xe7, 0x24, 0x3c,
	0xa8, 0x84, 0x03, 0x31, 0x65, 0x7d, 0x16, 0x53, 0x03, 0xe0, 0x0e, 0xd4, 0x76, 0x3a, 0xc4, 0xf1,
	0x51, 0x50, 0xef, 0x34, 0xca, 0xf0, 0x28, 0xef, 0x69, 0x7e, 0xcf, 0xe1, 0x73, 0xa8, 0x16, 0x35,
	0xb1, 0x7d, 0x14, 0xb8, 0x3d, 0x8b, 0x20, 0xba, 0xfb, 0xc2, 0x5d, 0xa8, 0x2e, 0x92, 0xb1, 0x8c,
	0xe4, 0x96, 0x58, 0x46, 0xf0, 0xb2, 0x2c, 0xf8, 0x52, 0xd8, 0xeb, 0xb3, 0xf8, 0x2d, 0xa7, 0xe8,
	0x0e, 0x6f, 0x7d, 0x21, 0xa8, 0x16, 0xdb, 0x61, 0x0c, 0x0e, 0x8f, 0x96, 0xda, 0x00, 0x0a, 0x3c,
	0x6a, 0x6a, 0xfc, 0x0f, 0x2c, 0x91, 0x1a, 0x51, 0x8f, 0x5a, 0x22, 0xc5, 0xa7, 0xe0, 0x26, 0x3c,
	0x5d, 0x29, 0x62, 0xfb, 0x76, 0xe0, 0xd1, 0xfc, 0x81, 0x1b, 0x50, 0x99, 0xb2, 0x75, 0x32, 0x61,
	0xc6, 0x8f, 0x47, 0x8b, 0x17, 0xbe, 0x07, 0x47, 0xe7, 0x40, 0x5c, 0x73, 0x92, 0x8b, 0x5f, 0x4e,
	0x12, 0x3e, 0x29, 0x25, 0x9f, 0xb9, 0x92, 0x5b, 0x6a, 0xd0, 0xe6, 0x00, 0xbc, 0xfd, 0x17, 0xfe,
	0x0f, 0xf6, 0x9c, 0x6d, 0x8b, 0x85, 0x74, 0x89, 0x6f, 0xc1, 0x35, 0xa1, 0x16, 0x3e, 0xcf, 0xca,
	0x92, 0x7a, 0x6e, 0xa4, 0x9b, 0x34, 0x67, 0x1e, 0xad, 0x2e, 0xea, 0xdd, 0x01, 0x11, 0x72, 0x56,
	0xc6, 0xf6, 0x69, 0xf6, 0xea, 0x26, 0xbb, 0xa1, 0x8e, 0x32, 0x1b, 0xa2, 0x6f, 0x84, 0xc6, 0x15,
	0x93, 0xeb, 0xc3, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x8d, 0x55, 0x7b, 0x8c, 0x02, 0x00,
	0x00,
}
