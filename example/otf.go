package main

import (
	"github.com/josscoder/bedrockpack/pack"
	"github.com/sandertv/gophertunnel/minecraft"
	"log/slog"
	"time"
)

func main() {
	log := slog.Default()

	conf := pack.OTFConfig{
		OrgName:        "Faithful-Resource-Pack",
		RepoName:       "Faithful-32x-Bedrock",
		Branch:         "bedrock-latest",
		PAT:            "",
		UpdateInterval: 1 * time.Minute,
	}
	otf := conf.New(log)
	if err := otf.Start(); err != nil {
		log.Error("failed to start otf", "error", err)
		return
	}

	listener, err := minecraft.Listen("raknet", "127.0.0.1:19132")
	if err != nil {
		panic(err)
	}

	otf.SetListener(listener)

	for {
		_, _ = listener.Accept()
	}
}
