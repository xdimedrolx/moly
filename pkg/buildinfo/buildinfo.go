package buildinfo

import "github.com/sagikazarmark/appkit/buildinfo"

// Provisioned by ldflags
// nolint: gochecknoglobals
var (
	version    string = "development"
	commitHash string
	buildDate  string
)

var BuildInfo = buildinfo.New(version, commitHash, buildDate)