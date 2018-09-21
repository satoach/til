#!/bin/bash

readonly TBL_CPU=cpu
readonly TBL_MEM=mem
readonly TBL_STORAGE=storage
readonly TBL_THERM=thermal
readonly TBL_PROC=ps

function exe_sql() {
    local readonly CONF=$HOME/.config/mysql/grafana.conf
    local readonly DBNAME=hw_stat
    mysql --defaults-extra-file=$CONF -e "$@" $DBNAME
}

function init_tbl() {
    exe_sql "truncate $1;"
}

function init_db() {
    init_tbl $TBL_CPU
    init_tbl $TBL_MEM
    init_tbl $TBL_STORAGE
    init_tbl $TBL_THERM
    init_tbl $TBL_PROC
}

function add_tbl() {
    tbl=$1
    val=$2
    exe_sql "insert into $tbl values(UNIX_TIMESTAMP(),'$tbl',$val);"
}

function add_db() {
    cpu_idle=$(vmstat | tail -n1 | tr -s ' ' '\t' | cut -f 16)
    cpu_per=$((100-cpu_idle))
    echo " cpu_per,$cpu_per"

    mem=$(free | grep Mem | tr -s ' '  '\t')
    mem_total=$(echo -e "$mem" | cut -f 2)
    mem_used=$(echo -e "$mem" | cut -f 3)
    mem_per=$((mem_used*100/mem_total))
    echo " mem_total,$mem_total"
    echo " mem_per,$mem_per"

    storage=$(df /dev/sda5 | tail -n1 | tr -s ' ' '\t')
    storage_total=$(echo -e "$storage" | cut -f 2)
    storage_used=$(echo -e "$storage" | cut -f 3)
    storage_per=$(echo -e "$storage" | cut -f 5 | tr -d '%')
    echo " storage_total,$storage_total"
    echo " storage_per,$storage_per"

    thermal=$(cat /sys/class/thermal/thermal_zone0/temp)
    thermal_deg=$(echo $((thermal/1000)))
    echo " thermal,$thermal"
    echo " thermal_deg,$thermal_deg"

    psnum=$(ps -ax | grep -v "ps -ax" | wc -l)
    echo " psnum,$psnum"

    add_tbl $TBL_CPU  $cpu_per
    add_tbl $TBL_MEM  $mem_per
    add_tbl $TBL_STORAGE $storage_per
    add_tbl $TBL_THERM $thermal_deg
    add_tbl $TBL_PROC $psnum
}

function get_latest_tbl() {
    tbl=$1
    exe_sql "select * from $tbl order by time_sec desc limit 1;"
}

function get_latest() {
    get_latest_tbl $TBL_CPU
    get_latest_tbl $TBL_MEM
    get_latest_tbl $TBL_STORAGE
    get_latest_tbl $TBL_THERM
    get_latest_tbl $TBL_PROC
}

case $1 in
    "add") add_db;;
    "init") init_db;;
    "latest") get_latest;;
    *) echo "USAGE: [init|add|latest]";exit 1;;
esac
