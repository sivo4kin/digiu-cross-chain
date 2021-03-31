package bridge

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"

	//"github.com/gorilla/mux"
	"github.com/linkpoolio/bridges"
)

type EthHealth struct {
	CLient *ethclient.Client
	Name   string
	Url    string
	Lambda bool
	Func   func() (*Output, error)
}

type Output struct {
	ChainId  string `json:"chainId"`
	BlockNum string `json:"blockNum"`
}

func NewEthHealth(ethClient *ethclient.Client, name, url string, f func() (*Output, error)) (a *EthHealth, err error) {
	a = &EthHealth{}
	a.CLient = ethClient
	a.Name = name
	a.Url = url
	a.Func = f
	return
}

func (ap *EthHealth) Opts() *bridges.Opts {
	return &bridges.Opts{
		Name:   ap.Name,
		Lambda: ap.Lambda,
		Path:   "/" + ap.Url,
	}
}

func (ap *EthHealth) Run(helper *bridges.Helper) (interface{}, error) {

	res, err := ap.Func()

	if err != nil {
		return nil, err
	}
	logrus.Print(res)
	return res, nil

}