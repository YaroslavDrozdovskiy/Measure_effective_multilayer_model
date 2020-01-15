
/////thickness of Co and Pd layers/////
thickness_Co := 1.4e-9
thickness_Pd := 2e-9
thickness := thickness_Co + thickness_Pd
c := thickness / thickness_Co

SetPBC(10,10,0)
SetGridsize(1000,1000,5)
SetCellsize(2e-9,2e-9,thickness)
SetGeom(layer(0).add(layer(1)).add(layer(2)).add(layer(3)).add(layer(4)))
//SetGeom(layer(0).add(layer(1)).add(layer(2)).add(layer(3)).add(layer(9)).add(layer(10)).add(layer(11)).add(layer(12)).add(layer(18)).add(layer(19)).add(layer(20)).add(layer(21)).add(layer(27)).add(layer(28)).add(layer(29)).add(layer(30)).add(layer(36)).add(layer(37)).add(layer(38)).add(layer(39)))

////////////////////INITIAL PARAMETERS///////////////////
Heff:=1400
MagCo:=1508
AexCo:=2.5e-11
Msat=MagCo*1e3/c
Aex=AexCo/c
Ku1=Heff*MagCo/20/c+2*Pi*MagCo*MagCo/10/c/c
print(Ku1)
AnisU = vector(0,0,1)

////////////////////BLOCK RKKY///////////////////
max_index_reg := 4
RKKY := 0.0e-3
scale := (RKKY * thickness) / (2 * Aex.Average())
for j := 1; j <= max_index_reg; j++ {
	i := j - 1
	defRegion(i, layer(i))
	defRegion(j, layer(j))
	ext_scaleExchange(i, j, scale)
}

////////////////////MAIN MEASURE/////////////////
S:=0.0
MinimizerStop = 5e-5
skyrm1 := cylinder(100e-9, thickness * 5).transl(-500e-9, -700e-9,0)
skyrm2 := cylinder(100e-9, thickness * 5).transl(400e-9, 300e-9,0)
both_skyrm := skyrm1.add(skyrm2)

for B:=0e-3; B<=4e-3; B+=0.5e-3{ 
   S=B*1e3
   Dind = B/c
   name:="0"
   name=sprintf("%.2f",S)
   m=uniform(0,0,1)
   m.setInShape(both_skyrm, uniform(0,0,-1))

   //defRegion(1,skyrm1.add(skyrm2))
   //m.setregion(1,uniform(0,0,-1))
   //frozenspins.setregion(1,1)
   
   print(name)
   minimize()
   snapshot(m)
   saveas(m,name)
   
   }