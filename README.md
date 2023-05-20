# Map Utils <a href="https://gitpod.io/#https://github.com/gouniverse/maputils" style="float:right:"><img src="https://gitpod.io/button/open-in-gitpod.svg" alt="Open in Gitpod" loading="lazy"></a>

![tests](https://github.com/gouniverse/maputils/workflows/tests/badge.svg)

## Description

Utility library for working with maps

## Installation

```
go get github.com/gouniverse/maputils
```

## Function

- AnyToArrayMapStringAny(data any) []map[string]any - converts an interface to array map[string]any
- AnyToArrayMapStringString(data any) []map[string]string - converts an interface to array of map[string]string

- AnyToMapStringAny(data any) map[string]any - converts an interface to map[string]any
- AnyToMapStringString(data any) map[string]string - converts an interface to map[string]string

- ArgsToMap([]string) map[string]string - converts an CLI arguments array to map

- MapAnyToArrayMapStringAny(valueAny any) []map[string]any - converts an interface to array map[string]any
- MapStringAnyToMapStringString(data map[string]any) map[string]string - converts a map[string]any to map[string]string
