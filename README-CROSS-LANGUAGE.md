# OpenTDF Cross-Language Native Architecture

This branch adds Go-side support for cross-language integration with Java via:

- **Apache Fory** - High-performance binary serialization for Go/Java interop
- **CGO (C-Go) Shared Library** - Exports OpenTDF platform services for Java FFM (Foreign Function & Memory API) consumption
- **UDS Service** - Go service for two-process container architecture

## New Modules

| Module | Description |
|--------|-------------|
| `lib/fory` | Apache Fory codec and DTOs for cross-language serialization |
| `cmd/libopentdf` | CGO shared library exporting platform services |
| `cmd/opentdf-service` | UDS-based Go service for two-process architecture |

---

## Two-Process UDS Architecture

The `cmd/opentdf-service` implements a Go service designed to run as PID 1 in a scratch container, communicating with a Java CLI client via Unix Domain Socket:

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                       Scratch Container (FROM scratch)                      │
│                                                                             │
│  ┌───────────────────────┐       UDS        ┌────────────────────────────┐  │
│  │    opentdf-client     │◄────────────────►│      opentdf-service       │  │
│  │   (Java, GraalVM)     │  /opentdf.sock   │      (this module)         │  │
│  │                       │                  │                            │  │
│  │  Responsibilities:    │    Protocol:     │  Responsibilities:         │  │
│  │  • PicoCLI frontend   │    [type:1]      │  • TDF policy enforcement  │  │
│  │  • User interaction   │    [length:4]    │  • KAS operations          │  │
│  │  • Fory serialize     │    [body:N]      │  • Authorization decisions │  │
│  │                       │                  │  • Policy storage          │  │
│  │                       │      Fory        │  • Entity resolution       │  │
│  │                       │    Serialized    │  • Fory serialization      │  │
│  │                       │                  │  • Child process reaping   │  │
│  │                       │                  │  • Signal handling         │  │
│  └───────────────────────┘                  └────────────────────────────┘  │
│                                                                             │
│  Container contents:           Security:                                    │
│  /opentdf-client   (5.7 MB)    • No network stack                           │
│  /opentdf-service  (3.5 MB)    • No shell                                   │
│  /opentdf.sock     (runtime)   • No libc (musl static)                      │
│                                • UDS namespace-isolated                     │
│  Total: ~9.2 MB                • Separate process memory spaces             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Why Two Processes?

| Approach | Problem |
|----------|---------|
| **Static linking Go + Java** | Go runtime conflicts with GraalVM runtime (GC, signals, stack) |
| **Dynamic linking (.so)** | Requires glibc, larger image, LD_PRELOAD attack surface |
| **Two-process UDS** | Clean separation, both fully static musl, minimal attack surface |

---

## Build Instructions

### Prerequisites

- Go 1.24+
- CGO enabled (`CGO_ENABLED=1`)

### Build Fory Module

```bash
cd lib/fory
go build ./...
go test ./...
```

### Build CGO Shared Library

```bash
cd cmd/libopentdf
go build -buildmode=c-shared -o libopentdf.so .
```

### Build UDS Service (Static musl)

```bash
cd cmd/opentdf-service
CGO_ENABLED=0 go build -ldflags="-s -w" -o opentdf-service .
```

### Build Static Archive (for static linking)

```bash
cd cmd/libopentdf
go build -buildmode=c-archive -o libopentdf.a .
```

---

## Exported Functions (CGO Library)

### Platform Lifecycle

| Function | Signature | Description |
|----------|-----------|-------------|
| `PlatformNew` | `(config *char) uint64` | Initialize platform, returns handle |
| `PlatformFree` | `(handle uint64)` | Release platform resources |
| `PlatformGetVersion` | `(buf *char, len *int) int` | Get SDK version |

### Key Access Service (KAS)

| Function | Signature | Description |
|----------|-----------|-------------|
| `KASGetPublicKey` | `(handle, algo, buf, len) int` | Get KAS public key |
| `KASRewrap` | `(handle, req, reqLen, resp, respLen) int` | Rewrap key |

### Authorization

| Function | Signature | Description |
|----------|-----------|-------------|
| `AuthGetDecision` | `(handle, req, reqLen, resp, respLen) int` | Get access decision |
| `AuthGetEntitlements` | `(handle, req, reqLen, resp, respLen) int` | Get entitlements |

### Policy

| Function | Signature | Description |
|----------|-----------|-------------|
| `PolicyListAttributes` | `(handle, resp, respLen) int` | List attributes |
| `PolicyGetAttributeValues` | `(handle, attrId, resp, respLen) int` | Get attribute values |

### Entity Resolution

| Function | Signature | Description |
|----------|-----------|-------------|
| `EntityResolve` | `(handle, req, reqLen, resp, respLen) int` | Resolve entity |

---

## UDS Service Message Types

The `cmd/opentdf-service` handles these message types over UDS:

