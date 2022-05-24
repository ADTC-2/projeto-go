package servidor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type produto struct {
	ID_PRODUTO uint32 `json:"id_produto"`
	DESCRICAO  string `json:"descricao"`
	QUANTIDADE int16  `json:"quantidade"`
}

func CriarProduto(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição!"))
		return
	}

	var produto produto
	if erro = json.Unmarshal(corpoRequisicao, &produto); erro != nil {
		w.Write([]byte("Erro ao converter produto para struct !"))
		return
	}
	fmt.Println(produto)
}
