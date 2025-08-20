package csharp

const messageTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
				if ({{ accessor . }} != null)
				{
{{- if $r.Skip }}
					// Message validation is skipped
{{- else }}
					Validate{{ simpleName .Field.Type.Embed }}({{ accessor . }});
{{- end }}
				}
{{- if $r.Required }}
				else
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "field is required");
				}
{{- end -}}
`
