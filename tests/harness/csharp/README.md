# C# Test Harness

This directory contains the C# test harness for protoc-gen-validate.

## Status

The C# test harness is currently a placeholder. To implement a full C# test harness:

1. Generate C# protobuf classes for test cases
2. Create a C# console application that:
   - Reads a serialized TestCase from stdin
   - Deserializes the message 
   - Runs validation using the generated validator methods
   - Outputs a TestResult

## Requirements

- .NET 6.0 or later
- Google.Protobuf NuGet package
- Generated C# protobuf classes for test cases

## Implementation Notes

The harness should follow the same pattern as other language harnesses:
- Read binary protobuf TestCase from stdin
- Execute validation 
- Write binary protobuf TestResult to stdout

Example structure:
```csharp
// Read test case
var input = Console.OpenStandardInput();
var testCase = TestCase.Parser.ParseFrom(input);

// Validate message
var result = new TestResult();
try 
{
    // Call generated validator
    SomeValidator.ValidateSomeMessage(message);
    result.Valid = true;
}
catch (ValidationException ex)
{
    result.Valid = false;
    result.Reasons.Add(ex.Message);
}

// Write result
Console.OpenStandardOutput().Write(result.ToByteArray());
```
