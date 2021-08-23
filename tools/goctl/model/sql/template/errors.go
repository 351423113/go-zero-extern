package template

// Error defines an error template
var Error = `package {{.pkg}}

import "github.com/351423113/go-zero-extern/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound
`
