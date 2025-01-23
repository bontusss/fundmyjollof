```toml
name = 'Creator Setup'
description = 'Setups creator initial data'
method = 'POST'
url = '{{creator}}/setup-creator'
sortWeight = 2000000
id = '931299b9-631d-4f39-bffe-7d5a25f8cf2e'

[body]
type = 'JSON'
raw = '''
{
  "email": "ukandu@fmj.com",
  "username": "bontussss",
  "name": "Ukandu Ikwechegh",
  "bio": "Wonderful bio",
  "country": "Nigeria",
  "payment_method": ["MTN"]
}'''
```
