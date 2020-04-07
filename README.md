# ldapsearch
LDAP Search

## GUI

![gui](gui.png)

## Config file

conf.json

```json
{
    "server": "your.ldap.server:389",
    "attributes": ["commonName","surName","givenName"],
    "email": "your@mail.address",
    "base_dn": "ou=users,o=hogehoge.com",
    "bind_dn": "uid=%EMAIL%,o=login",
    "filter": "(cn=jack*)"
}
```
