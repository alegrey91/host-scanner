package e2e_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/kubescape/host-scanner/sensor"
	ds "github.com/kubescape/host-scanner/sensor/datastructures"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CniInfo", func() {
	var (
		res     *http.Response
		err     error
		resBody []byte
	)

	Context("testing /cniinfo endpoint", func() {
		It("should respond to a GET request", func() {
			requestURL := url + "/cniinfo"
			res, err = http.Get(requestURL)
			Expect(err).ToNot(HaveOccurred())
		})
		It("should return a 200 status code", func() {
			Expect(res.StatusCode).To(BeEquivalentTo(200))
		})
		It("should return the expected value of CNIInfo", func() {
			jsonToCompare := &sensor.CNIInfo{
				CNIConfigFiles: []*ds.FileInfo{
					{
						Ownership: &ds.FileOwnership{
							Err:       "",
							UID:       0,
							GID:       0,
							Username:  "root",
							Groupname: "root",
						},
						Path:        "/etc/cni/net.d/10-kindnet.conflist",
						Permissions: 420,
					},
				},
				CNINames: []string{"Kindnet"},
			}
			jsonCniInfo := &sensor.CNIInfo{}

			resBody, err = ioutil.ReadAll(res.Body)
			Expect(err).ToNot(HaveOccurred())

			err = json.Unmarshal(resBody, jsonCniInfo)
			Expect(err).ToNot(HaveOccurred())

			Expect(jsonCniInfo.CNINames).To(Equal(jsonToCompare.CNINames))

			for i := range jsonCniInfo.CNIConfigFiles {
				Expect(jsonCniInfo.CNIConfigFiles[i]).To(Equal(jsonToCompare.CNIConfigFiles[i]))
			}
		})
	})
})
