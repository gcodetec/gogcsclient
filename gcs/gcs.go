// gcs is a package to send files to Google Cloud Storage
package gcs

import (
	"fmt"
	"io"
	"os"

	"cloud.google.com/go/storage"

	"golang.org/x/net/context"
)

// Upload a file to Google Cloud Storage
func Upload(origin string, dest string, bucketName string) (msg string, err error) {
	fmt.Println("Uploading " + origin + " to " + dest + "...")
	r, err := os.Open(origin)
	if err != nil {
		return "File not found", err
	}
	defer r.Close()
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return "Failed to create client: %v", err
	}

	bkt := client.Bucket(bucketName)

	obj := bkt.Object(dest)

	w := obj.NewWriter(ctx)

	if _, err := io.Copy(w, r); err != nil {
		return "", err
	}
	if err := w.Close(); err != nil {
		return "", err
	}

	objAttrs, err := obj.Attrs(ctx)

	if err != nil {
		return "", err
	}
	return "File " + origin + " uploaded to " + objAttrs.Name, nil
}
