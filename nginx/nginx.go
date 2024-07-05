package nginx

type NginxModule struct {
	Name      string `json:"name"`
	Version   string `json:"version"`
	Enable    bool   `json:"enable"`
	Installed bool   `json:"installed"`
}

func (nginx *NginxModule) Install() error {
	return nil
}
