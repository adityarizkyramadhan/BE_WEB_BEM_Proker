package handler

import (
	"BE_WEB_BEM_Proker/domain"
	"BE_WEB_BEM_Proker/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handlerProker struct {
	UseCase domain.ProkerService
}

func NewHandlerProker(useCase domain.ProkerService) domain.ProkerHandler {
	return handlerProker{
		UseCase: useCase,
	}
}

type getAllProker struct {
	NamaProker      string `json:"nama_proker"`
	WaktuTerlaksana string `json:"waktu_terlaksana"`
	Kementrian      string `json:"kementrian"`
}

func (h handlerProker) GetAll(c *gin.Context) {
	datas, err := h.UseCase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Fail to get proker", err.Error()))
	}
	datas = datas[len(datas)-12:]
	var dataOut []getAllProker
	for _, data := range datas {
		var dataIn getAllProker
		dataIn.NamaProker = data.NamaProker
		dataIn.WaktuTerlaksana = data.WaktuTerlaksana
		dataIn.Kementrian = data.Kementrian
		dataOut = append(dataOut, dataIn)
	}
	c.JSON(http.StatusOK, response.ResponseWhenSuccess("Success get proker", dataOut))
}

func (h handlerProker) GetByID(c *gin.Context) {
	var idpkr idProker
	if err := c.BindUri(&idpkr); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Errorr when bind URI", err.Error()))
		return
	}
	data, err := h.UseCase.GetByID(idpkr.Id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.ResponseWhenFail("Fail to connect database", err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.ResponseWhenSuccess("Success", data))
}

func (h handlerProker) Create(c *gin.Context) {
	var input domain.EntitasProkerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Error when bind JSON", err.Error()))
		return
	}
	if err := h.UseCase.Create(&domain.EntitasProker{
		NamaProker:      input.NamaProker,
		WaktuTerlaksana: input.WaktuTerlaksana,
		Deskripsi:       input.Deskripsi,
		PenanggungJawab: input.PenanggungJawab,
		Kementrian:      input.Kementrian,
		KontakPJ:        input.KontakPJ,
		LinkPDFProker:   input.LinkPDFProker,
		LinkDokumentasi: input.LinkDokumentasi,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Error when create data in database", err.Error()))
		return
	}
	c.JSON(http.StatusCreated, response.ResponseWhenSuccess("Success when add data to database", nil))
}

type idProker struct {
	Id uint `uri:"id"`
}

func (h handlerProker) Delete(c *gin.Context) {
	var idPrkr idProker
	if err := c.ShouldBindUri(&idPrkr); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Fail to bind uri", err.Error()))
		return
	}
	if err := h.UseCase.Delete(idPrkr.Id); err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Fail to delete data", err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.ResponseWhenSuccess("Success to delete data", nil))
}
