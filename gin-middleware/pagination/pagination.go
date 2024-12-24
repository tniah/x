package pagination

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tniah/iam-domain/errors"
	httperrors "github.com/tniah/x/errors/http"
	"net/http"
	"strconv"
)

func Paginator(customOpts ...Option) gin.HandlerFunc {
	opts := options{
		PageText:        PageText,
		PageSizeText:    PageSizeText,
		DefaultPage:     DefaultPage,
		DefaultPageSize: DefaultPageSize,
		MinPage:         MinPage,
		MinPageSize:     MinPageSize,
		MaxPageSize:     MaxPageSize,
		ErrReason:       ErrReason,
		ErrMsg:          ErrMsg,
	}
	for _, opt := range customOpts {
		opt(&opts)
	}

	return func(ctx *gin.Context) {
		p := &paginator{
			opts: opts,
			ctx:  ctx,
		}

		page, err := p.getPageFromQuery()
		if err != nil {
			p.abortWithError(p.opts.PageText, err)
			return
		}

		if err := p.validatePage(page); err != nil {
			p.abortWithError(p.opts.PageText, err)
			return
		}

		pageSize, err := p.getPageSizeFromQuery()
		if err != nil {
			p.abortWithError(p.opts.PageSizeText, err)
			return
		}

		if err := p.validatePageSize(pageSize); err != nil {
			p.abortWithError(p.opts.PageSizeText, err)
		}

		p.setPageAndPageSize(page, pageSize)
		p.next()
	}
}

type paginator struct {
	ctx  *gin.Context
	opts options
}

func (p *paginator) abortWithError(field string, err error) {
	he := httperrors.New(http.StatusBadRequest, p.opts.ErrMsg, p.opts.ErrReason)
	he.WithDetails(&domainerrors.InvalidArgument{
		Fields: []*domainerrors.FieldViolation{{
			Field:       field,
			Description: err.Error(),
		}},
	})
	p.ctx.Abort()
	_ = p.ctx.Error(he)
}

func (p *paginator) getPageFromQuery() (int, error) {
	return p.getIntValueWithDefault(p.opts.PageText, strconv.Itoa(p.opts.DefaultPage))
}

func (p *paginator) getPageSizeFromQuery() (int, error) {
	return p.getIntValueWithDefault(p.opts.PageSizeText, strconv.Itoa(p.opts.DefaultPageSize))
}

func (p *paginator) getIntValueWithDefault(key string, defaultValue string) (int, error) {
	valueStr := p.ctx.DefaultQuery(key, defaultValue)
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, fmt.Errorf("'%s' parameter must be an integer", key)
	}

	return value, nil
}

func (p *paginator) validatePage(page int) error {
	if page < p.opts.MinPage {
		return fmt.Errorf("%s must be greater or equal than %d", p.opts.PageText, p.opts.MinPage)
	}

	return nil
}

func (p *paginator) validatePageSize(pageSize int) error {
	if pageSize < p.opts.MinPageSize || pageSize > p.opts.MaxPageSize {
		return fmt.Errorf("%s must be between %d and %d", p.opts.PageSizeText, p.opts.MinPageSize, p.opts.MaxPageSize)
	}

	return nil
}

func (p *paginator) setPageAndPageSize(page, pageSize int) {
	p.ctx.Set(p.opts.PageText, page)
	p.ctx.Set(p.opts.PageSizeText, pageSize)
}

func (p *paginator) next() {
	p.ctx.Next()
}
