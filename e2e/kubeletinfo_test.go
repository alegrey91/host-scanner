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

var _ = Describe("Kubeletinfo", func() {
	var (
		res     *http.Response
		err     error
		resBody []byte
	)

	Context("testing /kubeletinfo endpoint", func() {
		It("should respond to a GET request", func() {
			requestURL := url + "/kubeletinfo"
			res, err = http.Get(requestURL)
			Expect(err).ToNot(HaveOccurred())
		})
		It("should return a 200 status code", func() {
			Expect(res.StatusCode).To(BeEquivalentTo(200))
		})
		It("should return the expected value of KubeletInfo", func() {
			jsonToCompare := &sensor.KubeletInfo{
				ServiceFiles: []ds.FileInfo{
					{
						Ownership: &ds.FileOwnership{
							Err:       "",
							UID:       0,
							GID:       0,
							Username:  "root",
							Groupname: "root",
						},
						Path:        "/etc/systemd/system/kubelet.service.d/10-kubeadm.conf",
						Permissions: 420,
					},
				},
				ConfigFile: &ds.FileInfo{
					Ownership: &ds.FileOwnership{
						Err:       "",
						UID:       0,
						GID:       0,
						Username:  "root",
						Groupname: "root",
					},
					Path:        "/var/lib/kubelet/config.yaml",
					Permissions: 420,
				},
				KubeConfigFile: &ds.FileInfo{
					Ownership: &ds.FileOwnership{
						Err:       "",
						UID:       0,
						GID:       0,
						Username:  "root",
						Groupname: "root",
					},
					Path:        "/etc/kubernetes/kubelet.conf",
					Permissions: 420,
				},
				ClientCAFile: &ds.FileInfo{
					Ownership: &ds.FileOwnership{
						Err:       "",
						UID:       0,
						GID:       0,
						Username:  "root",
						Groupname: "root",
					},
					Path:        "/etc/kubernetes/pki/ca.crt",
					Permissions: 420,
				},
			}
			jsonKubeletInfo := &sensor.KubeletInfo{}

			resBody, err = ioutil.ReadAll(res.Body)
			Expect(err).ToNot(HaveOccurred())

			err = json.Unmarshal(resBody, jsonKubeletInfo)
			Expect(err).ToNot(HaveOccurred())

			Expect(jsonKubeletInfo.ServiceFiles).
				To(Equal(jsonToCompare.ServiceFiles))
			// we checks only 'Path' in order to avoid fill 'Content' field too
			Expect(jsonKubeletInfo.ConfigFile.Path).
				To(Equal(jsonToCompare.ConfigFile.Path))
			Expect(jsonKubeletInfo.KubeConfigFile.Path).
				To(Equal(jsonToCompare.KubeConfigFile.Path))
			Expect(jsonKubeletInfo.ClientCAFile).
				To(Equal(jsonToCompare.ClientCAFile))
		})
	})
})
