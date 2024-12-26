package pagination

const (
	PageText         = "page"
	PageSizeText     = "page_size"
	DefaultPage      = 1
	DefaultPageSize  = 10
	MinPage          = 1
	MinPageSize      = 1
	MaxPageSize      = 100
	ErrReason        = "INVALID_ARGUMENT"
	ErrMsg           = "Invalid pagination request."
	FieldNameService = "service"
)

type options struct {
	PageText         string
	PageSizeText     string
	DefaultPage      int
	DefaultPageSize  int
	MinPage          int
	MinPageSize      int
	MaxPageSize      int
	ErrReason        string
	ErrMsg           string
	ErrInfoDomain    string
	ErrInfoService   string
	FieldNameService string
}

type Option func(*options)

func WithPageText(pageText string) Option {
	return func(opts *options) {
		opts.PageText = pageText
	}
}

func WithPageSizeText(pageSizeText string) Option {
	return func(opts *options) {
		opts.PageSizeText = pageSizeText
	}
}

func WithDefaultPage(page int) Option {
	return func(opts *options) {
		opts.DefaultPage = page
	}
}

func WithDefaultPageSize(pageSize int) Option {
	return func(opts *options) {
		opts.DefaultPageSize = pageSize
	}
}

func WithMinPage(minPage int) Option {
	return func(opts *options) {
		opts.MinPage = minPage
	}
}

func WithMinPageSize(minPageSize int) Option {
	return func(opts *options) {
		opts.MinPageSize = minPageSize
	}
}

func WithMaxPageSize(maxPageSize int) Option {
	return func(opts *options) {
		opts.MaxPageSize = maxPageSize
	}
}

func WithErrReason(reason string) Option {
	return func(opts *options) {
		opts.ErrReason = reason
	}
}

func WithErrMessage(msg string) Option {
	return func(opts *options) {
		opts.ErrMsg = msg
	}
}

func WithErrInfoDomain(domain string) Option {
	return func(opts *options) {
		opts.ErrInfoDomain = domain
	}
}

func WithErrInfoService(service string) Option {
	return func(opts *options) {
		opts.ErrInfoService = service
	}
}
