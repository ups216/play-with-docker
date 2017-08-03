package types

import "context"

type Instance struct {
	Image        string          `json:"image" bson:"image"`
	Name         string          `json:"name" bson:"name"`
	Hostname     string          `json:"hostname" bson:"hostname"`
	IP           string          `json:"ip" bson:"ip"`
	ServerCert   []byte          `json:"server_cert" bson:"server_cert"`
	ServerKey    []byte          `json:"server_key" bson:"server_key"`
	CACert       []byte          `json:"ca_cert" bson:"ca_cert"`
	Cert         []byte          `json:"cert" bson:"cert"`
	Key          []byte          `json:"key" bson:"key"`
	IsDockerHost bool            `json:"is_docker_host" bson:"is_docker_host"`
	SessionId    string          `json:"session_id" bson:"session_id"`
	ProxyHost    string          `json:"proxy_host" bson:"proxy_host"`
	SessionHost  string          `json:"session_host" bson:"session_host"`
	Type         string          `json:"type" bson:"type"`
	Session      *Session        `json:"-" bson:"-"`
	ctx          context.Context `json:"-" bson:"-"`
}

type InstanceConfig struct {
	ImageName  string
	Hostname   string
	ServerCert []byte
	ServerKey  []byte
	CACert     []byte
	Cert       []byte
	Key        []byte
	Host       string
	Type       string
}