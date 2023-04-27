package e2e_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	sensor "github.com/kubescape/host-scanner/sensor"
	ds "github.com/kubescape/host-scanner/sensor/datastructures"
)

var _ = Describe("ControlPlaneInfo", func() {
	var (
		res     *http.Response
		err     error
		resBody []byte
	)

	Context("testing /controlplaneinfo endpoint", func() {
		It("should respond to a GET request", func() {
			requestURL := url + "/controlplaneinfo"
			res, err = http.Get(requestURL)
			Expect(err).ToNot(HaveOccurred())
		})
		It("should return a 200 status code", func() {
			Expect(res.StatusCode).To(BeEquivalentTo(200))
		})
		It("should return the expected value of PKIDir and PKIFiles", func() {
			jsonToCompare := &sensor.ControlPlaneInfo{
				PKIDIr: &ds.FileInfo{Path: "/etc/kubernetes/pki"},
				PKIFiles: []*ds.FileInfo{
					{
						Path: "/etc/kubernetes/pki/ca.key",
					},
					{
						Path: "/etc/kubernetes/pki/ca.crt",
					},
					{
						Path: "/etc/kubernetes/pki/apiserver.key",
					},
					{
						Path: "/etc/kubernetes/pki/apiserver.crt",
					},
					{
						Path: "/etc/kubernetes/pki/apiserver-kubelet-client.key",
					},
					{
						Path: "/etc/kubernetes/pki/apiserver-kubelet-client.crt",
					},
					{
						Path: "/etc/kubernetes/pki/front-proxy-ca.key",
					},
					{
						Path: "/etc/kubernetes/pki/front-proxy-ca.crt",
					},
					{
						Path: "/etc/kubernetes/pki/front-proxy-client.key",
					},
					{
						Path: "/etc/kubernetes/pki/front-proxy-client.crt",
					},
					{
						Path: "/etc/kubernetes/pki/etcd/ca.key",
					},
					{
						Path: "/etc/kubernetes/pki/etcd/ca.crt",
					},
					{
						Path: "/etc/kubernetes/pki/etcd/server.key",
					},
					{
						Path: "/etc/kubernetes/pki/etcd/server.crt",
					},
					{
						Path: "/etc/kubernetes/pki/etcd/peer.key",
					},
					{
						Path: "/etc/kubernetes/pki/etcd/peer.crt",
					},
					{
						Path: "/etc/kubernetes/pki/etcd/healthcheck-client.key",
					},
					{
						Path: "/etc/kubernetes/pki/etcd/healthcheck-client.crt",
					},
					{
						Path: "/etc/kubernetes/pki/apiserver-etcd-client.key",
					},
					{
						Path: "/etc/kubernetes/pki/apiserver-etcd-client.crt",
					},
					{
						Path: "/etc/kubernetes/pki/sa.key",
					},
					{
						Path: "/etc/kubernetes/pki/sa.pub",
					},
				},
			}
			jsonControlPlaneInfo := &sensor.ControlPlaneInfo{}

			resBody, err = ioutil.ReadAll(res.Body)
			Expect(err).ToNot(HaveOccurred())

			err = json.Unmarshal(resBody, jsonControlPlaneInfo)
			Expect(err).ToNot(HaveOccurred())

			Expect(jsonControlPlaneInfo.PKIDIr.Path).To(Equal(jsonToCompare.PKIDIr.Path))

			for i := range jsonControlPlaneInfo.PKIFiles {
				Expect(jsonControlPlaneInfo.PKIFiles[i].Path).To(Equal(jsonToCompare.PKIFiles[i].Path))
			}
		})
		It("should return the expected value of ApiServerInfo", func() {
			jsonToCompare := &sensor.ControlPlaneInfo{
				APIServerInfo: &sensor.ApiServerInfo{
					K8sProcessInfo: &sensor.K8sProcessInfo{
						SpecsFile: &ds.FileInfo{
							Ownership: &ds.FileOwnership{
								Err:       "",
								UID:       0,
								GID:       0,
								Username:  "root",
								Groupname: "root",
							},
							Path:        "/etc/kubernetes/manifests/kube-apiserver.yaml",
							Permissions: 384,
						},
					},
				},
			}
			jsonControlPlaneInfo := &sensor.ControlPlaneInfo{}

			err = json.Unmarshal(resBody, jsonControlPlaneInfo)
			Expect(err).ToNot(HaveOccurred())

			Expect(jsonControlPlaneInfo.APIServerInfo.K8sProcessInfo.SpecsFile).
				To(Equal(jsonToCompare.APIServerInfo.K8sProcessInfo.SpecsFile))
		})
		It("should return the expected value of ControllerManagerInfo", func() {
			jsonToCompare := &sensor.ControlPlaneInfo{
				ControllerManagerInfo: &sensor.K8sProcessInfo{
					SpecsFile: &ds.FileInfo{
						Ownership: &ds.FileOwnership{
							Err:       "",
							UID:       0,
							GID:       0,
							Username:  "root",
							Groupname: "root",
						},
						Path:        "/etc/kubernetes/manifests/kube-controller-manager.yaml",
						Permissions: 384,
					},
					ConfigFile: &ds.FileInfo{
						Ownership: &ds.FileOwnership{
							Err:       "",
							UID:       0,
							GID:       0,
							Username:  "root",
							Groupname: "root",
						},
						Path:        "/etc/kubernetes/controller-manager.conf",
						Permissions: 384,
					},
				},
			}
			jsonControlPlaneInfo := &sensor.ControlPlaneInfo{}

			err = json.Unmarshal(resBody, jsonControlPlaneInfo)
			Expect(err).ToNot(HaveOccurred())

			Expect(jsonControlPlaneInfo.ControllerManagerInfo.SpecsFile).
				To(Equal(jsonToCompare.ControllerManagerInfo.SpecsFile))
			Expect(jsonControlPlaneInfo.ControllerManagerInfo.ConfigFile).
				To(Equal(jsonToCompare.ControllerManagerInfo.ConfigFile))
			Expect(jsonControlPlaneInfo.ControllerManagerInfo.KubeConfigFile).
				To(Equal(jsonToCompare.ControllerManagerInfo.KubeConfigFile))
			Expect(jsonControlPlaneInfo.ControllerManagerInfo.ClientCAFile).
				To(Equal(jsonToCompare.ControllerManagerInfo.ClientCAFile))
		})
		It("should return the expected value of SchedulerInfo", func() {
			jsonToCompare := &sensor.ControlPlaneInfo{
				SchedulerInfo: &sensor.K8sProcessInfo{
					SpecsFile: &ds.FileInfo{
						Ownership: &ds.FileOwnership{
							Err:       "",
							UID:       0,
							GID:       0,
							Username:  "root",
							Groupname: "root",
						},
						Path:        "/etc/kubernetes/manifests/kube-scheduler.yaml",
						Permissions: 384,
					},
					ConfigFile: &ds.FileInfo{
						Ownership: &ds.FileOwnership{
							Err:       "",
							UID:       0,
							GID:       0,
							Username:  "root",
							Groupname: "root",
						},
						Path:        "/etc/kubernetes/scheduler.conf",
						Permissions: 384,
					},
				},
			}
			jsonControlPlaneInfo := &sensor.ControlPlaneInfo{}

			err = json.Unmarshal(resBody, jsonControlPlaneInfo)
			Expect(err).ToNot(HaveOccurred())

			Expect(jsonControlPlaneInfo.SchedulerInfo.SpecsFile).
				To(Equal(jsonToCompare.SchedulerInfo.SpecsFile))
			Expect(jsonControlPlaneInfo.SchedulerInfo.ConfigFile).
				To(Equal(jsonToCompare.SchedulerInfo.ConfigFile))
			Expect(jsonControlPlaneInfo.SchedulerInfo.KubeConfigFile).
				To(Equal(jsonToCompare.SchedulerInfo.KubeConfigFile))
			Expect(jsonControlPlaneInfo.SchedulerInfo.ClientCAFile).
				To(Equal(jsonToCompare.SchedulerInfo.ClientCAFile))
		})
		It("should return the expected value of EtcdConfigFile and EtcdDataDir", func() {
			jsonToCompare := &sensor.ControlPlaneInfo{
				EtcdConfigFile: &ds.FileInfo{
					Ownership: &ds.FileOwnership{
						Err:       "",
						UID:       0,
						GID:       0,
						Username:  "root",
						Groupname: "root",
					},
					Path:        "/etc/kubernetes/manifests/etcd.yaml",
					Permissions: 384,
				},
				EtcdDataDir: &ds.FileInfo{
					Ownership: &ds.FileOwnership{
						Err:       "",
						UID:       0,
						GID:       0,
						Username:  "root",
						Groupname: "root",
					},
					Path:        "/var/lib/etcd",
					Permissions: 448,
				},
			}
			jsonControlPlaneInfo := &sensor.ControlPlaneInfo{}

			err = json.Unmarshal(resBody, jsonControlPlaneInfo)
			Expect(err).ToNot(HaveOccurred())

			Expect(jsonControlPlaneInfo.EtcdConfigFile).
				To(Equal(jsonToCompare.EtcdConfigFile))
			Expect(jsonControlPlaneInfo.EtcdDataDir).
				To(Equal(jsonToCompare.EtcdDataDir))
		})
		It("should return the expected value of EtcdConfigFile and EtcdDataDir", func() {
			jsonToCompare := &sensor.ControlPlaneInfo{
				AdminConfigFile: &ds.FileInfo{
					Ownership: &ds.FileOwnership{
						Err:       "",
						UID:       0,
						GID:       0,
						Username:  "root",
						Groupname: "root",
					},
					Path:        "/etc/kubernetes/admin.conf",
					Permissions: 384,
				},
			}
			jsonControlPlaneInfo := &sensor.ControlPlaneInfo{}

			err = json.Unmarshal(resBody, jsonControlPlaneInfo)
			Expect(err).ToNot(HaveOccurred())

			Expect(jsonControlPlaneInfo.AdminConfigFile).
				To(Equal(jsonToCompare.AdminConfigFile))
		})
	})
})
