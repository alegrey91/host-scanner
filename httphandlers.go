package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/armosec/host-sensor/sensor"
	"go.uber.org/zap"
)

func initHTTPHandlers() {
	// TODO: implement probe endpoint
	http.HandleFunc("/kubeletConfigurations", func(rw http.ResponseWriter, r *http.Request) {
		conf, err := sensor.SenseKubeletConfigurations()

		if err != nil {
			http.Error(rw, fmt.Sprintf("failed to sense kubelet conf: %v", err), http.StatusInternalServerError)
		} else {
			rw.WriteHeader(http.StatusOK)
			if _, err := rw.Write(conf); err != nil {
				zap.L().Error("In kubeletConfigurations handler failed to write", zap.Error(err))
			}
		}
	})
	http.HandleFunc("/kubeletCommandLine", func(rw http.ResponseWriter, r *http.Request) {
		proc, err := sensor.LocateKubeletProcess()

		cmdLine := strings.Join(proc.CmdLine, " ")
		if err != nil {
			http.Error(rw, fmt.Sprintf("failed to sense kubelet conf: %v", err), http.StatusInternalServerError)
		} else {
			rw.WriteHeader(http.StatusOK)
			if _, err := rw.Write([]byte(cmdLine)); err != nil {
				zap.L().Error("In kubeletConfigurations handler failed to write", zap.Error(err))
			}
		}
	})
	http.HandleFunc("/osRelease", osReleaseHandler)
	http.HandleFunc("/kernelVersion", kernelVersionHandler)
	http.HandleFunc("/linuxSecurityHardening", linuxSecurityHardeningHandler)
	http.HandleFunc("/openedPorts", openedPortsHandler)
	http.HandleFunc("/LinuxKernelVariables", LinuxKernelVariablesHandler)
}

func LinuxKernelVariablesHandler(rw http.ResponseWriter, r *http.Request) {
	resp, err := sensor.SenseKernelVariables()
	GenericSensorHandler(rw, r, resp, err, "SenseKernelVariables")
}

func openedPortsHandler(rw http.ResponseWriter, r *http.Request) {
	resp, err := sensor.SenseOpenPorts()
	GenericSensorHandler(rw, r, resp, err, "SenseOpenPorts")
}

func osReleaseHandler(rw http.ResponseWriter, r *http.Request) {
	fileContent, err := sensor.SenseOsRelease()
	if err != nil {
		http.Error(rw, fmt.Sprintf("failed to SenseOsRelease: %v", err), http.StatusInternalServerError)
	} else {
		rw.WriteHeader(http.StatusOK)
		if _, err := rw.Write(fileContent); err != nil {
			zap.L().Error("In SenseOsRelease handler failed to write", zap.Error(err))
		}
	}
}

func kernelVersionHandler(rw http.ResponseWriter, r *http.Request) {
	fileContent, err := sensor.SenseKernelVersion()
	if err != nil {
		http.Error(rw, fmt.Sprintf("failed to sense kernelVersionHandler: %v", err), http.StatusInternalServerError)
	} else {
		rw.WriteHeader(http.StatusOK)
		if _, err := rw.Write(fileContent); err != nil {
			zap.L().Error("In kernelVersionHandler handler failed to write", zap.Error(err))
		}
	}
}

func linuxSecurityHardeningHandler(rw http.ResponseWriter, r *http.Request) {
	resp, err := sensor.SenseLinuxSecurityHardening()
	GenericSensorHandler(rw, r, resp, err, "sense linuxSecurityHardeningHandler")
}

func GenericSensorHandler(w http.ResponseWriter, r *http.Request, respContent interface{}, err error, senseName string) {
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to %s: %v", senseName, err), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(respContent); err != nil {
			zap.L().Error(fmt.Sprintf("In %s handler failed to write", senseName), zap.Error(err))
		}
	}
}
