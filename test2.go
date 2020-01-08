// гистерезис

SetPBC(1,1,0)
SetGridsize(1000,1000,13)
SetCellsize(2e-9,2e-9,1e-9)
SetGeom(layer(0).add(layer(3)).add(layer(6)).add(layer(9)).add(layer(12)))
Heff:=6300 
MagCo:=1544
AexCo:=2.5e-11
c:=1

Msat=MagCo*1e3/c
Aex=AexCo/c
Ku1=Heff*MagCo/20/c+2*Pi*MagCo*MagCo/10/c/c
print(Ku1)
AnisU = vector(0,0,1)
Dind = 1.5e-3/c


   m=uniform(0,0,1)
   defregion(1,circle(50e-9).transl(-500e-9, -700e-9,0).add(circle(50e-9).transl(400e-9, 300e-9,0)))
   m.setregion(1,uniform(0,0,-1))
   frozenspins.setregion(1,1)

TableAdd(B_ext)

MinimizerStop = 5e-5
S:=0.00

for B:=10000e-4; B>=3000e-4; B-=1000e-4{ 
   B_ext = vector(B, 0, 0)
   S=B*1e4
   name:="0"
   name=sprintf("%.0f",S)
   print(name)
   minimize()
   tablesave()
   snapshot(m)
}

for B:=3000e-4; B>=0e-4; B-=60e-4{ 
   B_ext = vector(B, 0, 0)
   S=B*1e4
   name:="0"
   name=sprintf("%.0f",S)
   print(name)
    minimize()
 //  saveas(m, name)
 if S <= 60	{
	saveas(m, name)
		}
   tablesave()
   snapshot(m)
}