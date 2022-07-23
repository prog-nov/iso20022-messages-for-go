//go:build ignore
// +build ignore

// Copyright 2022 Tamás Gulácsi. All rights reserved.
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	mrand "math/rand"
	"mime"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"aqwari.net/xml/xsdgen"
	"github.com/google/renameio/v2"
	"golang.org/x/net/html"
	"golang.org/x/sync/errgroup"
)

func main() {
	if err := Main(); err != nil {
		log.Fatalf("ERROR: %+v", err)
	}
}

const DefaultSearchURL = "https://www.iso20022.org/iso-20022-message-definitions"

func GetXSDURLs(ctx context.Context, searchURL string) ([]string, error) {
	if searchURL == "" {
		searchURL = DefaultSearchURL
	}
	base, err := url.Parse(searchURL)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", searchURL, err)
	}
	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", searchURL, err)
	}
	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", searchURL, err)
	}
	defer resp.Body.Close()
	body := io.ReadCloser(resp.Body)
	if resp.StatusCode > 399 {
		var buf strings.Builder
		_, _ = io.Copy(&buf, resp.Body)
		resp.Body.Close()

		if body, err = os.Open(filepath.Join("xsd", "iso-20022-message-definitions")); err != nil {
			return nil, fmt.Errorf("%q: %s: %s", searchURL, resp.Status, buf.String())
		}
	}

	var results []string
	z := html.NewTokenizer(body)
Loop:
	for {
		switch tt := z.Next(); tt {
		case html.ErrorToken:
			err := z.Err()
			if errors.Is(err, io.EOF) {
				break Loop
			}
			return results, err
		case html.StartTagToken:
			if b, hasAttr := z.TagName(); hasAttr && bytes.Equal(b, []byte("a")) {
				var href, aria string
				for {
					k, v, more := z.TagAttr()
					if bytes.Equal(k, []byte("href")) {
						href = string(v)
					} else if bytes.Equal(k, []byte("aria-label")) {
						aria = string(v)
					}
					if !more {
						break
					}
				}
				if aria == "XSD" && href != "" && strings.HasPrefix(href, "/message/") && strings.HasSuffix(href, "/download") {
					U, err := url.Parse(href)
					if err != nil {
						return results, fmt.Errorf("%s: %w", href, err)
					}
					log.Printf("%q: %s -> %s", href, U, base.ResolveReference(U))
					results = append(results, base.ResolveReference(U).String())
				} else {
					log.Println("aria:", aria, "href:", href)
				}
			}
		default:
		}
		select {
		case <-ctx.Done():
			return results, nil
		default:
		}
	}
	if len(results) == 0 {
		return results, errors.New("not found")
	}
	return results, nil
}

func DownloadXSDs(ctx context.Context, destDir, searchURL string) error {
	_ = os.MkdirAll(destDir, 0755)
	urls, err := GetXSDURLs(ctx, searchURL)
	if err != nil {
		return err
	}
	mrand.Shuffle(len(urls), func(i, j int) { urls[i], urls[j] = urls[j], urls[i] })
	concCh := make(chan struct{}, 2)
	grp, grpCtx := errgroup.WithContext(ctx)
	for _, u := range urls {
		concCh <- struct{}{}
		grp.Go(func() error {
			defer func() { <-concCh }()
			fn, rc, err := DownloadFile(grpCtx, u)
			if err != nil || fn == "" {
				return err
			}
			log.Printf("Saving %q to %q", u, filepath.Join(destDir, fn))
			fh, err := renameio.NewPendingFile(filepath.Join(destDir, fn), renameio.WithTempDir(destDir))
			if err != nil {
				return err
			}
			defer fh.Cleanup()
			if _, err = io.Copy(fh, rc); err != nil {
				return err
			}
			return fh.CloseAtomicallyReplace()
		})
	}
	return grp.Wait()
}

var ErrStatusForbidden = errors.New("403 Forbidden")

func DownloadFile(ctx context.Context, dlURL string) (string, io.ReadCloser, error) {
	req, err := http.NewRequest("GET", dlURL, nil)
	if err != nil {
		return "", nil, fmt.Errorf("%s: %w", dlURL, err)
	}
	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return "", nil, fmt.Errorf("%s: %w", dlURL, err)
	}
	if resp.StatusCode > 399 {
		if resp.Body != nil {
			resp.Body.Close()
		}
		if resp.StatusCode == 401 || resp.StatusCode == 403 {
			return "", nil, ErrStatusForbidden
		}
		return "", nil, errors.New(resp.Status)
	}

	cd := resp.Header.Get("Content-Disposition")
	var filename string
	if _, params, err := mime.ParseMediaType(cd); err == nil {
		filename = params["filename"]
	}
	return filename, resp.Body, nil
}

func Main() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	_ = os.MkdirAll("xsd", 0755)

	files, err := filepath.Glob("xsd/*.xsd")
	if err != nil {
		return err
	}

	if len(files) < 491 {
		if err := DownloadXSDs(ctx, "xsd", ""); err != nil {
			return err
		}
		if files, err = filepath.Glob("xsd/*.xsd"); err != nil {
			return err
		}
	}
	mrand.Shuffle(len(files), func(i, j int) { files[i], files[j] = files[j], files[i] })

	var cfg xsdgen.Config
	cfg.Option(
		xsdgen.IgnoreAttributes("id", "href", "offset"),
		xsdgen.IgnoreElements("comment"),
		xsdgen.Replace("_", ""),
		xsdgen.HandleSOAPArrayType(),
		xsdgen.SOAPArrayAsSlice(),
	)

	concCh := make(chan struct{}, 16)
	grp, grpCtx := errgroup.WithContext(ctx)
	for _, f := range files {
		concCh <- struct{}{}
		cfg := cfg
		grp.Go(func() error {
			defer func() { <-concCh }()

			// File and folder name:
			bn := filepath.Base(f)
			nm := strings.ReplaceAll(bn, ".", "_")
			nm = strings.TrimRight(nm, "_xsd")
			pt := filepath.Join(nm, bn+".go")

			_ = os.Mkdir(nm, 0755)
			cfg.Option(xsdgen.PackageName(fmt.Sprintf("iso20022_%s", nm)))

			if err := grpCtx.Err(); err != nil {
				return nil
			}
			if err := cfg.GenCLI("-o", pt, f); err != nil {
				log.Println(f, err)
			}
			return nil
		})
	}
	return grp.Wait()
}
