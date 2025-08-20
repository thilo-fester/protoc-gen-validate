package csharp

const durationConstTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.Lt }}
		private static readonly Duration {{ constantName . "Lt" }} = {{ durLit $r.GetLt }};
{{- end -}}
{{- if $r.Lte }}
		private static readonly Duration {{ constantName . "Lte" }} = {{ durLit $r.GetLte }};
{{- end -}}
{{- if $r.Gt }}
		private static readonly Duration {{ constantName . "Gt" }} = {{ durLit $r.GetGt }};
{{- end -}}
{{- if $r.Gte }}
		private static readonly Duration {{ constantName . "Gte" }} = {{ durLit $r.GetGte }};
{{- end -}}
{{- if $r.In }}
		private static readonly HashSet<Duration> {{ constantName . "In" }} = new HashSet<Duration>
		{
			{{- range $r.In -}}
			{{ durLit . }},
			{{- end -}}
		};
{{- end -}}
{{- if $r.NotIn }}
		private static readonly HashSet<Duration> {{ constantName . "NotIn" }} = new HashSet<Duration>
		{
			{{- range $r.NotIn -}}
			{{ durLit . }},
			{{- end -}}
		};
{{- end -}}`

const durationTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
				if ({{ accessor . }} != null)
				{
{{- if $r.Required }}
					// Duration is present and required
{{- end -}}
{{- if $r.Const }}
					var constDuration = {{ durLit $r.GetConst }};
					if (!{{ accessor . }}.Equals(constDuration))
					{
						throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "duration must equal constant value");
					}
{{- end -}}
{{- if $r.In }}
					if (!{{ constantName . "In" }}.Contains({{ accessor . }}))
					{
						throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "duration must be in list {{ $r.In }}");
					}
{{- end -}}
{{- if $r.NotIn }}
					if ({{ constantName . "NotIn" }}.Contains({{ accessor . }}))
					{
						throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "duration must not be in list {{ $r.NotIn }}");
					}
{{- end -}}
{{- if $r.Lt }}
					if ({{ accessor . }}.CompareTo({{ constantName . "Lt" }}) >= 0)
					{
						throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "duration must be less than {{ $r.GetLt }}");
					}
{{- end -}}
{{- if $r.Lte }}
					if ({{ accessor . }}.CompareTo({{ constantName . "Lte" }}) > 0)
					{
						throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "duration must be less than or equal to {{ $r.GetLte }}");
					}
{{- end -}}
{{- if $r.Gt }}
					if ({{ accessor . }}.CompareTo({{ constantName . "Gt" }}) <= 0)
					{
						throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "duration must be greater than {{ $r.GetGt }}");
					}
{{- end -}}
{{- if $r.Gte }}
					if ({{ accessor . }}.CompareTo({{ constantName . "Gte" }}) < 0)
					{
						throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "duration must be greater than or equal to {{ $r.GetGte }}");
					}
{{- end -}}
				}
{{- if $r.Required }}
				else
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "duration field is required");
				}
{{- end -}}
`
