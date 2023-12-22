// @Time    : 11/22/2023 9:28 PM
// @Author  : Cillian
// @Email   : cilliandevops@gmail.com
// Website  : www.cillian.website
// Have a good day!

package config

import "os"

const (
	envKeyKubeconfigPath = "KUBECONFIG_PATH"
	ListenAddr           = "127.0.0.1:9091"
	Kubeconfig           = "E:\\Cillian\\VirtualMachine\\config"
	PodLogTailLine       = 2000
	AdminUser            = "admin"
	AdminPwd             = "123456"
)

func GetKubeconfigPath() string {
	// get kubeconfig path from env
	kubeconfigPath := os.Getenv(envKeyKubeconfigPath)
	if kubeconfigPath == "" {
		kubeconfigPath = Kubeconfig
	}

	if kubeconfigPath == "" {
		panic("kubeconfig path is empty")
	}

	return kubeconfigPath
}
