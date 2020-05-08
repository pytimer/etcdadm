/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"sigs.k8s.io/etcdadm/apis"
	"sigs.k8s.io/etcdadm/constants"
	log "sigs.k8s.io/etcdadm/pkg/logrus"
)

var etcdAdmConfig apis.EtcdAdmConfig

// LogLevel sets the log level for the logger
var LogLevel string

var (
	rootCmd = &cobra.Command{
		Use:  "etcdadm",
		Long: `Tool to bootstrap etcdadm on the host`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			logLevel, err := logrus.ParseLevel(LogLevel)
			if err != nil {
				log.Fatalf("Could not parse log level %v", logLevel)
			}
			log.SetLogLevel(logLevel)
		},
	}
)

// Execute executes the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&LogLevel, "log-level", "l", "info", "set log level for output, permitted values debug, info, warn, error, fatal and panic")
}

func generateInitOrJoinCommonFlags(flagSet *pflag.FlagSet) {
	flagSet.StringVar(&etcdAdmConfig.Name, "name", "", "etcd member name")
	flagSet.StringVar(&etcdAdmConfig.Version, "version", constants.DefaultVersion, "etcd version")
	flagSet.StringVar(&etcdAdmConfig.ReleaseURL, "release-url", constants.DefaultReleaseURL, "URL used to download etcd")
	flagSet.StringVar(&etcdAdmConfig.CertificatesDir, "certs-dir", constants.DefaultCertificateDir, "certificates directory")
	flagSet.StringVar(&etcdAdmConfig.InstallDir, "install-dir", constants.DefaultInstallDir, "install directory")
	flagSet.StringSliceVar(&etcdAdmConfig.ServerCertSANs, "server-cert-extra-sans", etcdAdmConfig.ServerCertSANs, "optional extra Subject Alternative Names for the etcd server signing cert, can be multiple comma separated DNS names or IPs")
	flagSet.StringArrayVar(&etcdAdmConfig.EtcdDiskPriorities, "disk-priorities", constants.DefaultEtcdDiskPriorities, "Setting etcd disk priority")
	flagSet.IntVar(&etcdAdmConfig.GOMAXPROCS, "max-procs", runtime.NumCPU(), "Setting etcd maximal OS threads")
}
