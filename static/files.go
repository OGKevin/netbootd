package static

import "embed"

//go:embed ipxe.efi undionly.kpxe ipxe_rpi4.efi
var Files embed.FS
