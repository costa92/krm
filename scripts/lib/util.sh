#!/usr/bin/env bash
# echo 密码 | sudo -S shell命令
function krm::sudo() {
    echo "${LINUX_PASSWORD}" | sudo -S $1
}