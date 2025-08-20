package csharp

const wrapperConstTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
// Wrapper constants would be defined here for {{ $f.Name }}
`

const wrapperTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
				if ({{ accessor . }} != null)
				{
					{{ render (unwrap .) }}
				}
{{- if $r.Required }}
				else
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "wrapper field is required");
				}
{{- end -}}
`
