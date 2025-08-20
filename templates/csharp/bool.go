package csharp

const boolTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.Const }}
				if ({{ accessor . }} != {{ $r.GetConst }})
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "value must equal {{ $r.GetConst }}");
				}
{{- end -}}
`
