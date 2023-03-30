#!/usr/bin/env bash

set -o nounset
set -o pipefail
set -o errexit

container_health_check() {
    local CHECK_COMMAND=("${@:2}") # store the command as an array and start accessing it from the second element of the array.
    local CONTAINER_NAME="$1"

    local CHECK_OK=0
    local MAX_CHECKS=10
    local CURRENT_CHECK=1

    echo ''
    echo "INFO: Container $CONTAINER_NAME check start !"

    set +o errexit

    while [ $CURRENT_CHECK -le $MAX_CHECKS ]; do
        if [ "$(docker inspect -f '{{.State.Running}}' "$CONTAINER_NAME")" = "true" ]; then
            "${CHECK_COMMAND[@]}" >/dev/null 2>&1 # call the command using the array form.
            local status=$?
            if [ $status -eq 0 ]; then
                CHECK_OK=1
                break
            fi
        else
            echo "Error: Container '$CONTAINER_NAME' is not running."
        fi

        CURRENT_CHECK=$((CURRENT_CHECK + 1))
        sleep 2
    done

    set -o errexit

    if [ "$CHECK_OK" -eq 1 ]; then
        echo "INFO: Container $CONTAINER_NAME check ok !"
        echo ''
    else
        # Output a message if the maximum number of checks was reached
        if [ $CURRENT_CHECK -gt $MAX_CHECKS ]; then
            echo "Error: The $CONTAINER_NAME could not be reached after $MAX_CHECKS checks."
        fi

        echo ''
        exit 1
    fi
}

normal_health_check() {
    local CHECK_COMMAND=("${@:1}") # store the command as an array and start accessing it from the second element of the array.
    local COMMAND_NAME="$1"

    local CHECK_OK=0
    local MAX_CHECKS=10
    local CURRENT_CHECK=1

    echo ''
    echo "INFO: COMMAND $COMMAND_NAME check start !"

    set +o errexit

    while [ $CURRENT_CHECK -le $MAX_CHECKS ]; do
        "${CHECK_COMMAND[@]}" # call the command using the array form.
        # "${CHECK_COMMAND[@]}" >/dev/null 2>&1 # call the command using the array form.
        local status=$?
        if [ $status -eq 0 ]; then
            CHECK_OK=1
            break
        fi
        CURRENT_CHECK=$((CURRENT_CHECK + 1))
        sleep 2
    done

    set -o errexit

    if [ "$CHECK_OK" -eq 1 ]; then
        echo "INFO: COMMAND $COMMAND_NAME check ok !"
        echo ''
    else
        # Output a message if the maximum number of checks was reached
        if [ $CURRENT_CHECK -gt $MAX_CHECKS ]; then
            echo "Error: The $COMMAND_NAME could not be reached after $MAX_CHECKS checks."
        fi

        echo ''
        exit 1
    fi
}

while getopts "cls" opt; do
    case $opt in
    c) # container
        docker-compose -p jubo up -d

        container_health_check \
            jubo-rdb \
            docker exec jubo-rdb psql -U caesar -d test -c '\l'

        docker exec -i jubo-rdb psql -U caesar -d test <./testdata.sql 2>&1

        echo -e "\n Project Website: \e[4m\e[33mhttp://localhost/\e[0m \n"
        ;;

    l) # local
        docker-compose -p jubo up -d

        normal_health_check \
            psql -U caesar -d test -c '\l'

        psql -U caesar -d test <./testdata.sql 2>&1
        ;;

    s) # stop
        docker-compose -p jubo down -v
        ;;

    *)
        echo "Invalid option: -$OPTARG" >&2
        exit 1
        ;;
    esac
done

# 將參數指標移回最初的位置，以便後續程式碼能夠正確解析其餘的命令列參數
# 使用 getopts 命令處理命令列參數時，$OPTIND 變數會記錄目前處理到的參數位置
shift $((OPTIND - 1))
