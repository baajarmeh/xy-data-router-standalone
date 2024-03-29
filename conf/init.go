package conf

import (
	"net"
	"path/filepath"

	"github.com/fufuok/utils"
	"github.com/imroc/req"
)

var (
	Debug     bool
	Version   = "v0.0.0"
	GoVersion = ""
	GitCommit = ""

	// RootPath 运行绝对路径
	RootPath = utils.ExecutableDir(true)

	// FilePath 配置文件绝对路径
	FilePath = filepath.Join(RootPath, "..", "etc")

	// ConfigFile 默认配置文件路径
	ConfigFile = filepath.Join(FilePath, ProjectName+".json")

	// LogDir 日志路径
	LogDir  = filepath.Join(RootPath, "..", "log")
	LogFile = filepath.Join(LogDir, ProjectName+".log")

	// LogDaemon 守护日志
	LogDaemon = filepath.Join(LogDir, "daemon.log")

	// Config 所有配置
	Config *tJSONConf

	// APIConfig 以接口名为键的配置
	APIConfig map[string]*TAPIConf

	// ESWhiteListConfig ES 查询接口 IP 白名单配置
	ESWhiteListConfig map[*net.IPNet]struct{}

	// ESBlackListConfig ES 上报接口 IP 黑名单配置
	ESBlackListConfig map[*net.IPNet]struct{}

	// ReqUserAgent 请求名称
	ReqUserAgent = req.Header{"User-Agent": APPName + "/" + Version}

	// ForwardTunnel 上联地址, 指定后将启动客户端, 将所有数据转交到 Tunnel Server
	ForwardTunnel = ""
)
