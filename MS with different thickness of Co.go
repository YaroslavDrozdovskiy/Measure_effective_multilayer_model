//установка эого значения похволяет выполнить циклическое повторения вдоль направлений x y z для размагничевающего поля(увеличивает время расчета) 
SetPBC(10,10,0)
// число ячеек вдоль x y z
SetGridsize(1000,1000,5)
thickness := 3.6e-9
//установка размеров ячеек в метрах
SetCellsize(2e-9,2e-9,thickness)

SetGeom(layer(0).add(layer(1)).add(layer(2)).add(layer(3)).add(layer(9)).add(layer(10)).add(layer(11)).add(layer(12)).add(layer(18)).add(layer(19)).add(layer(20)).add(layer(21)).add(layer(27)).add(layer(28)).add(layer(29)).add(layer(30)).add(layer(36)).add(layer(37)).add(layer(38)).add(layer(39)))
// SetMesh(int, int, int, float64, float64, float64, int, int, int) ... Sets GridSize, CellSize and PBC in once

Heff:=-300 
MagCo:=1497.5
AexCo:=2.5e-11
c:=2.25
// намагниченность наыщения
Msat=MagCo*1e3/c
// консанта обменного взаимодействия
Aex=AexCo/c
// Константа одноосной анизотропии 1-го порядка (Дж / м3).
Ku1=Heff*MagCo/20/c+2*Pi*MagCo*MagCo/10/c/c
print(Ku1)
//направление одноосной анизотропии
AnisU = vector(0,0,1)

// блок расчета косвенного обменного взаимодействия между словями
RKKY := 0.0e-3
scale := (RKKY * thickness) / (2 * Aex.Average())
for i := 0; i <= 3; i++ {
	j := i + 1
	defRegion(i, layer(i))
	defRegion(j, layer(j))
	// Масштабирование обменного взаимодействия между двумя регионами.
	ext_scaleExchange(i, j, scale)
}

S:=0.00
// Stopping max dM for Minimize
MinimizerStop = 5e-5
// блок пасчета системы при различных энергиях ВДМ
for B:=0e-3; B<=4e-3; B+=0.5e-3{ 
   S=B*1e3
//интеревейсная сила взаимодействия ВДМ, данном случае приведенная к эффективной модели
   Dind = B/c
   name:="0"
   name=sprintf("%.2f",S)

   m=uniform(0,0,1)
   //создание двух центров в 1-м слое с противоположно направленной намагниченностью
   defregion(1,circle(100e-9).transl(-500e-9, -700e-9,0).add(circle(100e-9).transl(400e-9, 300e-9,0)))
   m.setregion(1,uniform(0,0,-1))
   //frozenspins.setregion(1,1)
   print(name)
//    Использует  метод сопряженного градиента, чтобы минимизировать общую энергию
   minimize()
   //saveas(m,name)
   snapshot(m)
   }