#!/usr/bin/env sh

echo -e "\e[1;32m ************/etc/hosts************ \e[0m"

cat /etc/hosts

echo -e "\e[1;32m ********/etc/resolv.conf********** \e[0m"

cat /etc/resolv.conf

echo -e "\e[1;32m **************容器架构************** \e[0m"

uname -a

echo -e "\e[1;32m ************当前容器时间************* \e[0m"

date

echo -e "\e[1;32m ************容器Hostname*********** \e[0m"

hostname

echo -e "\e[1;32m ************开始运行**************** \e[0m"
