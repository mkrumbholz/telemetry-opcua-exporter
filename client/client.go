package client

import (
        "context"
        "time"

        "github.com/gopcua/opcua"
        "github.com/gopcua/opcua/ua"
        //"github.com/gopcua/opcua/id"
        "github.com/skilld-labs/telemetry-opcua-exporter/config"
        "github.com/skilld-labs/telemetry-opcua-exporter/log"
)

func NewClientFromServerConfig(c config.ServerConfig, l log.Logger, ctx context.Context) *opcua.Client {

        endpoints, err := opcua.GetEndpoints(ctx, c.Endpoint)
        if err != nil {
                l.Fatal("get endpoints failed: %v", err)
        }

        ep := opcua.SelectEndpoint(endpoints, c.SecPolicy, ua.MessageSecurityModeFromString(c.SecMode))
        if ep == nil {
                l.Fatal("Failed to find suitable endpoint")
        }

        opts := []opcua.Option{
                opcua.SecurityPolicy(c.SecPolicy),
                opcua.SecurityModeString(c.SecMode),
                opcua.CertificateFile(c.CertPath),
                opcua.PrivateKeyFile(c.KeyPath),
                opcua.AuthAnonymous(),
                opcua.SecurityFromEndpoint(ep, ua.UserTokenTypeAnonymous),
                opcua.Lifetime(1 * time.Minute),
                opcua.RequestTimeout(10 * time.Second),
                //opcua.SessionTimeout(10 * time.Second),
                //opcua.AutoReconnect(true),
                //opcua.ReconnectInterval(time.Second*20),
        }
	
        l.Info("client using config: Endpoint: %s, Security Mode: %s, %s, Authentication Mode : %s", ep.EndpointURL, ep.SecurityPolicyURI, ep.SecurityMode, c.AuthMode)

        opcua_c := opcua.NewClient(ep.EndpointURL, opts...)

        return opcua_c
}
