```toml
name = 'Verify'
description = 'Verifies users email'
method = 'POST'
url = '{{auth}}/verify'
sortWeight = 2000000
id = '92a061e7-08ae-4196-8f4a-d861c9816bfc'

[body]
type = 'JSON'
raw = '''
{
  "code": "24477"
}'''
```
