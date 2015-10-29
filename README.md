# goget

`goget` is designed to do just one thing: get a file from the
internet and write it to disk, having verified via a
user-supplied checksum that it is the right file.

```
goget -sha256 64_hexdigit_sha256sum URL
```

## Caution

It's not fully implemented yet.

## Why not use curl?

`curl` is big and bloated and almost certainly has bugs. You do
not want it on a production machine.
