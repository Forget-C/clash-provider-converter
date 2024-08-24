package clash

import (
	"gopkg.in/yaml.v3"
)

type Clash struct {
	Proxies        []*Proxies                `yaml:"proxies,omitempty"`
	ProxyGroup     []*ProxyGroup             `yaml:"proxy-groups,omitempty"`
	RuleProviders  map[string]*RuleProvider  `yaml:"rule-providers,omitempty"`
	ProxyProviders map[string]*ProxyProvider `yaml:"proxy-providers,omitempty"`
	Rules          []string                  `yaml:"rules,omitempty"`
}

type RuleProvider struct {
	Type     string `yaml:"type,omitempty"`
	Behavior string `yaml:"behavior,omitempty"`
	Url      string `yaml:"url,omitempty"`
	Interval MyInt  `yaml:"interval,omitempty"`
	Path     string `yaml:"path,omitempty"`
}

type ProxyProvider struct {
	Type        string `yaml:"type,omitempty"`
	Path        string `yaml:"path,omitempty"`
	HealthCheck struct {
		Enable   MyBool `yaml:"enable,omitempty"`
		Interval MyInt  `yaml:"interval,omitempty"`
		URL      string `yaml:"url,omitempty"`
	}
}

type Proxies struct {
	Name                 string            `yaml:"name,omitempty"`
	Type                 string            `yaml:"type,omitempty"`
	Server               string            `yaml:"server,omitempty"`
	Port                 string            `yaml:"port,omitempty"`
	Cipher               string            `yaml:"cipher,omitempty"`
	Uuid                 string            `yaml:"uuid,omitempty"`
	AlterId              MyInt             `yaml:"alterId,omitempty"`
	Udp                  MyBool            `yaml:"udp,omitempty"`
	Tls                  MyBool            `yaml:"tls,omitempty"`
	SkipCertVerify       MyBool            `yaml:"skip-cert-verify,omitempty"`
	Servername           string            `yaml:"servername,omitempty"`
	Network              string            `yaml:"network,omitempty"`
	WsOpts               wsOpts            `yaml:"ws-opts,omitempty"`
	WsHeaders            map[string]string `yaml:"ws-headers,omitempty"`
	H2Opts               h2Opts            `yaml:"h2-opts,omitempty"`
	HTTPOpts             hTTPOpts          `yaml:"http-opts,omitempty"`
	GrpcOpts             grpcOpts          `yaml:"grpc-opts,omitempty"`
	Username             string            `yaml:"username,omitempty"`
	Password             string            `yaml:"password,omitempty"`
	Sni                  string            `yaml:"sni,omitempty"`
	Alpn                 []string          `yaml:"alpn,omitempty"`
	Plugin               string            `yaml:"plugin,omitempty"`
	PluginOpts           yaml.Node         `yaml:"plugin-opts,omitempty"`
	Fingerprint          string            `yaml:"fingerprint,omitempty"`
	Obfs                 string            `yaml:"obfs,omitempty"`
	Protocol             string            `yaml:"protocol,omitempty"`
	ObfsParam            string            `yaml:"obfs-param,omitempty"`
	ProtocolParam        string            `yaml:"protocol-param,omitempty"`
	ClientFingerprint    string            `yaml:"client-fingerprint,omitempty"`
	Flow                 string            `yaml:"flow,omitempty"`
	PacketEncoding       string            `yaml:"packet_encoding,omitempty"`
	RealityOpts          realityOpts       `yaml:"reality-opts,omitempty"`
	AuthStr              string            `yaml:"auth-str,omitempty"`
	AuthStr1             string            `yaml:"auth_str,omitempty"`
	CaStr                string            `yaml:"ca-str,omitempty"`
	CaStr1               string            `yaml:"ca_str,omitempty"`
	DisableMtuDiscovery  MyBool            `yaml:"disable_mtu_discovery,omitempty"`
	Down                 string            `yaml:"down,omitempty"`
	FastOpen             MyBool            `yaml:"fast-open,omitempty"`
	RecvWindow           MyInt             `yaml:"recv-window,omitempty"`
	RecvWindowConn       MyInt             `yaml:"recv-window-conn,omitempty"`
	RecvWindow1          MyInt             `yaml:"recv_window,omitempty"`
	RecvWindowConn1      MyInt             `yaml:"recv_window_conn,omitempty"`
	Up                   string            `yaml:"up,omitempty"`
	Ports                string            `yaml:"ports,omitempty"`
	Smux                 smuxOpts          `yaml:"smux,omitempty"`
	UdpOverTcp           MyBool            `yaml:"udp-over-tcp,omitempty"`
	IP                   string            `yaml:"ip,omitempty"`
	IPv6                 string            `yaml:"ipv6,omitempty"`
	PublicKey            string            `yaml:"public-key,omitempty"`
	PreSharedKey         string            `yaml:"pre-shared-key,omitempty"`
	PrivateKey           string            `yaml:"private-key,omitempty"`
	Reserved             *wgReserved       `yaml:"reserved,omitempty"`
	DialerProxy          string            `yaml:"dialer-proxy,omitempty"`
	Peers                []wgPeer
	MTU                  MyInt  `yaml:"mtu,omitempty"`
	DisableSni           MyBool `yaml:"disable-sni,omitempty"`
	CongestionController string `yaml:"congestion-controller,omitempty"`
	UdpRelayMode         string `yaml:"udp-relay-mode,omitempty"`
	ReduceRtt            MyBool `yaml:"reduce-rtt,omitempty"`
	HeartbeatInterval    MyInt  `yaml:"heartbeat-interval,omitempty"`
	ObfsPassword         string `yaml:"obfs-password,omitempty"`
	Tfo                  bool   `yaml:"tfo,omitempty"`
	Mptcp                bool   `yaml:"mptcp,omitempty"`
}

type smuxOpts struct {
	Enabled        MyBool `yaml:"enabled"`
	MaxConnections MyInt  `yaml:"max-connections"`
	MaxStreams     MyInt  `yaml:"max-streams"`
	MinStreams     MyInt  `yaml:"min-streams"`
	Padding        MyBool `yaml:"padding"`
	Protocol       string `yaml:"protocol"`
}

type grpcOpts struct {
	GrpcServiceName string `yaml:"grpc-service-name"`
}

type hTTPOpts struct {
	Headers map[string][]string `yaml:"headers"`
	Method  string              `yaml:"method"`
	Path    []string            `yaml:"path"`
}

type h2Opts struct {
	Host []string `yaml:"host"`
	Path string   `yaml:"path"`
}

type wsOpts struct {
	EarlyDataHeaderName string            `yaml:"early-data-header-name"`
	Headers             map[string]string `yaml:"headers"`
	MaxEarlyData        MyInt             `yaml:"max-early-data"`
	Path                string            `yaml:"path"`
	V2rayHttpUpgrade    MyBool            `yaml:"v2ray-http-upgrade"`
}

type realityOpts struct {
	PublicKey string `yaml:"public-key"`
	ShortId   string `yaml:"short-id"`
}

type ProxyGroup struct {
	Name    string   `yaml:"name"`
	Type    string   `yaml:"type"`
	Proxies []string `yaml:"proxies"`
}
