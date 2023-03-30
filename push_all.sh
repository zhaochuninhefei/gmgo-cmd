#!/bin/bash

echo "----- 所有远程仓库:"
git remote -v

echo "----- 向所有远程仓库推送..."
git push origin master
git push github master

echo "----- 结束"