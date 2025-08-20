package csharp

const numConstTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.In }}
		private static readonly HashSet<{{ csharpTypeFor . }}> {{ constantName . "In" }} = new HashSet<{{ csharpTypeFor . }}>
		{
			{{- range $r.In -}}
			{{ . }},
			{{- end -}}
		};
{{- end -}}
{{- if $r.NotIn }}
		private static readonly HashSet<{{ csharpTypeFor . }}> {{ constantName . "NotIn" }} = new HashSet<{{ csharpTypeFor . }}>
		{
			{{- range $r.NotIn -}}
			{{ . }},
			{{- end -}}
		};
{{- end -}}`

const numTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.GetIgnoreEmpty }}
			if ({{ accessor . }} != 0)
			{
{{- end -}}
{{- if $r.Const }}
				if ({{ accessor . }} != {{ $r.GetConst }})
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "value must equal {{ $r.GetConst }}");
				}
{{- end -}}
{{- if $r.In }}
				if (!{{ constantName . "In" }}.Contains({{ accessor . }}))
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "value must be in list {{ $r.In }}");
				}
{{- end -}}
{{- if $r.NotIn }}
				if ({{ constantName . "NotIn" }}.Contains({{ accessor . }}))
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "value must not be in list {{ $r.NotIn }}");
				}
{{- end -}}
{{- if $r.Lt }}
				if ({{ accessor . }} >= {{ $r.GetLt }})
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "value must be less than {{ $r.GetLt }}");
				}
{{- end -}}
{{- if $r.Lte }}
				if ({{ accessor . }} > {{ $r.GetLte }})
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "value must be less than or equal to {{ $r.GetLte }}");
				}
{{- end -}}
{{- if $r.Gt }}
				if ({{ accessor . }} <= {{ $r.GetGt }})
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "value must be greater than {{ $r.GetGt }}");
				}
{{- end -}}
{{- if $r.Gte }}
				if ({{ accessor . }} < {{ $r.GetGte }})
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "value must be greater than or equal to {{ $r.GetGte }}");
				}
{{- end -}}
{{- if $r.GetIgnoreEmpty }}
			}
{{- end -}}
`
