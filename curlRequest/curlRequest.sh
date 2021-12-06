#!/bin/bash


echo "App №1"

echo "------------------------------------------------------"

echo "Первый Post запрос для добавления строки в БД"

curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"AplicationName":"1app","Param1":32,"Param2":"text"}' \
  http://localhost:8181/savestate

 echo "Смотрим сохранился ли файл"

 curl http://localhost:8181/getstate/1app

echo "Второй Post запрос для изменения той же строки в БД "

 curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"AplicationName":"1app","Param1":2,"Param2":"text"}' \
  http://localhost:8181/savestate

echo "Смотрим произошли ли изменения в БД"

 curl http://localhost:8181/getstate/1app

echo "Третий Post запрос без измененых данных для проверки того изменится ли в БД поле VERSION"

 curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"AplicationName":"1app","Param1":2,"Param2":"text"}' \
  http://localhost:8181/savestate

 echo "Смотрим произошли ли изменения в БД"

 curl http://localhost:8181/getstate/1app

echo "------------------------------------------------------"




echo "App №2"

echo "------------------------------------------------------"

echo "Первый Post запрос для добавления строки в БД"

curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"AplicationName":"2app","Param1":123,"Param2":"xD"}' \
  http://localhost:8181/savestate

 echo "Смотрим сохранился ли файл"

 curl http://localhost:8181/getstate/2app

echo "Второй Post запрос для изменения той же строки в БД "

curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"AplicationName":"2app","Param1":123,"Param2":"zxcvb"}' \
  http://localhost:8181/savestate

echo "Смотрим произошли ли изменения в БД"

 curl http://localhost:8181/getstate/2app

echo "Третий Post запрос без измененых данных для проверки того изменится ли в БД поле VERSION"

 curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"AplicationName":"2app","Param1":123,"Param2":"zxcvb"}' \
  http://localhost:8181/savestate

 echo "Смотрим произошли ли изменения в БД"

 curl http://localhost:8181/getstate/2app

echo "------------------------------------------------------"




echo "App №3"

echo "------------------------------------------------------"

echo "Первый Post запрос для добавления строки в БД"

curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"AplicationName":"3app","Param1":12,"Param2":"xD"}' \
  http://localhost:8181/savestate

 echo "Смотрим сохранился ли файл"

 curl http://localhost:8181/getstate/3app

echo "Второй Post запрос для изменения той же строки в БД "
  curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"AplicationName":"3app","Param1":1,"Param2":"zxcvb"}' \
  http://localhost:8181/savestate

echo "Смотрим произошли ли изменения в БД"

 curl http://localhost:8181/getstate/3app

echo "Третий Post запрос без измененых данных для проверки того изменится ли в БД поле VERSION"

  curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"AplicationName":"3app","Param1":1,"Param2":"zxcvb"}' \
  http://localhost:8181/savestate

 echo "Смотрим произошли ли изменения в БД"

 curl http://localhost:8181/getstate/3app

echo "------------------------------------------------------"


