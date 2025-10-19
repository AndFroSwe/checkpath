# CheckPath

Simple Go utility that checks that entries in the environment
path variable has a directory associated with them. Used to
cleanup the environment (typically needed on Windows where
installation/uninstallation is a complete mess...).

Usage:

``` bash
go run cmd/checkpath.go
```

Check the output, if it lists any entries, remove them from path.
Note that you may need to reload the environment or restart the 
terminal before running again to get the current path.

## Disclaimer

* This is a toy project and is not intended for production use (duh!)
* This was not vibe coded, despite the use of emojis in prints. I just thought
it looked nice :)
* Only tested on Windows
