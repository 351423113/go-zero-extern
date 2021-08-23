package template

var (
	// Imports defines a import template for model in cache case
	Imports = `import (
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

	"github.com/351423113/go-zero-extern/core/stores/cache"
	"github.com/351423113/go-zero-extern/core/stores/sqlc"
	"github.com/351423113/go-zero-extern/core/stores/sqlx"
	"github.com/351423113/go-zero-extern/core/stringx"
	"github.com/351423113/go-zero-extern/tools/goctl/model/sql/builderx"
)
`
	// ImportsNoCache defines a import template for model in normal case
	ImportsNoCache = `import (
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

	"github.com/351423113/go-zero-extern/core/stores/sqlc"
	"github.com/351423113/go-zero-extern/core/stores/sqlx"
	"github.com/351423113/go-zero-extern/core/stringx"
	"github.com/351423113/go-zero-extern/tools/goctl/model/sql/builderx"
)
`
)
