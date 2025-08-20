package csharp

const stringConstTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.In }}
		private static readonly HashSet<{{ csharpTypeFor . }}> {{ constantName . "In" }} = new HashSet<{{ csharpTypeFor . }}>
		{
			{{- range $r.In -}}
			{{ csharpStringEscape . }},
			{{- end -}}
		};
{{- end -}}
{{- if $r.NotIn }}
		private static readonly HashSet<{{ csharpTypeFor . }}> {{ constantName . "NotIn" }} = new HashSet<{{ csharpTypeFor . }}>
		{
			{{- range $r.NotIn -}}
			{{ csharpStringEscape . }},
			{{- end -}}
		};
{{- end -}}
{{- if $r.Pattern }}
		private static readonly Regex {{ constantName . "Pattern" }} = new Regex({{ csharpStringEscape $r.GetPattern }}, RegexOptions.Compiled);
{{- end -}}`

const stringTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.GetIgnoreEmpty }}
			if (!string.IsNullOrEmpty({{ accessor . }}))
			{
{{- end -}}
{{- if $r.Const }}
				if ({{ accessor . }} != {{ csharpStringEscape $r.GetConst }})
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
{{- if $r.Len }}
				if ({{ accessor . }}.Length != {{ $r.GetLen }})
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "value length must be {{ $r.GetLen }}");
				}
{{- end -}}
{{- if $r.MinLen }}
				if ({{ accessor . }}.Length < {{ $r.GetMinLen }})
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "value length must be at least {{ $r.GetMinLen }}");
				}
{{- end -}}
{{- if $r.MaxLen }}
				if ({{ accessor . }}.Length > {{ $r.GetMaxLen }})
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "value length must be at most {{ $r.GetMaxLen }}");
				}
{{- end -}}
{{- if $r.LenBytes }}
				if (System.Text.Encoding.UTF8.GetByteCount({{ accessor . }}) != {{ $r.GetLenBytes }})
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "value byte length must be {{ $r.GetLenBytes }}");
				}
{{- end -}}
{{- if $r.MinBytes }}
				if (System.Text.Encoding.UTF8.GetByteCount({{ accessor . }}) < {{ $r.GetMinBytes }})
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "value byte length must be at least {{ $r.GetMinBytes }}");
				}
{{- end -}}
{{- if $r.MaxBytes }}
				if (System.Text.Encoding.UTF8.GetByteCount({{ accessor . }}) > {{ $r.GetMaxBytes }})
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "value byte length must be at most {{ $r.GetMaxBytes }}");
				}
{{- end -}}
{{- if $r.Pattern }}
				if (!{{ constantName . "Pattern" }}.IsMatch({{ accessor . }}))
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "value does not match pattern {{ $r.GetPattern }}");
				}
{{- end -}}
{{- if $r.Prefix }}
				if (!{{ accessor . }}.StartsWith({{ csharpStringEscape $r.GetPrefix }}))
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "value does not have prefix {{ $r.GetPrefix }}");
				}
{{- end -}}
{{- if $r.Contains }}
				if (!{{ accessor . }}.Contains({{ csharpStringEscape $r.GetContains }}))
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "value does not contain required substring {{ $r.GetContains }}");
				}
{{- end -}}
{{- if $r.NotContains }}
				if ({{ accessor . }}.Contains({{ csharpStringEscape $r.GetNotContains }}))
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "value contains forbidden substring {{ $r.GetNotContains }}");
				}
{{- end -}}
{{- if $r.Suffix }}
				if (!{{ accessor . }}.EndsWith({{ csharpStringEscape $r.GetSuffix }}))
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "value does not have suffix {{ $r.GetSuffix }}");
				}
{{- end -}}
{{- if $r.GetEmail }}
				if (!IsValidEmail({{ accessor . }}))
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "value must be a valid email address");
				}
{{- end -}}
{{- if $r.GetAddress }}
				if (!IsValidAddress({{ accessor . }}))
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "value must be a valid address");
				}
{{- end -}}
{{- if $r.GetHostname }}
				if (!IsValidHostname({{ accessor . }}))
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "value must be a valid hostname");
				}
{{- end -}}
{{- if $r.GetIp }}
				if (!IsValidIP({{ accessor . }}))
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "value must be a valid IP address");
				}
{{- end -}}
{{- if $r.GetIpv4 }}
				if (!IsValidIPv4({{ accessor . }}))
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "value must be a valid IPv4 address");
				}
{{- end -}}
{{- if $r.GetIpv6 }}
				if (!IsValidIPv6({{ accessor . }}))
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "value must be a valid IPv6 address");
				}
{{- end -}}
{{- if $r.GetUri }}
				if (!IsValidUri({{ accessor . }}))
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "value must be a valid URI");
				}
{{- end -}}
{{- if $r.GetUriRef }}
				if (!IsValidUriRef({{ accessor . }}))
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "value must be a valid URI reference");
				}
{{- end -}}
{{- if $r.GetUuid }}
				if (!IsValidUuid({{ accessor . }}))
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "value must be a valid UUID");
				}
{{- end -}}
{{- if $r.GetIgnoreEmpty }}
			}
{{- end -}}
`
