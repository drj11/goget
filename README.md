# goget

```
goget --out destination-path --sha256 64_hexdigit_sha256sum URL
```

`goget` is designed to do just one thing: get a file from the
internet and write it to disk. The file is only written to the
destination-path when the SHA-256 checksum has been verified.

`goget` will only exit successfully (with 0 exit code) if the
file has been downloaded from the URL given, it's SHA-256
checksum matches, and the temporary file used during download
has been successfully renamed to the final target.

## Why not use curl?

`curl` is big and bloated and almost certainly has bugs. You do
not want it on a production machine.
