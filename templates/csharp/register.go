package csharp

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/envoyproxy/protoc-gen-validate/templates/shared"
)

func Register(tpl *template.Template, params pgs.Parameters) {
	fns := csharpFuncs{pgsgo.InitContext(params)}

	tpl.Funcs(map[string]interface{}{
		"accessor":                fns.accessor,
		"byteArrayLit":           fns.byteArrayLit,
		"camelCase":              fns.camelCase,
		"classNameFile":          classNameFile,
		"classNameMessage":       classNameMessage,
		"csharpNamespace":        csharpNamespace,
		"csharpStringEscape":     fns.csharpStringEscape,
		"csharpTypeFor":          fns.csharpTypeFor,
		"durLit":                 fns.durLit,
		"fieldName":              fns.fieldName,
		"hasAccessor":            fns.hasAccessor,
		"oneof":                  fns.oneofTypeName,
		"pascalCase":             fns.pascalCase,
		"simpleName":             fns.Name,
		"sprintf":                fmt.Sprintf,
		"tsLit":                  fns.tsLit,
		"qualifiedName":          fns.qualifiedName,
		"unwrap":                 fns.unwrap,
		"renderConstants":        fns.renderConstants(tpl),
		"constantName":           fns.constantName,
	})

	template.Must(tpl.Parse(fileTpl))
	template.Must(tpl.New("msg").Parse(msgTpl))
	template.Must(tpl.New("msgInner").Parse(msgInnerTpl))
	template.Must(tpl.New("utils").Parse(utilsTpl))
	template.Must(tpl.New("exception").Parse(exceptionTpl))

	template.Must(tpl.New("none").Parse(noneTpl))

	// Numeric types
	template.Must(tpl.New("float").Parse(numTpl))
	template.Must(tpl.New("floatConst").Parse(numConstTpl))
	template.Must(tpl.New("double").Parse(numTpl))
	template.Must(tpl.New("doubleConst").Parse(numConstTpl))
	template.Must(tpl.New("int32").Parse(numTpl))
	template.Must(tpl.New("int32Const").Parse(numConstTpl))
	template.Must(tpl.New("int64").Parse(numTpl))
	template.Must(tpl.New("int64Const").Parse(numConstTpl))
	template.Must(tpl.New("uint32").Parse(numTpl))
	template.Must(tpl.New("uint32Const").Parse(numConstTpl))
	template.Must(tpl.New("uint64").Parse(numTpl))
	template.Must(tpl.New("uint64Const").Parse(numConstTpl))
	template.Must(tpl.New("sint32").Parse(numTpl))
	template.Must(tpl.New("sint32Const").Parse(numConstTpl))
	template.Must(tpl.New("sint64").Parse(numTpl))
	template.Must(tpl.New("sint64Const").Parse(numConstTpl))
	template.Must(tpl.New("fixed32").Parse(numTpl))
	template.Must(tpl.New("fixed32Const").Parse(numConstTpl))
	template.Must(tpl.New("fixed64").Parse(numTpl))
	template.Must(tpl.New("fixed64Const").Parse(numConstTpl))
	template.Must(tpl.New("sfixed32").Parse(numTpl))
	template.Must(tpl.New("sfixed32Const").Parse(numConstTpl))
	template.Must(tpl.New("sfixed64").Parse(numTpl))
	template.Must(tpl.New("sfixed64Const").Parse(numConstTpl))

	// Other basic types
	template.Must(tpl.New("bool").Parse(boolTpl))
	template.Must(tpl.New("string").Parse(stringTpl))
	template.Must(tpl.New("stringConst").Parse(stringConstTpl))
	template.Must(tpl.New("bytes").Parse(bytesTpl))
	template.Must(tpl.New("bytesConst").Parse(bytesConstTpl))

	// Complex types
	template.Must(tpl.New("any").Parse(anyTpl))
	template.Must(tpl.New("anyConst").Parse(anyConstTpl))
	template.Must(tpl.New("enum").Parse(enumTpl))
	template.Must(tpl.New("enumConst").Parse(enumConstTpl))
	template.Must(tpl.New("message").Parse(messageTpl))
	template.Must(tpl.New("repeated").Parse(repeatedTpl))
	template.Must(tpl.New("repeatedConst").Parse(repeatedConstTpl))
	template.Must(tpl.New("map").Parse(mapTpl))
	template.Must(tpl.New("mapConst").Parse(mapConstTpl))
	template.Must(tpl.New("oneOf").Parse(oneOfTpl))
	template.Must(tpl.New("oneOfConst").Parse(oneOfConstTpl))

	// Well-known types
	template.Must(tpl.New("required").Parse(requiredTpl))
	template.Must(tpl.New("timestamp").Parse(timestampTpl))
	template.Must(tpl.New("timestampConst").Parse(timestampConstTpl))
	template.Must(tpl.New("duration").Parse(durationTpl))
	template.Must(tpl.New("durationConst").Parse(durationConstTpl))
	template.Must(tpl.New("wrapper").Parse(wrapperTpl))
	template.Must(tpl.New("wrapperConst").Parse(wrapperConstTpl))
}

