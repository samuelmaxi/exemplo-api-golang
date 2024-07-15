package controllers

import (
	"api_exemplo/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetListaUsuarios funcao que faz uma requisicao GET para um endpoint externo e cria um JSON
func ListaUsuarios(ctx *gin.Context) {
	req, err := http.Get("https://reqres.in/api/users?page=2")
	if err != nil {
		fmt.Println(err)
	}

	defer req.Body.Close()

	body, _ := io.ReadAll(req.Body)

	var usuarios models.UsuariosResponse
	err = json.Unmarshal(body, &usuarios)
	if err != nil {
		fmt.Println(err)
	}

	ctx.JSON(http.StatusOK, usuarios)
}

func NovoUsuario(ctx *gin.Context) {
	req, err := http.Post("https://reqres.in/api/users", "application/json", ctx.Request.Body)
	if err != nil {
		fmt.Println(err)
	}

	defer req.Body.Close()

	body, _ := io.ReadAll(req.Body)

	var novoUsuario models.NovoUsuarioResponse
	err = json.Unmarshal(body, &novoUsuario)
	if err != nil {
		fmt.Println(err)
	}

	ctx.JSON(http.StatusOK, novoUsuario)
}

func AtualizarUsuario(ctx *gin.Context) {
	id := ctx.Param("id") // id do usuario que sera editado
	req, err := http.NewRequest(http.MethodPut, "https://reqres.in/api/users/"+id, ctx.Request.Body)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Content-Type", "application/json") // Tipo do conteúdo enviado na requisição
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var usuarioAtualizado models.UsuarioEditadoReponse
	err = json.Unmarshal(body, &usuarioAtualizado)
	if err != nil {
		fmt.Println(err)
	}

	ctx.JSON(http.StatusOK, usuarioAtualizado)
}

func DeletaUsuario(ctx *gin.Context) {
	id := ctx.Param("id")
	req, err := http.NewRequest(http.MethodDelete, "https://reqres.in/api/users/"+id, nil)
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	ctx.JSON(http.StatusOK, gin.H{
		"messagem": "O cliente foi deletado com suceso!",
	})
}
