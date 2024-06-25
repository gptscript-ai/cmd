package init

import "github.com/gptscript-ai/cmd/pkg/logserver"

func init() {
	go logserver.StartServerWithDefaults()
}
