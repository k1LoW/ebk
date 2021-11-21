/*
Copyright © 2021 Ken'ichiro Oyama <k1lowxb@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/gabriel-vasile/mimetype"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/validate"
	"github.com/spf13/cobra"
	"github.com/taylorskalyo/goreader/epub"
)

var (
	withPath  bool
	delimiter string
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list directory ebook contents",
	Long:  `list directory ebook contents.`,
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := ""
		if len(args) == 0 {
			dir = "."
		} else {
			dir = args[0]
		}
		if delimiter == "\\t" {
			// tab
			delimiter = ""
		}

		fi, err := os.Stat(dir)
		if err != nil {
			return err
		}

		files := []string{}
		if fi.IsDir() {
			de, err := os.ReadDir(dir)
			if err != nil {
				return err
			}
			for _, d := range de {
				if d.IsDir() {
					continue
				}
				files = append(files, filepath.Join(dir, d.Name()))
			}
		} else {
			files = []string{dir}
		}

		tc := color.New(color.FgCyan).SprintFunc()
		ecf := color.New(color.FgRed).SprintfFunc()

		for _, p := range files {
			pp := filepath.Base(p)
			if withPath {
				pp = p
			}

			mtype, err := mimetype.DetectFile(p)
			if err != nil {
				cmd.Print(ecf("MIME type detect error \"%s\"%s%s\n", strings.ReplaceAll(err.Error(), "\n", ". "), delimiter, pp))
				continue
			}
			t := mtype.String()
			title := "???"
			switch t {
			case "application/pdf":
				f, err := os.Open(p)
				if err != nil {
					return err
				}
				pctx, err := api.ReadContext(f, pdfcpu.NewDefaultConfiguration())
				if err != nil {
					_ = f.Close()
					cmd.Print(ecf("PDF parse error \"%s\"%s%s\n", strings.ReplaceAll(err.Error(), "\n", ". "), delimiter, pp))
					continue
				}
				if err = validate.XRefTable(pctx.XRefTable); err != nil {
					_ = f.Close()
					cmd.Print(ecf("PDF parse error \"%s\"%s%s\n", strings.ReplaceAll(err.Error(), "\n", ". "), delimiter, pp))
					continue
				}
				if pctx.Title != "" {
					title = pctx.Title
				}
				_ = f.Close()
			case "application/epub+zip":
				rc, err := epub.OpenReader(p)
				if err != nil {
					cmd.Print(ecf("EPUB parse error \"%s\"%s%s\n", strings.ReplaceAll(err.Error(), "\n", ". "), delimiter, pp))
					continue
				}
				book := rc.Rootfiles[0]

				if book.Title != "" {
					title = book.Title
				}
				rc.Close()
			default:
				continue
			}

			cmd.Printf("%s%s%s\n", tc(title), delimiter, pp)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
	lsCmd.Flags().BoolVarP(&withPath, "with-path", "", false, "show with ebook file path")
	lsCmd.Flags().StringVarP(&delimiter, "delimiter", "d", ":", "delimiter")
}
