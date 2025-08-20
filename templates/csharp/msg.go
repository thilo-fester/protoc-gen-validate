package csharp

const msgTpl = `
{{ if not (ignored .) -}}
	/// <summary>
	/// Validates {{ simpleName . }} message
	/// </summary>
	/// <param name="message">The message to validate</param>
	/// <exception cref="ValidationException">Thrown when validation fails</exception>
	public static void Validate{{ simpleName . }}({{ qualifiedName . }} message)
	{
		{{- template "msgInner" . -}}
	}
{{- end -}}
`

const msgInnerTpl = `
	{{ if disabled . }}
		// Validation is disabled for {{ simpleName . }}
		return;
	{{- else -}}
		if (message == null)
		{
			throw new ValidationException("{{ simpleName . }}", "message", "cannot be null");
		}

		{{ range .NonOneOfFields }}
			{{ renderConstants (context .) }}
		{{ end }}
		{{ range .SyntheticOneOfFields }}
			{{ renderConstants (context .) }}
		{{ end }}
		{{ range .RealOneOfs }}
			{{ template "oneOfConst" . }}
		{{ end }}

		try
		{
		{{ range .NonOneOfFields -}}
			{{ render (context .) }}
		{{ end -}}
		{{ range .SyntheticOneOfFields }}
			if ({{ hasAccessor (context .) }})
			{
				{{ render (context .) }}
			}
		{{ end }}
		{{ range .RealOneOfs }}
			{{ template "oneOf" . }}
		{{- end -}}
		}
		catch (ValidationException)
		{
			throw;
		}
		catch (Exception ex)
		{
			throw new ValidationException("{{ simpleName . }}", "validation", "unexpected error during validation", ex);
		}
	{{- end }}
`
