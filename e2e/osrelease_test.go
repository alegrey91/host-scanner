package e2e_test

import (
	"io/ioutil"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OsRelease", func() {
	var (
		res     *http.Response
		err     error
		resBody []byte
		// this is usually the /etc/os-release content of github-actions workflows
		expectedResult = `PRETTY_NAME="Ubuntu 22.04.1 LTS"
NAME="Ubuntu"
VERSION_ID="22.04"
VERSION="22.04.1 LTS (Jammy Jellyfish)"
VERSION_CODENAME=jammy
ID=ubuntu
ID_LIKE=debian
HOME_URL="https://www.ubuntu.com/"
SUPPORT_URL="https://help.ubuntu.com/"
BUG_REPORT_URL="https://bugs.launchpad.net/ubuntu/"
PRIVACY_POLICY_URL="https://www.ubuntu.com/legal/terms-and-policies/privacy-policy"
UBUNTU_CODENAME=jammy
`
	)

	Context("testing /osrelease endpoint", func() {
		It("should respond to a GET request", func() {
			requestURL := url + "/osrelease"
			res, err = http.Get(requestURL)
			Expect(err).ToNot(HaveOccurred())
		})
		It("should return a 200 status code", func() {
			Expect(res.StatusCode).To(BeEquivalentTo(200))
		})
		It("should return the expected value", func() {
			resBody, err = ioutil.ReadAll(res.Body)
			Expect(string(resBody)).To(BeEquivalentTo(expectedResult))
		})
	})
})
