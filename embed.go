// GENERATED FILE

package jenny

import (
	"embed"
	"github.com/amitybell/piper-asset"
)

var (
	//go:embed dist.tar.zst hash.txt MODEL_CARD.txt
	fs embed.FS

	Asset = asset.Asset{Name: "jenny", FS: fs}
)
