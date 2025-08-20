package csharp

const repeatedConstTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.Items }}
	{{ render (context .Field $r.Items) }}
{{- end -}}`

const repeatedTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.GetIgnoreEmpty }}
			if ({{ accessor . }}.Count > 0)
			{
{{- end -}}
{{- if $r.MinItems }}
				if ({{ accessor . }}.Count < {{ $r.GetMinItems }})
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "repeated field must have at least {{ $r.GetMinItems }} items");
				}
{{- end -}}
{{- if $r.MaxItems }}
				if ({{ accessor . }}.Count > {{ $r.GetMaxItems }})
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "repeated field must have at most {{ $r.GetMaxItems }} items");
				}
{{- end -}}
{{- if $r.Unique }}
				if ({{ accessor . }}.Count != {{ accessor . }}.Distinct().Count())
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "repeated field must have unique values");
				}
{{- end -}}
{{- if $r.Items }}
				for (int i = 0; i < {{ accessor . }}.Count; i++)
				{
					try
					{
						var item = {{ accessor . }}[i];
						{{ render (context .Field $r.Items "item") }}
					}
					catch (ValidationException ex)
					{
						throw new ValidationException("{{ $f.FullyQualifiedName }}", $"{{ fieldName . }}[{i}]", ex.Reason, ex);
					}
				}
{{- end -}}
{{- if $r.GetIgnoreEmpty }}
			}
{{- end -}}
`