| Type | Value | Description |
|------|------:|-------------|
| `CREATE_POLICY` | 1 | Create TDF policy |
| `ENCRYPT` | 2 | Encrypt data |
| `DECRYPT` | 3 | Decrypt data |
| `VERIFY` | 4 | Verify policy binding |
| `GET_PUBLIC_KEY` | 5 | Get KAS public key |
| `REWRAP` | 6 | Rewrap key |
| `GET_DECISION` | 7 | Get authorization decision |
| `GET_ENTITLEMENTS` | 8 | Get entity entitlements |
| `LIST_ATTRIBUTES` | 9 | List policy attributes |
| `GET_VALUES` | 10 | Get attribute values |
| `RESOLVE_ENTITY` | 11 | Resolve entity chain |
| `VERSION` | 12 | Get service version |
| `SHUTDOWN` | 255 | Graceful shutdown |

---

## Benchmarks

### Run Fory Benchmarks

```bash
cd lib/fory
go test -bench=. -benchmem
```

### Benchmark Results (AMD Ryzen 7 9700X)

| Benchmark | Time | Memory | Allocations |
|-----------|-----:|-------:|------------:|
| EntityChain Serialize | 131 ns/op | 0 B/op | 0 allocs/op |
| EntityChain Deserialize | 185 ns/op | 80 B/op | 6 allocs/op |
| EntityChain RoundTrip | 405 ns/op | 80 B/op | 6 allocs/op |
| DecisionResponse Serialize | 105 ns/op | 80 B/op | 1 allocs/op |
| DecisionResponse Deserialize | 127 ns/op | 112 B/op | 4 allocs/op |
| Token Serialize | 57 ns/op | 0 B/op | 0 allocs/op |
| Token Deserialize | 102 ns/op | 485 B/op | 2 allocs/op |

**Key observations:**
- Sub-200ns latency for all operations
- Full round-trip under 500ns for most message types
- Optimized serializers with deferred error checking

### Fory vs Protobuf Comparison

Run comparison benchmarks:

```bash
cd lib/fory
go test -bench='Fory|Protobuf' -benchmem
```

**Optimizations applied:**
- Disabled ref tracking for DTOs without circular references (~15% gain)
- Hand-optimized serializers with deferred error checking (~9% gain)
- Object pooling with pre-allocated nested structs

#### Serialization (ns/op, lower is better)

| Message Type | Fory | Protobuf | Fory Advantage |
|--------------|-----:|---------:|----------------|
| EntityChain | 131 ns | 199 ns | **1.52x faster** |
| DecisionResponse | 105 ns | 144 ns | **1.37x faster** |
| Token | 57 ns | 90 ns | **1.58x faster** |

#### Deserialization - Realistic Usage (no pooling)

This represents typical developer code: `var result DecisionResponse; codec.Deserialize(data, &result)`

| Message Type | Fory | Protobuf | Fory Advantage |
|--------------|-----:|---------:|----------------|
| DecisionResponse | 193 ns, 7 allocs | 291 ns, 9 allocs | **1.51x faster, 22% fewer allocs** |

#### Deserialization - Optimized (with object pooling)

For high-throughput servers using `sync.Pool`:

| Message Type | Fory | Protobuf | Fory Advantage |
|--------------|-----:|---------:|----------------|
| EntityChain | 185 ns, 6 allocs | 335 ns, 11 allocs | **1.81x faster** |
| DecisionResponse | 127 ns, 4 allocs | 238 ns, 8 allocs | **1.87x faster, 50% fewer allocs** |
| Token | 102 ns, 2 allocs | 112 ns, 2 allocs | **1.10x faster** |

#### Round-Trip (serialize + deserialize)

| Message Type | Fory | Protobuf | Fory Advantage |
|--------------|-----:|---------:|----------------|
| EntityChain | 405 ns, 6 allocs | 596 ns, 12 allocs | **1.47x faster** |

**Key findings:**
- **Fory wins all operations** - Both serialize and deserialize
- **1.4x-1.9x faster than Protobuf** across all message types
- **Fewer allocations** - Less GC pressure in sustained workloads
- **Optimized serializers** - Hand-tuned with deferred error checking pattern

---

## Fory Configuration

Apache Fory uses struct tags for field mapping. Note: Fory was renamed from Fury, so tags use `fory:`:

```go
type Entity struct {
    EphemeralID string         `fory:"ephemeralId"`
    EntityType  EntityType     `fory:"entityType"`
    EntityValue string         `fory:"entityValue"`
    Category    EntityCategory `fory:"category"`
}
```

### Cross-Language Type Registration

Types must be registered with matching names in both Go and Java:

```go
// Go
codec.RegisterNamedStruct("io.opentdf.fory.dto.EntityDto", &dto.Entity{})
codec.RegisterNamedEnum("io.opentdf.fory.dto.EntityDto$EntityType", dto.EntityTypeEmailAddress)
```

```java
// Java
fory.register(EntityDto.class, "io.opentdf.fory.dto.EntityDto");
```

---

## Related Work

See the companion changes in [opentdf-java-sdk](https://github.com/morrismeyer/opentdf-java-sdk/tree/morrismeyer/opentdf-cross-language-native):
- `fory-serialization` - Java Fory DTOs and codec
- `ffm` - FFM bindings to load and call libopentdf.so
- `ffm/UDSClient.java` - Java UDS client for two-process architecture
- `ffm/TDFForyCli.java` - Java CLI using UDS transport
- `benchmarks` - JMH benchmarks comparing Fory vs Protobuf
- `tiny-container` - GraalVM native-image scratch container demos