type csharpFuncs struct{ pgsgo.Context }

func CsharpFilePath(f pgs.File, ctx pgsgo.Context, tpl *template.Template) *pgs.FilePath {
	// Don't generate validators for files that don't import PGV
	if !importsPvg(f) {
		return nil
	}

	fullPath := strings.ReplaceAll(csharpNamespace(f), ".", string(os.PathSeparator))
	fileName := classNameFile(f) + "Validator.cs"
	filePath := pgs.JoinPaths(fullPath, fileName)
	return &filePath
}

func importsPvg(f pgs.File) bool {
	for _, dep := range f.Descriptor().Dependency {
		if strings.HasSuffix(dep, "validate.proto") {
			return true
		}
	}
	return false
}

func classNameFile(f pgs.File) string {
	protoName := pgs.FilePath(f.Name().String()).BaseName()
	return sanitizeClassName(protoName)
}

func classNameMessage(m pgs.Message) string {
	className := m.Name().String()
	return sanitizeClassName(className)
}

func sanitizeClassName(className string) string {
	// Convert to PascalCase for C# conventions
	return strcase.ToCamel(className)
}

func csharpNamespace(file pgs.File) string {
	// Try to get C# namespace from options, fallback to package name
	options := file.Descriptor().GetOptions()
	if options != nil {
		// Check for csharp_namespace option (we'd need to extend this)
		// For now, convert package to C# style namespace
	}
	
	// Convert protobuf package to C# namespace
	pkg := file.Package().ProtoName().String()
	if pkg == "" {
		return "Global"
	}
	
	// Convert to PascalCase segments
	parts := strings.Split(pkg, ".")
	for i, part := range parts {
		parts[i] = strcase.ToCamel(part)
	}
	
	return strings.Join(parts, ".")
}

func (fns csharpFuncs) qualifiedName(entity pgs.Entity) string {
	switch e := entity.(type) {
	case pgs.File:
		return csharpNamespace(e) + "." + classNameFile(e)
	case pgs.Message:
		if e.Parent() != nil {
			return fns.qualifiedName(e.Parent()) + "." + entity.Name().String()
		}
		return entity.Name().String()
	case pgs.Enum:
		if e.Parent() != nil {
			return fns.qualifiedName(e.Parent()) + "." + entity.Name().String()
		}
		return entity.Name().String()
	}
	return entity.Name().String()
}

func (fns csharpFuncs) accessor(ctx shared.RuleContext) string {
	if ctx.AccessorOverride != "" {
		return ctx.AccessorOverride
	}
	return fns.fieldAccessor(ctx.Field)
}

func (fns csharpFuncs) fieldAccessor(f pgs.Field) string {
	fieldName := strcase.ToCamel(f.Name().String())
	return fmt.Sprintf("proto.%s", fieldName)
}

func (fns csharpFuncs) hasAccessor(ctx shared.RuleContext) string {
	if ctx.AccessorOverride != "" {
		return "true"
	}
	fieldName := strcase.ToCamel(ctx.Field.Name().String())
	return fmt.Sprintf("proto.Has%s", fieldName)
}

func (fns csharpFuncs) fieldName(ctx shared.RuleContext) string {
	return ctx.Field.Name().String()
}

