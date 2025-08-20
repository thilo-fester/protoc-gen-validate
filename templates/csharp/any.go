package csharp

const anyConstTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.In }}
		private static readonly HashSet<string> {{ constantName . "In" }} = new HashSet<string>
		{
			{{- range $r.In -}}
			{{ csharpStringEscape . }},
			{{- end -}}
		};
{{- end -}}
{{- if $r.NotIn }}
		private static readonly HashSet<string> {{ constantName . "NotIn" }} = new HashSet<string>
		{
			{{- range $r.NotIn -}}
			{{ csharpStringEscape . }},
			{{- end -}}
		};
{{- end -}}`

const anyTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
				if ({{ accessor . }} != null)
				{
{{- if $r.Required }}
					// Any field is present and required
{{- end -}}
{{- if $r.In }}
					if (!{{ constantName . "In" }}.Contains({{ accessor . }}.TypeUrl))
					{
						throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "any type URL must be in list {{ $r.In }}");
					}
{{- end -}}
{{- if $r.NotIn }}
					if ({{ constantName . "NotIn" }}.Contains({{ accessor . }}.TypeUrl))
					{
						throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "any type URL must not be in list {{ $r.NotIn }}");
					}
{{- end -}}
				}
{{- if $r.Required }}
				else
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "any field is required");
				}
{{- end -}}
`
