package models

import ()

type Baseline struct {
        Id   int
        Name string
}

func (b *Baseline) SetId(id int) {
        b.Id = id
}

func (b *Baseline) SetName(name string) {
        b.Name = name
}


type Command struct {
        Id              int
        Cmd             string
        ExeOrder        int
        ControlId       int
}
