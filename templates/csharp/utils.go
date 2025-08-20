package csharp

const utilsTpl = `
		#region Validation Helper Methods

		private static readonly Regex EmailRegex = new Regex(@"^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$", RegexOptions.Compiled);
		private static readonly Regex HostnameRegex = new Regex(@"^[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(\.[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$", RegexOptions.Compiled);
		private static readonly Regex UuidRegex = new Regex(@"^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$", RegexOptions.Compiled);

		private static bool IsValidEmail(string email)
		{
			if (string.IsNullOrEmpty(email))
				return false;
			
			return EmailRegex.IsMatch(email);
		}

		private static bool IsValidHostname(string hostname)
		{
			if (string.IsNullOrEmpty(hostname) || hostname.Length > 253)
				return false;
			
			return HostnameRegex.IsMatch(hostname);
		}

		private static bool IsValidAddress(string address)
		{
			return IsValidIP(address) || IsValidHostname(address);
		}

		private static bool IsValidIP(string ip)
		{
			return IsValidIPv4(ip) || IsValidIPv6(ip);
		}

		private static bool IsValidIPv4(string ip)
		{
			return System.Net.IPAddress.TryParse(ip, out var addr) && addr.AddressFamily == System.Net.Sockets.AddressFamily.InterNetwork;
		}

		private static bool IsValidIPv6(string ip)
		{
			return System.Net.IPAddress.TryParse(ip, out var addr) && addr.AddressFamily == System.Net.Sockets.AddressFamily.InterNetworkV6;
		}

		private static bool IsValidUri(string uri)
		{
			return System.Uri.TryCreate(uri, UriKind.Absolute, out _);
		}

		private static bool IsValidUriRef(string uri)
		{
			return System.Uri.TryCreate(uri, UriKind.RelativeOrAbsolute, out _);
		}

		private static bool IsValidUuid(string uuid)
		{
			if (string.IsNullOrEmpty(uuid))
				return false;
			
			return UuidRegex.IsMatch(uuid);
		}

		private static bool ContainsByteSequence(byte[] haystack, byte[] needle)
		{
			if (needle.Length == 0) return true;
			if (haystack.Length < needle.Length) return false;

			for (int i = 0; i <= haystack.Length - needle.Length; i++)
			{
				bool found = true;
				for (int j = 0; j < needle.Length; j++)
				{
					if (haystack[i + j] != needle[j])
					{
						found = false;
						break;
					}
				}
				if (found) return true;
			}
			return false;
		}

		#endregion
`
