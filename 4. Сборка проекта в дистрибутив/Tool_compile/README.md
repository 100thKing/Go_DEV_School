# Построение AST
После выполнения команды ниже было построенно абстрактное синтаксическое дерево
```console
go tool compile -W main.go
```
<details><summary>AST</summary>
  
    before walk main
    .   DCL # main.go:4:2
    .   .   NAME-main.rows esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:4:2
    .   AS Def tc(1) # main.go:4:7
    .   .   NAME-main.rows esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:4:2
    .   .   LITERAL-5 int tc(1) # main.go:4:10
    .   DCL # main.go:5:6
    .   .   NAME-main.i esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:5:6
    .   AS Def tc(1) # main.go:5:8
    .   .   NAME-main.i esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:5:6
    .   .   LITERAL-1 int tc(1) # main.go:5:11
    .   FOR tc(1) # main.go:5:2
    .   FOR-Cond
    .   .   LE bool tc(1) # main.go:5:16
    .   .   .   NAME-main.i esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:5:6
    .   .   .   NAME-main.rows esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:4:2
    .   FOR-Post
    .   .   BLOCK # main.go:5:26
    .   .   BLOCK-List
    .   .   .   ASOP-ADD AsOp:ADD IncDec tc(1) # main.go:5:26
    .   .   .   .   NAME-main.i esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:5:6
    .   .   .   .   LITERAL-1 int tc(1) # main.go:5:26
    .   FOR-Body
    .   .   DCL # main.go:6:7
    .   .   .   NAME-main.j esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:6:7
    .   .   AS Def tc(1) # main.go:6:9
    .   .   .   NAME-main.j esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:6:7
    .   .   .   LITERAL-1 int tc(1) # main.go:6:12
    .   .   FOR tc(1) # main.go:6:3
    .   .   FOR-Cond
    .   .   .   LE bool tc(1) # main.go:6:17
    .   .   .   .   NAME-main.j esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:6:7
    .   .   .   .   SUB int tc(1) # main.go:6:24
    .   .   .   .   .   NAME-main.rows esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:4:2
    .   .   .   .   .   NAME-main.i esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:5:6
    .   .   FOR-Post
    .   .   .   BLOCK # main.go:6:29
    .   .   .   BLOCK-List
    .   .   .   .   ASOP-ADD AsOp:ADD IncDec tc(1) # main.go:6:29
    .   .   .   .   .   NAME-main.j esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:6:7
    .   .   .   .   .   LITERAL-1 int tc(1) # main.go:6:29
    .   .   FOR-Body
    .   .   .   PRINT tc(1) # main.go:7:9
    .   .   .   PRINT-Args
    .   .   .   .   LITERAL-" " string tc(1) # main.go:7:10
    .   .   DCL # main.go:9:7
    .   .   .   NAME-main.k esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:9:7
    .   .   AS Def tc(1) # main.go:9:9
    .   .   .   NAME-main.k esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:9:7
    .   .   .   LITERAL-1 int tc(1) # main.go:9:12
    .   .   FOR tc(1) # main.go:9:3
    .   .   FOR-Cond
    .   .   .   LE bool tc(1) # main.go:9:17
    .   .   .   .   NAME-main.k esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:9:7
    .   .   .   .   SUB int tc(1) # main.go:9:23
    .   .   .   .   .   MUL int tc(1) # main.go:9:21
    .   .   .   .   .   .   LITERAL-2 int tc(1) # main.go:9:20
    .   .   .   .   .   .   NAME-main.i esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:5:6
    .   .   .   .   .   LITERAL-1 int tc(1) # main.go:9:24
    .   .   FOR-Post
    .   .   .   BLOCK # main.go:9:28
    .   .   .   BLOCK-List
    .   .   .   .   ASOP-ADD AsOp:ADD IncDec tc(1) # main.go:9:28
    .   .   .   .   .   NAME-main.k esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:9:7
    .   .   .   .   .   LITERAL-1 int tc(1) # main.go:9:28
    .   .   FOR-Body
    .   .   .   PRINT tc(1) # main.go:10:9
    .   .   .   PRINT-Args
    .   .   .   .   LITERAL-"*" string tc(1) # main.go:10:10
    .   .   PRINTN tc(1) # main.go:12:10
    .   DCL # main.go:15:6
    .   .   NAME-main.i esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:15:6
    .   AS Def tc(1) # main.go:15:8
    .   .   NAME-main.i esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:15:6
    .   .   SUB int tc(1) # main.go:15:16
    .   .   .   NAME-main.rows esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:4:2
    .   .   .   LITERAL-1 int tc(1) # main.go:15:18
    .   FOR tc(1) # main.go:15:2
    .   FOR-Cond
    .   .   GE bool tc(1) # main.go:15:23
    .   .   .   NAME-main.i esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:15:6
    .   .   .   LITERAL-1 int tc(1) # main.go:15:26
    .   FOR-Post
    .   .   BLOCK # main.go:15:30
    .   .   BLOCK-List
    .   .   .   ASOP-SUB AsOp:SUB IncDec tc(1) # main.go:15:30
    .   .   .   .   NAME-main.i esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:15:6
    .   .   .   .   LITERAL-1 int tc(1) # main.go:15:30
    .   FOR-Body
    .   .   DCL # main.go:16:7
    .   .   .   NAME-main.j esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:16:7
    .   .   AS Def tc(1) # main.go:16:9
    .   .   .   NAME-main.j esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:16:7
    .   .   .   LITERAL-1 int tc(1) # main.go:16:12
    .   .   FOR tc(1) # main.go:16:3
    .   .   FOR-Cond
    .   .   .   LE bool tc(1) # main.go:16:17
    .   .   .   .   NAME-main.j esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:16:7
    .   .   .   .   SUB int tc(1) # main.go:16:24
    .   .   .   .   .   NAME-main.rows esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:4:2
    .   .   .   .   .   NAME-main.i esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:15:6
    .   .   FOR-Post
    .   .   .   BLOCK # main.go:16:29
    .   .   .   BLOCK-List
    .   .   .   .   ASOP-ADD AsOp:ADD IncDec tc(1) # main.go:16:29
    .   .   .   .   .   NAME-main.j esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:16:7
    .   .   .   .   .   LITERAL-1 int tc(1) # main.go:16:29
    .   .   FOR-Body
    .   .   .   PRINT tc(1) # main.go:17:9
    .   .   .   PRINT-Args
    .   .   .   .   LITERAL-" " string tc(1) # main.go:17:10
    .   .   DCL # main.go:19:7
    .   .   .   NAME-main.k esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:19:7
    .   .   AS Def tc(1) # main.go:19:9
    .   .   .   NAME-main.k esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:19:7
    .   .   .   LITERAL-1 int tc(1) # main.go:19:12
    .   .   FOR tc(1) # main.go:19:3
    .   .   FOR-Cond
    .   .   .   LE bool tc(1) # main.go:19:17
    .   .   .   .   NAME-main.k esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:19:7
    .   .   .   .   SUB int tc(1) # main.go:19:23
    .   .   .   .   .   MUL int tc(1) # main.go:19:21
    .   .   .   .   .   .   LITERAL-2 int tc(1) # main.go:19:20
    .   .   .   .   .   .   NAME-main.i esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:15:6
    .   .   .   .   .   LITERAL-1 int tc(1) # main.go:19:24
    .   .   FOR-Post
    .   .   .   BLOCK # main.go:19:28
    .   .   .   BLOCK-List
    .   .   .   .   ASOP-ADD AsOp:ADD IncDec tc(1) # main.go:19:28
    .   .   .   .   .   NAME-main.k esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:19:7
    .   .   .   .   .   LITERAL-1 int tc(1) # main.go:19:28
    .   .   FOR-Body
    .   .   .   PRINT tc(1) # main.go:20:9
    .   .   .   PRINT-Args
    .   .   .   .   LITERAL-"*" string tc(1) # main.go:20:10
    .   .   PRINTN tc(1) # main.go:22:10
    after walk main
    .   DCL # main.go:4:2
    .   .   NAME-main.rows esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:4:2
    .   AS Def tc(1) # main.go:4:7
    .   .   NAME-main.rows esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:4:2
    .   .   LITERAL-5 int tc(1) # main.go:4:10
    .   DCL # main.go:5:6
    .   .   NAME-main.i esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:5:6
    .   AS Def tc(1) # main.go:5:8
    .   .   NAME-main.i esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:5:6
    .   .   LITERAL-1 int tc(1) # main.go:5:11
    .   FOR tc(1) # main.go:5:2
    .   FOR-Cond
    .   .   LE bool tc(1) # main.go:5:16
    .   .   .   NAME-main.i esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:5:6
    .   .   .   NAME-main.rows esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:4:2
    .   FOR-Post
    .   .   BLOCK # main.go:5:26
    .   .   BLOCK-List
    .   .   .   AS tc(1) # main.go:5:26
    .   .   .   .   NAME-main.i esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:5:6
    .   .   .   .   ADD int tc(1) # main.go:5:26
    .   .   .   .   .   NAME-main.i esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:5:6
    .   .   .   .   .   LITERAL-1 int tc(1) # main.go:5:26
    .   FOR-Body
    .   .   DCL # main.go:6:7
    .   .   .   NAME-main.j esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:6:7
    .   .   AS Def tc(1) # main.go:6:9
    .   .   .   NAME-main.j esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:6:7
    .   .   .   LITERAL-1 int tc(1) # main.go:6:12
    .   .   FOR tc(1) # main.go:6:3
    .   .   FOR-Cond
    .   .   .   LE bool tc(1) # main.go:6:17
    .   .   .   .   NAME-main.j esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:6:7
    .   .   .   .   SUB int tc(1) # main.go:6:24
    .   .   .   .   .   NAME-main.rows esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:4:2
    .   .   .   .   .   NAME-main.i esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:5:6
    .   .   FOR-Post
    .   .   .   BLOCK # main.go:6:29
    .   .   .   BLOCK-List
    .   .   .   .   AS tc(1) # main.go:6:29
    .   .   .   .   .   NAME-main.j esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:6:7
    .   .   .   .   .   ADD int tc(1) # main.go:6:29
    .   .   .   .   .   .   NAME-main.j esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:6:7
    .   .   .   .   .   .   LITERAL-1 int tc(1) # main.go:6:29
    .   .   FOR-Body
    .   .   .   BLOCK tc(1) # main.go:7:9
    .   .   .   BLOCK-List
    .   .   .   .   CALLFUNC Walked tc(1) # main.go:7:9
    .   .   .   .   .   NAME-runtime.printlock Class:PFUNC Offset:0 Used FUNC-func() tc(1)
    .   .   .   .   CALLFUNC Walked tc(1) # main.go:7:9
    .   .   .   .   .   NAME-runtime.printsp Class:PFUNC Offset:0 Used FUNC-func() tc(1)
    .   .   .   .   CALLFUNC Walked tc(1) # main.go:7:9
    .   .   .   .   .   NAME-runtime.printunlock Class:PFUNC Offset:0 Used FUNC-func() tc(1)
    .   .   DCL # main.go:9:7
    .   .   .   NAME-main.k esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:9:7
    .   .   AS Def tc(1) # main.go:9:9
    .   .   .   NAME-main.k esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:9:7
    .   .   .   LITERAL-1 int tc(1) # main.go:9:12
    .   .   FOR tc(1) # main.go:9:3
    .   .   FOR-Cond
    .   .   .   LE bool tc(1) # main.go:9:17
    .   .   .   .   NAME-main.k esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:9:7
    .   .   .   .   SUB int tc(1) # main.go:9:23
    .   .   .   .   .   MUL int tc(1) # main.go:9:21
    .   .   .   .   .   .   LITERAL-2 int tc(1) # main.go:9:20
    .   .   .   .   .   .   NAME-main.i esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:5:6
    .   .   .   .   .   LITERAL-1 int tc(1) # main.go:9:24
    .   .   FOR-Post
    .   .   .   BLOCK # main.go:9:28
    .   .   .   BLOCK-List
    .   .   .   .   AS tc(1) # main.go:9:28
    .   .   .   .   .   NAME-main.k esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:9:7
    .   .   .   .   .   ADD int tc(1) # main.go:9:28
    .   .   .   .   .   .   NAME-main.k esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:9:7
    .   .   .   .   .   .   LITERAL-1 int tc(1) # main.go:9:28
    .   .   FOR-Body
    .   .   .   BLOCK tc(1) # main.go:10:9
    .   .   .   BLOCK-List
    .   .   .   .   CALLFUNC Walked tc(1) # main.go:10:9
    .   .   .   .   .   NAME-runtime.printlock Class:PFUNC Offset:0 Used FUNC-func() tc(1)
    .   .   .   .   CALLFUNC Walked tc(1) # main.go:10:9
    .   .   .   .   .   NAME-runtime.printstring Class:PFUNC Offset:0 Used FUNC-func(string) tc(1)
    .   .   .   .   CALLFUNC-Args
    .   .   .   .   .   LITERAL-"*" string tc(1) # main.go:10:9
    .   .   .   .   CALLFUNC Walked tc(1) # main.go:10:9
    .   .   .   .   .   NAME-runtime.printunlock Class:PFUNC Offset:0 Used FUNC-func() tc(1)
    .   .   BLOCK tc(1) # main.go:12:10
    .   .   BLOCK-List
    .   .   .   CALLFUNC Walked tc(1) # main.go:12:10
    .   .   .   .   NAME-runtime.printlock Class:PFUNC Offset:0 Used FUNC-func() tc(1)
    .   .   .   CALLFUNC Walked tc(1) # main.go:12:10
    .   .   .   .   NAME-runtime.printnl Class:PFUNC Offset:0 Used FUNC-func() tc(1)
    .   .   .   CALLFUNC Walked tc(1) # main.go:12:10
    .   .   .   .   NAME-runtime.printunlock Class:PFUNC Offset:0 Used FUNC-func() tc(1)
    .   DCL # main.go:15:6
    .   .   NAME-main.i esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:15:6
    .   AS Def tc(1) # main.go:15:8
    .   .   NAME-main.i esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:15:6
    .   .   SUB int tc(1) # main.go:15:16
    .   .   .   NAME-main.rows esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:4:2
    .   .   .   LITERAL-1 int tc(1) # main.go:15:18
    .   FOR tc(1) # main.go:15:2
    .   FOR-Cond
    .   .   GE bool tc(1) # main.go:15:23
    .   .   .   NAME-main.i esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:15:6
    .   .   .   LITERAL-1 int tc(1) # main.go:15:26
    .   FOR-Post
    .   .   BLOCK # main.go:15:30
    .   .   BLOCK-List
    .   .   .   AS tc(1) # main.go:15:30
    .   .   .   .   NAME-main.i esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:15:6
    .   .   .   .   SUB int tc(1) # main.go:15:30
    .   .   .   .   .   NAME-main.i esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:15:6
    .   .   .   .   .   LITERAL-1 int tc(1) # main.go:15:30
    .   FOR-Body
    .   .   DCL # main.go:16:7
    .   .   .   NAME-main.j esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:16:7
    .   .   AS Def tc(1) # main.go:16:9
    .   .   .   NAME-main.j esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:16:7
    .   .   .   LITERAL-1 int tc(1) # main.go:16:12
    .   .   FOR tc(1) # main.go:16:3
    .   .   FOR-Cond
    .   .   .   LE bool tc(1) # main.go:16:17
    .   .   .   .   NAME-main.j esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:16:7
    .   .   .   .   SUB int tc(1) # main.go:16:24
    .   .   .   .   .   NAME-main.rows esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:4:2
    .   .   .   .   .   NAME-main.i esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:15:6
    .   .   FOR-Post
    .   .   .   BLOCK # main.go:16:29
    .   .   .   BLOCK-List
    .   .   .   .   AS tc(1) # main.go:16:29
    .   .   .   .   .   NAME-main.j esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:16:7
    .   .   .   .   .   ADD int tc(1) # main.go:16:29
    .   .   .   .   .   .   NAME-main.j esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:16:7
    .   .   .   .   .   .   LITERAL-1 int tc(1) # main.go:16:29
    .   .   FOR-Body
    .   .   .   BLOCK tc(1) # main.go:17:9
    .   .   .   BLOCK-List
    .   .   .   .   CALLFUNC Walked tc(1) # main.go:17:9
    .   .   .   .   .   NAME-runtime.printlock Class:PFUNC Offset:0 Used FUNC-func() tc(1)
    .   .   .   .   CALLFUNC Walked tc(1) # main.go:17:9
    .   .   .   .   .   NAME-runtime.printsp Class:PFUNC Offset:0 Used FUNC-func() tc(1)
    .   .   .   .   CALLFUNC Walked tc(1) # main.go:17:9
    .   .   .   .   .   NAME-runtime.printunlock Class:PFUNC Offset:0 Used FUNC-func() tc(1)
    .   .   DCL # main.go:19:7
    .   .   .   NAME-main.k esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:19:7
    .   .   AS Def tc(1) # main.go:19:9
    .   .   .   NAME-main.k esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:19:7
    .   .   .   LITERAL-1 int tc(1) # main.go:19:12
    .   .   FOR tc(1) # main.go:19:3
    .   .   FOR-Cond
    .   .   .   LE bool tc(1) # main.go:19:17
    .   .   .   .   NAME-main.k esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:19:7
    .   .   .   .   SUB int tc(1) # main.go:19:23
    .   .   .   .   .   MUL int tc(1) # main.go:19:21
    .   .   .   .   .   .   LITERAL-2 int tc(1) # main.go:19:20
    .   .   .   .   .   .   NAME-main.i esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:15:6
    .   .   .   .   .   LITERAL-1 int tc(1) # main.go:19:24
    .   .   FOR-Post
    .   .   .   BLOCK # main.go:19:28
    .   .   .   BLOCK-List
    .   .   .   .   AS tc(1) # main.go:19:28
    .   .   .   .   .   NAME-main.k esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:19:7
    .   .   .   .   .   ADD int tc(1) # main.go:19:28
    .   .   .   .   .   .   NAME-main.k esc(no) Class:PAUTO Offset:0 OnStack Used int tc(1) # main.go:19:7
    .   .   .   .   .   .   LITERAL-1 int tc(1) # main.go:19:28
    .   .   FOR-Body
    .   .   .   BLOCK tc(1) # main.go:20:9
    .   .   .   BLOCK-List
    .   .   .   .   CALLFUNC Walked tc(1) # main.go:20:9
    .   .   .   .   .   NAME-runtime.printlock Class:PFUNC Offset:0 Used FUNC-func() tc(1)
    .   .   .   .   CALLFUNC Walked tc(1) # main.go:20:9
    .   .   .   .   .   NAME-runtime.printstring Class:PFUNC Offset:0 Used FUNC-func(string) tc(1)
    .   .   .   .   CALLFUNC-Args
    .   .   .   .   .   LITERAL-"*" string tc(1) # main.go:20:9
    .   .   .   .   CALLFUNC Walked tc(1) # main.go:20:9
    .   .   .   .   .   NAME-runtime.printunlock Class:PFUNC Offset:0 Used FUNC-func() tc(1)
    .   .   BLOCK tc(1) # main.go:22:10
    .   .   BLOCK-List
    .   .   .   CALLFUNC Walked tc(1) # main.go:22:10
    .   .   .   .   NAME-runtime.printlock Class:PFUNC Offset:0 Used FUNC-func() tc(1)
    .   .   .   CALLFUNC Walked tc(1) # main.go:22:10
    .   .   .   .   NAME-runtime.printnl Class:PFUNC Offset:0 Used FUNC-func() tc(1)
    .   .   .   CALLFUNC Walked tc(1) # main.go:22:10
    .   .   .   .   NAME-runtime.printunlock Class:PFUNC Offset:0 Used FUNC-func() tc(1)

