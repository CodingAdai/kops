package service

type kubeconfig struct {
	Path    string
	Content string
}

var KubeConfig kubeconfig = kubeconfig{
	Path:    "",
	Content: "",
}

func (config *kubeconfig) SetKubeConfigPath(path string) {
	config.Path = path
}

func (config *kubeconfig) SetKubeConfigContent(content string) (bool, error) {
	config.Content = content
	return true, nil
}
