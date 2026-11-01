# Upgrading sqlite

Do not copy the official prebuilt amalgamation directly. It does not enable the `UPDATE ... LIMIT` / `DELETE ... LIMIT` grammar for this driver. Build the amalgamation from the canonical source with `SQLITE_ENABLE_UPDATE_DELETE_LIMIT=1` during code generation.

```sh
version=3540100
work=/tmp/sqlite-src-${version}
mkdir -p "$work"
curl -fL "https://www.sqlite.org/2026/sqlite-src-${version}.zip" -o "$work/sqlite-src.zip"
# verify the SHA3-256 from https://www.sqlite.org/download.html before unzip
unzip -q -o "$work/sqlite-src.zip" -d "$work"
cd "$work/sqlite-src-${version}"
CFLAGS='-DSQLITE_ENABLE_UPDATE_DELETE_LIMIT=1' ./configure
make sqlite3.c
cp sqlite3.c sqlite3.h /path/to/go-sqlite/sqlite3/
```

Before committing:

- check `sqlite3/sqlite3.h` has the expected `SQLITE_VERSION`
- run `go test -v ./sqlite3/... -count=1`
- keep a test that uses `DELETE ... LIMIT` or `UPDATE ... LIMIT`; copying the official amalgamation without regenerating from source makes that test fail even though cgo still passes `-DSQLITE_ENABLE_UPDATE_DELETE_LIMIT=1`
