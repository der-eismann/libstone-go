package cmd

import (
	"context"
	"embed"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"

	"github.com/der-eismann/libstone/pkg/header"
	"github.com/der-eismann/libstone/pkg/payload"
	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
	"github.com/rebuy-de/rebuy-go-sdk/v6/pkg/cmdutil"
	"github.com/rebuy-de/rebuy-go-sdk/v6/pkg/webutil"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

//go:embed templates/*
var templateFS embed.FS

func HTTPServer(ctx context.Context, cmd *cobra.Command, args []string) {
	group, ctx := errgroup.WithContext(ctx)
	templateFSSub, err := fs.Sub(templateFS, "templates")
	if err != nil {
		logrus.Fatal(err)
	}

	vh := webutil.NewViewHandler(templateFSSub)
	router := chi.NewRouter()
	router.Get("/", vh.Wrap(handleIndex))
	router.Get("/{filename}", vh.Wrap(handleStoneFile))
	group.Go(func() error {
		return webutil.ListenAndServeWithContext(
			ctx, "0.0.0.0:8080", router)
	})
	cmdutil.Must(errors.WithStack(group.Wait()))
}

func handleIndex(v *webutil.View, req *http.Request) webutil.Response {
	files := []string{}
	items, _ := os.ReadDir(".")
	for _, item := range items {
		if item.IsDir() {
			continue
		} else {
			if filepath.Ext(item.Name()) == ".stone" {
				files = append(files, item.Name())
			}
		}
	}
	return v.HTML(http.StatusOK, "index.html", files)
}

func handleStoneFile(v *webutil.View, req *http.Request) webutil.Response {
	var pos int64
	var stone Stone

	var filename = chi.URLParam(req, "filename")
	file, err := os.Open(filename)
	if err != nil {
		return v.Errorf(http.StatusInternalServerError,
			"failed to open file %s: %w", filename, err)
	}

	packageHeader, err := header.ReadHeader(io.NewSectionReader(file, 0, 32))
	if err != nil {
		return v.Errorf(http.StatusInternalServerError,
			"Failed to read package header: %w", err)
	}

	pos += 32

	for i := 0; i < int(packageHeader.Data.NumPayloads); i++ {
		payloadheader, err := payload.ReadPayloadHeader(io.NewSectionReader(file, pos, 32))
		if err != nil {
			return v.Errorf(http.StatusInternalServerError,
				"Failed to read payload header: %w", err)
		}

		pos += 32

		payloadReader, err := getCompressionReader(file, payloadheader.Compression, pos, int64(payloadheader.StoredSize))

		pos += int64(payloadheader.StoredSize)

		switch payloadheader.Kind {
		case payload.KindMeta:
			stone.Meta, err = payload.DecodeMetaPayload(payloadReader, int(payloadheader.NumRecords))
		//case payload.KindLayout:
		//err = payload.PrintLayoutPayload(payloadReader, int(payloadheader.NumRecords))
		//case payload.KindIndex:
		//err = payload.PrintIndexPayload(payloadReader, int(payloadheader.NumRecords))
		default:
			continue
		}
		if err != nil {
			return v.Errorf(http.StatusInternalServerError,
				"Failed to read payload: %w", err)
		}
	}
	return v.HTML(http.StatusOK, "info.html", stone)
}
