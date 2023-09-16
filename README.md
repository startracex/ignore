# ignore

Parse simple ignore rules

from content

```go
	content := `
...
`
	var rules []string = ignore.GetRules(content)
```

from file

```go
  bytes, _ :=os.ReadFIle("...")
  rules = ignore.GetRules(string(bytes))
```

check match

```go
  matched := false
  for _, rule := range rules {
    if Match(rule, "..."){
      matched = true
      break
    }
  }
```
