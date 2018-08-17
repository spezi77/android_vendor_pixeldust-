package android

// Global config used by Pixeldust soong additions
var PixeldustConfig = struct {
	// List of packages that are permitted
	// for java source overlays.
	JavaSourceOverlayModuleWhitelist []string
}{
	// JavaSourceOverlayModuleWhitelist
	[]string{
		"org.pixeldust.hardware",
	},
}
