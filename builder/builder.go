package builder

import (
    "fmt"
    "log"
    "os"
    "os/exec"
    "strings"
)

type Application struct {
    Name  string
    Root  string
    Key   string
    Steps []string
}

func (app *Application) StartBuild() {
    err := setupDirectory(app.Root)

    if err != nil {
        log.Fatal(err)
    }

    err = os.Chdir(app.Root)
    if err != nil {
        log.Fatal(err)
    }

    for _, v := range app.Steps {
        err := execute(v)
        if err != nil {
            log.Fatal(err)
        }
    }
}

func setupDirectory(path string) error {
    err := os.MkdirAll(path, 0750)
    if err != nil {
        switch {
        case os.IsExist(err):
            return nil
        default:
            return err
        }
    }

    return nil
}

func execute(step string) error {
    fmt.Println("Splitting command ", step)
    cmd_strings := strings.Split(step, " ")

    fmt.Println("Executing: %s with %v", cmd_strings[0], cmd_strings[1:])
    cmd := exec.Command(cmd_strings[0], strings.Join(cmd_strings[1:], " "))
    err := cmd.Run()

    return err
}
