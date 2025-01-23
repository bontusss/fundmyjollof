```toml
name = 'FundMyJollof'
description = 'Fundmyjollof API v1'
id = '62e02718-73f7-46b0-8fde-1b2b48b4b6e1'

[[environmentGroups]]
name = 'Default'
environments = ['dev']
```

#### Variables

```json5
{
  dev: {
    "auth": "http://localhost:7000/api/v1/auth",
    "creator": "http://localhost:7000/api/v1/creator"
  }
}
```
