package graph

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

func (r *Resolver) saveFile(upload *graphql.Upload, dir string) (string, error) {
	if upload == nil {
		return "", nil
	}

	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}

	ext := filepath.Ext(upload.Filename)
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	dstPath := filepath.Join(dir, filename)

	out, err := os.Create(dstPath)
	if err != nil {
		return "", fmt.Errorf("failed to create destination file: %w", err)
	}
	defer out.Close()

	if _, err = io.Copy(out, upload.File); err != nil {
		return "", fmt.Errorf("failed to copy file contents: %w", err)
	}

	return dstPath, nil
}

func (r *mutationResolver) saveProductImage(upload *graphql.Upload) (string, error) {
	return r.Resolver.saveFile(upload, "uploads/product")
}

func (r *mutationResolver) saveImageDetail(upload *graphql.Upload, prefix string, dir string) (string, error) {
	return r.Resolver.saveFile(upload, dir)
}

func (r *mutationResolver) saveSliderImage(upload *graphql.Upload) (string, error) {
	return r.Resolver.saveFile(upload, "uploads/slider")
}
