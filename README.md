# stockbit

Jawaban :

1 . SELECT a.id, a.name, b.name AS parent FROM `user` AS a JOIN `user` AS b ON a.id = b.parent

2 . clone repository dan run command go run ./app
    untuk menjalani unit test go tes 
    1 . cd app
    2 . go test ./...

3 . refactor code : go run ./refactor

4 . logic code : go ./logic