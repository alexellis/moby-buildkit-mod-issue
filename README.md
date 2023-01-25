# moby-buildkit-mod-issue

github.com/moby/moby/pkg/archive can't be referenced in go mod when buildkit is also referenced.

```
go build
# github.com/moby/moby/pkg/archive
../../../../pkg/mod/github.com/moby/moby@v20.10.23+incompatible/pkg/archive/archive.go:763:12: undefined: idtools.NewIDMappingsFromMaps
../../../../pkg/mod/github.com/moby/moby@v20.10.23+incompatible/pkg/archive/archive.go:923:23: undefined: idtools.NewIDMappingsFromMaps
../../../../pkg/mod/github.com/moby/moby@v20.10.23+incompatible/pkg/archive/changes.go:401:32: undefined: idtools.NewIDMappingsFromMaps
../../../../pkg/mod/github.com/moby/moby@v20.10.23+incompatible/pkg/archive/diff.go:36:23: undefined: idtools.NewIDMappingsFromMaps
../../../../pkg/mod/github.com/moby/moby@v20.10.23+incompatible/pkg/archive/diff.go:244:26: undefined: system.Umask
../../../../pkg/mod/github.com/moby/moby@v20.10.23+incompatible/pkg/archive/diff.go:248:16: undefined: system.Umask
```
