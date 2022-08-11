package handler

import (
	"BE_WEB_BEM_Proker/domain"
	"BE_WEB_BEM_Proker/implementation/proker/db"
	"BE_WEB_BEM_Proker/utils/response"
	"net/http"
	"strconv"

	"BE_WEB_BEM_Proker/utils"

	"github.com/gin-gonic/gin"
)

var handler *handlerProker

type handlerProker struct {
	db db.DatabaseService
}

type HandlerProker interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Delete(c *gin.Context)
	Create(c *gin.Context)
	Paging(c *gin.Context)
}

func NewHandlerProker(db db.DatabaseService) HandlerProker {
	if handler == nil {
		handler = &handlerProker{
			db: db,
		}
	}
	return handler
}
func (h *handlerProker) Paging(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	dataAll, err := h.db.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Fail to get data", err.Error()))
		return
	}
	len := len(dataAll)
	if page == 0 {
		page = 1
	}
	if limit == 0 || limit > len {
		limit = len
	}
	data, err := h.db.Paging(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Fail to get data", err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.ResponseWhenSuccess("Success to get data", gin.H{
		"data":        data,
		"banyak_data": len,
	}))
}

func (h *handlerProker) GetAll(c *gin.Context) {
	datas, err := h.db.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.ResponseWhenFail("Fail to get proker", err.Error()))
	}
	c.JSON(http.StatusOK, response.ResponseWhenSuccess("Success get proker", datas))
}

func (h handlerProker) GetByID(c *gin.Context) {
	var idpkr idProker
	if err := c.BindUri(&idpkr); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Errorr when bind URI", err.Error()))
		return
	}
	data, err := h.db.GetByID(idpkr.Id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, response.ResponseWhenFail("Fail to connect database", err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.ResponseWhenSuccess("Success", data))
}

func (h *handlerProker) Create(c *gin.Context) {
	idAdmin := c.MustGet("id").(uint)
	namaProker := c.Request.FormValue("nama_proker")
	waktuTerlaksana := c.Request.FormValue("waktu_terlaksana")
	deskripsi := c.Request.FormValue("deskripsi")
	penanggungJawab := c.Request.FormValue("penanggung_jawab")
	kontakPJ := c.Request.FormValue("kontak_pj")
	fileCover, err := c.FormFile("file_cover")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Errorr when bind file cover", err.Error()))
		return
	}
	fileKegiatan, err := c.FormFile("file_kegiatan_satu")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Errorr when bind file keg satu", err.Error()))
		return
	}
	fileKegiatanDua, err := c.FormFile("file_kegiatan_dua")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Errorr when bind file keg dua", err.Error()))
		return
	}
	fileKegiatanTiga, err := c.FormFile("file_kegiatan_tiga")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Error when get file keg tiga", err.Error()))
		return
	}
	if namaProker == "" || waktuTerlaksana == "" || deskripsi == "" || penanggungJawab == "" || kontakPJ == "" {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Error when get form", "Form tidak boleh kosong"))
		return
	}
	if fileCover == nil || fileKegiatan == nil || fileKegiatanDua == nil || fileKegiatanTiga == nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Error when get file", "File tidak boleh kosong"))
		return
	}
	if fileCover.Size > utils.MaxFileSize || fileKegiatan.Size > utils.MaxFileSize || fileKegiatanDua.Size > utils.MaxFileSize || fileKegiatanTiga.Size > utils.MaxFileSize {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Error when get file", "File cover tidak boleh lebih dari 5 MB"))
		return
	}
	linkCover, err := utils.UploadImage(fileCover)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Error when upload file", err.Error()))
		return
	}
	linkKegiatan, err := utils.UploadImage(fileKegiatan)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Error when upload file", err.Error()))
		return
	}
	linkKegiatanDua, err := utils.UploadImage(fileKegiatanDua)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Error when upload file", err.Error()))
		return
	}
	linkKegiatanTiga, err := utils.UploadImage(fileKegiatanTiga)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Error when upload file", err.Error()))
		return
	}
	data := domain.EntitasProker{
		NamaProker:       namaProker,
		WaktuTerlaksana:  waktuTerlaksana,
		Deskripsi:        deskripsi,
		PenanggungJawab:  penanggungJawab,
		KontakPJ:         kontakPJ,
		LinkCover:        linkCover,
		LinkKegiatan:     linkKegiatan,
		LinkKegiatanDua:  linkKegiatanDua,
		LinkKegiatanTiga: linkKegiatanTiga,
		EntitasAdminID:   uint(idAdmin),
	}
	if err := h.db.CreateProker(data); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Error when create proker", err.Error()))
		return
	}
	c.JSON(http.StatusCreated, response.ResponseWhenSuccess("Success create proker", data))

}

type idProker struct {
	Id uint `uri:"id"`
}

func (h *handlerProker) Delete(c *gin.Context) {
	var idPrkr idProker
	if err := c.ShouldBindUri(&idPrkr); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Fail to bind uri", err.Error()))
		return
	}
	if err := h.db.Delete(idPrkr.Id); err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Fail to delete data", err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.ResponseWhenSuccess("Success to delete data", nil))
}
