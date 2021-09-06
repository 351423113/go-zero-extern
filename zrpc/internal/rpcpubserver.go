package internal

import (
	"os"
	"strings"

	"github.com/lukebull/go-zero-extern/core/discov"
	"github.com/lukebull/go-zero-extern/core/netx"
)

const (
	allEths  = "0.0.0.0"
	envPodIp = "POD_IP"
)

func NewRpcPubServerExtern(etcd *discov.EtcdConf, listenOn string, opts ...ServerOption) (Server, error) {
	registerEtcd := func() error {
		pubListenOn := figureOutListenOn(listenOn)
		pubClient := discov.NewPublisher(etcd.Hosts, etcd.Key, pubListenOn, etcd.Tls, etcd.Cafile, etcd.Certfile, etcd.Keyfile)
		return pubClient.KeepAlive()
	}

	server := keepAliveServer{
		registerEtcd: registerEtcd,
		Server:       NewRpcServer(listenOn, opts...),
	}

	return server, nil
}

// NewRpcPubServer returns a Server.
func NewRpcPubServer(etcdEndpoints []string, etcdKey, listenOn string, opts ...ServerOption) (Server, error) {
	registerEtcd := func() error {
		pubListenOn := figureOutListenOn(listenOn)
		pubClient := discov.NewPublisher(etcdEndpoints, etcdKey, pubListenOn, false, "", "", "")
		return pubClient.KeepAlive()
	}
	server := keepAliveServer{
		registerEtcd: registerEtcd,
		Server:       NewRpcServer(listenOn, opts...),
	}

	return server, nil
}

type keepAliveServer struct {
	registerEtcd func() error
	Server
}

func (ags keepAliveServer) Start(fn RegisterFn) error {
	if err := ags.registerEtcd(); err != nil {
		return err
	}

	return ags.Server.Start(fn)
}

func figureOutListenOn(listenOn string) string {
	fields := strings.Split(listenOn, ":")
	if len(fields) == 0 {
		return listenOn
	}

	host := fields[0]
	if len(host) > 0 && host != allEths {
		return listenOn
	}

	ip := os.Getenv(envPodIp)
	if len(ip) == 0 {
		ip = netx.InternalIp()
	}
	if len(ip) == 0 {
		return listenOn
	}

	return strings.Join(append([]string{ip}, fields[1:]...), ":")
}
