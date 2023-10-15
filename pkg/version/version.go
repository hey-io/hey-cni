package version

import "runtime/debug"

var (
	HeyCNIVersion string = "latest"
	GitRevision   string
	GoVersion     string
)

func init() {
	buildInfo, ok := debug.ReadBuildInfo()

	if ok {
		GoVersion = buildInfo.GoVersion
		for _, v := range buildInfo.Settings {
			if v.Key == "vcs.revision" {
				GitRevision = v.Value
				break
			}
		}
	}
}
