```toml
name = 'Register'
description = 'Registers a user'
method = 'POST'
url = '{{auth}}/register'
sortWeight = 1000000
id = 'f7a910b0-848b-4981-a397-4e22232274d9'

[body]
type = 'JSON'
raw = '''
{
  "email": "ukandu@fmj.com",
  "pass": "ukandu"
}'''
```
