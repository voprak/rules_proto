package golang

import (
	"path"
	"strings"

	"github.com/bazelbuild/bazel-gazelle/label"
	"github.com/stackb/rules_proto/pkg/protoc"
)

const (
	// ProtocGenGoName is the name the plugin is registered under.
	ProtocGenGoName = "protoc-gen-go"
)

func init() {
	protoc.Plugins().MustRegisterPlugin(ProtocGenGoName, &ProtocGenGoPlugin{})
}

// ProtocGenGoPlugin implements Plugin for the the gogo_* family of plugins.
type ProtocGenGoPlugin struct{}

// Label implements part of the Plugin interface.
func (p *ProtocGenGoPlugin) Label() label.Label {
	return label.New("build_stack_rules_proto", "protocolbuffers/protobuf-go", "protoc-gen-go")
}

func (p *ProtocGenGoPlugin) ShouldApply(rel string, cfg protoc.PackageConfig, lib protoc.ProtoLibrary) bool {
	for _, f := range lib.Files() {
		if f.HasMessages() || f.HasEnums() {
			return true
		}
	}
	return false
}

// Outputs implements part of the Plugin interface
func (p *ProtocGenGoPlugin) Outputs(rel string, cfg protoc.PackageConfig, lib protoc.ProtoLibrary) []string {
	srcs := make([]string, 0)
	for _, f := range lib.Files() {
		if !(f.HasMessages() || f.HasEnums()) {
			continue
		}
		base := f.Name
		pkg := f.Package()
		// see https://github.com/gogo/protobuf/blob/master/protoc-gen-gogo/generator/generator.go#L347
		if goPackage, _, ok := protoc.GoPackageOption(f.Options()); ok {
			base = path.Join(goPackage, base)
		} else if pkg.Name != "" {
			base = path.Join(strings.ReplaceAll(pkg.Name, ".", "/"), base)
		}
		srcs = append(srcs, base+".pb.go")
	}
	return srcs
}
