package dishes

import "github.com/google/wire"

// Providers for dishes
var Providers = wire.NewSet(NewStore)
