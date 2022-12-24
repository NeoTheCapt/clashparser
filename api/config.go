package api

type ParserConfig struct {
	Tokens     []string `yaml:"tokens"`
	Rules4Win  []string `yaml:"rules4win"`
	Rules4Mac  []string `yaml:"rules4mac"`
	Rules4Ios  []string `yaml:"rules4ios"`
	Socks5IP   string   `yaml:"socks5ip"`
	Socks5Port int      `yaml:"socks5port"`
	Socks5User string   `yaml:"socks5user"`
	Socks5Pass string   `yaml:"socks5pass"`
}

type ClashConfig struct {
	Port               int    `yaml:"port"`
	SocksPort          int    `yaml:"socks-port"`
	MixedPort          int    `yaml:"mixed-port"`
	AllowLan           bool   `yaml:"allow-lan"`
	Mode               string `yaml:"mode"`
	LogLevel           string `yaml:"log-level"`
	ExternalController string `yaml:"external-controller"`
	Experimental       struct {
		IgnoreResolveFail bool `yaml:"ignore-resolve-fail"`
	} `yaml:"experimental"`
	DNS struct {
		Enable            bool     `yaml:"enable"`
		Listen            string   `yaml:"listen"`
		EnhancedMode      string   `yaml:"enhanced-mode"`
		FakeIPRange       string   `yaml:"fake-ip-range"`
		DefaultNameserver []string `yaml:"default-nameserver"`
		Nameserver        []string `yaml:"nameserver"`
		FakeIPFilter      []string `yaml:"fake-ip-filter"`
	} `yaml:"dns"`
	CfwLatencyTimeout    int    `yaml:"cfw-latency-timeout"`
	CfwLatencyURL        string `yaml:"cfw-latency-url"`
	CfwConnBreakStrategy bool   `yaml:"cfw-conn-break-strategy"`
	ClashForAndroid      struct {
		UISubtitlePattern string `yaml:"ui-subtitle-pattern"`
	} `yaml:"clash-for-android"`
	Proxies     []Proxy      `yaml:"proxies"`
	ProxyGroups []ProxyGroup `yaml:"proxy-groups"`
	Rules       []string     `yaml:"rules"`
	Script      struct {
		Code string `yaml:"code"`
	} `yaml:"script"`
	RuleProviders  map[string]map[string]any `yaml:"rule-providers"`
	ProxyProviders struct {
	} `yaml:"proxy-providers"`
}
type Proxy struct {
	Server         string   `yaml:"server"`
	Port           int      `yaml:"port"`
	Name           string   `yaml:"name"`
	Type           string   `yaml:"type"`
	Password       string   `yaml:"password"`
	UDP            bool     `yaml:"udp,omitempty"`
	Alpn           []string `yaml:"alpn,omitempty"`
	SkipCertVerify bool     `yaml:"skip-cert-verify,omitempty"`
	Cipher         string   `yaml:"cipher,omitempty"`
	Username       string   `yaml:"username,omitempty"`
	TLS            bool     `yaml:"tls,omitempty"`
}

type ProxyGroup struct {
	Name     string   `yaml:"name"`
	Type     string   `yaml:"type"`
	Proxies  []string `yaml:"proxies"`
	URL      string   `yaml:"url,omitempty"`
	Interval string   `yaml:"interval,omitempty"`
}
