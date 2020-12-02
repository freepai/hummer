package plugins

import (
	_ "github.com/freepai/hummer/plugins/admin"
	_ "github.com/freepai/hummer/plugins/base58"
	_ "github.com/freepai/hummer/plugins/grpc"
	_ "github.com/freepai/hummer/plugins/health"
	_ "github.com/freepai/hummer/plugins/http"
	_ "github.com/freepai/hummer/plugins/leveldb"
	_ "github.com/freepai/hummer/plugins/metrics"
	_ "github.com/freepai/hummer/plugins/snowflake"
)
