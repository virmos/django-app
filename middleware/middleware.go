package middleware

import (
	"myapp/data"

	"github.com/virmos/django"
)

type Middleware struct {
	App *django.Django
	Models data.Models
}