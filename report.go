package main

import (
  "github.com/jedib0t/go-pretty/table"
  "os"
)

func report(alives []Pong) {

  t := table.NewWriter()
  t.SetOutputMirror(os.Stdout)
  t.AppendHeader(table.Row{"ip address", "alive"})
  for _, a := range alives {
    t.AppendRows([]table.Row{
      {a.Ip, a.Alive},
    })
  }
  t.Render()

}
