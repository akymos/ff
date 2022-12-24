package internal

import (
	"bytes"
	"os"
	"text/template"
)

func InitAlias() error {
	file, err := os.OpenFile(BaseConfig.AliasFile, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return err
	}
	if stat.Size() == 0 {
		_, err = file.Write([]byte("#!/bin/bash\n"))
		if err != nil {
			return err
		}
	}
	return nil
}

func GetAlias() string {
	return BaseConfig.AliasFile
}

func GenerateAlias() error {
	file, err := os.OpenFile(BaseConfig.AliasFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	aliasTpl := template.Must(template.New("cd_override").Parse(`#!/bin/bash
cd() { 
	if [[ $# -eq 0 ]]; then builtin cd ~ || return;
	elif [[ -d $* ]]; then builtin cd "$@" || return; 
	else 
		if [[ $* == "ff-config" ]]; then builtin cd {{ .BasePath }} || return;
		{{- range $key, $val := .AliasList}}
		elif [[ $* == "{{$key}}" ]]; then builtin cd "{{$val}}" || return;
		{{- end}}
		else builtin cd "$*" || return; 
		fi; 
	fi;
}
`))
	list := FindAll()
	buffer := &bytes.Buffer{}
	err = aliasTpl.Execute(buffer, map[string]interface{}{
		"BasePath":  BaseConfig.FfDir,
		"AliasList": list,
	})
	_, err = file.Write(buffer.Bytes())
	if err != nil {
		return err
	}

	return nil
}
