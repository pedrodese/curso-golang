package enderecos

import "strings"

func tipoDeEndereco(endereco string) string {
	tiposValidados := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavraEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	for _, tipo := range tiposValidados {
		if tipo == primeiraPalavraEndereco {
			return primeiraPalavraEndereco
		}
	}

	return "Tipo inválido"

}
