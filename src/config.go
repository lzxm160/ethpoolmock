package proxy

import (
	"api"
	"payouts"
	"policy"
	"storage"
)

type ProxyConf struct {
	Coin    string `json:"coin"`
	Name    string `json:"name"`
	Threads int    `json:"threads"`

	LimitHeadersSize      int    `json:"limitHeadersSize"`
	LimitBodySize         int64  `json:"limitBodySize"`
	BehindReverseProxy    bool   `json:"behindReverseProxy"`
	BlockRefreshInterval  string `json:"blockRefreshInterval"`
	Difficulty            int64  `json:"difficulty"`
	StateUpdateInterval   string `json:"stateUpdateInterval"`
	HashrateExpiration    string `json:"hashrateExpiration"`
	UpstreamCheckInterval string `json:"upstreamCheckInterval"`

	Proxy Proxy2 `json:"proxy"`

	Stratum Stratum `json:"stratum"`

	Policy policy.Config `json:"policy"`

	MaxFails    int64 `json:"maxFails"`
	HealthCheck bool  `json:"healthCheck"`

	Upstream []Upstream `json:"upstream"`

	Redis storage.Config `json:"redis"`
}

type Config struct {
	Name                  string        `json:"name"`
	Proxy                 Proxy         `json:"proxy"`
	Api                   api.ApiConfig `json:"api"`
	Upstream              []Upstream    `json:"upstream"`
	UpstreamCheckInterval string        `json:"upstreamCheckInterval"`

	Threads int `json:"threads"`

	Coin  string         `json:"coin"`
	Redis storage.Config `json:"redis"`

	BlockUnlocker payouts.UnlockerConfig `json:"unlocker"`
	Payouts       payouts.PayoutsConfig  `json:"payouts"`

	NewrelicName    string `json:"newrelicName"`
	NewrelicKey     string `json:"newrelicKey"`
	NewrelicVerbose bool   `json:"newrelicVerbose"`
	NewrelicEnabled bool   `json:"newrelicEnabled"`
}

type Proxy struct {
	Enabled bool   `json:"enabled"`
	Listen  string `json:"listen"`

	BlockRefreshInterval string `json:"blockRefreshInterval"`
	Difficulty           int64  `json:"difficulty"`
	StateUpdateInterval  string `json:"stateUpdateInterval"`
	HashrateExpiration   string `json:"hashrateExpiration"`

	Policy policy.Config `json:"policy"`

	MaxFails    int64 `json:"maxFails"`
	HealthCheck bool  `json:"healthCheck"`

	Stratum Stratum `json:"stratum"`
}

type Proxy2 struct {
	Enabled            bool   `json:"enabled"`
	Listen             string `json:"listen"`
	LimitHeadersSize   int    `json:"limitHeadersSize"`
	LimitBodySize      int64  `json:"limitBodySize"`
	BehindReverseProxy bool   `json:"behindReverseProxy"`
}

type Stratum struct {
	Enabled bool   `json:"enabled"`
	Listen  string `json:"listen"`
	Timeout string `json:"timeout"`
	MaxConn int    `json:"maxConn"`
}

type Upstream struct {
	Name    string `json:"name"`
	Url     string `json:"url"`
	Timeout string `json:"timeout"`
}
