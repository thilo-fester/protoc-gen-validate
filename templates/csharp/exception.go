package csharp

const exceptionTpl = `
	/// <summary>
	/// Exception thrown when protobuf message validation fails
	/// </summary>
	public class ValidationException : System.Exception
	{
		public string FieldName { get; }
		public string FieldPath { get; }
		public string Reason { get; }

		public ValidationException(string fieldPath, string fieldName, string reason) 
			: base($"Validation failed for field '{fieldPath}.{fieldName}': {reason}")
		{
			FieldPath = fieldPath;
			FieldName = fieldName;
			Reason = reason;
		}

		public ValidationException(string fieldPath, string fieldName, string reason, System.Exception innerException) 
			: base($"Validation failed for field '{fieldPath}.{fieldName}': {reason}", innerException)
		{
			FieldPath = fieldPath;
			FieldName = fieldName;
			Reason = reason;
		}
	}
`
