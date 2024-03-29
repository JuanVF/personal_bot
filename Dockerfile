#Copyright 2023 Juan Jose Vargas Fletes
#
#This work is licensed under the Creative Commons Attribution-NonCommercial (CC BY-NC) license.
#To view a copy of this license, visit https://creativecommons.org/licenses/by-nc/4.0/
#
#Under the CC BY-NC license, you are free to:
#
#- Share: copy and redistribute the material in any medium or format
#- Adapt: remix, transform, and build upon the material
#
#Under the following terms:
#
#  - Attribution: You must give appropriate credit, provide a link to the license, and indicate if changes were made.
#    You may do so in any reasonable manner, but not in any way that suggests the licensor endorses you or your use.
#
#- Non-Commercial: You may not use the material for commercial purposes.
#
#You are free to use this work for personal or non-commercial purposes.
#If you would like to use this work for commercial purposes, please contact Juan Jose Vargas Fletes at juanvfletes@gmail.com.

FROM golang:1.20

WORKDIR /personal_bot

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 443

ENV ENVIRONMENT container

ENV AWS_ACCESS_KEY_ID set_your_key_id_access_key
ENV AWS_SECRET_ACCESS_KEY set_your_access_key


CMD ["./main"]