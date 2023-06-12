package rabbitmq

import "github.com/kmrhemant916/notification/utils"

type Rabbitmq struct {
    Username string
    Password string
    Host string
    Port string
    Exchange Exchange
}

type Exchange struct{
    Name string
    Kind string
}

func Setup() (*Rabbitmq) {
    var config utils.Config
    c, err:= config.ReadConf(Config)
    if err != nil {
        panic(err)
    }
    r := &Rabbitmq{
        Username: c.Rabbitmq.Username,
        Password: c.Rabbitmq.Password,
        Host: c.Rabbitmq.Host,
        Port: c.Rabbitmq.Port,
        Exchange: Exchange{
            Name: c.Rabbitmq.Exchange.Name,
            Kind: c.Rabbitmq.Exchange.Kind,
        },
    }
    return r
}