package csharp

const mapConstTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.Keys }}
	{{ render (context .Field $r.Keys) }}
{{- end -}}
{{- if $r.Values }}
	{{ render (context .Field $r.Values) }}
{{- end -}}`

const mapTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.GetIgnoreEmpty }}
			if ({{ accessor . }}.Count > 0)
			{
{{- end -}}
{{- if $r.MinPairs }}
				if ({{ accessor . }}.Count < {{ $r.GetMinPairs }})
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "map field must have at least {{ $r.GetMinPairs }} pairs");
				}
{{- end -}}
{{- if $r.MaxPairs }}
				if ({{ accessor . }}.Count > {{ $r.GetMaxPairs }})
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "map field must have at most {{ $r.GetMaxPairs }} pairs");
				}
{{- end -}}
{{- if $r.NoSparse }}
				// Note: C# protobuf maps don't have sparse values by default
{{- end -}}
{{- if or $r.Keys $r.Values }}
				foreach (var kvp in {{ accessor . }})
				{
					try
					{
{{- if $r.Keys }}
						var key = kvp.Key;
						{{ render (context .Field $r.Keys "key") }}
{{- end }}
{{- if $r.Values }}
						var value = kvp.Value;
						{{ render (context .Field $r.Values "value") }}
{{- end }}
					}
					catch (ValidationException ex)
					{
						throw new ValidationException("{{ $f.FullyQualifiedName }}", $"{{ fieldName . }}[{kvp.Key}]", ex.Reason, ex);
					}
				}
{{- end -}}
{{- if $r.GetIgnoreEmpty }}
			}
{{- end -}}
`
