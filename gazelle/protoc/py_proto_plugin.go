package protoc

func init() {
	MustRegisterProtoPlugin("py_proto", &PyProtoPlugin{})
}

// PyProtoPlugin implements ProtoPlugin for the built-in protoc python plugin.
type PyProtoPlugin struct{}

// ShouldApply implements part of the ProtoPlugin interface.
func (p *PyProtoPlugin) ShouldApply(rel string, cfg *ProtoPackageConfig, lib ProtoLibrary) bool {
	for _, f := range lib.Files() {
		if f.HasMessages() || f.HasEnums() {
			return true
		}
	}
	return false
}

// GeneratedSrcs implements part of the ProtoPlugin interface.
func (p *PyProtoPlugin) GeneratedSrcs(rel string, cfg *ProtoPackageConfig, lib ProtoLibrary) []string {
	srcs := make([]string, 0)
	for _, f := range lib.Files() {
		if f.HasMessages() || f.HasEnums() {
			srcs = append(srcs, f.Name+"_pb2.py")
		}
	}
	return srcs
}
