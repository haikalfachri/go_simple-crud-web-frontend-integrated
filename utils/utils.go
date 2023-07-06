package utils

import (
	"html/template"
	"io"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func GetConfig(key string) string {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error when reading configuration file: %s\n", err)
	}

	return viper.GetString(key)
}

func GenerateUniqueFileName(originalName string) string {
	ext := filepath.Ext(originalName)
	fileName := strings.TrimSuffix(originalName, ext)

	timestamp := time.Now().Unix()
	fileName = fileName + "_" + strconv.FormatInt(timestamp, 10)

	uniqueID := uuid.New().String()
	fileName = fileName + "_" + uniqueID

	fileName = fileName + ext

	return fileName
}

type Renderer struct {
    template *template.Template
    debug    bool
    location string
}

func NewRenderer(location string, debug bool) *Renderer {
    tpl := new(Renderer)
    tpl.location = location
    tpl.debug = debug

    tpl.ReloadTemplates()

    return tpl
}

func (t *Renderer) ReloadTemplates() {
    t.template = template.Must(template.ParseGlob(t.location))
}

func (t *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context,) error {
    if t.debug {
        t.ReloadTemplates()
    }

    return t.template.ExecuteTemplate(w, name, data)
}
