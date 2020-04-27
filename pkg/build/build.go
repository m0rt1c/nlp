package build

var gitCommitID = "dev"

// Version get current build version
func Version() string {
	return gitCommitID
}
