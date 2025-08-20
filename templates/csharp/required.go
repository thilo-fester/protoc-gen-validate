package csharp

const requiredTpl = `
				// Field {{ fieldName . }} is required
{{- if .Field.Type.ProtoType.IsNumeric }}
				// Note: Numeric fields in proto3 cannot distinguish between 0 and unset
{{- else if .Field.Type.ProtoType.IsString }}
				if (string.IsNullOrEmpty({{ accessor . }}))
				{
					throw new ValidationException("{{ .Field.FullyQualifiedName }}", "{{ fieldName . }}", "field is required");
				}
{{- else if .Field.Type.ProtoType.IsBytes }}
				if ({{ accessor . }} == null || {{ accessor . }}.Length == 0)
				{
					throw new ValidationException("{{ .Field.FullyQualifiedName }}", "{{ fieldName . }}", "field is required");
				}
{{- else if .Field.Type.IsEmbed }}
				if ({{ accessor . }} == null)
				{
					throw new ValidationException("{{ .Field.FullyQualifiedName }}", "{{ fieldName . }}", "field is required");
				}
{{- else if .Field.Type.IsRepeated }}
				if ({{ accessor . }}.Count == 0)
				{
					throw new ValidationException("{{ .Field.FullyQualifiedName }}", "{{ fieldName . }}", "field is required");
				}
{{- else if .Field.Type.IsMap }}
				if ({{ accessor . }}.Count == 0)
				{
					throw new ValidationException("{{ .Field.FullyQualifiedName }}", "{{ fieldName . }}", "field is required");
				}
{{- else }}
				// Required validation for {{ .Field.Type.ProtoType }} fields
{{- end }}
`
