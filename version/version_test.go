package version_test

import (
	"github.com/PredixDev/cf-predix/version"
	"github.com/cloudfoundry/cli/plugin"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Version", func() {
	Describe("VersionInfo", func() {
		It("Parses the version string to form the plugin version info", func() {
			version.VersionString = "1.2.3"

			versionInfo := version.GetVersion()
			Expect(versionInfo).To(Equal(plugin.VersionType{
				Major: 1,
				Minor: 2,
				Build: 3,
			}))
		})
	})
})
