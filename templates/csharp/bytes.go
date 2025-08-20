package csharp

const bytesConstTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.In }}
		private static readonly HashSet<string> {{ constantName . "In" }} = new HashSet<string>
		{
			{{- range $r.In -}}
			System.Convert.ToBase64String({{ byteArrayLit . }}),
			{{- end -}}
		};
{{- end -}}
{{- if $r.NotIn }}
		private static readonly HashSet<string> {{ constantName . "NotIn" }} = new HashSet<string>
		{
			{{- range $r.NotIn -}}
			System.Convert.ToBase64String({{ byteArrayLit . }}),
			{{- end -}}
		};
{{- end -}}
{{- if $r.Pattern }}
		private static readonly Regex {{ constantName . "Pattern" }} = new Regex({{ csharpStringEscape $r.GetPattern }}, RegexOptions.Compiled);
{{- end -}}`

const bytesTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.GetIgnoreEmpty }}
			if ({{ accessor . }}.Length > 0)
			{
{{- end -}}
{{- if $r.Const }}
				if (!{{ accessor . }}.Equals(ByteString.CopyFrom({{ byteArrayLit $r.GetConst }})))
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "value must equal constant");
				}
{{- end -}}
{{- if $r.In }}
				var {{ fieldName . }}Base64 = System.Convert.ToBase64String({{ accessor . }}.ToByteArray());
				if (!{{ constantName . "In" }}.Contains({{ fieldName . }}Base64))
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "value must be in list {{ $r.In }}");
				}
{{- end -}}
{{- if $r.NotIn }}
				var {{ fieldName . }}Base64 = System.Convert.ToBase64String({{ accessor . }}.ToByteArray());
				if ({{ constantName . "NotIn" }}.Contains({{ fieldName . }}Base64))
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
{{- if $r.Pattern }}
				var {{ fieldName . }}String = System.Text.Encoding.UTF8.GetString({{ accessor . }}.ToByteArray());
				if (!{{ constantName . "Pattern" }}.IsMatch({{ fieldName . }}String))
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "value does not match pattern {{ $r.GetPattern }}");
				}
{{- end -}}
{{- if $r.Prefix }}
				var {{ fieldName . }}Bytes = {{ accessor . }}.ToByteArray();
				var prefixBytes = {{ byteArrayLit $r.GetPrefix }};
				if ({{ fieldName . }}Bytes.Length < prefixBytes.Length || !{{ fieldName . }}Bytes.Take(prefixBytes.Length).SequenceEqual(prefixBytes))
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "value does not have required prefix");
				}
{{- end -}}
{{- if $r.Contains }}
				var {{ fieldName . }}Bytes = {{ accessor . }}.ToByteArray();
				var containsBytes = {{ byteArrayLit $r.GetContains }};
				if (!ContainsByteSequence({{ fieldName . }}Bytes, containsBytes))
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "value does not contain required byte sequence");
				}
{{- end -}}
{{- if $r.Suffix }}
				var {{ fieldName . }}Bytes = {{ accessor . }}.ToByteArray();
				var suffixBytes = {{ byteArrayLit $r.GetSuffix }};
				if ({{ fieldName . }}Bytes.Length < suffixBytes.Length || !{{ fieldName . }}Bytes.Skip({{ fieldName . }}Bytes.Length - suffixBytes.Length).SequenceEqual(suffixBytes))
				{
					throw new ValidationException("{{ $f.FullyQualifiedName }}", "{{ fieldName . }}", "value does not have required suffix");
				}
{{- end -}}
{{- if $r.GetIgnoreEmpty }}
			}
{{- end -}}
`
