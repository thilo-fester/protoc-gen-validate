package csharp

const timestampConstTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.Lt }}
		private static readonly Timestamp {{ constantName . "Lt" }} = {{ tsLit $r.GetLt }};
{{- end -}}
{{- if $r.Lte }}
		private static readonly Timestamp {{ constantName . "Lte" }} = {{ tsLit $r.GetLte }};
{{- end -}}
{{- if $r.Gt }}
		private static readonly Timestamp {{ constantName . "Gt" }} = {{ tsLit $r.GetGt }};
{{- end -}}
{{- if $r.Gte }}
		private static readonly Timestamp {{ constantName . "Gte" }} = {{ tsLit $r.GetGte }};
{{- end -}}`

const timestampTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
				if ({{ accessor . }} != null)
				{
{{- if $r.Required }}
					// Timestamp is present and required
{{- end -}}
{{- if $r.Const }}
					var constTime = {{ tsLit $r.GetConst }};
					if (!{{ accessor . }}.Equals(constTime))
					{
						throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "timestamp must equal constant value");
					}
{{- end -}}
{{- if $r.Lt }}
					if ({{ accessor . }}.CompareTo({{ constantName . "Lt" }}) >= 0)
					{
						throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "timestamp must be less than {{ $r.GetLt }}");
					}
{{- end -}}
{{- if $r.Lte }}
					if ({{ accessor . }}.CompareTo({{ constantName . "Lte" }}) > 0)
					{
						throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "timestamp must be less than or equal to {{ $r.GetLte }}");
					}
{{- end -}}
{{- if $r.Gt }}
					if ({{ accessor . }}.CompareTo({{ constantName . "Gt" }}) <= 0)
					{
						throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "timestamp must be greater than {{ $r.GetGt }}");
					}
{{- end -}}
{{- if $r.Gte }}
					if ({{ accessor . }}.CompareTo({{ constantName . "Gte" }}) < 0)
					{
						throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "timestamp must be greater than or equal to {{ $r.GetGte }}");
					}
{{- end -}}
{{- if $r.LtNow }}
					if ({{ accessor . }}.CompareTo(Timestamp.FromDateTimeOffset(DateTimeOffset.UtcNow)) >= 0)
					{
						throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "timestamp must be less than now");
					}
{{- end -}}
{{- if $r.GtNow }}
					if ({{ accessor . }}.CompareTo(Timestamp.FromDateTimeOffset(DateTimeOffset.UtcNow)) <= 0)
					{
						throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "timestamp must be greater than now");
					}
{{- end -}}
{{- if $r.Within }}
					var withinDuration = {{ durLit $r.GetWithin }};
					var now = Timestamp.FromDateTimeOffset(DateTimeOffset.UtcNow);
					var diff = {{ accessor . }} - now;
					if (System.Math.Abs(diff.ToTimeSpan().TotalMilliseconds) > withinDuration.ToTimeSpan().TotalMilliseconds)
					{
						throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "timestamp must be within {{ $r.GetWithin }} of now");
					}
{{- end -}}
				}
{{- if $r.Required }}
				else
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "timestamp field is required");
				}
{{- end -}}
`
