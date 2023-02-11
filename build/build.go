package main

import (
	builder "github.com/soulteary/flare/build/builder"
)

func main() {
	builder.TaskForMdi(
		"embed/assets/vendor/mdi-cheat-sheets", "internal/resources/mdi/mdi-cheat-sheets",
		"embed/assets/vendor/mdi/mdi.js", "internal/resources/mdi/icons.go",
	)
	builder.TaskForSimpleIcons("internal/resources/simpleicon")
	builder.TaskForGuideAssets("embed/assets/vendor/guide-assets", "internal/guide/guide-assets")
	builder.TaskForEditorAssets("embed/assets/vendor/editor-assets", "internal/editor/editor-assets")
	builder.TaskForStyles("internal/state/style.go")
	builder.TaskForFavicon("embed/assets/favicon.ico", "internal/resources/assets/favicon.ico")
	builder.TaskForFavicon("embed/assets/icons/favicon/android-chrome-192x192.png", "internal/resources/assets/android-chrome-192x192.png")
	builder.TaskForFavicon("embed/assets/icons/favicon/android-chrome-512x512.png", "internal/resources/assets/android-chrome-512x512.png")
	builder.TaskForFavicon("embed/assets/icons/favicon/manifest.json", "internal/resources/assets/manifest.json")
	builder.TaskForTemplates("embed/templates", "internal/resources/templates/html")
}
