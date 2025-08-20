package csharp

const oneOfConstTpl = `{{ range .Fields }}
			{{ renderConstants (context .) }}
		{{ end }}`

const oneOfTpl = `
			switch (message.{{ pascalCase .Name }}Case)
			{
				{{ range .Fields -}}
				case {{ qualifiedName .Message }}.{{ pascalCase .OneOf.Name }}OneofCase.{{ pascalCase .Name }}:
					{{ render (context .) }}
					break;
				{{ end -}}
				case {{ qualifiedName .Message }}.{{ pascalCase .Name }}OneofCase.None:
				{{- if required . }}
					throw new ValidationException("{{ .Message.FullyQualifiedName }}", "{{ .Name }}", "oneof field is required");
				{{- else }}
					// No field set in oneof
				{{- end }}
					break;
				default:
					throw new ValidationException("{{ .Message.FullyQualifiedName }}", "{{ .Name }}", "unknown oneof case");
			}
`
