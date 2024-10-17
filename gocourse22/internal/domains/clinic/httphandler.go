package clinic

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"

	chttp "prjctr.com/gocourse22/internal/interface/http"
)

func NewClinicHandler(inj *do.Injector) (*ClinicHandler, error) {
	return &ClinicHandler{
		s: do.MustInvoke[*Service](inj),
	}, nil
}

type ClinicHandler struct {
	s *Service
}

func (h *ClinicHandler) RegisterRoutes(g *gin.Engine) {
	api := g.Group("/api")
	{
		group := api.Group("/clinic")
		{
			group.GET("",
				h.viewAll,
			)

			group.POST("",
				h.create,
			)

			group.GET("/:clinic_id",
				h.view,
			)

			group.PUT("/:clinic_id",
				h.edit,
			)

			group.DELETE("/:clinic_id",
				h.delete,
			)

			group.GET("/:clinic_id/config",
				h.config,
			)

			group.GET("/visits/",
				h.groupVisits,
			)
		}
	}
}

func (h *ClinicHandler) groupVisits(c *gin.Context) {
	res := h.s.GroupPatientsVisits()

	chttp.NewResponse(c, res)
}

func (h *ClinicHandler) viewAll(c *gin.Context) {
	res, err := h.s.GetAll(c.Request.Context())
	if err != nil {
		log.Println(err)
		c.Status(http.StatusNoContent)
		return
	}

	chttp.NewResponse(c, res)
}

func (h *ClinicHandler) view(c *gin.Context) {
	// do view

	chttp.NewResponse(c, []byte{})
}

func (h *ClinicHandler) create(c *gin.Context) {
	// do create

	chttp.NewResponse(c, []byte{}, chttp.WithStatusCode(http.StatusCreated))
}

func (h *ClinicHandler) edit(c *gin.Context) {
	// do edit

	chttp.NewResponse(c, []byte{})
}

func (h *ClinicHandler) delete(c *gin.Context) {
	if err := h.s.DeleteClinic(); err != nil {
		fmt.Println(err)
		c.Status(http.StatusNoContent)
		return
	}

	c.Status(http.StatusOK)
}

func (h *ClinicHandler) config(c *gin.Context) {
	// do config creation

	c.FileAttachment("/tmp/", "tmp.zip")
}