func (fns csharpFuncs) csharpTypeFor(ctx shared.RuleContext) string {
	t := ctx.Field.Type()

	// Map key and value types
	if t.IsMap() {
		switch ctx.AccessorOverride {
		case "key":
			return fns.csharpTypeForProtoType(t.Key().ProtoType())
		case "value":
			return fns.csharpTypeForProtoType(t.Element().ProtoType())
		}
	}

	if t.IsEmbed() {
		if embed := t.Embed(); embed.IsWellKnown() {
			switch embed.WellKnownType() {
			case pgs.AnyWKT:
				return "string"
			case pgs.DurationWKT:
				return "Google.Protobuf.WellKnownTypes.Duration"
			case pgs.TimestampWKT:
				return "Google.Protobuf.WellKnownTypes.Timestamp"
			case pgs.Int32ValueWKT, pgs.UInt32ValueWKT:
				return "int"
			case pgs.Int64ValueWKT, pgs.UInt64ValueWKT:
				return "long"
			case pgs.DoubleValueWKT:
				return "double"
			case pgs.FloatValueWKT:
				return "float"
			}
		}
	}

	if t.IsRepeated() {
		if t.ProtoType() == pgs.MessageT {
			return fns.qualifiedName(t.Element().Embed())
		} else if t.ProtoType() == pgs.EnumT {
			return fns.qualifiedName(t.Element().Enum())
		}
	}

	if t.IsEnum() {
		return fns.qualifiedName(t.Enum())
	}

	return fns.csharpTypeForProtoType(t.ProtoType())
}

func (fns csharpFuncs) csharpTypeForProtoType(t pgs.ProtoType) string {
	switch t {
	case pgs.Int32T, pgs.SInt32, pgs.SFixed32:
		return "int"
	case pgs.Int64T, pgs.SInt64, pgs.SFixed64:
		return "long"
	case pgs.UInt32T, pgs.Fixed32T:
		return "uint"
	case pgs.UInt64T, pgs.Fixed64T:
		return "ulong"
	case pgs.DoubleT:
		return "double"
	case pgs.FloatT:
		return "float"
	case pgs.BoolT:
		return "bool"
	case pgs.StringT:
		return "string"
	case pgs.BytesT:
		return "Google.Protobuf.ByteString"
	default:
		return "object"
	}
}

func (fns csharpFuncs) csharpStringEscape(s string) string {
	s = fmt.Sprintf("%q", s)
	s = s[1 : len(s)-1]
	s = strings.ReplaceAll(s, `\"`, `"`)
	return `"` + s + `"`
}

func (fns csharpFuncs) camelCase(name pgs.Name) string {
	return strcase.ToLowerCamel(name.String())
}

func (fns csharpFuncs) pascalCase(name pgs.Name) string {
	return strcase.ToCamel(name.String())
}

func (fns csharpFuncs) byteArrayLit(bytes []uint8) string {
	var sb strings.Builder
	sb.WriteString("new byte[] { ")
	for i, b := range bytes {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("0x%02X", b))
	}
	sb.WriteString(" }")
	return sb.String()
}

func (fns csharpFuncs) durLit(dur *durationpb.Duration) string {
	return fmt.Sprintf(
		"Google.Protobuf.WellKnownTypes.Duration.FromTimeSpan(System.TimeSpan.FromSeconds(%d).Add(System.TimeSpan.FromTicks(%d * 100)))",
		dur.GetSeconds(), dur.GetNanos()/100)
}

func (fns csharpFuncs) tsLit(ts *timestamppb.Timestamp) string {
	return fmt.Sprintf(
		"Google.Protobuf.WellKnownTypes.Timestamp.FromDateTimeOffset(System.DateTimeOffset.FromUnixTimeSeconds(%d).AddTicks(%d * 100))",
		ts.GetSeconds(), ts.GetNanos()/100)
}

func (fns csharpFuncs) oneofTypeName(f pgs.Field) pgsgo.TypeName {
	return pgsgo.TypeName(strcase.ToCamel(f.Name().String()))
}

func (fns csharpFuncs) unwrap(ctx shared.RuleContext) (shared.RuleContext, error) {
	ctx, err := ctx.Unwrap("wrapped")
	if err != nil {
		return ctx, err
	}
	ctx.AccessorOverride = fmt.Sprintf("%s.%s", fns.fieldAccessor(ctx.Field),
		fns.pascalCase(ctx.Field.Type().Embed().Fields()[0].Name()))
	return ctx, nil
}

func (fns csharpFuncs) renderConstants(tpl *template.Template) func(ctx shared.RuleContext) (string, error) {
	return func(ctx shared.RuleContext) (string, error) {
		var b bytes.Buffer
		var err error

		hasConstTemplate := false
		for _, t := range tpl.Templates() {
			if t.Name() == ctx.Typ+"Const" {
				hasConstTemplate = true
			}
		}

		if hasConstTemplate {
			err = tpl.ExecuteTemplate(&b, ctx.Typ+"Const", ctx)
		}

		return b.String(), err
	}
}

func (fns csharpFuncs) constantName(ctx shared.RuleContext, rule string) string {
	return strcase.ToCamel(ctx.Field.Name().String() + "_" + ctx.Index + "_" + rule)
}
