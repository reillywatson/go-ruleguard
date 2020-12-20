package dslgen

var Fluent = []byte("package fluent\n\n// Bundle is a rules file export manifest.\ntype Bundle struct {\n\t// Version is a bundle version.\n\t// It's preferred to use a semver format.\n\t// Examples: \"0.5.1\", \"1.0.0\".\n\tVersion string\n\n\t// TODO: what else do we need here?\n}\n\n// ImportRules imports all rules from the bundle and prefixes them with a specified string.\n// Only packages that have an exported Bundle variable can be imported.\nfunc ImportRules(prefix string, bundle Bundle) {}\n\n\n\n// Matcher is a main API group-level entry point.\n// It's used to define and configure the group rules.\n// It also represents a map of all rule-local variables.\ntype Matcher map[string]Var\n\n// Import loads given package path into a rule group imports table.\n//\n// That table is used during the rules compilation.\n//\n// The table has the following effect on the rules:\n//\t* For type expressions, it's used to resolve the\n//\t  full package paths of qualified types, like `foo.Bar`.\n//\t  If Import(`a/b/foo`) is called, `foo.Bar` will match\n//\t  `a/b/foo.Bar` type during the pattern execution.\nfunc (m Matcher) Import(pkgPath string) {}\n\n// Match specifies a set of patterns that match a rule being defined.\n// Pattern matching succeeds if at least 1 pattern matches.\n//\n// If none of the given patterns matched, rule execution stops.\nfunc (m Matcher) Match(pattern string, alternatives ...string) Matcher {\n\treturn m\n}\n\n// Where applies additional constraint to a match.\n// If a given cond is not satisfied, a match is rejected and\n// rule execution stops.\nfunc (m Matcher) Where(cond bool) Matcher {\n\treturn m\n}\n\n// Report prints a message if associated rule match is successful.\n//\n// A message is a string that can contain interpolated expressions.\n// For every matched variable it's possible to interpolate\n// their printed representation into the message text with $<name>.\n// An entire match can be addressed with $$.\nfunc (m Matcher) Report(message string) Matcher {\n\treturn m\n}\n\n// Suggest assigns a quickfix suggestion for the matched code.\nfunc (m Matcher) Suggest(suggestion string) Matcher {\n\treturn m\n}\n\n// At binds the reported node to a named submatch.\n// If no explicit location is given, the outermost node ($$) is used.\nfunc (m Matcher) At(v Var) Matcher {\n\treturn m\n}\n\n// File returns the current file context.\nfunc (m Matcher) File() File { return File{} }\n\n// Var is a pattern variable that describes a named submatch.\ntype Var struct {\n\t// Pure reports whether expr matched by var is side-effect-free.\n\tPure bool\n\n\t// Const reports whether expr matched by var is a constant value.\n\tConst bool\n\n\t// Value is a compile-time computable value of the expression.\n\tValue ExprValue\n\n\t// Addressable reports whether the corresponding expression is addressable.\n\t// See https://golang.org/ref/spec#Address_operators.\n\tAddressable bool\n\n\t// Type is a type of a matched expr.\n\t//\n\t// For function call expressions, a type is a function result type,\n\t// but for a function expression itself it's a *types.Signature.\n\t//\n\t// Suppose we have a `a.b()` expression:\n\t//\t`$x()` m[\"x\"].Type is `a.b` function type\n\t//\t`$x` m[\"x\"].Type is `a.b()` function call result type\n\tType ExprType\n\n\t// Text is a captured node text as in the source code.\n\tText MatchedText\n\n\t// Node is a captured AST node.\n\tNode MatchedNode\n}\n\n// MatchedNode represents an AST node associated with a named submatch.\ntype MatchedNode struct{}\n\n// Is reports whether a matched node AST type is compatible with the specified type.\n// A valid argument is a ast.Node implementing type name from the \"go/ast\" package.\n// Examples: \"BasicLit\", \"Expr\", \"Stmt\", \"Ident\", \"ParenExpr\".\n// See https://golang.org/pkg/go/ast/.\nfunc (MatchedNode) Is(typ string) bool { return boolResult }\n\n// ExprValue describes a compile-time computable value of a matched expr.\ntype ExprValue struct{}\n\n// Int returns compile-time computable int value of the expression.\n// If value can't be computed, condition will fail.\nfunc (ExprValue) Int() int { return intResult }\n\n// ExprType describes a type of a matcher expr.\ntype ExprType struct {\n\t// Size represents expression type size in bytes.\n\tSize int\n}\n\n// Underlying returns expression type underlying type.\n// See https://golang.org/pkg/go/types/#Type Underlying() method documentation.\n// Read https://golang.org/ref/spec#Types section to learn more about underlying types.\nfunc (ExprType) Underlying() ExprType { return underlyingType }\n\n// AssignableTo reports whether a type is assign-compatible with a given type.\n// See https://golang.org/pkg/go/types/#AssignableTo.\nfunc (ExprType) AssignableTo(typ string) bool { return boolResult }\n\n// ConvertibleTo reports whether a type is conversible to a given type.\n// See https://golang.org/pkg/go/types/#ConvertibleTo.\nfunc (ExprType) ConvertibleTo(typ string) bool { return boolResult }\n\n// Implements reports whether a type implements a given interface.\n// See https://golang.org/pkg/go/types/#Implements.\nfunc (ExprType) Implements(typ string) bool { return boolResult }\n\n// Is reports whether a type is identical to a given type.\nfunc (ExprType) Is(typ string) bool { return boolResult }\n\n// MatchedText represents a source text associated with a matched node.\ntype MatchedText string\n\n// Matches reports whether the text matches the given regexp pattern.\nfunc (MatchedText) Matches(pattern string) bool { return boolResult }\n\n// String represents an arbitrary string-typed data.\ntype String string\n\n// Matches reports whether a string matches the given regexp pattern.\nfunc (String) Matches(pattern string) bool { return boolResult }\n\n// File represents the current Go source file.\ntype File struct {\n\t// Name is a file base name.\n\tName String\n\n\t// PkgPath is a file package path.\n\t// Examples: \"io/ioutil\", \"strings\", \"github.com/quasilyte/go-ruleguard/dsl/fluent\".\n\tPkgPath String\n}\n\n// Imports reports whether the current file imports the given path.\nfunc (File) Imports(path string) bool { return boolResult }\n\n\n\nvar boolResult bool\nvar intResult int\n\nvar underlyingType ExprType\n\n")
