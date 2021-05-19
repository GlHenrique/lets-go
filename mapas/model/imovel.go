package model

//Imovel é uma struct que armazena dados de um imvóvel
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

//SetValor define o valor do imóvel
func (i *Imovel) SetValor(valor int) {
	i.valor = valor
}

//GetValor retorna o valor do imóvel
func (i *Imovel) GetValor() int {
	return i.valor
}
