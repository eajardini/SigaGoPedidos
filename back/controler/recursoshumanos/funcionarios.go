package recursoshumanos

import (
	_ "log"

	bancodedados "github.com/eajardini/SigaGoPedidos/back/bancodedados"
	//modelfuncionarios  "github.com/eajardini/SigaGoPedidos/back/controler/recursoshumanos/model"
	_ "github.com/gin-gonic/gin"
)

var (
	bd bancodedados.BDCon
)
