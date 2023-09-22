package enderecos

import (
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}
// Rode o comando go test -v para ver o resultado dos testes com mais detalhes
// Rode o comando go ./... para rodar todos os testes do projeto
// Rode em paralelo utilizando a funcao t.Parallel() dentro do teste
// Rode o comando go test -cover para ver a cobertura de testes do projeto
// Rode o comando go test -coverprofile cobertura.txt para ver a cobertura de testes do projeto e gerar um arquivo com o resultado
// E depois go tool cover --html=resultado.txt para ver o resultado em html
func TestTipoDeEndereco(t *testing.T) { // Pacote padrão de testes
	t.Parallel() // Roda em paralelo
	cenarioDeTestes := []cenarioDeTeste {
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praça das Rosas", "Tipo Inválido"},
		{"Estrada Qualquer", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenarioDeTestes {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				retornoRecebido,
				cenario.retornoEsperado)
			
		}
	}

	

}