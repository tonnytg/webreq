# Security Policy

## Vulnerability Assessment

This document outlines the security vulnerability assessment performed on the webreq library.

### Assessment Date
January 2026

### Dependencies Check
✅ **No external dependencies detected** - The webreq library uses only the Go standard library, which significantly reduces the attack surface.

### Code Security Analysis

#### Vulnerability Found and Fixed

**Vulnerability:** Unbounded Response Body Reading (DoS/Memory Exhaustion)
- **Severity:** HIGH
- **Status:** ✅ FIXED
- **CWE:** CWE-770 (Allocation of Resources Without Limits or Throttling)

**Description:**
The original implementation used `io.ReadAll(response.Body)` without any size limit, which could lead to:
- Memory exhaustion if a malicious server returns extremely large responses
- Denial of Service (DoS) attacks
- Application crashes due to Out of Memory errors

**Fix Implemented:**
1. Added `MaxResponseSize` field to the `Request` struct with a default limit of 100MB
2. Implemented `io.LimitReader` to enforce the maximum response size
3. Added `SetMaxResponseSize()` method to allow users to customize the limit
4. Added comprehensive tests to verify the fix

**Code Changes:**
```go
// Before (VULNERABLE)
responseBody, err := io.ReadAll(response.Body)

// After (SECURE)
limitedReader := io.LimitReader(response.Body, request.MaxResponseSize)
responseBody, err := io.ReadAll(limitedReader)
```

**Usage:**
```go
// Uses default 100MB limit
request := webreq.NewRequest("GET")
request.SetURL("https://api.example.com/data")
response, err := request.Execute()

// Custom limit (e.g., 10MB for smaller responses)
request := webreq.NewRequest("GET")
request.SetURL("https://api.example.com/data")
request.SetMaxResponseSize(10 * 1024 * 1024) // 10MB
response, err := request.Execute()
```

### Additional Security Considerations

While no other vulnerabilities were found, users should follow these best practices:

1. **HTTPS/TLS**: Always use HTTPS URLs when making requests to protect data in transit
2. **Input Validation**: Validate and sanitize any user input before using it in URLs or headers
3. **Timeouts**: Use appropriate timeout values via `SetTimeout()` to prevent hanging requests
4. **Error Handling**: Always check and properly handle errors returned by `Execute()`
5. **Response Size**: Consider the expected response size and set `MaxResponseSize` accordingly

### Reporting Security Issues

If you discover a security vulnerability in this project, please report it by:
1. Opening a security advisory on GitHub
2. Emailing the maintainer at tonnytg@gmail.com

Please do not open public issues for security vulnerabilities.

### Security Update History

- **2026-01**: Fixed unbounded response body reading vulnerability (CWE-770)
