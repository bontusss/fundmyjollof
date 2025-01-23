```toml
name = 'Login'
description = 'Logs a user in and starts a session'
method = 'POST'
url = '{{auth}}/login'
sortWeight = 3000000
id = '01b18356-6e29-4ed5-8e2e-d1fd079e47f9'

[body]
type = 'JSON'
raw = '''
{
  "email": "ukandu@fmj.com",
  "pass": "ukandu"
}'''
```
