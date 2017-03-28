package main

import "fmt"

type Kg float64
type Lb float64
type Kan float64

const kgVSLb float64 = 2.204
const kgVSKan float64 = 4.0 / 15.0

func (kg Kg) String() string   { return fmt.Sprintf("%gkg", kg) }
func (lb Lb) String() string   { return fmt.Sprintf("%glb", lb) }
func (kan Kan) String() string { return fmt.Sprintf("%g貫", kan) }

func KgToLb(kg Kg) Lb    { return Lb(float64(kg) * kgVSLb) }
func LbToKg(lb Lb) Kg    { return Kg(float64(lb) / kgVSLb) }
func KgToKan(kg Kg) Kan  { return Kan(float64(kg) * kgVSKan) }
func KanToKg(kan Kan) Kg { return Kg(float64(kan) / kgVSKan) }
func LbToKan(lb Lb) Kan  { return KgToKan(LbToKg(lb)) }
func KanToLb(kan Kan) Lb { return KgToLb(KanToKg(kan)) }
