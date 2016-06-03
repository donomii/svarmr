#!/bin/bash
TORCH=/Users/jprice/torch/

IMG="$(cd "$(dirname "$1")"; pwd)/$(basename "$1")"
export LUA_PATH='/Users/jprice/.luarocks/share/lua/5.1/?.lua;/Users/jprice/.luarocks/share/lua/5.1/?/init.lua;$TORCH/install/share/lua/5.1/?.lua;$TORCH/install/share/lua/5.1/?/init.lua;./?.lua;$TORCH/install/share/luajit-2.1.0-beta1/?.lua;/usr/local/share/lua/5.1/?.lua;/usr/local/share/lua/5.1/?/init.lua'
export LUA_CPATH='/Users/jprice/.luarocks/lib/lua/5.1/?.so;/Users/jprice/torch/install/lib/lua/5.1/?.so;./?.so;/usr/local/lib/lua/5.1/?.so;/usr/local/lib/lua/5.1/loadall.so'
export PATH=$TORCH/install/bin:$PATH
export LD_LIBRARY_PATH=$TORCH/install/lib:$LD_LIBRARY_PATH
export DYLD_LIBRARY_PATH=$TORCH/install/lib:$DYLD_LIBRARY_PATH
export LUA_CPATH='$TORCH/install/lib/?.dylib;'$LUA_CPATH
cd $TORCH/overfeat-torch && th run.lua --network big --img $IMG
