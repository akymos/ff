package internal

import (
	"bytes"
	"os"
	"text/template"
)

func InitAlias() error {
	file, err := os.OpenFile(baseConfig.aliasFile, os.O_CREATE|os.O_WRONLY, 0666)
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
	return baseConfig.aliasFile
}

func PopulateAlias() error {
	file, err := os.OpenFile(baseConfig.aliasFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	aliasTpl := template.Must(template.New("cd_override").Parse(`#!/bin/bash
cd() { 
	if [[ -d $@ ]]; then builtin cd "$@"; 
	else 
		if [[ $@ == "ff-config" ]]; then builtin cd {{ .BasePath }};
		{{- range $key, $val := .AliasList}}
		elif [[ $@ == "{{$key}}" ]]; then builtin cd {{$val}};
		{{- end}}
		else builtin cd "$@"; 
		fi; 
	fi;
}
`))
	list := LocalDb.FindAll()
	buffer := &bytes.Buffer{}
	err = aliasTpl.Execute(buffer, map[string]interface{}{
		"BasePath":  baseConfig.ffDir,
		"AliasList": list,
	})
	_, err = file.Write(buffer.Bytes())
	if err != nil {
		return err
	}

	return nil
}