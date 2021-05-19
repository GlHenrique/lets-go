package model

import "errors"

//Imovel é uma struct que armazena dados de um imvóvel
type Imovel struct {
	X     int    `json:"coordenada_x"`
	Y     int    `json:"coordenada_y"`
	Nome  string `json:"nome"`
	valor int
}

//ErrValorImovelBaixo é um erro quando o valor inserido está inválido
var ErrValorImovelBaixo = errors.New(
	"o valor informado é muito baixo para um imóvel",
)

//ErrValorImovelAlto é um erro quando o valor inserido está inválido
var ErrValorImovelAlto = errors.New(
	"o valor informado é muito alto p/ um imóvel",
)

//SetValor define o valor do imóvel
func (i *Imovel) SetValor(valor int) (err error) {
	err = nil
	if valor < 50000 {
		err = ErrValorImovelBaixo
		return
	} else if valor > 600000 {
		err = ErrValorImovelAlto
		return
	}
	i.valor = valor
	return err
}

//GetValor retorna o valor do imóvel
func (i *Imovel) GetValor() int {
	return i.valor
}
