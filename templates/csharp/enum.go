package csharp

const enumConstTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.In }}
		private static readonly HashSet<{{ csharpTypeFor . }}> {{ constantName . "In" }} = new HashSet<{{ csharpTypeFor . }}>
		{
			{{- range $r.In -}}
			({{ csharpTypeFor $ }}){{ . }},
			{{- end -}}
		};
{{- end -}}
{{- if $r.NotIn }}
		private static readonly HashSet<{{ csharpTypeFor . }}> {{ constantName . "NotIn" }} = new HashSet<{{ csharpTypeFor . }}>
		{
			{{- range $r.NotIn -}}
			({{ csharpTypeFor $ }}){{ . }},
			{{- end -}}
		};
{{- end -}}`

const enumTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.GetIgnoreEmpty }}
			if ({{ accessor . }} != 0)
			{
{{- end -}}
{{- if $r.Const }}
				if ({{ accessor . }} != ({{ csharpTypeFor . }}){{ $r.GetConst }})
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
{{- if $r.DefinedOnly }}
				if (!System.Enum.IsDefined(typeof({{ csharpTypeFor . }}), {{ accessor . }}))
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "value must be a defined enum value");
				}
{{- end -}}
{{- if $r.GetIgnoreEmpty }}
			}
{{- end -}}
`
