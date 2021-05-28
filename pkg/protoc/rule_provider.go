package protoc

import (
	"github.com/bazelbuild/bazel-gazelle/config"
	"github.com/bazelbuild/bazel-gazelle/label"
	"github.com/bazelbuild/bazel-gazelle/rule"
)

// RuleProvider implementations are capable of providing a rule and import list
// to the gazelle GenerateArgs response.
type RuleProvider interface {
	// Kind of rule e.g. 'proto_library'
	Kind() string
	// Name provides the name of the rule.
	Name() string
	// Rule provides the gazelle rule implementation.
	Rule() *rule.Rule
	// Imports provides the list of imported symbols for the rule.
	Imports() []string
	// Visibility provides the visibility list for the rule.
	Visibility() []string
	// Resolve performs deps resolution, similar to the gazelle Resolver interface.
	Resolve(c *config.Config, r *rule.Rule, importsRaw interface{}, from label.Label)
}

// FileVisitor is an optional interface for RuleProvider implementations.  It
// will get called back with the rule.File of the package being visited by
// gazelle (it may be nil if no build file already exists). This exists to allow
// RuleProvider implementations to modify the file directly (e.g. adding
// additional load statements).
type FileVisitor interface {
	VisitFile(*rule.File) *rule.File
}