</details>

# Построение SSA
```console
go tool compile -S main.go
```
<details><summary>SSA</summary>
  
    main.main STEXT size=389 args=0x0 locals=0x50 funcid=0x0 align=0x0
            0x0000 00000 (/home/maks/Desktop/GO/task_4/main.go:3)   TEXT    main.main(SB), ABIInternal, $80-0
            0x0000 00000 (/home/maks/Desktop/GO/task_4/main.go:3)   CMPQ    SP, 16(R14)
            0x0004 00004 (/home/maks/Desktop/GO/task_4/main.go:3)   PCDATA  $0, $-2
            0x0004 00004 (/home/maks/Desktop/GO/task_4/main.go:3)   JLS     378
            0x000a 00010 (/home/maks/Desktop/GO/task_4/main.go:3)   PCDATA  $0, $-1
            0x000a 00010 (/home/maks/Desktop/GO/task_4/main.go:3)   SUBQ    $80, SP
            0x000e 00014 (/home/maks/Desktop/GO/task_4/main.go:3)   MOVQ    BP, 72(SP)
            0x0013 00019 (/home/maks/Desktop/GO/task_4/main.go:3)   LEAQ    72(SP), BP
            0x0018 00024 (/home/maks/Desktop/GO/task_4/main.go:3)   FUNCDATA        $0, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
            0x0018 00024 (/home/maks/Desktop/GO/task_4/main.go:3)   FUNCDATA        $1, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
            0x0018 00024 (/home/maks/Desktop/GO/task_4/main.go:3)   MOVL    $1, AX
            0x001d 00029 (/home/maks/Desktop/GO/task_4/main.go:5)   JMP     55
            0x001f 00031 (/home/maks/Desktop/GO/task_4/main.go:12)  PCDATA  $1, $0
            0x001f 00031 (/home/maks/Desktop/GO/task_4/main.go:12)  NOP
            0x0020 00032 (/home/maks/Desktop/GO/task_4/main.go:12)  CALL    runtime.printlock(SB)
            0x0025 00037 (/home/maks/Desktop/GO/task_4/main.go:12)  CALL    runtime.printnl(SB)
            0x002a 00042 (/home/maks/Desktop/GO/task_4/main.go:12)  CALL    runtime.printunlock(SB)
            0x002f 00047 (/home/maks/Desktop/GO/task_4/main.go:5)   MOVQ    main.i+56(SP), AX
            0x0034 00052 (/home/maks/Desktop/GO/task_4/main.go:5)   INCQ    AX
            0x0037 00055 (/home/maks/Desktop/GO/task_4/main.go:5)   CMPQ    AX, $5
            0x003b 00059 (/home/maks/Desktop/GO/task_4/main.go:5)   JGT     86
            0x003d 00061 (/home/maks/Desktop/GO/task_4/main.go:5)   MOVQ    AX, main.i+56(SP)
            0x0042 00066 (/home/maks/Desktop/GO/task_4/main.go:5)   MOVL    $1, CX
            0x0047 00071 (/home/maks/Desktop/GO/task_4/main.go:6)   JMP     299
            0x004c 00076 (/home/maks/Desktop/GO/task_4/main.go:6)   MOVL    $1, CX
            0x0051 00081 (/home/maks/Desktop/GO/task_4/main.go:6)   JMP     361
            0x0056 00086 (/home/maks/Desktop/GO/task_4/main.go:15)  MOVL    $4, AX
            0x005b 00091 (/home/maks/Desktop/GO/task_4/main.go:15)  JMP     118
            0x005d 00093 (/home/maks/Desktop/GO/task_4/main.go:19)  MOVQ    DX, main..autotmp_7+64(SP)
            0x0062 00098 (/home/maks/Desktop/GO/task_4/main.go:22)  CALL    runtime.printlock(SB)
            0x0067 00103 (/home/maks/Desktop/GO/task_4/main.go:22)  CALL    runtime.printnl(SB)
            0x006c 00108 (/home/maks/Desktop/GO/task_4/main.go:22)  CALL    runtime.printunlock(SB)
            0x0071 00113 (/home/maks/Desktop/GO/task_4/main.go:15)  MOVQ    main..autotmp_7+64(SP), AX
            0x0076 00118 (/home/maks/Desktop/GO/task_4/main.go:15)  CMPQ    AX, $1
            0x007a 00122 (/home/maks/Desktop/GO/task_4/main.go:15)  JLT     143
            0x007c 00124 (/home/maks/Desktop/GO/task_4/main.go:15)  MOVQ    AX, main.i+48(SP)
            0x0081 00129 (/home/maks/Desktop/GO/task_4/main.go:15)  MOVL    $1, CX
            0x0086 00134 (/home/maks/Desktop/GO/task_4/main.go:16)  JMP     188
            0x0088 00136 (/home/maks/Desktop/GO/task_4/main.go:16)  MOVL    $1, CX
            0x008d 00141 (/home/maks/Desktop/GO/task_4/main.go:16)  JMP     247
            0x008f 00143 (/home/maks/Desktop/GO/task_4/main.go:24)  PCDATA  $1, $-1
            0x008f 00143 (/home/maks/Desktop/GO/task_4/main.go:24)  MOVQ    72(SP), BP
            0x0094 00148 (/home/maks/Desktop/GO/task_4/main.go:24)  ADDQ    $80, SP
            0x0098 00152 (/home/maks/Desktop/GO/task_4/main.go:24)  RET
            0x0099 00153 (/home/maks/Desktop/GO/task_4/main.go:16)  MOVQ    CX, main.j+32(SP)
            0x009e 00158 (/home/maks/Desktop/GO/task_4/main.go:17)  PCDATA  $1, $0
            0x009e 00158 (/home/maks/Desktop/GO/task_4/main.go:17)  NOP
            0x00a0 00160 (/home/maks/Desktop/GO/task_4/main.go:17)  CALL    runtime.printlock(SB)
            0x00a5 00165 (/home/maks/Desktop/GO/task_4/main.go:17)  CALL    runtime.printsp(SB)
            0x00aa 00170 (/home/maks/Desktop/GO/task_4/main.go:17)  CALL    runtime.printunlock(SB)
            0x00af 00175 (/home/maks/Desktop/GO/task_4/main.go:16)  MOVQ    main.j+32(SP), CX
            0x00b4 00180 (/home/maks/Desktop/GO/task_4/main.go:16)  INCQ    CX
            0x00b7 00183 (/home/maks/Desktop/GO/task_4/main.go:16)  MOVQ    main.i+48(SP), AX
            0x00bc 00188 (/home/maks/Desktop/GO/task_4/main.go:16)  LEAQ    -5(AX), DX
            0x00c0 00192 (/home/maks/Desktop/GO/task_4/main.go:16)  NEGQ    DX
            0x00c3 00195 (/home/maks/Desktop/GO/task_4/main.go:16)  CMPQ    CX, DX
            0x00c6 00198 (/home/maks/Desktop/GO/task_4/main.go:16)  JLE     153
            0x00c8 00200 (/home/maks/Desktop/GO/task_4/main.go:16)  JMP     136
            0x00ca 00202 (/home/maks/Desktop/GO/task_4/main.go:19)  MOVQ    CX, main.k+16(SP)
            0x00cf 00207 (/home/maks/Desktop/GO/task_4/main.go:20)  CALL    runtime.printlock(SB)
            0x00d4 00212 (/home/maks/Desktop/GO/task_4/main.go:20)  LEAQ    go:string."*"(SB), AX
            0x00db 00219 (/home/maks/Desktop/GO/task_4/main.go:20)  MOVL    $1, BX
            0x00e0 00224 (/home/maks/Desktop/GO/task_4/main.go:20)  CALL    runtime.printstring(SB)
            0x00e5 00229 (/home/maks/Desktop/GO/task_4/main.go:20)  CALL    runtime.printunlock(SB)
            0x00ea 00234 (/home/maks/Desktop/GO/task_4/main.go:19)  MOVQ    main.k+16(SP), CX
            0x00ef 00239 (/home/maks/Desktop/GO/task_4/main.go:19)  INCQ    CX
            0x00f2 00242 (/home/maks/Desktop/GO/task_4/main.go:19)  MOVQ    main.i+48(SP), AX
            0x00f7 00247 (/home/maks/Desktop/GO/task_4/main.go:19)  LEAQ    -1(AX), DX
            0x00fb 00251 (/home/maks/Desktop/GO/task_4/main.go:19)  LEAQ    (AX)(DX*1), BX
            0x00ff 00255 (/home/maks/Desktop/GO/task_4/main.go:19)  NOP
            0x0100 00256 (/home/maks/Desktop/GO/task_4/main.go:19)  CMPQ    CX, BX
            0x0103 00259 (/home/maks/Desktop/GO/task_4/main.go:19)  JLE     202
            0x0105 00261 (/home/maks/Desktop/GO/task_4/main.go:19)  JMP     93
            0x010a 00266 (/home/maks/Desktop/GO/task_4/main.go:6)   MOVQ    CX, main.j+40(SP)
            0x010f 00271 (/home/maks/Desktop/GO/task_4/main.go:7)   CALL    runtime.printlock(SB)
            0x0114 00276 (/home/maks/Desktop/GO/task_4/main.go:7)   CALL    runtime.printsp(SB)
            0x0119 00281 (/home/maks/Desktop/GO/task_4/main.go:7)   CALL    runtime.printunlock(SB)
            0x011e 00286 (/home/maks/Desktop/GO/task_4/main.go:6)   MOVQ    main.j+40(SP), CX
            0x0123 00291 (/home/maks/Desktop/GO/task_4/main.go:6)   INCQ    CX
            0x0126 00294 (/home/maks/Desktop/GO/task_4/main.go:6)   MOVQ    main.i+56(SP), AX
            0x012b 00299 (/home/maks/Desktop/GO/task_4/main.go:6)   LEAQ    -5(AX), DX
            0x012f 00303 (/home/maks/Desktop/GO/task_4/main.go:6)   NEGQ    DX
            0x0132 00306 (/home/maks/Desktop/GO/task_4/main.go:6)   CMPQ    CX, DX
            0x0135 00309 (/home/maks/Desktop/GO/task_4/main.go:6)   JLE     266
            0x0137 00311 (/home/maks/Desktop/GO/task_4/main.go:6)   JMP     76
            0x013c 00316 (/home/maks/Desktop/GO/task_4/main.go:9)   MOVQ    CX, main.k+24(SP)
            0x0141 00321 (/home/maks/Desktop/GO/task_4/main.go:10)  CALL    runtime.printlock(SB)
            0x0146 00326 (/home/maks/Desktop/GO/task_4/main.go:10)  LEAQ    go:string."*"(SB), AX
            0x014d 00333 (/home/maks/Desktop/GO/task_4/main.go:10)  MOVL    $1, BX
            0x0152 00338 (/home/maks/Desktop/GO/task_4/main.go:10)  CALL    runtime.printstring(SB)
            0x0157 00343 (/home/maks/Desktop/GO/task_4/main.go:10)  CALL    runtime.printunlock(SB)
            0x015c 00348 (/home/maks/Desktop/GO/task_4/main.go:9)   MOVQ    main.k+24(SP), CX
            0x0161 00353 (/home/maks/Desktop/GO/task_4/main.go:9)   INCQ    CX
            0x0164 00356 (/home/maks/Desktop/GO/task_4/main.go:9)   MOVQ    main.i+56(SP), AX
            0x0169 00361 (/home/maks/Desktop/GO/task_4/main.go:9)   LEAQ    -1(AX), DX
            0x016d 00365 (/home/maks/Desktop/GO/task_4/main.go:9)   ADDQ    AX, DX
            0x0170 00368 (/home/maks/Desktop/GO/task_4/main.go:9)   CMPQ    CX, DX
            0x0173 00371 (/home/maks/Desktop/GO/task_4/main.go:9)   JLE     316
            0x0175 00373 (/home/maks/Desktop/GO/task_4/main.go:9)   JMP     31
            0x017a 00378 (/home/maks/Desktop/GO/task_4/main.go:9)   NOP
            0x017a 00378 (/home/maks/Desktop/GO/task_4/main.go:3)   PCDATA  $1, $-1
            0x017a 00378 (/home/maks/Desktop/GO/task_4/main.go:3)   PCDATA  $0, $-2
            0x017a 00378 (/home/maks/Desktop/GO/task_4/main.go:3)   CALL    runtime.morestack_noctxt(SB)
            0x017f 00383 (/home/maks/Desktop/GO/task_4/main.go:3)   PCDATA  $0, $-1
            0x017f 00383 (/home/maks/Desktop/GO/task_4/main.go:3)   NOP
            0x0180 00384 (/home/maks/Desktop/GO/task_4/main.go:3)   JMP     0
            0x0000 49 3b 66 10 0f 86 70 01 00 00 48 83 ec 50 48 89  I;f...p...H..PH.
            0x0010 6c 24 48 48 8d 6c 24 48 b8 01 00 00 00 eb 18 90  l$HH.l$H........
            0x0020 e8 00 00 00 00 e8 00 00 00 00 e8 00 00 00 00 48  ...............H
            0x0030 8b 44 24 38 48 ff c0 48 83 f8 05 7f 19 48 89 44  .D$8H..H.....H.D
            0x0040 24 38 b9 01 00 00 00 e9 df 00 00 00 b9 01 00 00  $8..............
            0x0050 00 e9 13 01 00 00 b8 04 00 00 00 eb 19 48 89 54  .............H.T
            0x0060 24 40 e8 00 00 00 00 e8 00 00 00 00 e8 00 00 00  $@..............
            0x0070 00 48 8b 44 24 40 48 83 f8 01 7c 13 48 89 44 24  .H.D$@H...|.H.D$
            0x0080 30 b9 01 00 00 00 eb 34 b9 01 00 00 00 eb 68 48  0......4......hH
            0x0090 8b 6c 24 48 48 83 c4 50 c3 48 89 4c 24 20 66 90  .l$HH..P.H.L$ f.
            0x00a0 e8 00 00 00 00 e8 00 00 00 00 e8 00 00 00 00 48  ...............H
            0x00b0 8b 4c 24 20 48 ff c1 48 8b 44 24 30 48 8d 50 fb  .L$ H..H.D$0H.P.
            0x00c0 48 f7 da 48 39 d1 7e d1 eb be 48 89 4c 24 10 e8  H..H9.~...H.L$..
            0x00d0 00 00 00 00 48 8d 05 00 00 00 00 bb 01 00 00 00  ....H...........
            0x00e0 e8 00 00 00 00 e8 00 00 00 00 48 8b 4c 24 10 48  ..........H.L$.H
            0x00f0 ff c1 48 8b 44 24 30 48 8d 50 ff 48 8d 1c 10 90  ..H.D$0H.P.H....
            0x0100 48 39 d9 7e c5 e9 53 ff ff ff 48 89 4c 24 28 e8  H9.~..S...H.L$(.
            0x0110 00 00 00 00 e8 00 00 00 00 e8 00 00 00 00 48 8b  ..............H.
            0x0120 4c 24 28 48 ff c1 48 8b 44 24 38 48 8d 50 fb 48  L$(H..H.D$8H.P.H
            0x0130 f7 da 48 39 d1 7e d3 e9 10 ff ff ff 48 89 4c 24  ..H9.~......H.L$
            0x0140 18 e8 00 00 00 00 48 8d 05 00 00 00 00 bb 01 00  ......H.........
            0x0150 00 00 e8 00 00 00 00 e8 00 00 00 00 48 8b 4c 24  ............H.L$
            0x0160 18 48 ff c1 48 8b 44 24 38 48 8d 50 ff 48 01 c2  .H..H.D$8H.P.H..
            0x0170 48 39 d1 7e c7 e9 a5 fe ff ff e8 00 00 00 00 90  H9.~............
            0x0180 e9 7b fe ff ff                                   .{...
            rel 33+4 t=7 runtime.printlock+0
            rel 38+4 t=7 runtime.printnl+0
            rel 43+4 t=7 runtime.printunlock+0
            rel 99+4 t=7 runtime.printlock+0
            rel 104+4 t=7 runtime.printnl+0
            rel 109+4 t=7 runtime.printunlock+0
            rel 161+4 t=7 runtime.printlock+0
            rel 166+4 t=7 runtime.printsp+0
            rel 171+4 t=7 runtime.printunlock+0
            rel 208+4 t=7 runtime.printlock+0
            rel 215+4 t=14 go:string."*"+0
            rel 225+4 t=7 runtime.printstring+0
            rel 230+4 t=7 runtime.printunlock+0
            rel 272+4 t=7 runtime.printlock+0
            rel 277+4 t=7 runtime.printsp+0
            rel 282+4 t=7 runtime.printunlock+0
            rel 322+4 t=7 runtime.printlock+0
            rel 329+4 t=14 go:string."*"+0
            rel 339+4 t=7 runtime.printstring+0
            rel 344+4 t=7 runtime.printunlock+0
            rel 379+4 t=7 runtime.morestack_noctxt+0
    go:cuinfo.producer.<unlinkable> SDWARFCUINFO dupok size=0
            0x0000 72 65 67 61 62 69                                regabi
    go:cuinfo.packagename.main SDWARFCUINFO dupok size=0
            0x0000 6d 61 69 6e                                      main
    main..inittask SNOPTRDATA size=24
            0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
            0x0010 00 00 00 00 00 00 00 00                          ........
    go:string." " SRODATA dupok size=1
            0x0000 20                                                
    go:string."*" SRODATA dupok size=1
            0x0000 2a                                               *
    gclocals·g2BeySu+wFnoycgXfElmcg== SRODATA dupok size=8
            0x0000 01 00 00 00 00 00 00 00                          ........  
</details>


# Генерация объектного файла
После выполнения команды ниже программа была скомпилирована в объектный файл main.o
```console
go tool compile main.go
```

# Генерация исполняемого файла
После выполнения команды ниже программа была скомпилирована и собрана в исполняемый файл main.exe
```console
go build main.go
```
