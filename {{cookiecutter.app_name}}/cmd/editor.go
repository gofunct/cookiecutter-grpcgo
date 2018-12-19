package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	editorPort string
)
func init() {
	rootCmd.AddCommand(webEditorCmd)
	webEditorCmd.Flags().StringVarP(&editorPort, "port", "p", ":8080", "the port for the web editor to listen on")
}

var webEditorCmd = &cobra.Command{
	Use: "edit",
	Short: "edit templates",
	Run: 	func(cmd *cobra.Command, args []string) {
		r := mux.NewRouter()

		r.Handle("/", http.FileServer(http.Dir("static")))
		r.HandleFunc("/generate", Generate)
		addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
		if addr == ":" {
			addr = ":8080"
		}

		fmt.Printf("Listening on %s...\n", addr)
		h := handlers.LoggingHandler(os.Stderr, r)
		h = handlers.CompressHandler(h)
		h = handlers.RecoveryHandler()(h)
		if err := http.ListenAndServe(addr, h); err != nil {
			panic(err)
		}
	},
}

func Generate(w http.ResponseWriter, r *http.Request) {
	// read input
	decoder := json.NewDecoder(r.Body)
	type Input struct {
		Protobuf string `json:"protobuf"`
		Template string `json:"template"`
	}
	var input Input
	if err := decoder.Decode(&input); err != nil {
		ReturnError(w, err)
		return
	}

	// create workspace
	dir, err := ioutil.TempDir("", "pggt")
	if err != nil {
		ReturnError(w, err)
	}
	// clean up
	defer func() {
		if err = os.RemoveAll(dir); err != nil {
			logrus.Printf("error: failed to remove temporary directory: %v", err)
		}
	}()
	if err = ioutil.WriteFile(filepath.Join(dir, "example.proto"), []byte(input.Protobuf), 0644); err != nil {
		ReturnError(w, err)
		return
	}
	if err = ioutil.WriteFile(filepath.Join(dir, "example.output.tmpl"), []byte(input.Template), 0644); err != nil {
		ReturnError(w, err)
		return
	}

	// generate
	cmd := exec.Command("protoc", "-I"+dir, "--gotemplate_out=template_dir="+dir+",debug=true:"+dir, filepath.Join(dir, "example.proto")) // #nosec
	out, err := cmd.CombinedOutput()
	if err != nil {
		ReturnError(w, errors.New(string(out)))
		return
	}

	// read output
	content, err := ioutil.ReadFile(filepath.Join(dir, "example.output")) // #nosec
	if err != nil {
		ReturnError(w, err)
		return
	}

	ReturnContent(w, content)
}

func ReturnContent(w http.ResponseWriter, output interface{}) {
	payload := map[string]interface{}{
		"output": fmt.Sprintf("%s", output),
	}
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ReturnError(w http.ResponseWriter, err error) {
	payload := map[string]interface{}{
		"error": fmt.Sprintf("%v", err),
	}
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	if _, err := w.Write(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

