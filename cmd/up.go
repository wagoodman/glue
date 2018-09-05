// Copyright © 2018 Alex Goodman
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wagoodman/stitch/core"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	DisableFlagParsing: true,
	Run: composeUp,
}

func init() {
	rootCmd.AddCommand(upCmd)
}

func composeUp(cmd *cobra.Command, args []string) {
	_, project, err := core.LoadCurrentProject()
	core.Check(err, "unable to load current project")

	// var paths []string
	// for _, repository := range project.Repositories {
	// 	path, found := repository.GetComposePath()
	// 	if found {
	// 		paths = append(paths, path)
	// 	}
	// }
	//
	// project.Compose.Merge(paths)

	project.Compose.Merge(project.Repositories)

	// project.Compose.Up(args)
}